package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	task  = "Нужно преобразовать целое число в прописной вариант: 12 – двенадцать."
	howto = "http://localhost:8080/number/123 #должно вывести результат 'сто двадцать три'"
)

type NumberResponse struct {
	Success bool   `json:"success"`
	Number  string `json:"number"`
	Result  string `json:"result"`
}

type Instruction struct {
	Task  string `json:"task"`
	Howto string `json:"howto"`
}

func main() {
	defer CloseDB()

	router := gin.Default()
	router.GET("/", getInstructionPage)
	router.GET("/number/:num", getTextRepresentation)
	//router.Run("0.0.0.0:8080")
	router.Run("localhost:8080")
}

func getTextRepresentation(c *gin.Context) {
	success := true
	number := c.Param("num")
	lang := c.Request.URL.Query().Get("lang")
	if lang == "" {
		lang = "RU"
	}

	resText, err := GetStringRepresentation(number, lang)
	if err != nil {
		success = false
		resText = err.Error()
	}

	result := NumberResponse{
		Success: success,
		Number:  number,
		Result:  resText,
	}

	c.JSON(http.StatusOK, result)
}

func getInstructionPage(c *gin.Context) {
	c.JSON(http.StatusOK, Instruction{Task: task, Howto: howto})
}
