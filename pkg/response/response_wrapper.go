package response

import (
	"encoding/json"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter, data Response) {
	response := httpResponse{
		Data: data.Data,
		Header: headerResponse{
			Error: data.Error,
		},
	}
	responseByte, _ := json.Marshal(response)

	statusCode := http.StatusOK
	if data.HttpStatusCode != 0 {
		statusCode = data.HttpStatusCode
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseByte)
}

func WriteErrorResponse(w http.ResponseWriter, data Response) {
	response := httpResponse{
		Data: data.Data,
		Header: headerResponse{
			Error: data.Error,
		},
	}
	responseByte, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.HttpStatusCode)
	w.Write(responseByte)
}
