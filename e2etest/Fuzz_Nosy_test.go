package e2etest

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon-staking-indexer/e2etest/container"
	"github.com/babylonlabs-io/babylon/btcstaking"
	"github.com/babylonlabs-io/babylon/testutil/datagen"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_BitcoindTestHandler_CreateWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var manager *container.Manager
		fill_err = tp.Fill(&manager)
		if fill_err != nil {
			return
		}
		var walletName string
		fill_err = tp.Fill(&walletName)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if t1 == nil || manager == nil {
			return
		}

		h := NewBitcoindHandler(t1, manager)
		h.CreateWallet(walletName, passphrase)
	})
}

func Fuzz_Nosy_BitcoindTestHandler_GenerateBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var manager *container.Manager
		fill_err = tp.Fill(&manager)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if t1 == nil || manager == nil {
			return
		}

		h := NewBitcoindHandler(t1, manager)
		h.GenerateBlocks(count)
	})
}

func Fuzz_Nosy_BitcoindTestHandler_GetBlockCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var manager *container.Manager
		fill_err = tp.Fill(&manager)
		if fill_err != nil {
			return
		}
		if t1 == nil || manager == nil {
			return
		}

		h := NewBitcoindHandler(t1, manager)
		h.GetBlockCount()
	})
}

func Fuzz_Nosy_BitcoindTestHandler_ImportDescriptors__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var manager *container.Manager
		fill_err = tp.Fill(&manager)
		if fill_err != nil {
			return
		}
		var descriptor string
		fill_err = tp.Fill(&descriptor)
		if fill_err != nil {
			return
		}
		if t1 == nil || manager == nil {
			return
		}

		h := NewBitcoindHandler(t1, manager)
		h.ImportDescriptors(descriptor)
	})
}

func Fuzz_Nosy_BitcoindTestHandler_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var manager *container.Manager
		fill_err = tp.Fill(&manager)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if t1 == nil || manager == nil || t3 == nil {
			return
		}

		h := NewBitcoindHandler(t1, manager)
		h.Start(t3)
	})
}

func Fuzz_Nosy_BitcoindTestHandler_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var manager *container.Manager
		fill_err = tp.Fill(&manager)
		if fill_err != nil {
			return
		}
		if t1 == nil || manager == nil {
			return
		}

		h := NewBitcoindHandler(t1, manager)
		h.Stop()
	})
}

func Fuzz_Nosy_TestManager_CatchUpBTCLightClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.CatchUpBTCLightClient(t4)
	})
}

func Fuzz_Nosy_TestManager_CheckNextActiveStakingEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.CheckNextActiveStakingEvent(t4, stakingTxHashHex)
	})
}

func Fuzz_Nosy_TestManager_CheckNextUnbondingStakingEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.CheckNextUnbondingStakingEvent(t4, stakingTxHashHex)
	})
}

func Fuzz_Nosy_TestManager_CreateBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || fpSK == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.CreateBTCDelegation(t4, fpSK)
	})
}

func Fuzz_Nosy_TestManager_CreateBTCDelegationWithoutIncl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || fpSK == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.CreateBTCDelegationWithoutIncl(t4, fpSK)
	})
}

func Fuzz_Nosy_TestManager_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.CreateFinalityProvider(t4)
	})
}

func Fuzz_Nosy_TestManager_InsertBTCHeadersToBabylon__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var headers []*wire.BlockHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.InsertBTCHeadersToBabylon(headers)
	})
}

func Fuzz_Nosy_TestManager_MustGetBabylonSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.MustGetBabylonSigner()
	})
}

func Fuzz_Nosy_TestManager_RetrieveTransactionFromMempool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var hashes []*chainhash.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.RetrieveTransactionFromMempool(t4, hashes)
	})
}

func Fuzz_Nosy_TestManager_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.Stop(t4)
	})
}

func Fuzz_Nosy_TestManager_SubmitInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var txInfo *types.TransactionInfo
		fill_err = tp.Fill(&txInfo)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || txInfo == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.SubmitInclusionProof(t4, stakingTxHash, txInfo)
	})
}

// skipping Fuzz_Nosy_TestManager_Undelegate__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_TestManager_WaitForDelegationStored__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
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
		var expectedState types.DelegationState
		fill_err = tp.Fill(&expectedState)
		if fill_err != nil {
			return
		}
		var expectedSubState *types.DelegationSubState
		fill_err = tp.Fill(&expectedSubState)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || expectedSubState == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.WaitForDelegationStored(t4, ctx, stakingTxHashHex, expectedState, expectedSubState)
	})
}

func Fuzz_Nosy_TestManager_WaitForFinalityProviderStored__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpPKHex string
		fill_err = tp.Fill(&fpPKHex)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.WaitForFinalityProviderStored(t4, ctx, fpPKHex)
	})
}

