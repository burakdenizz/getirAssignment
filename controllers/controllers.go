package controllers

import (
	"encoding/json"
	"fmt"
	"getirAssignment/database"
	"getirAssignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var KeyValueStore map[string]string

func GetData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("Content type must be 'application/json, but you provide '%s'", ct)))
		}
		var rawFilter map[string]string
		if err := json.Unmarshal(bodyBytes, &rawFilter); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		startDate, startDateError := time.Parse("2006-01-02", rawFilter["startDate"])
		endDate, endDateError := time.Parse("2006-01-02", rawFilter["endDate"])
		gte, _ := strconv.Atoi(rawFilter["gte"])
		lte, _ := strconv.Atoi(rawFilter["lte"])
		fmt.Println(gte, lte)
		if startDateError != nil || endDateError != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Date input is wrong please input date like this : 2006-12-30"))
			return
		}

		// Match filter to filter data according to totalCount and created at
		matchStage := bson.D{{"$match", bson.D{
			{"totalCount", bson.D{
				{"$gte", gte},
				{"$lte", lte}},
			},
			{"createdAt", bson.D{
				{"$gte", startDate},
				{"$lte", endDate}},
			}}}}

		// Group stage to aggregate over the table and create the totalCount by summing the integers in counts field
		groupStage := bson.D{
			{"$project", bson.D{
				{"key", "$_id"},
				{"totalCount", bson.D{
					{"$sum", "$counts"}}},
				{"createdAt", "$createdAt"}}}}

		cursor, err := database.Query("getir-case-study", "records", matchStage, groupStage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There is a problem creating the cursor please check the inputs"))
			w.Write([]byte(err.Error()))
			return
		}
		records := database.FetchAllData(cursor)

		response := models.GetResponse{
			Code: 0, Msg: "Success", Records: records,
		}
		jsonBytes, _ := json.Marshal(response)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

}

func InMemory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("Content type must be 'application/json, but you provide '%s'", ct)))
		}
		var cKey map[string]string
		if err := json.Unmarshal(bodyBytes, &cKey); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		key := cKey["key"]
		value := KeyValueStore[key]
		if value == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("The key value pair you are looking for is not present in the store"))
			return
		}
		keyValue := models.KeyValue{Key: key, Value: KeyValueStore[key]}
		responseBytes, err := json.Marshal(keyValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)

	case "POST":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("Content type must be 'application/json, but you provide '%s'", ct)))
		}
		var cKey map[string]string
		if err := json.Unmarshal(bodyBytes, &cKey); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		key := cKey["key"]
		value := cKey["value"]
		KeyValueStore[key] = value

		keyValue := models.KeyValue{Key: key, Value: KeyValueStore[key]}
		responseBytes, err := json.Marshal(keyValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
