package types

import (
	"testing"

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

func Fuzz_Nosy_Error_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statusCode int
		fill_err = tp.Fill(&statusCode)
		if fill_err != nil {
			return
		}
		var errorCode ErrorCode
		fill_err = tp.Fill(&errorCode)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		e := NewErrorWithMsg(statusCode, errorCode, msg)
		e.Error()
	})
}

func Fuzz_Nosy_DelegationState_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s DelegationState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_DelegationSubState_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DelegationSubState
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.String()
	})
}

func Fuzz_Nosy_ErrorCode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e ErrorCode
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.String()
	})
}

func Fuzz_Nosy_EventType_ShortName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e EventType
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.ShortName()
	})
}

func Fuzz_Nosy_EventType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e EventType
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.String()
	})
}

func Fuzz_Nosy_QualifiedStatesForCovenantQuorumReached__(f *testing.F) {
	f.Fuzz(func(t *testing.T, babylonState string) {
		QualifiedStatesForCovenantQuorumReached(babylonState)
	})
}

func Fuzz_Nosy_QualifiedStatesForInclusionProofReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, babylonState string) {
		QualifiedStatesForInclusionProofReceived(babylonState)
	})
}

func Fuzz_Nosy_QualifiedStatesForWithdrawable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var subState DelegationSubState
		fill_err = tp.Fill(&subState)
		if fill_err != nil {
			return
		}

		QualifiedStatesForWithdrawable(subState)
	})
}
