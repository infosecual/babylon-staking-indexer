package mocks

import (
	context "context"
	"testing"
	time "time"

	bbnclient "github.com/babylonlabs-io/babylon-staking-indexer/internal/clients/bbnclient"
	model "github.com/babylonlabs-io/babylon-staking-indexer/internal/db/model"
	types "github.com/babylonlabs-io/babylon-staking-indexer/internal/types"
	wire "github.com/btcsuite/btcd/wire"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_BbnInterface_GetAllStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetAllStakingParams(ctx)
	})
}

func Fuzz_Nosy_BbnInterface_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHeight *int64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		if _m == nil || blockHeight == nil {
			return
		}

		_m.GetBlock(ctx, blockHeight)
	})
}

func Fuzz_Nosy_BbnInterface_GetBlockResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHeight *int64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		if _m == nil || blockHeight == nil {
			return
		}

		_m.GetBlockResults(ctx, blockHeight)
	})
}

func Fuzz_Nosy_BbnInterface_GetCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetCheckpointParams(ctx)
	})
}

func Fuzz_Nosy_BbnInterface_GetLatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLatestBlockNumber(ctx)
	})
}

func Fuzz_Nosy_BbnInterface_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IsRunning()
	})
}

func Fuzz_Nosy_BbnInterface_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Start()
	})
}

func Fuzz_Nosy_BbnInterface_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		var query string
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var healthCheckInterval time.Duration
		fill_err = tp.Fill(&healthCheckInterval)
		if fill_err != nil {
			return
		}
		var maxEventWaitInterval time.Duration
		fill_err = tp.Fill(&maxEventWaitInterval)
		if fill_err != nil {
			return
		}
		var outCapacity []int
		fill_err = tp.Fill(&outCapacity)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Subscribe(subscriber, query, healthCheckInterval, maxEventWaitInterval, outCapacity...)
	})
}

func Fuzz_Nosy_BbnInterface_UnsubscribeAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BbnInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.UnsubscribeAll(subscriber)
	})
}

func Fuzz_Nosy_BtcInterface_GetBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBlockTimestamp(height)
	})
}

func Fuzz_Nosy_BtcInterface_GetTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetTipHeight()
	})
}

func Fuzz_Nosy_BtcNotifier_RegisterSpendNtfn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcNotifier
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var outpoint *wire.OutPoint
		fill_err = tp.Fill(&outpoint)
		if fill_err != nil {
			return
		}
		var pkScript []byte
		fill_err = tp.Fill(&pkScript)
		if fill_err != nil {
			return
		}
		var heightHint uint32
		fill_err = tp.Fill(&heightHint)
		if fill_err != nil {
			return
		}
		if _m == nil || outpoint == nil {
			return
		}

		_m.RegisterSpendNtfn(outpoint, pkScript, heightHint)
	})
}

func Fuzz_Nosy_BtcNotifier_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcNotifier
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Start()
	})
}

func Fuzz_Nosy_DbInterface_DeleteExpiredDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.DeleteExpiredDelegation(ctx, stakingTxHashHex)
	})
}

func Fuzz_Nosy_DbInterface_FindExpiredDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var btcTipHeight uint64
		fill_err = tp.Fill(&btcTipHeight)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindExpiredDelegations(ctx, btcTipHeight, limit)
	})
}

func Fuzz_Nosy_DbInterface_GetBTCDelegationByStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBTCDelegationByStakingTxHash(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_DbInterface_GetBTCDelegationState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBTCDelegationState(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_DbInterface_GetBTCDelegationsByStates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var states []types.DelegationState
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBTCDelegationsByStates(ctx, states)
	})
}

func Fuzz_Nosy_DbInterface_GetDelegationsByFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetDelegationsByFinalityProvider(ctx, fpBtcPkHex)
	})
}

func Fuzz_Nosy_DbInterface_GetFinalityProviderByBtcPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var btcPk string
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetFinalityProviderByBtcPk(ctx, btcPk)
	})
}

func Fuzz_Nosy_DbInterface_GetLastProcessedBbnHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetLastProcessedBbnHeight(ctx)
	})
}

func Fuzz_Nosy_DbInterface_GetStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetStakingParams(ctx, version)
	})
}

func Fuzz_Nosy_DbInterface_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Ping(ctx)
	})
}

func Fuzz_Nosy_DbInterface_SaveBTCDelegationUnbondingCovenantSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var covenantBtcPkHex string
		fill_err = tp.Fill(&covenantBtcPkHex)
		if fill_err != nil {
			return
		}
		var signatureHex string
		fill_err = tp.Fill(&signatureHex)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveBTCDelegationUnbondingCovenantSignature(ctx, stakingTxHash, covenantBtcPkHex, signatureHex)
	})
}

func Fuzz_Nosy_DbInterface_SaveCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params *bbnclient.CheckpointParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _m == nil || params == nil {
			return
		}

		_m.SaveCheckpointParams(ctx, params)
	})
}

func Fuzz_Nosy_DbInterface_SaveNewBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegationDoc *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegationDoc)
		if fill_err != nil {
			return
		}
		if _m == nil || delegationDoc == nil {
			return
		}

		_m.SaveNewBTCDelegation(ctx, delegationDoc)
	})
}

func Fuzz_Nosy_DbInterface_SaveNewFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpDoc *model.FinalityProviderDetails
		fill_err = tp.Fill(&fpDoc)
		if fill_err != nil {
			return
		}
		if _m == nil || fpDoc == nil {
			return
		}

		_m.SaveNewFinalityProvider(ctx, fpDoc)
	})
}

func Fuzz_Nosy_DbInterface_SaveNewTimeLockExpire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var expireHeight uint32
		fill_err = tp.Fill(&expireHeight)
		if fill_err != nil {
			return
		}
		var subState types.DelegationSubState
		fill_err = tp.Fill(&subState)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveNewTimeLockExpire(ctx, stakingTxHashHex, expireHeight, subState)
	})
}

func Fuzz_Nosy_DbInterface_SaveStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _m == nil || params == nil {
			return
		}

		_m.SaveStakingParams(ctx, version, params)
	})
}

// skipping Fuzz_Nosy_DbInterface_UpdateBTCDelegationState__ because parameters include func, chan, or unsupported interface: []github.com/babylonlabs-io/babylon-staking-indexer/internal/db.UpdateOption

func Fuzz_Nosy_DbInterface_UpdateFinalityProviderDetailsFromEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var detailsToUpdate *model.FinalityProviderDetails
		fill_err = tp.Fill(&detailsToUpdate)
		if fill_err != nil {
			return
		}
		if _m == nil || detailsToUpdate == nil {
			return
		}

		_m.UpdateFinalityProviderDetailsFromEvent(ctx, detailsToUpdate)
	})
}

func Fuzz_Nosy_DbInterface_UpdateFinalityProviderState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var btcPk string
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var newState string
		fill_err = tp.Fill(&newState)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.UpdateFinalityProviderState(ctx, btcPk, newState)
	})
}

func Fuzz_Nosy_DbInterface_UpdateLastProcessedBbnHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.UpdateLastProcessedBbnHeight(ctx, height)
	})
}
