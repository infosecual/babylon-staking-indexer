package services

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon-staking-indexer/internal/clients/bbnclient"
	"github.com/babylonlabs-io/babylon-staking-indexer/internal/db/model"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/chainntnfs"
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

func Fuzz_Nosy_Service_ResubscribeToMissedBtcNotifications__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ResubscribeToMissedBtcNotifications(ctx)
	})
}

func Fuzz_Nosy_Service_StartBbnBlockProcessor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartBbnBlockProcessor(ctx)
	})
}

func Fuzz_Nosy_Service_StartExpiryChecker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartExpiryChecker(ctx)
	})
}

func Fuzz_Nosy_Service_StartIndexerSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartIndexerSync(ctx)
	})
}

func Fuzz_Nosy_Service_SubscribeToBbnEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SubscribeToBbnEvents(ctx)
	})
}

func Fuzz_Nosy_Service_SyncGlobalParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SyncGlobalParams(ctx)
	})
}

func Fuzz_Nosy_Service_checkExpiry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.checkExpiry(ctx)
	})
}

func Fuzz_Nosy_Service_doProcessEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event BbnEvent
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var blockHeight int64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.doProcessEvent(ctx, event, blockHeight)
	})
}

func Fuzz_Nosy_Service_emitActiveDelegationEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || delegation == nil {
			return
		}

		s.emitActiveDelegationEvent(ctx, delegation)
	})
}

func Fuzz_Nosy_Service_emitUnbondingDelegationEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || delegation == nil {
			return
		}

		s.emitUnbondingDelegationEvent(ctx, delegation)
	})
}

func Fuzz_Nosy_Service_emitWithdrawableDelegationEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || delegation == nil {
			return
		}

		s.emitWithdrawableDelegationEvent(ctx, delegation)
	})
}

func Fuzz_Nosy_Service_emitWithdrawnDelegationEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || delegation == nil {
			return
		}

		s.emitWithdrawnDelegationEvent(ctx, delegation)
	})
}

func Fuzz_Nosy_Service_fetchAndSaveParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.fetchAndSaveParams(ctx)
	})
}

func Fuzz_Nosy_Service_getEventsFromBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHeight int64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getEventsFromBlock(ctx, blockHeight)
	})
}

func Fuzz_Nosy_Service_getLatestHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var initialHeight int64
		fill_err = tp.Fill(&initialHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getLatestHeight(initialHeight)
	})
}

func Fuzz_Nosy_Service_handleSpendingStakingTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var spendingTx *wire.MsgTx
		fill_err = tp.Fill(&spendingTx)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil || spendingTx == nil {
			return
		}

		s.handleSpendingStakingTransaction(ctx, spendingTx, spendingInputIdx, spendingHeight, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Service_handleSpendingUnbondingTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var spendingTx *wire.MsgTx
		fill_err = tp.Fill(&spendingTx)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || spendingTx == nil || delegation == nil {
			return
		}

		s.handleSpendingUnbondingTransaction(ctx, spendingTx, spendingHeight, spendingInputIdx, delegation)
	})
}

func Fuzz_Nosy_Service_handleWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var subState types.DelegationSubState
		fill_err = tp.Fill(&subState)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}
		if s == nil || delegation == nil {
			return
		}

		s.handleWithdrawal(ctx, delegation, subState, spendingHeight)
	})
}

func Fuzz_Nosy_Service_isSpendingStakingTxSlashingPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.isSpendingStakingTxSlashingPath(tx, spendingInputIdx, delegation, params)
	})
}

func Fuzz_Nosy_Service_isSpendingStakingTxTimeLockPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.isSpendingStakingTxTimeLockPath(tx, spendingInputIdx, delegation, params)
	})
}

func Fuzz_Nosy_Service_isSpendingStakingTxUnbondingPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.isSpendingStakingTxUnbondingPath(tx, delegation, params)
	})
}

func Fuzz_Nosy_Service_isSpendingUnbondingTxSlashingPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.isSpendingUnbondingTxSlashingPath(tx, delegation, spendingInputIdx, params)
	})
}

func Fuzz_Nosy_Service_isSpendingUnbondingTxTimeLockPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.isSpendingUnbondingTxTimeLockPath(tx, delegation, spendingInputIdx, params)
	})
}

func Fuzz_Nosy_Service_processBTCDelegationExpiredEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight int64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processBTCDelegationExpiredEvent(ctx, event, bbnBlockHeight)
	})
}

func Fuzz_Nosy_Service_processBTCDelegationInclusionProofReceivedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight int64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processBTCDelegationInclusionProofReceivedEvent(ctx, event, bbnBlockHeight)
	})
}

func Fuzz_Nosy_Service_processBTCDelegationUnbondedEarlyEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight int64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processBTCDelegationUnbondedEarlyEvent(ctx, event, bbnBlockHeight)
	})
}

func Fuzz_Nosy_Service_processBlocksSequentially__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processBlocksSequentially(ctx)
	})
}

func Fuzz_Nosy_Service_processCovenantQuorumReachedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight int64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processCovenantQuorumReachedEvent(ctx, event, bbnBlockHeight)
	})
}

func Fuzz_Nosy_Service_processCovenantSignatureReceivedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processCovenantSignatureReceivedEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_processEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event BbnEvent
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var blockHeight int64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processEvent(ctx, event, blockHeight)
	})
}

func Fuzz_Nosy_Service_processFinalityProviderEditedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processFinalityProviderEditedEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_processFinalityProviderStateChangeEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processFinalityProviderStateChangeEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_processNewBTCDelegationEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight int64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processNewBTCDelegationEvent(ctx, event, bbnBlockHeight)
	})
}

