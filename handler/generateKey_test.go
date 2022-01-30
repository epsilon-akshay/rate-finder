package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateAPIKey(t *testing.T) {
	t.Run("should return 5xx response when failed to generate key", func(t *testing.T) {
		m := mockKeyGen(func(ctx context.Context) (string, error) {
			return "", errors.New("failed to generate key")
		})

		ts := httptest.NewServer(GenerateAPIKey(m))
		defer ts.Close()

		url := fmt.Sprintf("%s", ts.URL)
		res, err := http.Get(url)
		require.NoError(t, err)

		expectedResponse, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		require.NoError(t, err)

		res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusInternalServerError)
		assert.Equal(t, string(expectedResponse), `{"success":false,"err":"failed to generate key"}`)
	})

	t.Run("should return 2xx response when succeed to generate key", func(t *testing.T) {
		m := mockKeyGen(func(ctx context.Context) (string, error) {
			return "random", nil
		})

		ts := httptest.NewServer(GenerateAPIKey(m))
		defer ts.Close()

		url := fmt.Sprintf("%s", ts.URL)
		res, err := http.Get(url)
		require.NoError(t, err)

		expectedResponse, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		require.NoError(t, err)

		res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusOK)
		assert.Equal(t, string(expectedResponse), `{"success":true,"api_key":"random"}`)
	})
}
