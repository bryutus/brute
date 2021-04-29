package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bryutus/brute/app/infrastructure"
	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	db := Setup()
	defer db.Close()
	m.Run()
}

func Setup() *gorm.DB {
	db := infrastructure.Init()
	return db
}

func TestGetBrute(t *testing.T) {
	ts := httptest.NewServer(router())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/brute", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %v", http.StatusOK, resp.StatusCode)
	}

	var response map[string]string
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	err = json.Unmarshal([]byte(buf.String()), &response)

	if response["language_code"] != "la" {
		t.Fatalf("Expected response language_code la, got %s", response["language_code"])
	}

	if response["phrase"] != "et tu" {
		t.Fatalf("Expected response phrase et tu, got %s", response["phrase"])
	}
}
