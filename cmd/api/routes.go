package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// دالة معالج للمسار الجذري
func (app *application) rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the API!"))
}

// دالة لتكوين المسارات
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// تعريف المسارات المتاحة
	router.HandlerFunc(http.MethodGet, "/", app.rootHandler) // إضافة المسار الجذري
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	return router
}
