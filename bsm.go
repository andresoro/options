// Package bsm implements the Black-Scholes-Merton options pricing algorithm
package bsm

import (
	"math"
	"time"

	"github.com/chobie/go-gaussian"
)

// risk free rate
var rate float64

// Option contract
type Option struct {
	//current price
	I float64
	//strike price
	S float64
	// implited volatility
	V float64

	//expiration
	E time.Time
	// time till expiration in years
	T float64

	G Greek

	//value of the option
	val float64

	//false if put
	call bool
}

// New returns a new option
func New(i, s, v float64, call bool, e time.Time) *Option {
	opt := &Option{
		I:    i,
		S:    s,
		V:    v,
		E:    e,
		call: call,
	}
	return opt
}

func (o *Option) calculate() {
	//normal distribution
	gauss := gaussian.NewGaussian(0, 1)
	//time to expiry in years
	o.T = (o.E.Sub(time.Now()).Hours() / 24) / 365

	//calculate d1
	d1 := (math.Log(o.I/o.S) + (rate+math.Pow(o.V, 2)/2)*o.T) / (o.V * math.Sqrt(o.T))

	//calculate d2
	d2 := d1 - (o.V * math.Sqrt(o.T))

	if o.call {
		o.val = (o.I * gauss.Cdf(d1)) - (o.S * gauss.Cdf(d2))
	} else {
		o.val = (o.S * gauss.Cdf(-d2)) - (o.I * gauss.Cdf(-d1))
	}

}

// Value returns option value at current time
func (o *Option) Value() float64 {
	o.calculate()
	return o.val
}

// Greek holds the partial derivative values of an option
type Greek struct {
	Delta float64
	Theta float64
	Gamma float64
	Vega  float64
}

func (g *Greek) set(d float64, t float64, gm float64, v float64) {
	g.Delta = d
	g.Theta = t
	g.Gamma = gm
	g.Vega = v
}
