package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	RootDir            string = getProjectRoot()
	Environment        string
	DBConnection       string
	DBConnectionString string
	ThirdPartyAPIURL   string
)

func getProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get root workdir:", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("Faild to find root workdir")
		}
		dir = parent
	}
}

func LoadEnvs() {
	envFile := filepath.Join(RootDir, ".env")
	if flag.Lookup("test.v") != nil {
		envFile = filepath.Join(RootDir, ".env.testing")
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Failed to load env file %s: %v", envFile, err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	DBConnection = os.Getenv("DB_CONNECTION")
	DBConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	ThirdPartyAPIURL = os.Getenv("THIRD_PARTY_API_URL")
}
