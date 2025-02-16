package bbnclient

import (
	"context"
	"testing"
	"time"

	"github.com/babylonlabs-io/babylon-staking-indexer/internal/config"
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

func Fuzz_Nosy_BBNClient_GetAllStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetAllStakingParams(ctx)
	})
}

func Fuzz_Nosy_BBNClient_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
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
		if c == nil || blockHeight == nil {
			return
		}

		c.GetBlock(ctx, blockHeight)
	})
}

func Fuzz_Nosy_BBNClient_GetBlockResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
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
		if c == nil || blockHeight == nil {
			return
		}

		c.GetBlockResults(ctx, blockHeight)
	})
}

func Fuzz_Nosy_BBNClient_GetCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetCheckpointParams(ctx)
	})
}

func Fuzz_Nosy_BBNClient_GetLatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetLatestBlockNumber(ctx)
	})
}

func Fuzz_Nosy_BBNClient_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IsRunning()
	})
}

func Fuzz_Nosy_BBNClient_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Start()
	})
}

func Fuzz_Nosy_BBNClient_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
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
		if c == nil {
			return
		}

		c.Subscribe(subscriber, query, healthCheckInterval, maxEventWaitInterval, outCapacity...)
	})
}

func Fuzz_Nosy_BBNClient_UnsubscribeAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BBNClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnsubscribeAll(subscriber)
	})
}

func Fuzz_Nosy_BbnInterface_GetAllStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.GetAllStakingParams(ctx)
	})
}

func Fuzz_Nosy_BbnInterface_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
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
		if cfg == nil || blockHeight == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.GetBlock(ctx, blockHeight)
	})
}

func Fuzz_Nosy_BbnInterface_GetBlockResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
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
		if cfg == nil || blockHeight == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.GetBlockResults(ctx, blockHeight)
	})
}

func Fuzz_Nosy_BbnInterface_GetCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.GetCheckpointParams(ctx)
	})
}

func Fuzz_Nosy_BbnInterface_GetLatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.GetLatestBlockNumber(ctx)
	})
}

func Fuzz_Nosy_BbnInterface_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.IsRunning()
	})
}

func Fuzz_Nosy_BbnInterface_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.Start()
	})
}

func Fuzz_Nosy_BbnInterface_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
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
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.Subscribe(subscriber, query, healthCheckInterval, maxEventWaitInterval, outCapacity...)
	})
}

func Fuzz_Nosy_BbnInterface_UnsubscribeAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BBNConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewBBNClient(cfg)
		_x1.UnsubscribeAll(subscriber)
	})
}

// skipping Fuzz_Nosy_clientCallWithRetry__ because parameters include func, chan, or unsupported interface: github.com/avast/retry-go/v4.RetryableFuncWithData[*T]
