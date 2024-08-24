package routes

import (
	"net/http"
	"time"
)

//type Handler struct {
//	Service facade.ServiceI
//}

func InitRoutes() http.Handler {

	mux := http.NewServeMux()
	//routes files
	mux.Handle("/stream", dataStream)

	//later on middleware for logging purposes
	return mux
}

func dataStream(w http.ResponseWriter, r *http.Request) {
	sw := time.Now()

	apiResponse := h.Ha
}
