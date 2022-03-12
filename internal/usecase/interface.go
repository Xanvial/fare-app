package usecase

import "github.com/Xanvial/fare-app/internal/model"

//go:generate mockgen -source=interface.go -destination=mock/fare_mock.go
type FareUsecase interface {
	// Setter
	SetFareData(input model.FareData)
	ResetFareData()

	// Getter data
	GetLatestData() model.FareData
	GetDataCount() int
	GetMinDataCount() int

	// Check validity of data,
	// total recorded data is more than minimum, and last distance is  more than zero
	IsDataValid() bool

	// Validate fare data
	ValidateFareData(input model.FareData) error

	// Calculate fare data
	CalculateFare() (int, error)
}
