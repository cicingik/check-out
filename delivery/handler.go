package delivery

import (
	"fmt"
	"net/http"

	"github.com/cicingik/check-out/config"
	"github.com/cicingik/check-out/pkg/httputils"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	httputils.JsonResponse(w, http.StatusTeapot, nil, struct {
		Message string `json:"message"`
	}{
		Message: http.StatusText(http.StatusTeapot),
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	httputils.JsonResponse(w, http.StatusOK, nil, struct {
		Version    string `json:"version"`
		CommitHash string `json:"commit_hash"`
		CommitMsg  string `json:"commit_msg"`
	}{
		Version:    config.AppVersion,
		CommitHash: config.CommitHash,
		CommitMsg:  config.CommitMsg,
	})
}

func HealthZX(w http.ResponseWriter, r *http.Request) {
	httputils.JsonResponse(w, http.StatusOK, nil, struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf(`I am Healthy`),
	})
}
