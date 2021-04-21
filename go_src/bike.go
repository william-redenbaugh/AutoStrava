package main

type Tire struct {
	Brand string
	Model string

	distance_mi        float64
	lifespan_mi        float64
	power_transfer_kwh float64
}

type BottomBracket struct {
	Brand string
	Model string

	distance_mi         float64
	elevation_ft        float64
	lifespan_wear_ft_mi float64
	power_transfer_kwh  float64
}

type Crankset struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Cassette struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Brakes struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64

	is_rim       bool
	is_hydrualic bool
}

type Handlebar struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Stem struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Saddle struct {
	Brand     string
	Model     string
	age_years float32
}

type SeatPost struct {
	Brand              string
	Model              string
	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Frame struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Chain struct {
	Brand string
	Model string

	distance_mi           float64
	elevation_ft          float64
	livespan_wear_ft_mile float64
	power_transfer_kwh    float64
}

type Derailiure struct {
	Brand string
	Model string

	is_rear               bool
	distance_mi           float64
	elevation_ft          float64
	lifespan_wear_ft_mile float64
	power_transfer_kwh    float64
}

type WheelHub struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type WheelRim struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type WheelSpokesDescripter struct {
	Brand string
	Model string

	distance_mi        float64
	elevation_ft       float64
	power_transfer_kwh float64
}

type Wheel struct {
	rim    WheelRim
	spokes WheelSpokesDescripter
	hub    WheelHub
}

type Bike struct {
	Name  string
	frame Frame

	// Mechanicals here
	chain         Chain
	bottomBracket BottomBracket
	cassete       Cassette
	crankset      Crankset
	front_brakes  Brakes
	rear_brakes   Brakes

	// For Single Speed and RD only bike support
	has_front_derailiure bool
	frontDerailiure      *Derailiure
	has_rear_derailiure  bool
	rearDerailiure       *Derailiure

	// Wheels and tires here
	frontTire  Tire
	rearTire   Tire
	frontWheel Wheel
	rearWheel  Wheel

	// Finishing Kit here
	handleBar Handlebar
	stem      Stem
	seatPost  SeatPost
	saddle    Saddle
}
