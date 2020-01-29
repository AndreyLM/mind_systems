package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

// func TestShortener(t *testing.T) {
// 	postData := map[string]string{
// 		"url": "http://google.com",
// 	}
// 	jsonData, _ := json.Marshal(postData)
// 	reader := bytes.NewReader(jsonData)
// 	req, err := http.NewRequest("POST", "/create-url", reader)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(Shortener("localhost:8000"))
// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: %v", status)
// 	}
// }

func TestFullProcess(t *testing.T) {
	client := http.Client{}
	postData := map[string]string{
		"url": "http://google.com",
	}
	jsonData, _ := json.Marshal(postData)
	reader := bytes.NewReader(jsonData)

	resp, err := client.Post("http://localhost:8000/create-url", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	result := struct {
		URL string `json:"shortened_url"`
	}{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(body, &result)

	resp2, err := client.Get(result.URL)
	if err != nil {
		t.Fatal(err)
	}

	respStatus := resp2.StatusCode
	if respStatus != http.StatusOK {
		t.Fatalf("Invalid status: want -> %v, get -> %v", http.StatusOK, respStatus)
	}

	resp3, err := client.Get(result.URL + "invalid_url")
	if err != nil {
		t.Fatal(err)
	}
	respStatus = resp3.StatusCode
	if respStatus != http.StatusNotFound {
		t.Fatalf("Invalid status: want -> %v, get -> %v", http.StatusNotFound, respStatus)
	}
}
