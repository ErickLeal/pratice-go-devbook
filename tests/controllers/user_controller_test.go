package tests

import (
	models "api/app/models/user"
	"api/config"
	"api/tests"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.LoadEnvs()

	code := m.Run()

	os.Exit(code)
}

func setupTest(t *testing.T) {
	err := tests.RunMigrations()
	if err != nil {
		t.Fatalf("Erro ao rodar migrações: %v", err)
	}

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
		t.Fatalf("Erro ao criar request: %v", err)
	}

	response := tests.MakeRequest(t, req)

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Erro ao ler resposta: %v", err)
	}
	println("responseData")
	println(string(responseData))

	err = tests.CleanDb()
	if err != nil {
		t.Fatalf("Erro to clean db: %v", err)
	}

}

func TestListUsers(t *testing.T) {
	setupTest(t)
	err := tests.CreateUser()
	if err != nil {
		t.Fatalf("Erro ao criar user: %v", err)
	}

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Erro ao criar request: %v", err)
	}

	response := tests.MakeRequest(t, req)

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Erro ao ler resposta: %v", err)
	}
	println("responseData")
	println(string(responseData))

}
