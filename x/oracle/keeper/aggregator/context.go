package aggregator

import (
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/ExocoreNetwork/exocore/x/oracle/keeper/cache"
	"github.com/ExocoreNetwork/exocore/x/oracle/keeper/common"
	"github.com/ExocoreNetwork/exocore/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PriceItemKV struct {
	TokenID uint64
	PriceTR types.PriceTimeRound
}

type roundInfo struct {
	// this round of price will start from block basedBlock+1, the basedBlock served as a trigger to notify validators to submit prices
	basedBlock uint64
	// next round id of the price oracle service, price with the id will be record on block basedBlock+1 if all prices submitted by validators(for v1, validators serve as oracle nodes) get to consensus immediately
	nextRoundID uint64
	// indicate if this round is open for collecting prices or closed in either condition that success with a consensused price or not
	// 1: open, 2: closed
	status roundStatus
}

// roundStatus is an enum type to indicate the status of a roundInfo
type roundStatus int32

const (
	// roundStatusOpen indicates the round is open for collecting prices
	roundStatusOpen roundStatus = iota + 1
	// roundStatusClosed indicates the round is closed, either success with a consensused price or not
	roundStatusClosed
)

// AggregatorContext keeps memory cache for state params, validatorset, and updatedthese values as they updated on chain. And it keeps the information to track all tokenFeeders' status and data collection
// nolint
type AggregatorContext struct {
	params *types.Params

	// validator->power
	validatorsPower map[string]*big.Int
	totalPower      *big.Int

	// each active feederToken has a roundInfo
	rounds map[uint64]*roundInfo

	// each roundInfo has a worker
	aggregators map[uint64]*worker
}

func (agc *AggregatorContext) Copy4CheckTx() *AggregatorContext {
	ret := &AggregatorContext{
		// params, validatorsPower, totalPower, these values won't change during block executing
		params:          agc.params,
		validatorsPower: agc.validatorsPower,
		totalPower:      agc.totalPower,

		rounds:      make(map[uint64]*roundInfo),
		aggregators: make(map[uint64]*worker),
	}

	for k, v := range agc.rounds {
		vTmp := *v
		ret.rounds[k] = &vTmp
	}

	for k, v := range agc.aggregators {
		w := newWorker(k, ret)
		w.sealed = v.sealed
		w.price = v.price

		w.f = v.f.copy4CheckTx()
		w.c = v.c.copy4CheckTx()
		w.a = v.a.copy4CheckTx()
	}

	return ret
}

// sanity check for the msgCreatePrice
func (agc *AggregatorContext) sanityCheck(msg *types.MsgCreatePrice) error {
	// sanity check
	// TODO: check the msgCreatePrice's Decimal is correct with params setting
	// TODO: check len(price.prices)>0, len(price.prices._range_eachPriceSource.Prices)>0, at least has one source, and for each source has at least one price

	if accAddress, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errors.New("invalid address")
	} else if _, ok := agc.validatorsPower[sdk.ConsAddress(accAddress).String()]; !ok {
		return errors.New("signer is not validator")
	}

	if len(msg.Prices) == 0 {
		return errors.New("msg should provide at least one price")
	}

	for _, pSource := range msg.Prices {
		if len(pSource.Prices) == 0 || len(pSource.Prices) > int(common.MaxDetID) || !agc.params.IsValidSource(pSource.SourceID) {
			return errors.New("source should be valid and provide at least one price")
		}
		// check with params is coressponding source is deteministic
		if agc.params.IsDeterministicSource(pSource.SourceID) {
			for _, pDetID := range pSource.Prices {
				// TODO: verify the format of DetId is correct, since this is string, and we will make consensus with validator's power, so it's ok not to verify the format
				// just make sure the DetId won't mess up with NS's placeholder id, the limitation of maximum count one validator can submit will be check by filter
				if len(pDetID.DetID) == 0 {
					// deterministic must have specified deterministicId
					return errors.New("ds should have roundid")
				}
				// DS's price value will go through consensus process, so it's safe to skip the check here
			}
			// sanity check: NS submit only one price with detId==""
		} else if len(pSource.Prices) > 1 || len(pSource.Prices[0].DetID) > 0 {
			return errors.New("ns should not have roundid")
		}
	}
	return nil
}

