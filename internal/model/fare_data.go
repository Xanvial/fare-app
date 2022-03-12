package model

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

type FareData struct {
	ElapsedTime time.Duration // use time.Duration, because default time.Time can't handle 99:99:99.999 format
	Distance    float64
}

func (fd FareData) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("elapsed_time", fd.ElapsedTime.String())
	enc.AddString("distance", strconv.FormatFloat(fd.Distance, 'f', -1, 64))
	return nil
}

func ParseInputText(input string) (FareData, error) {
	if input == "" {
		return FareData{}, errors.New("empty input")
	}

	// split input by space
	data := strings.Split(input, " ")
	if len(data) != 2 {
		return FareData{}, errors.New("invalid input count")
	}

	// Parse time portion of the data
	timeData := strings.Split(data[0], ":")
	if len(timeData) != 3 {
		return FareData{}, errors.New("invalid time format")
	}

	var elapsedTime time.Duration
	hour, err := strconv.Atoi(timeData[0])
	if err != nil {
		return FareData{}, errors.New("invalid hour format")
	}

	minute, err := strconv.Atoi(timeData[1])
	if err != nil {
		return FareData{}, errors.New("invalid minute format")
	}

	second, err := strconv.ParseFloat(timeData[2], 64) // Can be changed to bitSize 32 if additional precision not needed
	if err != nil {
		return FareData{}, errors.New("invalid second format")
	}

	elapsedTime = time.Duration(hour)*time.Hour +
		time.Duration(minute)*time.Minute +
		time.Duration(second*1000)*time.Millisecond

	// Parse distance portion of the data
	dist, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return FareData{}, errors.New("invalid distance format")
	}

	return FareData{
		ElapsedTime: elapsedTime,
		Distance:    dist,
	}, nil
}
