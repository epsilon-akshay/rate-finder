package handler

import (
	"context"
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
			return
		}

		convAmount, err := converter.GetConvertedAmountFrom(r.Context(), base, amount, "")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			r := &Response{
				Success: false,
				Err:     err.Error(),
			}
			resp, _ := json.Marshal(r)
			w.Write(resp)
			return
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
	GetConvertedAmountFrom(ctx context.Context, base string, amount float64, key string) (float64, error)
}
