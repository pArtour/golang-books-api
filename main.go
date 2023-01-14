package main

import (
	"golang-books-app/controller"
	"golang-books-app/model"
	"log"
)

func main() {
	repository, err := model.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	controller.Init(repository)
}
