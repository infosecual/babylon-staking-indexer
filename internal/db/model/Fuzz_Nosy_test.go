package model

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon-staking-indexer/internal/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.mongodb.org/mongo-driver/mongo"
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

func Fuzz_Nosy_BTCDelegationDetails_HasInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var event *types.EventBTCDelegationCreated
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight int64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		var bbnBlockTime int64
		fill_err = tp.Fill(&bbnBlockTime)
		if fill_err != nil {
			return
		}
		if event == nil {
			return
		}

		d, err := FromEventBTCDelegationCreated(event, bbnBlockHeight, bbnBlockTime)
		if err != nil {
			return
		}
		d.HasInclusionProof()
	})
}

func Fuzz_Nosy_ToStateStrings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stateHistory []StateRecord
		fill_err = tp.Fill(&stateHistory)
		if fill_err != nil {
			return
		}

		ToStateStrings(stateHistory)
	})
}

func Fuzz_Nosy_createCollection__(f *testing.F) {
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
		var database *mongo.Database
		fill_err = tp.Fill(&database)
		if fill_err != nil {
			return
		}
		var collectionName string
		fill_err = tp.Fill(&collectionName)
		if fill_err != nil {
			return
		}
		if database == nil {
			return
		}

		createCollection(ctx, database, collectionName)
	})
}

func Fuzz_Nosy_createIndex__(f *testing.F) {
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
		var database *mongo.Database
		fill_err = tp.Fill(&database)
		if fill_err != nil {
			return
		}
		var collectionName string
		fill_err = tp.Fill(&collectionName)
		if fill_err != nil {
			return
		}
		var idx index
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if database == nil {
			return
		}

		createIndex(ctx, database, collectionName, idx)
	})
}
