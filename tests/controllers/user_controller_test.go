package tests

import (
	models "api/app/models/user"
	"api/config"
	"api/tests"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadEnvs()

	code := m.Run()

	os.Exit(code)
}

func setupTest(t *testing.T) {
	tests.RunMigrations()
	t.Cleanup(func() {
		tests.CleanDb()
	})

}

func TestCreateUser(t *testing.T) {
	setupTest(t)
	user := models.UserCreateRequest{
		Name:     "Name Test",
		Nick:     "Nick Test",
		Email:    "test@example.com",
		Password: "Password Test",
	}
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	response := tests.MakeRequest(t, req)

	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)
	userResponse := responseData["user"].(map[string]interface{})

	assert.Equal(t, http.StatusCreated, response.Code)

	assert.Equal(t, user.Name, userResponse["name"])
	assert.Equal(t, user.Nick, userResponse["nick"])
	assert.Equal(t, user.Email, userResponse["email"])
	assert.NotEmpty(t, userResponse["id"])
	assert.NotEmpty(t, userResponse["created_at"])
}

func TestListUsers(t *testing.T) {
	setupTest(t)
	user, err := tests.CreateUser()
	if err != nil {
		t.Fatalf("failed do create user: %v", err)
	}

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	response := tests.MakeRequest(t, req)

	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	users := responseData["data"].([]interface{})

	firstUser := users[0].(map[string]interface{})

	assert.Equal(t, http.StatusOK, response.Code)

	assert.Len(t, users, 1)

	assert.Equal(t, user.Name, firstUser["name"])
	assert.Equal(t, user.Nick, firstUser["nick"])
	assert.Equal(t, user.Email, firstUser["email"])
	assert.Equal(t, user.ID, uint64(firstUser["id"].(float64)))
	assert.NotEmpty(t, firstUser["created_at"])
}

func TestShowUser(t *testing.T) {
	setupTest(t)
	user, err := tests.CreateUser()
	if err != nil {
		t.Fatalf("failed do create user: %v", err)
	}

	req, err := http.NewRequest("GET", "/users/"+strconv.FormatUint(user.ID, 10), nil)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	response := tests.MakeRequest(t, req)

	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	userResponse := responseData["user"].(map[string]interface{})

	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, user.Name, userResponse["name"])
	assert.Equal(t, user.Nick, userResponse["nick"])
	assert.Equal(t, user.Email, userResponse["email"])
	assert.Equal(t, user.ID, uint64(userResponse["id"].(float64)))
	assert.NotEmpty(t, userResponse["created_at"])
}

func TestUpdateUser(t *testing.T) {
	setupTest(t)

	user, err := tests.CreateUser()
	if err != nil {
		t.Fatalf("failed do create user: %v", err)
	}

	userUpdate := models.UserUpdateRequest{
		Name:  "Name Update",
		Nick:  "Nick Update",
		Email: "Update@Update.Update",
	}
	body, _ := json.Marshal(userUpdate)

	req, err := http.NewRequest("PUT", "/users/"+strconv.FormatUint(user.ID, 10), bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	response := tests.MakeRequest(t, req)

	var responseData map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseData)
	userResponse := responseData["user"].(map[string]interface{})

	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, userUpdate.Name, userResponse["name"])
	assert.Equal(t, userUpdate.Nick, userResponse["nick"])
	assert.Equal(t, userUpdate.Email, userResponse["email"])
	assert.NotEmpty(t, userResponse["id"])
	assert.NotEmpty(t, userResponse["created_at"])
}

func TestDeleteUser(t *testing.T) {
	setupTest(t)
	user, err := tests.CreateUser()
	if err != nil {
		t.Fatalf("failed do create user: %v", err)
	}

	req, err := http.NewRequest("DELETE", "/users/"+strconv.FormatUint(user.ID, 10), nil)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	response := tests.MakeRequest(t, req)

	assert.Equal(t, http.StatusNoContent, response.Code)

	searchUser, _ := tests.FindUserById(user.ID)

	assert.Empty(t, searchUser.Email)
	assert.Empty(t, searchUser.ID)
}
