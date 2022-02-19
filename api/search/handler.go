package search

import (
	"encoding/json"
	"getir-case/api/data/database"
	"getir-case/api/model"
	"getir-case/api/util"
	"go.mongodb.org/mongo-driver/bson"
	"io"
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			util.Error(err)
		}
	}(request.Body)

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
	rw.WriteHeader(http.StatusAccepted)
	rslt := json.NewEncoder(rw).Encode(result)
	if rslt != nil {
		util.Error(err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

}
