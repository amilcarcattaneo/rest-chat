package check

import (
	"encoding/json"
	"net/http"

	"rest-chat/src/api/controllers"
	checkService "rest-chat/src/api/services/check"
)

type checkResult struct {
	Health string `json:"health"`
}

func Check(w http.ResponseWriter, r *http.Request) {
	ok, err := checkService.Check()
	if err != nil {
		controllers.HandleError(w, http.StatusInternalServerError, err)
		return
	}

	if ok {
		healthCheck := checkResult{
			Health: "ok",
		}
		if err = json.NewEncoder(w).Encode(healthCheck); err != nil {
			controllers.HandleError(w, http.StatusInternalServerError, err)
			return
		}
	}
}
