package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	envfile := ".env.test"

	if err := infrastructure.Refresh(envfile); err != nil {
		log.Fatalf("Failed to refresh database %v", err)
	}

	db := infrastructure.Init(envfile)

	db.Exec("insert into aphorisms (phrase, language_code) values (?, ?) ", "et tu", "la")
	db.Exec("insert into aphorisms (phrase, language_code) values (?, ?) ", "お前もか", "ja")
	db.Exec("insert into aphorisms (phrase, language_code) values (?, ?) ", "even you", "en")

	return db
}

func TestShowBruteOK(t *testing.T) {
	testCases := []struct {
		in            string
		language_code string
		phrase        string
	}{
		{"", "la", "et tu"},
		{"ja", "ja", "お前もか"},
	}

	ts := httptest.NewServer(router())
	defer ts.Close()

	for _, test := range testCases {
		url := fmt.Sprintf("%s/brute?language_code=%s", ts.URL, test.in)
		if test.in == "" {
			url = fmt.Sprintf("%s/brute", ts.URL)
		}

		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got %v", http.StatusOK, resp.StatusCode)
		}

		response := getResponse(resp.Body, t)

		if response["language_code"] != test.language_code {
			t.Fatalf("Expected response language_code %s, got %s", test.language_code, response["language_code"])
		}

		if response["phrase"] != test.phrase {
			t.Fatalf("Expected response phrase %s, got %s", test.phrase, response["phrase"])
		}
	}
}

func TestShowBruteNG(t *testing.T) {
	testCases := []struct {
		in      string
		status  int
		message string
	}{
		{"", 404, "record not found: language_code="},
		{"isl", 400, "Key: 'requestShowAphorism.LanguageCode' Error:Field validation for 'LanguageCode' failed on the 'len' tag"},
	}

	ts := httptest.NewServer(router())
	defer ts.Close()

	for _, test := range testCases {
		resp, err := http.Get(fmt.Sprintf("%s/brute?language_code=%s", ts.URL, test.in))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != test.status {
			t.Fatalf("Expected status code %d, got %v", test.status, resp.StatusCode)
		}

		response := getResponse(resp.Body, t)

		if response["message"] != test.message {
			t.Fatalf("Expected response error message %s, got %s", test.message, response["message"])
		}
	}
}

func getResponse(r io.Reader, t *testing.T) map[string]string {
	var response map[string]string
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	if err := json.Unmarshal([]byte(buf.String()), &response); err != nil {
		t.Fatalf("Failed unmarshal response")
	}

	return response
}
