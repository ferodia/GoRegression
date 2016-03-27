package regression


import (
	"github.com/ferodia/GoRegression/utils"
	"github.com/ferodia/GoRegression/model"
	"fmt"
	"log"
)

func main() {

	//Linear regression

	// load data from a csv file
	fmt.Println("Loading training set.. ")
	X_linear, Y_linear, err := utils.ReadCSV("data/foodTrack.txt")
	if err != nil {
		log.Fatal(err)
	}

	// plot the data
	fmt.Println("Ploting training set.. ")
	utils.Plot(X_linear, Y_linear, "plot/plot.png")
	fmt.Println("Running linear regression for one feature problem.. ")
	ln := model.NewModel(model.Linear)
	ln.Initialise(X_linear,Y_linear)
	ln.Train(1500,0.01)

	fmt.Println("Predicting.. ")
	predictions := ln.Predict(X_linear)
	fmt.Println("values", Y_linear)
	fmt.Println("predictions", predictions)
	// plot the data with prediction
	fmt.Println("Ploting training set with predictions.. ")
	utils.PlotWithPredictions(X_linear, Y_linear, predictions, "plot/plotpredictions.png")


	//Logistic regression

	fmt.Println("Loading training set.. ")
	X_logistic, Y_logistic, err := utils.ReadCSV("data/students.txt")
	if err != nil {
		log.Fatal(err)
	}

	// plot the data
	fmt.Println("Ploting training set.. ")
	utils.Plot(X_logistic, Y_logistic, "plot/plot_exams.png")
	lg := model.NewModel(model.Logistic)
	lg.Initialise(X_logistic,Y_logistic)
	lg.Train(2000,0.0001)
	fmt.Println("Predicting.. ")
	logistic_preds := lg.Predict(X_logistic)
	fmt.Println("accuracy", utils.Accuracy(Y_logistic,logistic_preds))



}

