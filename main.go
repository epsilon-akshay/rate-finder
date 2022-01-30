package main

import (
	"currency_converter/handler"
	"currency_converter/pkg/currency_rate"
	"currency_converter/repository"
	"currency_converter/usecase"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
)

var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const Size = 32
const RedisHost = "localhost:6379"
const FixerAccessKey = "d8c40d4421250bc8dbfde033a9272f40"
const FixerURL = "http://data.fixer.io/api/latest"
const Port = 8000

func main() {
	redisOpt := &redis.Options{
		Addr:     RedisHost,
		Password: "",
		DB:       0,
	}

	httpClient := http.DefaultClient

	currencyConverter := usecase.ConversionCalculator{
		RateFinder: currency_rate.FixerClient{AccessKey: FixerAccessKey, Url: FixerURL, HttpClient: httpClient},
	}

	keyGenClient := usecase.ProtectApiClient{
		KeyHolder: repository.GoRedis{
			Client: redis.NewClient(redisOpt),
		},
		KeyGenerator: usecase.RandomGen{LetterRunes: Letters, Size: Size},
	}

	r := handler.Router(currencyConverter, keyGenClient)

	fmt.Println("Server started listening to port 8000")
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), r)
	if err != nil {
		panic(err)
	}

}
