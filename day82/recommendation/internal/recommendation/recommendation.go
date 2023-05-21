package recommendation

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

type Recommendation struct {
	TripStart time.Time
	TripEnd   time.Time
	HotelName string
	Location  string
	TripPrice money.Money
}

type Option struct {
	HotelName     string
	Location      string
	PricePerNight money.Money
}

type AvailabilityGetter interface {
	GetAvailability(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string) ([]Option, error)
}

type Service struct {
	availability AvailabilityGetter
}

func NewService(availability AvailabilityGetter) (*Service, error) {
	if availability == nil {
		return nil, errors.New("Availability must not be nil")
	}

	return &Service{availability: availability}, nil
}

func (svc *Service) Get(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string, budget money.Money) (*Recommendation, error) {

	switch {
	case tripStart.IsZero():
		return nil, errors.New("trip start cannot be empty")
	case tripEnd.IsZero():
		return nil, errors.New("trip end cannot be empty")
	case location == "":
		return nil, errors.New("location cannot be empty")
	}

	opts, err := svc.availability.GetAvailability(ctx, tripStart, tripEnd, location)
	if err != nil {
		return nil, fmt.Errorf("error getting availability: %v", err)
	}

	tripDuration := math.Round(float64(tripEnd.Sub(tripStart).Hours()) / 24)
	lowerPrice := money.NewFormFloat(999999999, "USD")
	var cheapestTrip *Option

	for _, option := range opts {
		price := option.PricePerNight.Multiply(int64(tripDuration))
		if ok, _ := price.GreaterThan(budget); ok {
			continue
		}
		if ok, _ := price.LessThan(lowerPrice); ok {
			lowerPrice = price
			cheapestTrip = &option
		}
	}

	if cheapestTrip == nil {
		return nil, errors.New("No trips within budget")
	}
	return &Recommendation{
		TripStart: tripStart,
		TripEnd:   tripEnd,
		HotelName: cheapestTrip.HotelName,
		Location:  cheapestTrip.Location,
		TripPrice: *lowerPrice,
	}, nil
}
