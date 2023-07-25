package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const QUERY = `
query ($id: Int) {
  Media (id: $id, type: ANIME) {
    id
    title {
      romaji
      english
      native
    }
  }
}`

func GetData() {
	queryData := map[string]interface{}{
		"query":     QUERY,
		"variables": map[string]int{"id": 15125},
	}

	jsonData, err := json.Marshal(queryData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://graphql.anilist.co/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
}
