package api

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/darenliang/jikan-go"
)

func TestGetData(t *testing.T) {
	result, err := GetData(jikan.ScheduleFilter("monday"))
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	var animeInfos []struct {
		Title         string  `json:"title"`
		TitleEnglish  string  `json:"title_english"`
		TitleJapanese string  `json:"title_japanese"`
		Type          string  `json:"type"`
		Episodes      int     `json:"episodes"`
		Status        string  `json:"status"`
		Airing        bool    `json:"airing"`
		AiredFrom     string  `json:"aired_from"`
		AiredTo       string  `json:"aired_to"`
		Score         float64 `json:"score"`
	}

	err = json.Unmarshal([]byte(result), &animeInfos)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	for _, info := range animeInfos {
		if info.Title == "" {
			t.Errorf("Expected a non-empty title, got an empty string")
		}
		if info.Type == "" {
			t.Errorf("Expected a non-empty type, got an empty string")
		}
	}
}

func TestResponseTime(t *testing.T) {
	startTime := time.Now()
	_, err := GetData(jikan.ScheduleFilter("monday"))
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	elapsed := time.Since(startTime)
	if elapsed > 5*time.Second {
		t.Errorf("Expected response time to be less than 5 seconds, got: %v", elapsed)
	}
}
