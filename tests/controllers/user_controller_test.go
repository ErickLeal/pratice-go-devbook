package tests

import (
	models "api/app/models/user"
	"api/config"
	"api/tests"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	config.LoadEnvs()
	db, err := config.ConnectDatabase()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	err = tests.RunMigrations(db)
	if err != nil {
		t.Fatalf("Erro ao rodar migrações: %v", err)
	}

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

	err = tests.CleanDb(db)
	if err != nil {
		t.Fatalf("Erro to clean db: %v", err)
	}

}

func TestListUsers(t *testing.T) {
	config.LoadEnvs()
	db, err := config.ConnectDatabase()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	err = tests.RunMigrations(db)
	if err != nil {
		t.Fatalf("Erro ao rodar migrações: %v", err)
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
