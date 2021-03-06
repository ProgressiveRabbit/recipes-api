package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"io/ioutil"
	"net/http"
	"time"
)


var recipes []Recipe //list of recipes

func init()  {
	recipes = make([]Recipe,0)
	file,_ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file),&recipes)  //convert json file to recipes array
}

type Recipe struct {
	ID string `json:"id"` //unique identifier to differentiate each recipe in tge database
	Name string `json:"name"`//recipe name
	Tag []string `json:"tags"` // recipe category
	Ingredients []string `json:"ingredients"` //components of recipe
	Instruction []string `json:"instruction"` //step by  step
	PublishedAt time.Time `json:"published_at"` //publication date
}


//POST
func NewRecipeHandler(ctx *gin.Context)  {
	var recipe Recipe
	if err := ctx.ShouldBindJSON(&recipe);err != nil { //parse the request body into Recipe struct
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String() //generate a unique identifier
	recipe.PublishedAt = time.Now()
	recipes = append(recipes,recipe)
	ctx.JSON(http.StatusOK,recipe) //set recipe type into json

}


//GET
func ListRecipeHandler(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,recipes)
}

func UpdateRecipeHandler(ctx *gin.Context)  {
	id := ctx.Param("id")
	var recipe Recipe
	if err := ctx.ShouldBindJSON(&recipe);err!=nil {
		ctx.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}


	index := -1 //the position to recipes
	for i := 0;i < len(recipes);i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error" : "recipe not found",
		})
		return
	}

	recipes[index] = recipe
	ctx.JSON(http.StatusOK,recipe)
}

func main()  {
	router := gin.Default()
	router.POST("/recipes",NewRecipeHandler)
	router.GET("/recipes",ListRecipeHandler)
	router.PUT("/recipes/:id",UpdateRecipeHandler)
	router.Run()
}

