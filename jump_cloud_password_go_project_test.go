package main

import (
		"testing"
		"encoding/json"
		"net/http"
		"fmt"
		"os"
		"net/http/httptest"
		"strings"
		)

func TestEncodePassword(t *testing.T) {
	encode := EncodePassword("angryMonkey")
	expected_encode := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	if encode != expected_encode {
		t.Errorf("Encoded password was incorrect, got: %s, want: %s.", encode, expected_encode)
	}
}

func TestUnsupportedURLEndpoint(t *testing.T) {
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodGet, "/garbage", nil)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusBadRequest)
    }

}

func TestEncodePasswordEndpoint(t *testing.T) {
	if _, existserr := os.Stat(filepath); !os.IsNotExist(existserr) {
		err := os.Remove(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	expected_encode := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodPost, "/hash", strings.NewReader("password=angryMonkey"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusOK)
    }

    var m map[string]interface{}
    json.Unmarshal(rr.Body.Bytes(), &m)

    if m["password"] != expected_encode {
		t.Errorf("Encoded password was incorrect, got: %s, want: %s.", m, expected_encode)
	}
}

func TestGetErrorEncodePasswordEndpoint(t *testing.T) {
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodGet, "/hash", strings.NewReader("password=angryMonkey"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusBadRequest)
    }
}

func TestPutErrorEncodePasswordEndpoint(t *testing.T) {
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodPut, "/hash", strings.NewReader("password=angryMonkey"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusBadRequest)
    }
}

func TestDeleteErrorEncodePasswordEndpoint(t *testing.T) {
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/hash", strings.NewReader("password=angryMonkey"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusBadRequest)
    }
}

func TestEncodeStatsEndpoint(t *testing.T) {
	expected_total := 1
	if _, existserr := os.Stat(filepath); !os.IsNotExist(existserr) {
		err := os.Remove(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodPost, "/hash", strings.NewReader("password=angryMonkey"))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusOK)
    }

    req, _ = http.NewRequest(http.MethodGet, "/stats", nil)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr = httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusOK)
    }

    var m Stats
    json.Unmarshal(rr.Body.Bytes(), &m)

    if m.Total != expected_total {
		t.Errorf("Encoded password was incorrect, got: %d, want: %d.", m.Total, expected_total)
	}
}

func TestPostErrorEncodeStatsEndpoint(t *testing.T) {
	if _, existserr := os.Stat(filepath); !os.IsNotExist(existserr) {
		err := os.Remove(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodPost, "/stats", nil)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusBadRequest)
    }

}

func TestDeleteErrorEncodeStatsEndpoint(t *testing.T) {
	if _, existserr := os.Stat(filepath); !os.IsNotExist(existserr) {
		err := os.Remove(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	router := NewRouter()

    req, _ := http.NewRequest(http.MethodDelete, "/stats", nil)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req)

    if rr.Code != http.StatusBadRequest {
    	t.Errorf("Status code was incorrect, got: %v, want: %v.", rr.Code, http.StatusBadRequest)
    }

}