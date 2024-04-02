package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/ExocoreNetwork/exocore/testutil/network"
	"github.com/ExocoreNetwork/exocore/testutil/nullify"
	"github.com/ExocoreNetwork/exocore/x/oracle/client/cli"
	"github.com/ExocoreNetwork/exocore/x/oracle/types"
)

func networkWithIndexRecentMsgObjects(t *testing.T) (*network.Network, types.IndexRecentMsg) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	indexRecentMsg := &types.IndexRecentMsg{}
	nullify.Fill(&indexRecentMsg)
	state.IndexRecentMsg = indexRecentMsg
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.IndexRecentMsg
}

func TestShowIndexRecentMsg(t *testing.T) {
	net, obj := networkWithIndexRecentMsgObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.IndexRecentMsg
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowIndexRecentMsg(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetIndexRecentMsgResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.IndexRecentMsg)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.IndexRecentMsg),
				)
			}
		})
	}
}
