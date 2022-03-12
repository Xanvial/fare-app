package presenter

import (
	"bufio"
	"errors"
	"io"
	"log"

	"github.com/Xanvial/fare-app/internal/model"
	"github.com/Xanvial/fare-app/internal/usecase"
	"go.uber.org/zap"
)

type App struct {
	fare usecase.FareUsecase
}

func New(cfg model.Config) *App {
	return &App{
		fare: usecase.NewFare(cfg),
	}
}

func (a *App) Run(r io.Reader, isConsole bool) {
	scanner := bufio.NewScanner(r)

	if isConsole {
		// on console print this to ask user to input data
		log.Print("Enter data:")
	}
	// forever loop this until empty line found
	for scanner.Scan() {
		inp := scanner.Text()

		// stop waiting for input if found empty line on console
		if isConsole && len(inp) == 0 {
			break
		}

		err := a.ProcessInput(inp)
		if err != nil {
			// just continue the loop, the error already logged inside the function
			// adding this in case there's additional process on input if no error found
			continue
		}
	}

	// process recorded data after new empty line encountered
	a.ProcessData()
}

func (a *App) ProcessInput(inp string) error {
	// parse input
	data, err := model.ParseInputText(inp)
	if err != nil {
		// log the error and continue the loop skipping this
		zap.L().Error("invalid input",
			zap.Error(err),
			zap.String("input", inp),
		)
		return err
	}

	// validate input with existing data
	err = a.fare.ValidateFareData(data)
	if err != nil {
		// log the error and continue the loop skipping this
		zap.L().Error("invalid fare data",
			zap.Error(err),
			zap.Any("latest_fare_data", a.fare.GetLatestData()),
			zap.Any("cur_fare_data", data),
		)
		return err
	}

	// update latest fare with current data
	a.fare.SetFareData(data)

	return nil
}

func (a *App) ProcessData() error {
	if !a.fare.IsDataValid() {
		// fare data is lower than minimum, log the error
		zap.L().Error("input data is less than minimum",
			zap.Any("latest_fare_data", a.fare.GetLatestData()),
			zap.Int("data_count:", a.fare.GetDataCount()),
			zap.Int("min_data:", a.fare.GetMinDataCount()),
		)

		// No need to show output based on requirement
		// log.Println("total cost: 0")
		return errors.New("input data is less than minimum")
	}

	// calculate fare
	cost, err := a.fare.CalculateFare()
	if err != nil {
		zap.L().Error("error calculating fare",
			zap.Any("latest_fare_data", a.fare.GetLatestData()),
			zap.Int("data_count:", a.fare.GetDataCount()),
		)
		return err
	}

	// output the result
	log.Println("Output:", cost)
	return nil
}
