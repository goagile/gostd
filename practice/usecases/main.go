package main

import (
	"fmt"
)

func main() {
	request := map[string]interface{}{
		"action": "rateColor",
		"colorName": "red",
		"rating": 5,
	}
	controller := new(UseCaseController)
	controller.handleRequest(request)
}

type UseCaseController struct {}

func (uc UseCaseController) handleRequest(request map[string]interface{}) {
	action := request["action"]
	switch action {
		
	}
}

func (uc UseCaseController) rateColor() {
	fmt.Println("RATE")
}
