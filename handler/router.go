package handler

import (
	"github.com/gorilla/mux"
)

func Router(converter CurrencyConverter, keyGen KeyGenerator) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/currency_converter/currency/convert", ConvertCurrency(converter)).Methods("GET")
	router.HandleFunc("/v1/currency_converter/api_key", GenerateAPIKey(keyGen)).Methods("GET")
	return router
}
