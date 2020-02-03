package main

type Unit struct {
	Key        string
	isMain     bool
	toMain     float64
	unitType   int
	unitSystem int
}

func (u *Unit) ToMain(v float64) float64 {
	return v * u.toMain
}

func (u *Unit) FromMain(v float64) float64 {
	return v / u.toMain
}
