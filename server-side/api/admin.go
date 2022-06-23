package api

import (
	"encoding/json"
	"final-project-engineering-75/server-side/repository"
	"net/http"
)

type AdminErrorResponse struct {
    Error string `json:"error"`
}

type AdminResponse struct {
    Creations []repository.Journal `json:"journals"`
}

func (api *API) getAdminDashboard(w http.ResponseWriter, req *http.Request) {
    api.AllowOrigin(w, req)

    encoder := json.NewEncoder(w)
    journals, err := api.journalRepo.FetchJournals()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        encoder.Encode(AdminErrorResponse{Error: err.Error()})
        return
    }

    encoder.Encode(journals)

    return
}