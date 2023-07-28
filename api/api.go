package api

import (
	"encoding/json"
	"fmt"

	jikan "github.com/darenliang/jikan-go"
)

func GetData() {
	filter := jikan.ScheduleFilterMonday
	schedules, err := jikan.GetSchedules(filter)
	if err != nil {
		fmt.Println("[ERROR] error getting schedules:", err)
		return
	}

	type AnimeInfo struct {
		Title           string `json:"title"`
		TitleEnglish    string `json:"title_english"`
		TitleJapanese   string `json:"title_japanese"`
		Type            string `json:"type"`
		Episodes        int    `json:"episodes"`
		Status          string `json:"status"`
		Airing          bool   `json:"airing"`
		AiredFrom       string `json:"aired_from"`
		AiredTo         string `json:"aired_to"`
		Score           float64 `json:"score"`
	}

	var animeInfos []AnimeInfo

	for _, anime := range schedules.Data {
		info := AnimeInfo{
			Title:         anime.Title,
			TitleEnglish:  anime.TitleEnglish,
			TitleJapanese: anime.TitleJapanese,
			Type:          anime.Type,
			Episodes:      anime.Episodes,
			Status:        anime.Status,
			Airing:        anime.Airing,
			AiredFrom:     anime.Aired.From.Format("2000-02-02"),
			AiredTo:       anime.Aired.To.Format("2000-02-02"),
			Score:         anime.Score,
		}

		animeInfos = append(animeInfos, info)
	}

	prettyJSON, err := json.MarshalIndent(animeInfos, "", "  ")
	if err != nil {
		fmt.Println("[ERROR] error formatting data:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
