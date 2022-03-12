package model

import (
	"reflect"
	"testing"
	"time"

	mock_zapcore "github.com/Xanvial/fare-app/internal/model/mock"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
)

func TestFareData_MarshalLogObject(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockZapObj := mock_zapcore.NewMockEncoder(ctl)

	type fields struct {
		ElapsedTime time.Duration
		Distance    float64
	}
	type args struct {
		enc zapcore.ObjectEncoder
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		expects *gomock.Call
		wantErr bool
	}{
		{
			name: "function call",
			args: args{
				enc: mockZapObj,
			},
			expects: mockZapObj.EXPECT().AddString(gomock.Any(), gomock.Any()).
				Times(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := FareData{
				ElapsedTime: tt.fields.ElapsedTime,
				Distance:    tt.fields.Distance,
			}
			if err := fd.MarshalLogObject(tt.args.enc); (err != nil) != tt.wantErr {
				t.Errorf("FareData.MarshalLogObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseInputText(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    FareData
		wantErr bool
	}{
		{
			name:    "empty input",
			wantErr: true,
		},
		{
			name: "only one part of data",
			args: args{
				input: "test",
			},
			wantErr: true,
		},
		{
			name: "time part is invalid",
			args: args{
				input: "12:01 123.4",
			},
			wantErr: true,
		},
		{
			name: "hour part is invalid",
			args: args{
				input: "xx:01:01.123 123.4",
			},
			wantErr: true,
		},
		{
			name: "minute part is invalid",
			args: args{
				input: "12:xx:01.123 123.4",
			},
			wantErr: true,
		},
		{
			name: "second part is invalid",
			args: args{
				input: "12:11:xx.123 123.4",
			},
			wantErr: true,
		},
		{
			name: "distance part is invalid",
			args: args{
				input: "12:11:22.123 xxx",
			},
			wantErr: true,
		},
		{
			name: "valid data",
			args: args{
				input: "12:11:56.123 123.4",
			},
			want: FareData{
				ElapsedTime: 12*time.Hour + 11*time.Minute + 56123*time.Millisecond,
				Distance:    123.4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInputText(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInputText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInputText() = %v, want %v", got, tt.want)
			}
		})
	}
}
