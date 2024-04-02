package keeper_test

import (
	"math/big"
	reflect "reflect"

	"bou.ke/monkey"
	math "cosmossdk.io/math"
	"github.com/ExocoreNetwork/exocore/x/oracle/keeper"
	"github.com/ExocoreNetwork/exocore/x/oracle/keeper/cache"
	"github.com/ExocoreNetwork/exocore/x/oracle/keeper/common"
	"github.com/ExocoreNetwork/exocore/x/oracle/keeper/testdata"
	"github.com/ExocoreNetwork/exocore/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//go:generate mockgen -destination mock_validator_test.go -package keeper_test github.com/cosmos/cosmos-sdk/x/staking/types ValidatorI

var _ = Describe("MsgCreatePrice", func() {
	var operator1, operator2, operator3 sdk.ValAddress
	var c *cache.Cache
	BeforeEach(func() {
		ks.Reset()
		Expect(ks.ms).ToNot(BeNil())

		validatorC := NewMockValidatorI(ks.ctrl)
		validatorC.EXPECT().GetBondedTokens().Return(math.NewInt(1))
		validatorC.EXPECT().GetBondedTokens().Return(math.NewInt(1))
		validatorC.EXPECT().GetBondedTokens().Return(math.NewInt(1))

		validatorC.EXPECT().GetConsensusPower(gomock.Any()).Return(int64(1))
		validatorC.EXPECT().GetConsensusPower(gomock.Any()).Return(int64(1))
		validatorC.EXPECT().GetConsensusPower(gomock.Any()).Return(int64(1))

		privVal1 := mock.NewPV()
		pubKey1, _ := privVal1.GetPubKey()
		operator1 = sdk.ValAddress(pubKey1.Address())

		privVal2 := mock.NewPV()
		pubKey2, _ := privVal2.GetPubKey()
		operator2 = sdk.ValAddress(pubKey2.Address())

		privVal3 := mock.NewPV()
		pubKey3, _ := privVal3.GetPubKey()
		operator3 = sdk.ValAddress(pubKey3.Address())

		validatorC.EXPECT().GetOperator().Return(operator1)
		validatorC.EXPECT().GetOperator().Return(operator2)
		validatorC.EXPECT().GetOperator().Return(operator3)

		monkey.PatchInstanceMethod(reflect.TypeOf(keeper.Keeper{}), "IterateBondedValidatorsByPower", func(k keeper.Keeper, ctx sdk.Context, f func(index int64, validator stakingtypes.ValidatorI) bool) {
			f(0, validatorC)
			f(0, validatorC)
			f(0, validatorC)
		})
		monkey.PatchInstanceMethod(reflect.TypeOf(keeper.Keeper{}), "GetLastTotalPower", func(k keeper.Keeper, ctx sdk.Context) *big.Int { return big.NewInt(3) })

		Expect(ks.ctx.BlockHeight()).To(Equal(int64(2)))
	})

	AfterEach(func() {
		ks.ctrl.Finish()
	})

	Context("3 validators with 1 voting power each", func() {
		BeforeEach(func() {
			ks.ms.CreatePrice(ks.ctx, &types.MsgCreatePrice{
				Creator:    operator1.String(),
				FeederId:   1,
				Prices:     testdata.PS1,
				BasedBlock: 1,
				Nonce:      1,
			})

			c = keeper.GetCaches()
			pRes := &common.Params{}
			c.GetCache(cache.CacheItemP(pRes))
			Expect(*pRes).Should(BeEquivalentTo(types.DefaultParams()))
		})

		It("success on 3rd message", func() {

			iRes := make([]*cache.CacheItemM, 0)
			c.GetCache(&iRes)
			Expect(iRes[0].Validator).Should(Equal(operator1.String()))

			ks.ms.CreatePrice(ks.ctx, &types.MsgCreatePrice{
				Creator:    operator2.String(),
				FeederId:   1,
				Prices:     testdata.PS2,
				BasedBlock: 1,
				Nonce:      1,
			},
			)
			ks.ms.CreatePrice(ks.ctx, &types.MsgCreatePrice{})
			c.GetCache(&iRes)
			Expect(len(iRes)).Should(Equal(2))

			ks.ms.CreatePrice(ks.ctx, &types.MsgCreatePrice{
				Creator:    operator3.String(),
				FeederId:   1,
				Prices:     testdata.PS4,
				BasedBlock: 1,
				Nonce:      1,
			},
			)
			c.GetCache(&iRes)
			Expect(len(iRes)).Should(Equal(0))
			prices := ks.k.GetAllPrices(sdk.UnwrapSDKContext(ks.ctx))
			//fmt.Println("GetAllPrices", prices[0])
			Expect(prices[0]).Should(BeEquivalentTo(types.Prices{
				TokenId:     1,
				NextRountId: 2,
				PriceList: []*types.PriceWithTimeAndRound{
					{
						Price:     testdata.PTD2.Price,
						Decimal:   testdata.PTD2.Decimal,
						Timestamp: prices[0].PriceList[0].Timestamp,
						RoundId:   1,
					},
				},
			}))
		})
	})
})