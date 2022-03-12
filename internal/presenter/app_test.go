package presenter

import (
	"bytes"
	"errors"
	"io"
	"testing"

	mock_usecase "github.com/Xanvial/fare-app/internal/usecase/mock"
	"github.com/golang/mock/gomock"
)

func TestApp_ProcessInput(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	fareMock := mock_usecase.NewMockFareUsecase(ctl)

	type args struct {
		inp string
	}
	tests := []struct {
		name    string
		args    args
		expects []*gomock.Call
		wantErr bool
	}{
		{
			name:    "invalid input",
			wantErr: true,
		},
		{
			name: "error on calling validate fare",
			args: args{
				inp: "12:11:56.123 123.4",
			},
			expects: []*gomock.Call{
				fareMock.EXPECT().ValidateFareData(gomock.Any()).
					Return(errors.New("error unit test")),
				fareMock.EXPECT().GetLatestData(),
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				inp: "12:11:56.123 123.4",
			},
			expects: []*gomock.Call{
				fareMock.EXPECT().ValidateFareData(gomock.Any()).
					Return(nil),
				fareMock.EXPECT().SetFareData(gomock.Any()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				fare: fareMock,
			}
			if err := a.ProcessInput(tt.args.inp); (err != nil) != tt.wantErr {
				t.Errorf("App.ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_ProcessData(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	fareMock := mock_usecase.NewMockFareUsecase(ctl)

	tests := []struct {
		name    string
		expects []*gomock.Call
		wantErr bool
	}{
		{
			name: "invalid data",
			expects: []*gomock.Call{
				fareMock.EXPECT().IsDataValid().Return(false),
				fareMock.EXPECT().GetLatestData(),
				fareMock.EXPECT().GetDataCount(),
				fareMock.EXPECT().GetMinDataCount(),
			},
			wantErr: true,
		},
		{
			name: "valid data, error on calculation",
			expects: []*gomock.Call{
				fareMock.EXPECT().IsDataValid().Return(true),
				fareMock.EXPECT().CalculateFare().
					Return(0, errors.New("error unit test")),
				fareMock.EXPECT().GetLatestData(),
				fareMock.EXPECT().GetDataCount(),
			},
			wantErr: true,
		},
		{
			name: "valid data",
			expects: []*gomock.Call{
				fareMock.EXPECT().IsDataValid().Return(true),
				fareMock.EXPECT().CalculateFare().Return(123, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				fare: fareMock,
			}
			if err := a.ProcessData(); (err != nil) != tt.wantErr {
				t.Errorf("App.ProcessData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	fareMock := mock_usecase.NewMockFareUsecase(ctl)
	// fareMock.EXPECT().ResetFareData().Times(2)

	// this will simulate:
	// first line is "testdata"
	// second line is "test2"
	// third line is empty line which will stop the process
	var r io.Reader = bytes.NewBufferString("testdata\ntest2\n\n")
	// process the data
	fareMock.EXPECT().IsDataValid().Return(true)
	fareMock.EXPECT().CalculateFare().Return(123, nil)

	a := &App{
		fare: fareMock,
	}
	a.Run(r, true)
}
