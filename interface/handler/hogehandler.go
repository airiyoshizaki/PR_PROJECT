package handler

import (
	"fmt"
	"net/http"
)

// @Summary      Show an account
// @Description  get string by ID
// @Tags         home
// @Accept       json
// @Produce      json
// @Success 200 "Success"
// @Error 400 {object} apiError.ApiError
// @Error 401 {object} apiError.ApiError
// @Error 403 {object} apiError.ApiError
// @Error 500 {object} apiError.ApiError
// @Router       / [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         about
// @Accept       json
// @Produce      json
// @Success 200 "Success"
// @Error 400 {object} apiError.ApiError
// @Error 401 {object} apiError.ApiError
// @Error 403 {object} apiError.ApiError
// @Error 500 {object} apiError.ApiError
// @Router       /about [get]
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the About Page.")
}
