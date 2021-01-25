package controllers

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

// CovarCorrelMatrix - produces a covariance and correlation matrix
// func CovarCorrelMatrix(slice []float64) (covar *mat.SymDense, covarmap map[string][]float64, correl *mat.SymDense, correlmap map[string][]float64) {
// 	cols := len(SeriesSlice)
// 	rows := len(slice) / cols
// 	srcmatrix := mat.NewDense(rows, cols, slice)
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		covar = mat.NewSymDense(len(SeriesSlice), nil)
// 		stat.CovarianceMatrix(covar, srcmatrix, nil)
// 		// covar = mat.Formatted(covartemp, mat.Prefix(""), mat.Squeeze())
// 		covarmap = make(map[string][]float64, cols)
// 		for i := range SeriesSlice {
// 			covarmap[SeriesSlice[i]] = make([]float64, cols)
// 			for ii := range SeriesSlice {
// 				covarmap[SeriesSlice[i]][ii] = covar.At(i, ii)
// 			}
// 		}
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		correltemp := mat.NewSymDense(len(SeriesSlice), nil)
// 		stat.CorrelationMatrix(correltemp, srcmatrix, nil)
// 		// covar = mat.Formatted(covartemp, mat.Prefix(""), mat.Squeeze())
// 		correlmap = make(map[string][]float64, cols)
// 		for i := range SeriesSlice {
// 			correlmap[SeriesSlice[i]] = make([]float64, cols)
// 			for ii := range SeriesSlice {
// 				correlmap[SeriesSlice[i]][ii] = correltemp.At(i, ii)
// 			}
// 		}
// 	}()
// 	wg.Wait()
// 	return covar, covarmap, correl, correlmap
// }

func valueatrisk(signif int, mean []float64, stdev []float64) (valueatrisk []float64) {
	vatr := 0.0
	c := 0.0
	switch signif {
	case 5:
		c = 1.65
	case 1:
		c = 2.33
	}
	valueatrisk = make([]float64, 0)
	for h := 0; h <= 7; h++ {
		vatr = (mean[h] - stdev[h]*c) * 100
		valueatrisk = append(valueatrisk, vatr)
	}
	return valueatrisk
}

func roundslice(precision float64, scale float64, source []float64) []float64 {
	precision = math.Pow(10, precision)
	scale = math.Pow(10, scale)
	target := make([]float64, 0)
	for _, v := range source {
		v = v * precision * scale
		v = math.Round(v)
		v = v / precision
		target = append(target, v)
	}
	return target
}