func (agc *AggregatorContext) checkMsg(msg *types.MsgCreatePrice) error {
	if err := agc.sanityCheck(msg); err != nil {
		return err
	}

	// check feeder is active
	feederContext := agc.rounds[msg.FeederID]
	if feederContext == nil || feederContext.status != roundStatusOpen {
		// feederId does not exist or not alive
		return errors.New("context not exist or not available")
	}
	// senity check on basedBlock
	if msg.BasedBlock != feederContext.basedBlock {
		return errors.New("baseblock not match")
	}

	// check sources rule matches
	if ok, err := agc.params.CheckRules(msg.FeederID, msg.Prices); !ok {
		return err
	}

	for _, pSource := range msg.Prices {
		for _, pTimeDetID := range pSource.Prices {
			if ok := agc.params.CheckDecimal(msg.FeederID, pTimeDetID.Decimal); !ok {
				return fmt.Errorf("decimal not match for source ID %d and price ID %s", pSource.SourceID, pTimeDetID.DetID)
			}
		}
	}
	return nil
}

func (agc *AggregatorContext) FillPrice(msg *types.MsgCreatePrice) (*PriceItemKV, *cache.ItemM, error) {
	feederWorker := agc.aggregators[msg.FeederID]
	// worker initialzed here reduce workload for Endblocker
	if feederWorker == nil {
		feederWorker = newWorker(msg.FeederID, agc)
		agc.aggregators[msg.FeederID] = feederWorker
	}

	// if feederWorker.sealed {
	// 	// record this message for performance evaluate(used for slashing)
	// 	return nil, nil, types.ErrPriceProposalIgnored.Wrap("price aggregation for this round has sealed")
	// }

	// list4Calculator, list4Aggregator := w.f.filtrate(msg)
	if feederWorker.sealed {
		if _, list4Aggregator := feederWorker.filtrate(msg); list4Aggregator != nil {
			// record this message for performance evaluation(used for slashing)
			feederWorker.recordMessage(msg.Creator, msg.FeederID, list4Aggregator)
		}
		return nil, nil, types.ErrPriceProposalIgnored.Wrap("price aggregation for this round has sealed")
	}

	if listFilled := feederWorker.do(msg); listFilled != nil {
		// record this message for performance evaluation(used for slashing)
		feederWorker.recordMessage(msg.Creator, msg.FeederID, listFilled)
		if finalPrice := feederWorker.aggregate(); finalPrice != nil {
			agc.rounds[msg.FeederID].status = roundStatusClosed
			feederWorker.seal()
			return &PriceItemKV{agc.params.GetTokenFeeder(msg.FeederID).TokenID, types.PriceTimeRound{
				Price:   finalPrice.String(),
				Decimal: agc.params.GetTokenInfo(msg.FeederID).Decimal,
				// TODO: check the format
				Timestamp: msg.Prices[0].Prices[0].Timestamp,
				RoundID:   agc.rounds[msg.FeederID].nextRoundID,
			}}, &cache.ItemM{FeederID: msg.FeederID}, nil
		}
		return nil, &cache.ItemM{FeederID: msg.FeederID, PSources: listFilled, Validator: msg.Creator}, nil
	}

	// return nil, nil, errors.New("no valid price proposal to add for aggregation")
	return nil, nil, types.ErrPriceProposalIgnored
}

// NewCreatePrice receives msgCreatePrice message, and goes process: filter->aggregator, filter->calculator->aggregator
// non-deterministic data will goes directly into aggregator, and deterministic data will goes into calculator first to get consensus on the deterministic id.
func (agc *AggregatorContext) NewCreatePrice(_ sdk.Context, msg *types.MsgCreatePrice) (*PriceItemKV, *cache.ItemM, error) {
	if err := agc.checkMsg(msg); err != nil {
		return nil, nil, types.ErrInvalidMsg.Wrap(err.Error())
	}
	return agc.FillPrice(msg)
}

