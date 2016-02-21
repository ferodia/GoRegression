package main

//io and data processing imports
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// plotting imports
import (
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
	"image/color"
)
import (
	"github.com/ferodia/GoRegression/linear"
)

func main() {
	// load data from a csv file
	fmt.Println("Loading training set.. ")
	X, Y, err := ReadCSV("data/foodTrack.txt")
	if err != nil {
		log.Fatal(err)
	}

	// plot the data
	fmt.Println("Ploting training set.. ")
	plotData(X, Y, "plot/plot.png")
	//Linear regression
	fmt.Println("Running linear regression for one feature problem.. ")
	predictions := linear.Run(X, Y)

	// plot the data with prediction
	fmt.Println("Ploting training set with predictions.. ")
	plotPredictions(X, Y, predictions, "plot/plotpredictions.png")

}

func ReadCSV(fileName string) (x, y []float64, err error) {
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
		feature, err := strconv.ParseFloat(record[0], 64)
		label, err := strconv.ParseFloat(record[1], 64)
		x = append(x, feature)
		y = append(y, label)
	}
	return
}

func plotData(x, y []float64, plotFile string) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Profit of a food track depending on the population size"
	p.X.Label.Text = "population"
	p.Y.Label.Text = "profit"

	xys := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		xys[i].X, xys[i].Y = x[i], y[i]
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

}

func plotPredictions(x, y, pred []float64, plotFile string) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Profit of a food track depending on the population size"
	p.X.Label.Text = "population"
	p.Y.Label.Text = "profit"

	//scatter data points
	xys := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		xys[i].X, xys[i].Y = x[i], y[i]
	}

	// Create a scatter and set properties
	s, err := plotter.NewScatter(xys)
	s.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	// predictions data points
	preds := make(plotter.XYs, len(x))
	for i := 0; i < len(x); i++ {
		preds[i].X, preds[i].Y = x[i], pred[i]
	}
	//Create line for predictions
	d, err := plotter.NewLine(preds)
	d.Color = color.RGBA{B: 255, A: 255}
	p.Add(d)

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, plotFile); err != nil {
		panic(err)
	}

}
