package model

import (
	"reflect"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestReadConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "function call",
			want: Config{
				LogConfig: Log{
					LogLevel: int(zapcore.InfoLevel),
					OutputPaths: []string{
						"./fare-app.log",
					},
				},
				Fare: FareConfig{
					MinData:        2,
					MaxIntervalSec: 600,

					BaseFare:         400,
					BaseFareDistance: 1000,

					LimitDistance:          10000,
					UnderLimitFare:         40,
					UnderLimitFareDistance: 400,
					OverLimitFare:          40,
					OverLimitFareDistance:  350,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
