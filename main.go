package main

import (
	"fmt"
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

const (
	from = -1.0
	to   = 6
)

var n int = 6
var axX []float64
var axY []float64

func main() {
	NewTons := []NewTon{
		NewTon{6, []float64{3.1}, []float64{f(3.1)}},
		NewTon{6, []float64{3.1, 3.0}, []float64{f(3.1), f(3.0)}},
		NewTon{6, []float64{3.1, 3.0, 3.3}, []float64{f(3.1), f(3.0), f(3.3)}},
		NewTon{6, []float64{3.1, 3.0, 3.3, 2.8}, []float64{f(3.1), f(3.0), f(3.3), f(2.8)}},
	}
	for _, NewTon := range NewTons {
		var Koefs []float64
		for i := 0; i < len(NewTon.uzly); i++ {
			Koefs = append(Koefs, NewTon.UzelCalculate(i))
		}

		//for i := from; i <= to; i += 0.1 {

		axX = append(axX, math.Pi)
		axY = append(axY, NewTon.CalcInPoint(Koefs, math.Pi))
		fmt.Println(axY)
		//}
		//http.HandleFunc("/", drawChart)
		//http.ListenAndServe(":8000", nil)
	}
}

func (f *NewTon) CalcInPoint(koefs []float64, x float64) float64 {
	point := 0.0
	for i, val := range koefs {
		//if i == 3 {
		kef := val
		for j := 0; j < i; j++ {
			kef *= x - f.uzly[j]
		}
		point += kef
		//}
	}
	return point
}

func (f *NewTon) UzelCalculate(iter int) float64 {
	if iter == 0 {
		return f.funcs[0]
	}

	for i := 0; i < len(f.funcs)-1; i++ {
		f.funcs[i] = (f.funcs[i+1] - f.funcs[i]) / (f.uzly[i+iter] - f.uzly[i])
	}
	answer := f.funcs[0]
	f.funcs = f.funcs[:len(f.funcs)-1]

	return answer
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
					StrokeWidth: 5.0,
					//FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: axX,
				YValues: axY,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}
