package database

import (
	"context"
	"fmt"
	"github.com/beyzaekici/getir-case-study/data/store"
	"github.com/beyzaekici/getir-case-study/model"
	"github.com/beyzaekici/getir-case-study/util"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type mongodb struct {
	collection *mongo.Collection
}

const records = "records"

func ConnectMongo() store.DataManager {
	Ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(viper.GetString("MongoServer")))
	if err != nil {
		panic(err)
	}
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	database := session.Client().Database(viper.GetString("MongoDatabase"))

	recordsCollection := database.Collection(records)
	return &mongodb{collection: recordsCollection}
}

var mongoIns = ConnectMongo()

func MongoManager() store.DataManager { return mongoIns }

func (m *mongodb) Retrieve(input interface{}) (out interface{}, err error) {
	var rData []bson.M
	var Resp model.Response
	var Req model.Request

	Req, _ = input.(model.Request)

	Resp.Code = http.StatusBadRequest
	Resp.Records = rData

	start, err := time.Parse("2006-01-02", Req.StartDate)
	if err != nil {
		Resp.Msg = err.Error()
		return Resp, err
	}

	end, err := time.Parse("2006-01-02", Req.EndDate)
	if err != nil {
		Resp.Msg = err.Error()
		return Resp, err
	}

	datas := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": start,
					"$lt": end,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": Req.MinCount,
					"$lt": Req.MaxCount,
				},
			},
		},
	}

	cursor, err := m.collection.Aggregate(context.TODO(), datas)
	if err != nil {
		Resp.Msg = err.Error()
		return Resp, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			util.Error(err)
		}
	}(cursor, context.TODO())

	if err = cursor.All(context.TODO(), &rData); err != nil {
		Resp.Msg = err.Error()
		return Resp, err
	}

	if len(rData) > 0 {
		Resp.Code = 0
		Resp.Msg = "Success"
		Resp.Records = rData
		return Resp, nil
	}

	Resp.Code = http.StatusNoContent
	Resp.Msg = "Data not Found"
	Resp.Records = rData
	err = fmt.Errorf(Resp.Msg)
	return Resp, err
}
