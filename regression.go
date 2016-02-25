package main


import (
	"github.com/ferodia/GoRegression/regression"
	"github.com/ferodia/GoRegression/utils"
	"fmt"
	"log"
)

func main() {

	// load data from a csv file
	fmt.Println("Loading training set.. ")
	X, Y, err := utils.ReadCSV("data/foodTrack.txt")
	if err != nil {
		log.Fatal(err)
	}

	// plot the data
	fmt.Println("Ploting training set.. ")
	utils.Plot(X, Y, "plot/plot.png")

	//Linear regression

	fmt.Println("Running linear regression for one feature problem.. ")
	model := new(regression.Linear)
	model.Initialise(X,Y)
	model.Train(1500,0.01)

	fmt.Println("Predicting.. ")
	predictions := model.Predict(X)

	// plot the data with prediction
	fmt.Println("Ploting training set with predictions.. ")
	utils.PlotWithPredictions(X, Y, predictions, "plot/plotpredictions.png")

}

