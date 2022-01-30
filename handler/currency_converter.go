package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Success bool    `json:"success"`
	Err     string  `json:"err,omitempty"`
	Val     float64 `json:"converted_amount,omitempty"`
}

func ConvertCurrency(converter CurrencyConverter) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := r.URL.Query()
		base := params.Get("base")
		amountStr := params.Get("amount")

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			r := &Response{
				Success: false,
				Err:     err.Error(),
			}
			resp, _ := json.Marshal(r)
			w.Write(resp)
		}

		convAmount, err := converter.GetConvertedAmountFrom(base, amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			r := &Response{
				Success: false,
				Err:     err.Error(),
			}
			resp, _ := json.Marshal(r)
			w.Write(resp)
		}

		w.WriteHeader(http.StatusOK)
		res := &Response{
			Success: true,
			Val:     convAmount,
		}
		resp, _ := json.Marshal(res)
		w.Write(resp)
	})
}

type CurrencyConverter interface {
	GetConvertedAmountFrom(base string, amount float64) (float64, error)
}