// prepare for new roundInfo, just update the status kept in memory
// executed at EndBlock stage, seall all success or expired roundInfo
// including possible aggregation and state update
// when validatorSet update, set force to true, to seal all alive round
// returns: 1st successful sealed, need to be written to KVStore, 2nd: failed sealed tokenID, use previous price to write to KVStore
func (agc *AggregatorContext) SealRound(ctx sdk.Context, force bool) (success []*PriceItemKV, failed []uint64, sealed []uint64, windowClosed []uint64) {
	feederIDs := make([]uint64, 0, len(agc.rounds))
	for fID, _ := range agc.rounds {
		feederIDs = append(feederIDs, fID)
	}
	sort.Slice(feederIDs, func(i, j int) bool {
		return feederIDs[i] < feederIDs[j]
	})
	// make sure feederIDs are accessed in order to calculate the indexOffset for slashing
	for _, feederID := range feederIDs {
		round := agc.rounds[feederID]
		if round.status == roundStatusOpen {
			feeder := agc.params.GetTokenFeeder(feederID)
			// TODO: for mode=1, we don't do aggregate() here, since if it donesn't success in the transaction execution stage, it won't success here
			// but it's not always the same for other modes, switch modes
			switch common.Mode {
			case types.ConsensusModeASAP:
				offset := uint64(ctx.BlockHeight()) - round.basedBlock
				expired := feeder.EndBlock > 0 && uint64(ctx.BlockHeight()) >= feeder.EndBlock
				outOfWindow := offset >= uint64(common.MaxNonce)

				// an open round reach its end of window, increase offsetIndex for active valdiator and chech the performance(missing/malicious)
				if round.status == roundStatusOpen && offset == uint64(common.MaxNonce) {
					windowClosed = append(windowClosed, feederID)
				}

				if expired || outOfWindow || force {
					failed = append(failed, feeder.TokenID)
					if expired {
						delete(agc.rounds, feederID)
					} else {
						round.status = roundStatusClosed
					}
					// TODO: optimize operformance
					sealed = append(sealed, feederID)
					delete(agc.aggregators, feederID)
				}
			default:
				ctx.Logger().Info("mode other than 1 is not support now")
			}
		}
		// all status: 1->2, remove its aggregator
		if agc.aggregators[feederID] != nil && agc.aggregators[feederID].sealed {
			delete(agc.aggregators, feederID)
			sealed = append(sealed, feederID)
		}
	}
	return success, failed, sealed, windowClosed
}

// PrepareEndBlock is called at EndBlock stage, to prepare the roundInfo for the next block(of input block)
// func (agc *AggregatorContext) PrepareRoundEndBlock(ctx sdk.Context, block uint64) {
func (agc *AggregatorContext) PrepareRoundEndBlock(block uint64) (newRoundFeederIDs []uint64) {
	if block < 1 {
		return newRoundFeederIDs
	}

	for feederID, feeder := range agc.params.GetTokenFeeders() {
		if feederID == 0 {
			continue
		}
		if (feeder.EndBlock > 0 && feeder.EndBlock <= block) || feeder.StartBaseBlock > block {
			// this feeder is inactive
			continue
		}

		delta := block - feeder.StartBaseBlock
		left := delta % feeder.Interval
		count := delta / feeder.Interval
		latestBasedblock := block - left
		latestNextRoundID := feeder.StartRoundID + count

		feederIDUint64 := uint64(feederID)
		round := agc.rounds[feederIDUint64]
		if round == nil {
			round = &roundInfo{
				basedBlock:  latestBasedblock,
				nextRoundID: latestNextRoundID,
			}
			if left >= uint64(common.MaxNonce) {
				// since do sealround properly before prepareRound, this only possible happens in node restart, and nonce has been taken care of in kvStore
				round.status = roundStatusClosed
			} else {
				round.status = roundStatusOpen
				if left == 0 {
					// set nonce for corresponding feederID for new roud start
					newRoundFeederIDs = append(newRoundFeederIDs, feederIDUint64)
				}
			}
			agc.rounds[feederIDUint64] = round
		} else {
			// prepare a new round for exist roundInfo
			if left == 0 {
				round.basedBlock = latestBasedblock
				round.nextRoundID = latestNextRoundID
				round.status = roundStatusOpen
				// set nonce for corresponding feederID for new roud start
				newRoundFeederIDs = append(newRoundFeederIDs, feederIDUint64)
				// drop previous worker
				delete(agc.aggregators, feederIDUint64)
			} else if round.status == roundStatusOpen && left >= uint64(common.MaxNonce) {
				// this shouldn't happen, if do sealround properly before prepareRound, basically for test only
				round.status = roundStatusClosed
				// TODO: just modify the status here, since sealRound should do all the related seal actions already when parepare invoked
			}
		}
	}
	return newRoundFeederIDs
}

