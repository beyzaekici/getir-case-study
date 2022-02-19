package data

import (
	"encoding/json"
	"fmt"
	"getir-case/api/model"
	"getir-case/api/util"
	"net/http"
)

func (d *DataStore) SetInMemory(w http.ResponseWriter, r *http.Request) {
	var input model.DataInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		util.Error(err)
		return
	}

	_ = d.SetKey(input.Key, input.Value)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&input)
	if err != nil {
		util.Error(err)
		return
	}
}

func (d *DataStore) GetInMemory(w http.ResponseWriter, r *http.Request) {
	HeaderKey := r.Header.Get("key")
	value, err := d.GetKey(HeaderKey)
	if err != nil {
		_, err := fmt.Fprintf(w, "%+v", err.Error())
		if err != nil {
			util.Error(err)
			return
		}
	} else {
		out := model.DataInput{Key: HeaderKey, Value: value}
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			util.Error(err)
			return
		}
	}
}
