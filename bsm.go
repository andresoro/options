// Implementation of the Black-Scholes-Merton options pricing algorithm
package bsm

import (
	"time"
)

type Option struct {
	//initial strike price
	I float64
	//strike price
	K float64
	//volatility
	V float64

	//expiration
	E time.Time
	// time till expiration
	T time.Time

	G Greek

	//false if put
	call bool
}

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
