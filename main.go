package main

import (
	"github.com/wcharczuk/go-chart"
	"math"
	"net/http"
)

func f(x float64) float64 {
	return math.Sin(x)
}

type NewTon struct {
	n     int
	uzly  []float64
	funcs []float64
}

func NewNewTon(n int, uzly []float64, funcs []float64) *NewTon {
	return &NewTon{n: n, uzly: uzly, funcs: funcs}
}

var n int = 6
var axX []float64
var axY []float64

func main() {
	NewTon := NewNewTon(6, []float64{0, 1, 2, 3, 4, 5}, []float64{f(0), f(1), f(2), f(3), f(4), f(5)})
	var Koefs []float64
	for i := 0; i < n; i++ {
		Koefs = append(Koefs, NewTon.uzelCalculate(i))
	}

	for i := 0.0; i < 5; i += 0.1 {
		axX = append(axX, i)
		axY = append(axY, NewTon.CalcInPoint(Koefs, i))
	}

	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}

func (f *NewTon) CalcInPoint(koefs []float64, x float64) float64 {
	point := 0.0
	for i, val := range koefs {
		kef := val
		for j := 0; j < i; j++ {
			kef *= x - f.uzly[j]
		}
		point += kef
	}
	return point
}

func (f *NewTon) uzelCalculate(iter int) float64 {
	if iter == 0 {
		return f.funcs[0]
	}

	for i := 0; i < len(f.funcs)-1; i++ {
		f.funcs[i] = (f.funcs[i+1] - f.funcs[i]) / (f.uzly[i+iter] - f.uzly[i])
	}

	f.funcs = f.funcs[:len(f.funcs)-1]

	return f.funcs[0]
}

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	*/

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: axX,
				YValues: axY,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}
