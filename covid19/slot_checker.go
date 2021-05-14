package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/imshakthi/lets-go/covid19/constants"
	"github.com/imshakthi/lets-go/covid19/models"
	"os"
	"time"
)

func CheckSlots() {
	slotResp, err := getSlotResponse()
	if err != nil {
		fmt.Printf("error in getting response: %+v", err)
		return
	}
	slotResp.GetSlotAvailHospital()
}

func getSlotResponse() (models.CheckSlotResponse, error) {
	client := resty.New()

	url := getURL()
	headers := getHeaders()
	fmt.Println("Hitting URL::", url)

	response, err := client.R().
		SetHeaders(headers).
		SetHeader("Accept", "application/json").
		Get(url)
	if err != nil {
		fmt.Errorf("error in getting response: %+v", err)
		return models.CheckSlotResponse{}, err
	}
	fmt.Println("got slot availability response")

	var slotResp models.CheckSlotResponse
	err = json.Unmarshal(response.Body(), &slotResp)
	if err != nil {
		fmt.Errorf("error in unmarshalling response: %+v", err)
		return models.CheckSlotResponse{}, err
	}
	fmt.Printf("response is marshalled")
	fmt.Println()
	return slotResp, nil
}

func getURL() string {
	baseURL := getEnv(constants.EnvBaseUrl)
	now := time.Now()
	date := now.Format("2-01-2006")
	fmt.Println("checking for the date::", date)

	return fmt.Sprintf("%s%s", baseURL, date)
}

func getHeaders() map[string]string {
	bearerToken := getEnv(constants.EnvBearerToken)
	originReferer := getEnv(constants.EnvOriginReferer)

	return map[string]string{
		"Authorization":   bearerToken,
		"Origin":          originReferer,
		"Referer":         originReferer,
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "en-US,en;q=0.5' --compressed",
		"Connection":      "keep-alive",
		"TE":              "Trailers",
	}
}

func getEnv(variable string) string {
	value, ok := os.LookupEnv(variable)
	if !ok {
		panic("PLACE SET VALUE for :: " + variable)
	}
	return value
}
