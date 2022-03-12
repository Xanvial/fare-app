package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/Xanvial/fare-app/internal/model"
)

func TestNewFare(t *testing.T) {
	type args struct {
		cfg model.Config
	}
	tests := []struct {
		name string
		args args
		want FareUsecase
	}{
		{
			name: "function call",
			want: &fare{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFare(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fare_SetFareData(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	type args struct {
		input model.FareData
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "function call",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			f.SetFareData(tt.args.input)
		})
	}
}

func Test_fare_ResetFareData(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "function call",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			f.ResetFareData()
		})
	}
}

func Test_fare_GetLatestData(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   model.FareData
	}{
		{
			name: "function call",
			fields: fields{
				latestFare: model.FareData{
					Distance: 123,
				},
				dataCount: 5,
			},
			want: model.FareData{
				Distance: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			if got := f.GetLatestData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fare.GetLatestData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fare_GetDataCount(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "function call",
			fields: fields{
				latestFare: model.FareData{
					Distance: 123,
				},
				dataCount: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			if got := f.GetDataCount(); got != tt.want {
				t.Errorf("fare.GetDataCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fare_GetMinDataCount(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "function call",
			fields: fields{
				cfg: model.FareConfig{
					MinData: 7,
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			if got := f.GetMinDataCount(); got != tt.want {
				t.Errorf("fare.GetMinDataCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fare_IsDataValid(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "datacount is less than min and distance is zero",
			fields: fields{
				cfg: model.FareConfig{
					MinData: 7,
				},
				latestFare: model.FareData{
					Distance: 0,
				},
				dataCount: 3,
			},
			want: false,
		},
		{
			name: "datacount is less than min and distance is more than zero",
			fields: fields{
				cfg: model.FareConfig{
					MinData: 7,
				},
				latestFare: model.FareData{
					Distance: 66,
				},
				dataCount: 3,
			},
			want: false,
		},
		{
			name: "datacount is more than min and distance is zero",
			fields: fields{
				cfg: model.FareConfig{
					MinData: 7,
				},
				latestFare: model.FareData{
					Distance: 0,
				},
				dataCount: 13,
			},
			want: false,
		},
		{
			name: "datacount is more than min and distance is more than zero",
			fields: fields{
				cfg: model.FareConfig{
					MinData: 7,
				},
				latestFare: model.FareData{
					Distance: 10,
				},
				dataCount: 13,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			if got := f.IsDataValid(); got != tt.want {
				t.Errorf("fare.IsDataValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fare_ValidateFareData(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	type args struct {
		input model.FareData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "input is older than latest data",
			fields: fields{
				latestFare: model.FareData{
					ElapsedTime: time.Minute,
				},
			},
			args: args{
				input: model.FareData{
					ElapsedTime: time.Second,
				},
			},
			wantErr: true,
		},
		{
			name: "input and latest data interval is more than maximum",
			fields: fields{
				latestFare: model.FareData{
					ElapsedTime: time.Minute,
				},
				cfg: model.FareConfig{
					MaxIntervalSec: 100,
				},
			},
			args: args{
				input: model.FareData{
					ElapsedTime: 3 * time.Minute,
				},
			},
			wantErr: true,
		},
		{
			name: "input distance is less than latest data",
			fields: fields{
				latestFare: model.FareData{
					ElapsedTime: time.Minute,
					Distance:    100,
				},
				cfg: model.FareConfig{
					MaxIntervalSec: 200,
				},
			},
			args: args{
				input: model.FareData{
					ElapsedTime: 3 * time.Minute,
					Distance:    90,
				},
			},
			wantErr: true,
		},
		{
			name: "valid data",
			fields: fields{
				latestFare: model.FareData{
					ElapsedTime: time.Minute,
					Distance:    100,
				},
				cfg: model.FareConfig{
					MaxIntervalSec: 200,
				},
			},
			args: args{
				input: model.FareData{
					ElapsedTime: 3 * time.Minute,
					Distance:    190,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			if err := f.ValidateFareData(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("fare.ValidateFareData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fare_CalculateFare(t *testing.T) {
	type fields struct {
		latestFare model.FareData
		dataCount  int
		cfg        model.FareConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "empty data",
			wantErr: true,
		},
		{
			name: "calculation sample 1",
			fields: fields{
				latestFare: model.FareData{
					Distance: 1800,
				},
				dataCount: 5,
				cfg: model.FareConfig{
					BaseFare:               400,
					BaseFareDistance:       1000,
					LimitDistance:          10000,
					UnderLimitFare:         40,
					UnderLimitFareDistance: 400,
					OverLimitFare:          40,
					OverLimitFareDistance:  350,
				},
			},
			want: 480,
		},
		{
			name: "calculation sample 2, with over limit distance",
			fields: fields{
				latestFare: model.FareData{
					Distance: 11200,
				},
				dataCount: 5,
				cfg: model.FareConfig{
					BaseFare:               400,
					BaseFareDistance:       1000,
					LimitDistance:          10000,
					UnderLimitFare:         40,
					UnderLimitFareDistance: 400,
					OverLimitFare:          40,
					OverLimitFareDistance:  350,
				},
			},
			want: 1440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fare{
				latestFare: tt.fields.latestFare,
				dataCount:  tt.fields.dataCount,
				cfg:        tt.fields.cfg,
			}
			got, err := f.CalculateFare()
			if (err != nil) != tt.wantErr {
				t.Errorf("fare.CalculateFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fare.CalculateFare() = %v, want %v", got, tt.want)
			}
		})
	}
}
