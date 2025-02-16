package db

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon-staking-indexer/internal/clients/bbnclient"
	"github.com/babylonlabs-io/babylon-staking-indexer/internal/config"
	"github.com/babylonlabs-io/babylon-staking-indexer/internal/db/model"
	"github.com/babylonlabs-io/babylon-staking-indexer/internal/types"
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

func Fuzz_Nosy_Database_DeleteExpiredDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.DeleteExpiredDelegation(c3, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Database_FindExpiredDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.FindExpiredDelegations(c3, btcTipHeight, limit)
	})
}

func Fuzz_Nosy_Database_GetBTCDelegationByStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetBTCDelegationByStakingTxHash(c3, stakingTxHash)
	})
}

func Fuzz_Nosy_Database_GetBTCDelegationState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetBTCDelegationState(c3, stakingTxHash)
	})
}

func Fuzz_Nosy_Database_GetBTCDelegationsByStates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var states []types.DelegationState
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetBTCDelegationsByStates(c3, states)
	})
}

func Fuzz_Nosy_Database_GetDelegationsByFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var fpBTCPKHex string
		fill_err = tp.Fill(&fpBTCPKHex)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetDelegationsByFinalityProvider(c3, fpBTCPKHex)
	})
}

func Fuzz_Nosy_Database_GetFinalityProviderByBtcPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var btcPk string
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetFinalityProviderByBtcPk(c3, btcPk)
	})
}

func Fuzz_Nosy_Database_GetLastProcessedBbnHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetLastProcessedBbnHeight(c3)
	})
}

func Fuzz_Nosy_Database_GetStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetStakingParams(c3, version)
	})
}

func Fuzz_Nosy_Database_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.Ping(c3)
	})
}

func Fuzz_Nosy_Database_SaveBTCDelegationSlashingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var slashingTxHex string
		fill_err = tp.Fill(&slashingTxHex)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveBTCDelegationSlashingTxHex(c3, stakingTxHash, slashingTxHex, spendingHeight)
	})
}

func Fuzz_Nosy_Database_SaveBTCDelegationUnbondingCovenantSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveBTCDelegationUnbondingCovenantSignature(c3, stakingTxHash, covenantBtcPkHex, signatureHex)
	})
}

func Fuzz_Nosy_Database_SaveBTCDelegationUnbondingSlashingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var unbondingSlashingTxHex string
		fill_err = tp.Fill(&unbondingSlashingTxHex)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveBTCDelegationUnbondingSlashingTxHex(c3, stakingTxHash, unbondingSlashingTxHex, spendingHeight)
	})
}

func Fuzz_Nosy_Database_SaveCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var params *bbnclient.CheckpointParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveCheckpointParams(c3, params)
	})
}

func Fuzz_Nosy_Database_SaveNewBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var delegationDoc *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegationDoc)
		if fill_err != nil {
			return
		}
		if delegationDoc == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveNewBTCDelegation(c3, delegationDoc)
	})
}

func Fuzz_Nosy_Database_SaveNewFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var fpDoc *model.FinalityProviderDetails
		fill_err = tp.Fill(&fpDoc)
		if fill_err != nil {
			return
		}
		if fpDoc == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveNewFinalityProvider(c3, fpDoc)
	})
}

func Fuzz_Nosy_Database_SaveNewTimeLockExpire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveNewTimeLockExpire(c3, stakingTxHashHex, expireHeight, subState)
	})
}

func Fuzz_Nosy_Database_SaveStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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
		if params == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveStakingParams(c3, version, params)
	})
}

// skipping Fuzz_Nosy_Database_UpdateBTCDelegationState__ because parameters include func, chan, or unsupported interface: []github.com/babylonlabs-io/babylon-staking-indexer/internal/db.UpdateOption

func Fuzz_Nosy_Database_UpdateFinalityProviderDetailsFromEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var detailsToUpdate *model.FinalityProviderDetails
		fill_err = tp.Fill(&detailsToUpdate)
		if fill_err != nil {
			return
		}
		if detailsToUpdate == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.UpdateFinalityProviderDetailsFromEvent(c3, detailsToUpdate)
	})
}

func Fuzz_Nosy_Database_UpdateFinalityProviderState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.UpdateFinalityProviderState(c3, btcPk, newState)
	})
}

func Fuzz_Nosy_Database_UpdateLastProcessedBbnHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.UpdateLastProcessedBbnHeight(c3, height)
	})
}

func Fuzz_Nosy_Database_collection__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		d1, err := New(ctx, cfg)
		if err != nil {
			return
		}
		d1.collection(name)
	})
}

func Fuzz_Nosy_DuplicateKeyError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *DuplicateKeyError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_DuplicateKeyError_Is__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_InvalidPaginationTokenError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *InvalidPaginationTokenError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_InvalidPaginationTokenError_Is__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_NotFoundError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *NotFoundError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_NotFoundError_Is__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DbInterface_DeleteExpiredDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_FindExpiredDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetBTCDelegationByStakingTxHash__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetBTCDelegationState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetBTCDelegationsByStates__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetDelegationsByFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetFinalityProviderByBtcPk__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetLastProcessedBbnHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetStakingParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_Ping__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveBTCDelegationUnbondingCovenantSignature__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveCheckpointParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveNewBTCDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveNewFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveNewTimeLockExpire__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveStakingParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_UpdateBTCDelegationState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_UpdateFinalityProviderDetailsFromEvent__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_UpdateFinalityProviderState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_UpdateLastProcessedBbnHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/db.DbInterface

// skipping Fuzz_Nosy_IsDuplicateKeyError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_IsInvalidPaginationTokenError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_IsNotFoundError__ because parameters include func, chan, or unsupported interface: error
