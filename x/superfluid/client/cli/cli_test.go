package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"

	superfluidtypes "github.com/osmosis-labs/osmosis/v12/x/superfluid/types"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/osmosis-labs/osmosis/v12/app"
	lockuptestutil "github.com/osmosis-labs/osmosis/v12/x/lockup/client/testutil"
	"github.com/osmosis-labs/osmosis/v12/x/superfluid/client/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	s.cfg = app.DefaultConfig()

	genesisState := app.ModuleBasics.DefaultGenesis(s.cfg.Codec)
	superfluidGen := superfluidtypes.DefaultGenesis()
	minimum_risk_factor, err := sdk.NewDecFromStr("0.01")
	s.Require().NoError(err)

	superfluidGen.Params.MinimumRiskFactor = minimum_risk_factor
	superfluidGenJson := s.cfg.Codec.MustMarshalJSON(superfluidGen)
	genesisState[superfluidtypes.ModuleName] = superfluidGenJson
	s.cfg.GenesisState = genesisState

	s.network = network.New(s.T(), s.cfg)
	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	val := s.network.Validators[0]
	clientContext := val.ClientCtx

	info, _, err := val.ClientCtx.Keyring.NewMnemonic("TestingAcc",
		keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	newAddr := sdk.AccAddress(info.GetPubKey().Address())
	newAddr, err = sdk.AccAddressFromBech32(newAddr.String())
	s.Require().NoError(err)

	dayLockAmt, err := sdk.ParseCoinNormalized(fmt.Sprintf("200%s", s.network.Config.BondDenom))
	fmt.Println("DEBUG:  ", newAddr.String())

	lockuptestutil.MsgLockTokens(clientContext, val.Address, dayLockAmt, "24h")
	defaultLockId := "1"
	out, err := clitestutil.ExecTestCLICmd(clientContext, cli.NewSuperfluidDelegateCmd(), []string{
		defaultLockId,
		val.ValAddress.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr.String()),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	})
	fmt.Println("DEBUG:  ", newAddr, out)
}
func (s *IntegrationTestSuite) TestGetCmdQueryParams() {
	val := s.network.Validators[0]

	cmd := cli.GetCmdQueryParams()
	clientCtx := val.ClientCtx

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{})
	s.Require().NoError(err)
	s.Require().NotNil(out)
}

func (s *IntegrationTestSuite) TestGetCmdAllSuperfluidAssets() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"query all superfluid assets",
			[]string{
				fmt.Sprintf("--%s=%s", tmcli.OutputFlag, "json"),
			},
			false,
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.GetCmdAllSuperfluidAssets()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{})
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				fmt.Println(out.String())
				s.Require().NoError(err, out.String())
				s.Require().NotNil(out.String())
			}
		})
	}
}

// func (s *IntegrationTestSuite) TestGetCmdAssetMultiplier() {
// 	val := s.network.Validators[0]

// 	cmd := cli.GetCmdAssetMultiplier()
// 	// clientCtx := val.ClientCtx

//		// out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{})
//	}
func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
