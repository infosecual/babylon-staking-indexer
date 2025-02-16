package metrics

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

func Fuzz_Nosy_Outcome_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var O Outcome
		fill_err = tp.Fill(&O)
		if fill_err != nil {
			return
		}

		O.String()
	})
}

func Fuzz_Nosy_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, metricsPort int) {
		Init(metricsPort)
	})
}

// skipping Fuzz_Nosy_RecordBtcClientMetrics__ because parameters include func, chan, or unsupported interface: func() (T, error)

func Fuzz_Nosy_StartClientRequestDurationTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, baseUrl string, method string, path string) {
		StartClientRequestDurationTimer(baseUrl, method, path)
	})
}

func Fuzz_Nosy_initMetricsRouter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, metricsPort int) {
		initMetricsRouter(metricsPort)
	})
}
