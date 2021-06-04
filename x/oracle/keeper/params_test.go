package keeper_test

import (
	"testing"

	"github.com/bandprotocol/chain/testing/testapp"
	"github.com/bandprotocol/chain/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetSetParams(t *testing.T) {
	_, ctx, k := testapp.CreateTestInput(true)
	expectedParams := types.Params{
		MaxRawRequestCount:      1,
		MaxAskCount:             10,
		ExpirationBlockCount:    30,
		BaseOwasmGas:            50000,
		PerValidatorRequestGas:  3000,
		SamplingTryCount:        3,
		OracleRewardPercentage:  50,
		InactivePenaltyDuration: 1000,
		IBCRequestEnabled:       true,
	}
	k.SetParams(ctx, expectedParams)
	require.Equal(t, expectedParams, k.GetParams(ctx))
	require.Equal(t, expectedParams.MaxRawRequestCount, k.MaxRawRequestCount(ctx))
	require.Equal(t, expectedParams.MaxAskCount, k.MaxAskCount(ctx))
	require.Equal(t, expectedParams.ExpirationBlockCount, k.ExpirationBlockCount(ctx))
	require.Equal(t, expectedParams.BaseOwasmGas, k.BaseOwasmGas(ctx))
	require.Equal(t, expectedParams.PerValidatorRequestGas, k.PerValidatorRequestGas(ctx))
	require.Equal(t, expectedParams.SamplingTryCount, k.SamplingTryCount(ctx))
	require.Equal(t, expectedParams.OracleRewardPercentage, k.OracleRewardPercentage(ctx))
	require.Equal(t, expectedParams.InactivePenaltyDuration, k.InactivePenaltyDuration(ctx))
	require.Equal(t, expectedParams.IBCRequestEnabled, k.IBCRequestEnabled(ctx))
	expectedParams = types.Params{
		MaxRawRequestCount:      2,
		MaxAskCount:             20,
		ExpirationBlockCount:    40,
		BaseOwasmGas:            150000,
		PerValidatorRequestGas:  30000,
		SamplingTryCount:        5,
		OracleRewardPercentage:  80,
		InactivePenaltyDuration: 10000,
		IBCRequestEnabled:       false,
	}
	k.SetParams(ctx, expectedParams)
	require.Equal(t, expectedParams, k.GetParams(ctx))
	require.Equal(t, expectedParams.MaxRawRequestCount, k.MaxRawRequestCount(ctx))
	require.Equal(t, expectedParams.MaxAskCount, k.MaxAskCount(ctx))
	require.Equal(t, expectedParams.ExpirationBlockCount, k.ExpirationBlockCount(ctx))
	require.Equal(t, expectedParams.BaseOwasmGas, k.BaseOwasmGas(ctx))
	require.Equal(t, expectedParams.PerValidatorRequestGas, k.PerValidatorRequestGas(ctx))
	require.Equal(t, expectedParams.SamplingTryCount, k.SamplingTryCount(ctx))
	require.Equal(t, expectedParams.OracleRewardPercentage, k.OracleRewardPercentage(ctx))
	require.Equal(t, expectedParams.InactivePenaltyDuration, k.InactivePenaltyDuration(ctx))
	require.Equal(t, expectedParams.IBCRequestEnabled, k.IBCRequestEnabled(ctx))
}
