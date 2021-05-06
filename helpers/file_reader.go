package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Prayer map[int]string

func getPrayerByName(prayerName string) (Prayer, error) {
	jsonFile, err := os.Open("prayers.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var result map[string]map[int]string
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}

	return result[prayerName], nil
}

func (p Prayer) getPart(part int) (string, bool, error) {
	prayerPart, ok := p[part]
	if !ok {
		return "", true, fmt.Errorf("couldn`t find part %d for prayer", part)
	}

	var isLastPart bool
	_, ok = p[part+1]
	if !ok {
		isLastPart = true
	}

	return prayerPart, isLastPart, nil
}