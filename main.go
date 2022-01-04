package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type operation struct {
	Operands []int `json:"operands"`
}

func add(c *gin.Context) {
	var op operation

	if err := c.BindJSON(&op); err != nil {
		return
	}

	result := 0

	for i := 0; i < len(op.Operands); i++ {
		result += op.Operands[i]
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func main() {
	router := gin.Default()
	router.POST("/add", add)

	router.Run()
}
