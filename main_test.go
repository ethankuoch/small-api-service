package main

import (
	"bytes"
	"encoding/json"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBaseballPlayers(t *testing.T) {
	r := SetUpRouter()
	r.GET("/baseballPlayers", getBaseballPlayers)
	req, _ := http.NewRequest("GET", "/baseballPlayers", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var baseballPlayers []baseballPlayer
	err := json.Unmarshal(w.Body.Bytes(), &baseballPlayers)
	if err != nil {
		return
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, baseballPlayers)
}

func TestPostBaseballPlayers(t *testing.T) {
	r := SetUpRouter()
	r.POST("/baseballPlayers", postBaseballPlayers)
	baseballPlayerId := xid.New().String()
	baseballPlayer := baseballPlayer{
		ID:                 baseballPlayerId,
		Name:               "Nick Castellanos",
		Team:               "Philadelphia Phillies",
		BattingAverage:     0.244,
		OnBasePercentage:   0.301,
		OnBasePlusSlugging: 0.708,
	}
	jsonValue, _ := json.Marshal(baseballPlayer)
	req, _ := http.NewRequest("POST", "/baseballPlayers", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
