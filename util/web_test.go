package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TestResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func TestGetRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		if r.URL.Query().Get("id") != "123" {
			t.Errorf("Expected id=123, got %s", r.URL.Query().Get("id"))
		}
		if r.URL.Query().Get("name") != "test" {
			t.Errorf("Expected name=test, got %s", r.URL.Query().Get("name"))
		}

		resp := TestResp{Success: true, Message: "ok"}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	req := &TestReq{ID: 123, Name: "test"}
	resp, err := GetRequest[TestReq, TestResp](server.URL, req)
	if err != nil {
		t.Fatalf("GetRequest failed: %v", err)
	}
	if !resp.Success {
		t.Error("Expected Success=true")
	}
}

func TestPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		var reqBody TestReq
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatal(err)
		}
		if reqBody.ID != 456 {
			t.Errorf("Expected ID 456, got %d", reqBody.ID)
		}

		resp := TestResp{Success: true, Message: "created"}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	req := &TestReq{ID: 456, Name: "post"}
	resp, err := PostRequest[TestReq, TestResp](server.URL, req)
	if err != nil {
		t.Fatalf("PostRequest failed: %v", err)
	}
	if resp.Message != "created" {
		t.Errorf("Expected created, got %s", resp.Message)
	}
}

func TestDeleteRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}
		var reqBody TestReq
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatal(err)
		}
		if reqBody.ID != 789 {
			t.Errorf("Expected ID 789, got %d", reqBody.ID)
		}

		resp := TestResp{Success: true, Message: "deleted"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	req := &TestReq{ID: 789, Name: "delete"}
	resp, err := DeleteRequest[TestReq, TestResp](server.URL, req)
	if err != nil {
		t.Fatalf("DeleteRequest failed: %v", err)
	}
	if resp.Message != "deleted" {
		t.Errorf("Expected deleted, got %s", resp.Message)
	}
}

func TestRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	req := &TestReq{}
	_, err := GetRequest[TestReq, TestResp](server.URL, req)
	if err == nil {
		t.Error("Expected error for 500")
	}

	_, err = PostRequest[TestReq, TestResp](server.URL, req)
	if err == nil {
		t.Error("Expected error for 500")
	}

	_, err = DeleteRequest[TestReq, TestResp](server.URL, req)
	if err == nil {
		t.Error("Expected error for 500")
	}
}
