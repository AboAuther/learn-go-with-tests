package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"lrg": 20,
			"chj": 10,
		},
		nil,
	}
	server := &PlayerServer{&store}
	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "return lrg's score",
			player:             "lrg",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		}, {
			name:               "return chj's score",
			player:             "chj",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "10",
		}, {
			name:               "return 404 on missing players",
			player:             "lyl",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetScoreRequest(tt.player)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			assertStatus(t, response.Code, tt.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}
func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}
	t.Run("it returns accepted on Post", func(t *testing.T) {
		player := "lrg"
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", player), nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}
func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return req
}
func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
