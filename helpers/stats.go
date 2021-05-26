package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const statsFileName = "stats/stats.json"

type Stats struct {
	FiveMin    int `json:"five_min"`
	FifteenMin int `json:"fifteen_min"`
	ThirtyMin  int `json:"thirty_min"`
	OneHour    int `json:"one_hour"`
}

func GetStatsText() (string, error) {
	numUsers, err := getTotalNumOfUsers()
	if err != nil {
		return "", err
	}

	stats, err := getStats()
	if err != nil {
		return "", err
	}

	statsText := fmt.Sprintf(GetText(StatsText),
		numUsers,
		stats.FiveMin,
		stats.FifteenMin,
		stats.ThirtyMin,
		stats.OneHour)

	return statsText, nil
}

func getStats() (stats Stats, err error) {
	jsonFile, err := os.Open(statsFileName)
	if err != nil {
		log.Println(err.Error())
		return stats, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(byteValue, &stats)
	if err != nil {
		return stats, err
	}

	return stats, nil
}

var mu sync.Mutex

func saveStats(stats Stats) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := json.MarshalIndent(stats, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(statsFileName, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStats(state State) error {
	stats, err := getStats()
	if err != nil {
		return err
	}

	switch state {
	case StFiveMins:
		stats.FiveMin++
	case StFifteenMins:
		stats.FifteenMin++
	case StThirtyMins:
		stats.ThirtyMin++
	case StOneHour:
		stats.OneHour++
	}

	err = saveStats(stats)
	if err != nil {
		return err
	}

	return nil
}