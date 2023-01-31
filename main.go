package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Any("/when/:year", func(ctx *gin.Context) {
		yearParam := ctx.Param("year")
		yearInt, err := strconv.Atoi(yearParam)
		yearTime := time.Date(yearInt, time.January, 1, 0, 0, 0, 0, time.UTC)
		sinceYear := yearTime.Sub(time.Now().UTC())
		if err != nil {
			log.Println(err)
		}	
		daysResponse := int(math.Abs(sinceYear.Hours() / 24))
		ctx.String(http.StatusOK, "%d", daysResponse)
	})

	router.Run()
}