func Fuzz_Nosy_Service_processNewFinalityProviderEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processNewFinalityProviderEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_quitContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.quitContext()
	})
}

func Fuzz_Nosy_Service_registerStakingSpendNotification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
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
		var stakingTxHex string
		fill_err = tp.Fill(&stakingTxHex)
		if fill_err != nil {
			return
		}
		var stakingOutputIdx uint32
		fill_err = tp.Fill(&stakingOutputIdx)
		if fill_err != nil {
			return
		}
		var stakingStartHeight uint32
		fill_err = tp.Fill(&stakingStartHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.registerStakingSpendNotification(ctx, stakingTxHashHex, stakingTxHex, stakingOutputIdx, stakingStartHeight)
	})
}

func Fuzz_Nosy_Service_registerUnbondingSpendNotification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || delegation == nil {
			return
		}

		s.registerUnbondingSpendNotification(ctx, delegation)
	})
}

func Fuzz_Nosy_Service_startWatchingSlashingChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var slashingTx *wire.MsgTx
		fill_err = tp.Fill(&slashingTx)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var subState types.DelegationSubState
		fill_err = tp.Fill(&subState)
		if fill_err != nil {
			return
		}
		if s == nil || slashingTx == nil || delegation == nil {
			return
		}

		s.startWatchingSlashingChange(ctx, slashingTx, spendingHeight, delegation, subState)
	})
}

func Fuzz_Nosy_Service_validateBTCDelegationCreatedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var event *types.EventBTCDelegationCreated
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil || event == nil {
			return
		}

		s.validateBTCDelegationCreatedEvent(event)
	})
}

func Fuzz_Nosy_Service_validateBTCDelegationExpiredEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event *types.EventBTCDelegationExpired
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil || event == nil {
			return
		}

		s.validateBTCDelegationExpiredEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_validateBTCDelegationInclusionProofReceivedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event *types.EventBTCDelegationInclusionProofReceived
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil || event == nil {
			return
		}

		s.validateBTCDelegationInclusionProofReceivedEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_validateBTCDelegationUnbondedEarlyEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event *types.EventBTCDelgationUnbondedEarly
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil || event == nil {
			return
		}

		s.validateBTCDelegationUnbondedEarlyEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_validateCovenantQuorumReachedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var event *types.EventCovenantQuorumReached
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if s == nil || event == nil {
			return
		}

		s.validateCovenantQuorumReachedEvent(ctx, event)
	})
}

func Fuzz_Nosy_Service_validateFinalityProviderCreatedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var fpCreated *types.EventFinalityProviderCreated
		fill_err = tp.Fill(&fpCreated)
		if fill_err != nil {
			return
		}
		if s == nil || fpCreated == nil {
			return
		}

		s.validateFinalityProviderCreatedEvent(fpCreated)
	})
}

func Fuzz_Nosy_Service_validateFinalityProviderEditedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var fpEdited *types.EventFinalityProviderEdited
		fill_err = tp.Fill(&fpEdited)
		if fill_err != nil {
			return
		}
		if s == nil || fpEdited == nil {
			return
		}

		s.validateFinalityProviderEditedEvent(fpEdited)
	})
}

func Fuzz_Nosy_Service_validateFinalityProviderStateChangeEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpStateChange *types.EventFinalityProviderStatusChange
		fill_err = tp.Fill(&fpStateChange)
		if fill_err != nil {
			return
		}
		if s == nil || fpStateChange == nil {
			return
		}

		s.validateFinalityProviderStateChangeEvent(ctx, fpStateChange)
	})
}

func Fuzz_Nosy_Service_validateUnbondingTxOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var params *bbnclient.StakingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.validateUnbondingTxOutput(tx, delegation, params)
	})
}

func Fuzz_Nosy_Service_watchForSpendSlashingChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var spendEvent *chainntnfs.SpendEvent
		fill_err = tp.Fill(&spendEvent)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var subState types.DelegationSubState
		fill_err = tp.Fill(&subState)
		if fill_err != nil {
			return
		}
		if s == nil || spendEvent == nil || delegation == nil {
			return
		}

		s.watchForSpendSlashingChange(spendEvent, delegation, subState)
	})
}

func Fuzz_Nosy_Service_watchForSpendStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var spendEvent *chainntnfs.SpendEvent
		fill_err = tp.Fill(&spendEvent)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil || spendEvent == nil {
			return
		}

		s.watchForSpendStakingTx(spendEvent, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Service_watchForSpendUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var spendEvent *chainntnfs.SpendEvent
		fill_err = tp.Fill(&spendEvent)
		if fill_err != nil {
			return
		}
		var delegation *model.BTCDelegationDetails
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		if s == nil || spendEvent == nil || delegation == nil {
			return
		}

		s.watchForSpendUnbondingTx(spendEvent, delegation)
	})
}

func Fuzz_Nosy_btcNotifierWithRetries_RegisterSpendNtfn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *btcNotifierWithRetries
		fill_err = tp.Fill(&b)
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
		if b == nil || outpoint == nil {
			return
		}

		b.RegisterSpendNtfn(outpoint, pkScript, heightHint)
	})
}

func Fuzz_Nosy_btcNotifierWithRetries_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *btcNotifierWithRetries
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Start()
	})
}

// skipping Fuzz_Nosy_BtcNotifier_RegisterSpendNtfn__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/services.BtcNotifier

// skipping Fuzz_Nosy_BtcNotifier_Start__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/services.BtcNotifier

func Fuzz_Nosy_parseEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var expectedType types.EventType
		fill_err = tp.Fill(&expectedType)
		if fill_err != nil {
			return
		}
		var event types.Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}

		parseEvent(expectedType, event)
	})
}
