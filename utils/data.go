package utils

import (
	"log"
	"os"
	"encoding/csv"
	"io"
	"strconv"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
"github.com/gonum/plot/vg/draw"
	"image/color"
	"github.com/gonum/plot/vg"
	"fmt"
)

func ReadCSV(fileName string) (x [][]float64, y []float64, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		//read values
		var feature []float64
		for i:=0; i<len(record)-1;i++ {
			v, _ := strconv.ParseFloat(record[i], 64)
			feature = append(feature, v)

		}
		x = append(x,feature)
		label, err := strconv.ParseFloat(record[len(record)-1], 64)
		y = append(y, label)
	}
	return
}



func Plot(x [][]float64, y []float64, plotFile string) (err error) {
	features := len(x[0])
	switch features {
	case 1:
		err = plot2D(x, y, plotFile )
	case 2:
		err = plot3D(x, y, plotFile )
	default:
		err = fmt.Errorf("Too many dimensions to plot")
	}
	return err
}

func plot2D(x [][]float64, y []float64, plotFile string) (err error) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	xys := make(plotter.XYs, len(y))
	for i := 0; i < len(y); i++ {
		xys[i].X, xys[i].Y = x[i][0], y[i]
	}

	// Create a scatter and set properties
	s, err := plotter.NewScatter(xys)
	s.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)
	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, plotFile); err != nil {
		panic(err)
	}
	return
}

func plot3D(x [][]float64, y []float64, plotFile string) (err error) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	//p.Title.Text = "Profit of a food track depending on the population size"
	//p.X.Label.Text = "population"
	//p.Y.Label.Text = "profit"

	xys0 := make(plotter.XYs, len(y))
	xys1 := make(plotter.XYs, len(y))
	for i := 0; i < len(x); i++ {
		if y[i]==1 {
			xys1[i].X, xys1[i].Y = x[i][0], x[i][1]
		} else {
			xys0[i].X, xys0[i].Y = x[i][0], x[i][1]
		}

	}

	// Create a scatter and set properties for y=0
	s0, err := plotter.NewScatter(xys0)
	s0.Shape = draw.CircleGlyph{}
	s0.Color = color.RGBA{R: 255, A: 255}
	p.Add(s0)

	// Create a scatter and set properties for y=1
	s1, err := plotter.NewScatter(xys1)
	s1.Shape = draw.CrossGlyph{}
	s1.Color = color.RGBA{B: 25, A: 255}
	p.Add(s1)
	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, plotFile); err != nil {
		panic(err)
	}
	return
}

func PlotWithPredictions(x [][]float64, y, pred []float64, plotFile string) (err error) {
	features := len(x[0])
	switch features {
	case 1:
		err = plot2DWithPredictions(x, y,pred,  plotFile )
	case 2:
		err = plot3DWithPredictions(x, y,pred,  plotFile )
	default:
		err = fmt.Errorf("Too many dimensions to plot")
	}
	return err
}

func plot2DWithPredictions(x [][]float64, y, pred []float64, plotFile string) (err error){
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	//scatter data points
	xys := make(plotter.XYs, len(y))
	for i := 0; i < len(y); i++ {
		xys[i].X, xys[i].Y = x[i][0], y[i]
	}

	// Create a scatter and set properties
	s, err := plotter.NewScatter(xys)
	s.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	// predictions data points
	preds := make(plotter.XYs, len(y))
	for i := 0; i < len(y); i++ {
		preds[i].X, preds[i].Y = x[i][0], pred[i]
	}
	//Create line for predictions
	d, err := plotter.NewLine(preds)
	d.Color = color.RGBA{B: 255, A: 255}
	p.Add(d)

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, plotFile); err != nil {
		panic(err)
	}
	return
}


func plot3DWithPredictions(x [][]float64, y, pred []float64, plotFile string) (err error) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	//p.Title.Text = "Profit of a food track depending on the population size"
	//p.X.Label.Text = "population"
	//p.Y.Label.Text = "profit"

	xys0 := make(plotter.XYs, len(y))
	xys1 := make(plotter.XYs, len(y))
	for i := 0; i < len(x); i++ {
		if y[i]==1 {
			xys1[i].X, xys1[i].Y = x[i][0], x[i][1]
		} else {
			xys0[i].X, xys0[i].Y = x[i][0], x[i][1]
		}

	}

	// Create a scatter and set properties for y=0
	s0, err := plotter.NewScatter(xys0)
	s0.Shape = draw.CircleGlyph{}
	s0.Color = color.RGBA{R: 255, A: 255}
	p.Add(s0)

	// Create a scatter and set properties for y=1
	s1, err := plotter.NewScatter(xys1)
	s1.Shape = draw.CrossGlyph{}
	s1.Color = color.RGBA{B: 25, A: 255}
	p.Add(s1)

	// predictions data points

	preds := make(plotter.XYs, len(y))
	for i := 0; i < len(y); i++ {
		preds[i].X, preds[i].Y = x[i][0], pred[i]
	}
	//Create line for predictions
	d, err := plotter.NewLine(preds)
	d.Color = color.RGBA{B: 255, A: 255}
	p.Add(d)
	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, plotFile); err != nil {
		panic(err)
	}
	return
}