// SetParams sets the params field of aggregatorContext“
func (agc *AggregatorContext) SetParams(p *types.Params) {
	agc.params = p
}

// SetValidatorPowers sets the map of validator's power for aggreagtorContext
func (agc *AggregatorContext) SetValidatorPowers(vp map[string]*big.Int) {
	//	t := big.NewInt(0)
	agc.totalPower = big.NewInt(0)
	agc.validatorsPower = make(map[string]*big.Int)
	for addr, power := range vp {
		agc.validatorsPower[addr] = power
		agc.totalPower = new(big.Int).Add(agc.totalPower, power)
	}
}

// GetValidatorPowers returns the map of validator's power stored in aggregatorContext
func (agc *AggregatorContext) GetValidatorPowers() (vp map[string]*big.Int) {
	return agc.validatorsPower
}

func (agc *AggregatorContext) GetValidators() (validators []string) {
	for k := range agc.validatorsPower {
		validators = append(validators, k)
	}
	return
}

// GetTokenIDFromAssetID returns tokenID for corresponding tokenID, it returns 0 if agc.params is nil or assetID not found in agc.params
func (agc *AggregatorContext) GetTokenIDFromAssetID(assetID string) int {
	if agc.params == nil {
		return 0
	}
	return agc.params.GetTokenIDFromAssetID(assetID)
}

// GetParams returns the params field of aggregatorContext
func (agc *AggregatorContext) GetParams() types.Params {
	return *agc.params
}

func (agc *AggregatorContext) GetParamsMaxSizePrices() uint64 {
	return uint64(agc.params.MaxSizePrices)
}

// GetFinalPriceListForFeederIDs get final price list for required feederIDs in format []{feederID, sourceID, detID, price} with asc of {feederID, sourceID}
// feederIDs is required to be ordered asc
func (agc *AggregatorContext) GetFinalPriceListForFeederIDs(feederIDs []uint64) []*types.AggFinalPrice {
	ret := make([]*types.AggFinalPrice, 0, len(feederIDs))
	for _, feederID := range feederIDs {
		feederWorker := agc.aggregators[feederID]
		if feederWorker != nil {
			if pList := feederWorker.getFinalPriceList(feederID); len(pList) > 0 {
				ret = append(ret, pList...)
			}
		}
	}
	return ret
}

func (agc *AggregatorContext) PerformanceReview(finalPrice *types.AggFinalPrice, validator string) (exist, matched bool) {
	feederWorker := agc.aggregators[finalPrice.FeederID]
	if feederWorker == nil {
		// should not happen
		// something wrong happened in the agc, so we just skip this round by treat the validator as effective&honest
		exist = true
		matched = true
		return
	}
	exist, matched = feederWorker.check(validator, finalPrice.FeederID, finalPrice.SourceID, finalPrice.Price, finalPrice.DetID)
	return
}

// NewAggregatorContext returns a new instance of AggregatorContext
func NewAggregatorContext() *AggregatorContext {
	return &AggregatorContext{
		validatorsPower: make(map[string]*big.Int),
		totalPower:      big.NewInt(0),
		rounds:          make(map[uint64]*roundInfo),
		aggregators:     make(map[uint64]*worker),
	}
}
