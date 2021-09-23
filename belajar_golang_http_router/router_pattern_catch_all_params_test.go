package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/julienschmidt/httprouter"
)

func TestRouterPatternCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprintf(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}
