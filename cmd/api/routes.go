package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// application هي هيكل يحتوي على معالجات التطبيق.
type application struct{}

// rootHandler يعالج الطلبات على المسار الجذري (/).
func (app *application) rootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the API!"))
}

// healthcheckHandler يعالج الطلبات على مسار الصحة (/v1/healthcheck).
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","version":"0.0.1"}`))
}

// routes يقوم بتحديد المسارات ومعالجاتها.
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// تعريف المسارات المتاحة
	router.GET("/", app.rootHandler)          // المسار الجذري
	router.GET("/v1/healthcheck", app.healthcheckHandler) // مسار الصحة

	return router
}
