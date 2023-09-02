package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	pb "github.com/divyam-rai/personal-finance-tracker/packages/models/protos"
)

func (a *apiRouter) handleCreateDeposit(rw http.ResponseWriter, req *http.Request) {
	payload, response := &pb.FixedDeposit{}, ApiResponse{}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.Error = err.Error()
		a.WriteJson(rw, http.StatusBadRequest, response)
		return
	}

	if err := json.Unmarshal(body, payload); err != nil {
		response.Error = err.Error()
		a.WriteJson(rw, http.StatusBadRequest, response)
		return
	}

	response.Success = true
	response.Data = payload
	a.WriteJson(rw, http.StatusOK, response)
	return
}

func (a *apiRouter) depositRoutes() []Route {
	return []Route{
		{
			Method:      "POST",
			Path:        "/deposits/create",
			HandlerFunc: a.handleCreateDeposit,
		},
	}
}
