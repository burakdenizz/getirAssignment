package tests

import (
	"bytes"
	"encoding/json"
	"getirAssignment/controllers"
	"getirAssignment/database"
	"getirAssignment/models"
	"getirAssignment/routes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestSetKeyValue(t *testing.T) {
	keyValueStore := setupApp(false)
	var pBody models.KeyValue
	pBody.Key = "TestKey"
	pBody.Value = "TestValue"
	body, _ := json.Marshal(pBody)
	req := httptest.NewRequest("POST", "/api/in-memory", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	controllers.InMemory(w, req)
	assert.NotNil(t, keyValueStore["TestKey"])
	assert.Equal(t, "TestValue", keyValueStore["TestKey"])
}
func TestGetKeyValue(t *testing.T) {
	keyValueStore := setupApp(false)
	keyValueStore["TestKey"] = "TestValue"
	var pBody models.KeyValue
	pBody.Key = "TestKey"
	body, _ := json.Marshal(pBody)
	req := httptest.NewRequest("GET", "/api/in-memory", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	controllers.InMemory(w, req)

	resp := w.Result()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	respString := string(respBytes)
	checkString := "{\"key\":\"TestKey\",\"value\":\"TestValue\"}"
	assert.NotNil(t, respString)
	assert.Equal(t, checkString, respString)
}
func TestGetDataFromMongo(t *testing.T) {
	setupApp(true)
	var getResponse models.GetResponse
	var pBody models.GetRequest
	pBody.StartDate = "2016-01-26"
	pBody.EndDate = "2021-10-12"
	pBody.Gte = "2000"
	pBody.Lte = "3000"
	body, _ := json.Marshal(pBody)
	req := httptest.NewRequest("GET", "/api/in-memory", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	controllers.GetData(w, req)
	resp := w.Result()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBytes, &getResponse)
	assert.Equal(t, 0, getResponse.Code)
	assert.Equal(t, "Success", getResponse.Msg)
	assert.NotNil(t, getResponse.Records)
}

var routesCheck bool

func setupApp(dbNeed bool) map[string]string {
	if !routesCheck {
		routes.Setup()
		routesCheck = true
	}
	if dbNeed {
		database.Connect()
	}
	keyValueStore := map[string]string{}
	controllers.KeyValueStore = keyValueStore
	return keyValueStore
}
