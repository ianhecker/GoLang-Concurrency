package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type ScatterPlot struct {
	points plotter.XYs
}

func MakeScatterPlot(count int) ScatterPlot {
	points := make(plotter.XYs, count)
	return ScatterPlot{points: points}
}

func (p *ScatterPlot) AddPoint(x int, y int64, index int) {
	p.points[index].X = float64(x)
	p.points[index].Y = float64(y)
}

func (p *ScatterPlot) Plot(plotName string) {
	plot, err := plot.New()
	if err != nil {
		panic(err)
	}

	plot.Title.Text = "Routine Count vs Time in ns"
	plot.X.Label.Text = "Routine Count"
	plot.Y.Label.Text = "Time in nanoseconds"

	scatter, err := plotter.NewScatter(p.points)
	if err != nil {
		panic(err)
	}

	plot.Add(scatter)
	plot.Legend.Add("Routine vs Time", scatter)

	if err := plot.Save(30*vg.Inch, 5*vg.Inch, plotName+".png"); err != nil {
		panic(err)
	}
}
