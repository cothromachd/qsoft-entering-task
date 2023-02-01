package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {

		headersMap, breakFlag := ctx.Request.Header, false

		for key, value := range headersMap {
			if strings.ToLower(key) == "x-ping" {
				for _, valueVal := range value {
					if valueVal == "ping" {
						ctx.Header("X-PONG", "pong")
						breakFlag = true
						break
					}

				}
			}

			if breakFlag {
				break
			}
		}
	})

	router.Any("/when/:year", func(ctx *gin.Context) {

		// Получем year из запроса.
		yearParam := ctx.Param("year")

		// Кастуем из строки в целое число, чтобы
		// вставить в параметры time.Date()
		yearTypeInt, err := strconv.Atoi(yearParam)
		if err != nil {
			log.Println(err)
		}

		// Образуем дату time.Time с заданнным year.
		yearTypeTime := time.Date(yearTypeInt, time.January, 1, 0, 0, 0, 0, time.UTC)

		// Сохраняем только актуальный год, месяц и день, чтобы часы, минуты и.т.д
		// не мешали исчислению прошедших дней.
		Now := time.Now().UTC()
		curTimestamp := time.Date(Now.Year(), Now.Month(), Now.Day(), 0, 0, 0, 0, time.UTC)

		// Находим разницу между заданной датой и настоящей датой.
		sinceYear := yearTypeTime.Sub(curTimestamp)

		// Находим количество прошедших/предстоящих дней от заданной даты.
		// Для этого переведем часы в дни поделив часы на 24, так как
		// в днях (имеется ввиду в сутках) 24 часа.
		// В случае, если значение отрицательное - берём модуль числа.
		daysResponse := int(math.Abs(sinceYear.Hours() / 24))

		// Отправляем ответ.
		if yearTypeInt > curTimestamp.Year() {
			ctx.String(http.StatusOK, "Days left: %d", daysResponse)
		} else if yearTypeInt < curTimestamp.Year() || yearTypeInt == curTimestamp.Year() {
			ctx.String(http.StatusOK, "Days gone: %d", daysResponse)
		}
	})

	// Запускаем сервер на порту по умолчанию.
	router.Run()
}
