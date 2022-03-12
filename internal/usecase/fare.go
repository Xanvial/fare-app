package usecase

import (
	"errors"
	"time"

	"github.com/Xanvial/fare-app/internal/model"
)

type fare struct {
	latestFare model.FareData
	dataCount  int

	cfg model.FareConfig
}

func NewFare(cfg model.Config) FareUsecase {
	return &fare{
		cfg: cfg.Fare,
	}
}

// Setter
func (f *fare) SetFareData(input model.FareData) {
	f.latestFare = input
	f.dataCount++
}

func (f *fare) ResetFareData() {
	f.latestFare = model.FareData{}
	f.dataCount = 0
}

// Getter, latest data
func (f *fare) GetLatestData() model.FareData {
	return f.latestFare
}

// Getter, data count
func (f *fare) GetDataCount() int {
	return f.dataCount
}

// Getter, min data
func (f *fare) GetMinDataCount() int {
	return f.cfg.MinData
}

func (f *fare) IsDataValid() bool {
	return f.dataCount >= f.cfg.MinData && f.latestFare.Distance > 0
}

func (f *fare) ValidateFareData(input model.FareData) error {
	// check if input is older than latest data
	if input.ElapsedTime < f.latestFare.ElapsedTime {
		return errors.New("fare time data is older than previous data")
	}

	// check if input is more than 5 minutes from latest data
	if input.ElapsedTime-f.latestFare.ElapsedTime > time.Duration(f.cfg.MaxIntervalSec)*time.Second {
		return errors.New("fare time interval is more than 5 minutes")
	}

	// check if input has lower distance than latest data
	if input.Distance < f.latestFare.Distance {
		return errors.New("fare distance is lower than previous data")
	}

	return nil
}

func (f *fare) CalculateFare() (int, error) {
	dist := f.latestFare.Distance

	// sanity check, if distance is zero or negative, directly return error
	if dist <= 0 {
		return 0, errors.New("distance is zero or less")
	}

	// base cost is 400
	cost := f.cfg.BaseFare
	// reduce to distance of base cost
	dist -= f.cfg.BaseFareDistance

	// loop until reaching total 10 km (which means 9km after reducing base distance)
	maxReduce := f.cfg.LimitDistance - f.cfg.BaseFareDistance
	for reduce := float64(0); dist > 0 && reduce < maxReduce; reduce += f.cfg.UnderLimitFareDistance {
		// note: this is possible to break past the limit distance without using overlimit fare price
		// considering there's no requirement for this

		dist -= f.cfg.UnderLimitFareDistance
		cost += f.cfg.UnderLimitFare
	}

	// loop until all distance calculated
	for dist > 0 {
		dist -= f.cfg.OverLimitFareDistance
		cost += f.cfg.OverLimitFare
	}

	return cost, nil
}
