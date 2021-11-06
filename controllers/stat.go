package controllers

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

// BetaSample - give mu and stdev, alpha and beta are calculated, and then used to randomly sample the beta distribution
func BetaSample(mu, stdev float64) float64 {
	if stdev == 0 || mu == 0 {
		return mu
	}
	v := math.Pow(stdev, 2)
	alpha := ((1-mu)/v - 1/mu) * math.Pow(mu, 2)
	beta := ((1-mu)/v - 1/mu) * mu * (1 - mu)
	x := distuv.Beta{
		Alpha: alpha,
		Beta:  beta,
	}
	return x.Rand()
}

func AlphaCheck(mu, sigma float64) (alphaZero float64) {
	var alpha = ((1-mu)/math.Pow(sigma, 2) - 1/mu) * math.Pow(mu, 2)
	alphaZero = sigma
	if alpha <= 0 {
		mtwo := math.Pow(mu, 2)
		mthree := math.Pow(mu, 3)
		x := (mtwo - mthree)
		alphaZero = math.Pow(x/mu, .5) - .00001
		return alphaZero
	}
	return alphaZero
}

// BinomialSample - Skew must be between 0 and 1. a value of .5 indicates no Skew, <.5 is positive Skew and >.5 is negative Skew.
func BinomialSample(expval float64, Skew float64) float64 {
	if Skew == 0 {
		return expval
	}
	n := expval / Skew
	x := distuv.Binomial{
		N: n,
		P: Skew,
	}
	return math.Round(x.Rand())
}

// NormalSample -
func NormalSample(mean float64, stdev float64, min float64) (val float64) {
	if stdev == 0 {
		return mean
	}
	x := distuv.Normal{
		Mu:    mean,
		Sigma: stdev,
	}
	val = x.Rand()
	if val < min {
		val = min
	}
	return val
}

// LogNormalSample -
func LogNormalSample(mean float64, stdev float64) float64 {
	if stdev == 0 {
		return mean
	}
	x := distuv.LogNormal{
		Mu:    mean,
		Sigma: stdev,
	}
	return x.Rand()
}

// TriangularSample -
func TriangularSample(mean float64) float64 {
	x := distuv.Triangle{}
	return x.Rand()
}