func Fuzz_Nosy_TestManager_addCovenantSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var signerAddr string
		fill_err = tp.Fill(&signerAddr)
		if fill_err != nil {
			return
		}
		var stakingMsgTx *wire.MsgTx
		fill_err = tp.Fill(&stakingMsgTx)
		if fill_err != nil {
			return
		}
		var stakingMsgTxHash *chainhash.Hash
		fill_err = tp.Fill(&stakingMsgTxHash)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		var slashingSpendPath *btcstaking.SpendInfo
		fill_err = tp.Fill(&slashingSpendPath)
		if fill_err != nil {
			return
		}
		var stakingSlashingInfo *datagen.TestStakingSlashingInfo
		fill_err = tp.Fill(&stakingSlashingInfo)
		if fill_err != nil {
			return
		}
		var unbondingSlashingInfo *datagen.TestUnbondingSlashingInfo
		fill_err = tp.Fill(&unbondingSlashingInfo)
		if fill_err != nil {
			return
		}
		var unbondingSlashingPathSpendInfo *btcstaking.SpendInfo
		fill_err = tp.Fill(&unbondingSlashingPathSpendInfo)
		if fill_err != nil {
			return
		}
		var stakingOutIdx uint32
		fill_err = tp.Fill(&stakingOutIdx)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || stakingMsgTx == nil || stakingMsgTxHash == nil || fpSK == nil || slashingSpendPath == nil || stakingSlashingInfo == nil || unbondingSlashingInfo == nil || unbondingSlashingPathSpendInfo == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.addCovenantSig(t4, signerAddr, stakingMsgTx, stakingMsgTxHash, fpSK, slashingSpendPath, stakingSlashingInfo, unbondingSlashingInfo, unbondingSlashingPathSpendInfo, stakingOutIdx)
	})
}

func Fuzz_Nosy_TestManager_createStakingAndSlashingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		var bsParams *types.QueryParamsResponse
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var covenantBtcPks []*btcec.PublicKey
		fill_err = tp.Fill(&covenantBtcPks)
		if fill_err != nil {
			return
		}
		var topUTXO *types.UTXO
		fill_err = tp.Fill(&topUTXO)
		if fill_err != nil {
			return
		}
		var stakingValue int64
		fill_err = tp.Fill(&stakingValue)
		if fill_err != nil {
			return
		}
		var stakingTimeBlocks uint32
		fill_err = tp.Fill(&stakingTimeBlocks)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || fpSK == nil || bsParams == nil || topUTXO == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.createStakingAndSlashingTx(t4, fpSK, bsParams, covenantBtcPks, topUTXO, stakingValue, stakingTimeBlocks)
	})
}

func Fuzz_Nosy_TestManager_createUnbondingData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var fpPK *btcec.PublicKey
		fill_err = tp.Fill(&fpPK)
		if fill_err != nil {
			return
		}
		var bsParams *types.QueryParamsResponse
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var covenantBtcPks []*btcec.PublicKey
		fill_err = tp.Fill(&covenantBtcPks)
		if fill_err != nil {
			return
		}
		var stakingSlashingInfo *datagen.TestStakingSlashingInfo
		fill_err = tp.Fill(&stakingSlashingInfo)
		if fill_err != nil {
			return
		}
		var stakingMsgTxHash *chainhash.Hash
		fill_err = tp.Fill(&stakingMsgTxHash)
		if fill_err != nil {
			return
		}
		var stakingOutIdx uint32
		fill_err = tp.Fill(&stakingOutIdx)
		if fill_err != nil {
			return
		}
		var stakingTimeBlocks uint32
		fill_err = tp.Fill(&stakingTimeBlocks)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil || fpPK == nil || bsParams == nil || stakingSlashingInfo == nil || stakingMsgTxHash == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.createUnbondingData(t4, fpPK, bsParams, covenantBtcPks, stakingSlashingInfo, stakingMsgTxHash, stakingOutIdx, stakingTimeBlocks)
	})
}

func Fuzz_Nosy_TestManager_getBTCUnbondingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.getBTCUnbondingTime(t4)
	})
}

func Fuzz_Nosy_TestManager_getHighUTXOAndSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.getHighUTXOAndSum()
	})
}

func Fuzz_Nosy_TestManager_mineBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var numMatureOutputsInWallet uint32
		fill_err = tp.Fill(&numMatureOutputsInWallet)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		tm := StartManager(t1, numMatureOutputsInWallet, epochInterval)
		tm.mineBlock(t4)
	})
}

func Fuzz_Nosy_bbnPksToBtcPks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []types.BIP340PubKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}

		bbnPksToBtcPks(pks)
	})
}

func Fuzz_Nosy_importPrivateKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcHandler *BitcoindTestHandler
		fill_err = tp.Fill(&btcHandler)
		if fill_err != nil {
			return
		}
		if btcHandler == nil {
			return
		}

		importPrivateKey(btcHandler)
	})
}

func Fuzz_Nosy_outIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var candOut *wire.TxOut
		fill_err = tp.Fill(&candOut)
		if fill_err != nil {
			return
		}
		if tx == nil || candOut == nil {
			return
		}

		outIdx(tx, candOut)
	})
}

func Fuzz_Nosy_tempDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		tempDir(t1)
	})
}
