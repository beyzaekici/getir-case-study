package search

import (
	"encoding/json"
	"getir-case/api/model"
	"getir-case/api/store/database"
	"getir-case/api/util"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
)

type MongoDb struct{}

func (m *MongoDb) ServeMongo(rw http.ResponseWriter, request *http.Request) {
	var result interface{}
	var mongoResponse model.Response
	var data []bson.M

	mongoResponse.Code = http.StatusBadRequest
	mongoResponse.Records = data

	if request.Method != "POST" {
		mongoResponse.Msg = "Method not allowed"
		rw.WriteHeader(500)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			util.Error(err)
			return
		}
		return
	}
	if nil == request.Body {
		mongoResponse.Msg = "No request content to process"
		rw.WriteHeader(500)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			util.Error(err)
			return
		}
		return
	}

	defer request.Body.Close()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		mongoResponse.Msg = err.Error()
		rw.WriteHeader(500)
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			util.Error(err)
			return
		}
		return
	}

	var content model.Request

	if err = json.Unmarshal(body, &content); err != nil {
		rw.WriteHeader(500)
		mongoResponse.Msg = err.Error()
		err := json.NewEncoder(rw).Encode(mongoResponse)
		if err != nil {
			util.Error(err)
			return
		}
		return
	}

	result, err = database.MongoManager().Retrieve(content)
	if err != nil {
		util.Error(err)
		return
	}
	err = json.NewEncoder(rw).Encode(result)
	if err != nil {
		util.Error(err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
