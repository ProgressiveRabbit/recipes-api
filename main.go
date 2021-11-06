package main

import "github.com/gin-gonic/gin"


func main()  {
	router := gin.Default()
	router.Run()
}

type Recipe struct {
	Name string `json:"Name"`
	Tag []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instruction []string `json:"instruction"`
	PublishedAt []string `json:"published_at"`
}