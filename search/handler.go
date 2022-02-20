package search

import (
	"encoding/json"
	"github.com/beyzaekici/getir-case-study/data/database"
	"github.com/beyzaekici/getir-case-study/model"
	"net/http"
)

type MongoDb struct{}

func (m *MongoDb) SearchInMongo(rw http.ResponseWriter, request *http.Request) {
	var result interface{}
	if request.URL.Path == "/records" && request.Method == http.MethodPost {
		var recordQuery model.Request

		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&recordQuery)

		result, _ = database.MongoManager().Retrieve(recordQuery)

		jData, _ := json.Marshal(result)

		rw.WriteHeader(http.StatusOK)
		rw.Write(jData)
		return
	}
}
