package btcclient

import (
	"testing"

	"github.com/babylonlabs-io/babylon-staking-indexer/internal/config"
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

func Fuzz_Nosy_BTCClient_GetBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := NewBTCClient(cfg)
		if err != nil {
			return
		}
		c.GetBlockTimestamp(height)
	})
}

func Fuzz_Nosy_BTCClient_GetTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := NewBTCClient(cfg)
		if err != nil {
			return
		}
		c.GetTipHeight()
	})
}

func Fuzz_Nosy_EmptyHintCache_CommitConfirmHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var confRequests []chainntnfs.ConfRequest
		fill_err = tp.Fill(&confRequests)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CommitConfirmHint(height, confRequests...)
	})
}

func Fuzz_Nosy_EmptyHintCache_CommitSpendHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var spendRequests []chainntnfs.SpendRequest
		fill_err = tp.Fill(&spendRequests)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CommitSpendHint(height, spendRequests...)
	})
}

func Fuzz_Nosy_EmptyHintCache_PurgeConfirmHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var confRequests []chainntnfs.ConfRequest
		fill_err = tp.Fill(&confRequests)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PurgeConfirmHint(confRequests...)
	})
}

func Fuzz_Nosy_EmptyHintCache_PurgeSpendHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var spendRequests []chainntnfs.SpendRequest
		fill_err = tp.Fill(&spendRequests)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PurgeSpendHint(spendRequests...)
	})
}

func Fuzz_Nosy_EmptyHintCache_QueryConfirmHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var confRequest chainntnfs.ConfRequest
		fill_err = tp.Fill(&confRequest)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.QueryConfirmHint(confRequest)
	})
}

func Fuzz_Nosy_EmptyHintCache_QuerySpendHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var spendRequest chainntnfs.SpendRequest
		fill_err = tp.Fill(&spendRequest)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.QuerySpendHint(spendRequest)
	})
}

// skipping Fuzz_Nosy_BtcInterface_GetBlockTimestamp__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/clients/btcclient.BtcInterface

// skipping Fuzz_Nosy_BtcInterface_GetTipHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon-staking-indexer/internal/clients/btcclient.BtcInterface

func Fuzz_Nosy_BuildDialer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rpcHost string) {
		BuildDialer(rpcHost)
	})
}

// skipping Fuzz_Nosy_clientCallWithRetry__ because parameters include func, chan, or unsupported interface: github.com/avast/retry-go/v4.RetryableFuncWithData[*T]
