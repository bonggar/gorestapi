package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bonggar/gorestapi/database"

	"github.com/gin-gonic/gin"
)

func TestCreateUserWIthUnvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]string{
		"email":   "bonggar2@mail.com",
		"phone":   "081212340002",
		"dob":     "2000-02-02",
		"address": "jakarta",
		"gender":  "m",
	})
	if err != nil {
		t.Fail()
	}
	database.SQLiteDBConnect()
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	CreateUser(c)
	bodySb, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Error reading body: %v\n", err)
	}
	body := string(bodySb)
	t.Logf("Body: %v\n", body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}
	if w.Code == 422 {
		t.Logf("Create User Success")
	} else {
		t.Fail()
	}
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	jsonData, err := json.Marshal(map[string]string{
		"name":    "Bonggar",
		"email":   "bonggar@gmail.com",
		"phone":   "081212340222",
		"dob":     "2000-02-02",
		"address": "jakarta",
		"gender":  "m",
	})
	if err != nil {
		t.Fail()
	}
	database.SQLiteDBConnect()
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fail()
	}
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	CreateUser(c)
	bodySb, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Error reading body: %v\n", err)
	}
	body := string(bodySb)
	t.Logf("Body: %v\n", body)
	var decodedResponse interface{}
	err = json.Unmarshal(bodySb, &decodedResponse)
	if err != nil {
		t.Fatalf("Cannot decode response <%p> from server. Err: %v", bodySb, err)
	}
	if w.Code == 201 {
		t.Logf("Create User Success")
	} else if w.Code == 409 {
		t.Logf("Prevent duplicated user creation success")
	} else if w.Code == 422 {
		t.Logf("Prevent duplicated user creation success")
	} else {
		t.Fail()
	}
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.SQLiteDBConnect()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	GetUsers(c)
	if w.Code == 200 {
		t.Logf("Get Users List Success")
	}
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	database.SQLiteDBConnect()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("id", "0")
	GetUsers(c)
	if w.Code == 200 {
		t.Logf("Get User Success")
	} else if w.Code == 404 {
		t.Logf("User not found in DB")
	} else {
		t.Fail()
	}
}
