package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/imshakthi/lets-go/covid19/models"
	"github.com/imshakthi/lets-go/covid19/service"
)

func Handler() (models.Result, error) {
	fmt.Println("Check for available slots started...")
	results := service.GetResults()
	fmt.Println("Check for available slots ended :)")
	return results, nil
}

func main() {
	lambda.Start(Handler)
}
