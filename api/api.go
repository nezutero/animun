package api

import (
  "fmt"

	jikan "github.com/darenliang/jikan-go"
)

func GetData() {
  fmt.Println(jikan.GetSchedules("monday"))
}
