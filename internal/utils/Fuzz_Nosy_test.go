package utils

import (
	"testing"

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

func Fuzz_Nosy_SupportedBtcNetwork_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c SupportedBtcNetwork
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.String()
	})
}

func Fuzz_Nosy_GetFunctionName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth int) {
		GetFunctionName(depth)
	})
}

func Fuzz_Nosy_GetWrappedTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *wire.MsgBlock
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		GetWrappedTxs(msg)
	})
}

func Fuzz_Nosy_ParseUint32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		ParseUint32(s)
	})
}

func Fuzz_Nosy_RandomAlphaNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, length int) {
		RandomAlphaNum(length)
	})
}

func Fuzz_Nosy_SafeUnescape__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		SafeUnescape(s)
	})
}

func Fuzz_Nosy_SerializeBtcTransaction__(f *testing.F) {
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
		if tx == nil {
			return
		}

		SerializeBtcTransaction(tx)
	})
}

func Fuzz_Nosy_shortFuncName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fullName string) {
		shortFuncName(fullName)
	})
}
