package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

type UserInfo struct {
	CurrentPrayer  string      `json:"current_prayer"` // the prayer user is reading currently
	PrayerPart     int         `json:"prayer_part"`    // part of prayer being read
	PrayerCount    int         `json:"prayer_count"`   // num of prayer
	UserState      State       `json:"user_state"`
	PrayersInState int         `json:"prayers_in_state"`
	LastMsg        *tb.Message `json:"last_message"`
}

type State string

const (
	StEmpty       State = "empty"
	StFiveMins    State = "5min"
	StFifteenMins State = "15min"
	StThirtyMins  State = "30min"
	StOneHour     State = "1h"

	fileNameRaw = "users/%d.json"
)

var InitialUserInfo = UserInfo{
	CurrentPrayer:  "",
	PrayerPart:     1,
	PrayerCount:    1,
	UserState:      StEmpty,
	PrayersInState: 1,
}

func GetPrayersInState(state State) int {
	switch state {
	case StFifteenMins, StThirtyMins:
		return 3
	case StOneHour:
		return 4
	default:
		return 1
	}
}


func createInitialUserInfo(userID int) (userInfo UserInfo, err error) {
	fileName := fmt.Sprintf(fileNameRaw, userID)

	userInfo = InitialUserInfo

	file, err := json.MarshalIndent(userInfo, "", " ")
	if err != nil {
		return
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		return
	}

	log.Printf("added new user with ID %v", userID)
	return
}

func GetUserInfo(userID int) (userInfo UserInfo, err error) {
	fileName := fmt.Sprintf(fileNameRaw, userID)

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return createInitialUserInfo(userID)
	}

	err = json.Unmarshal(file, &userInfo)
	if err != nil {
		return
	}

	return
}

func UpdateUserInfo(userID int, userInfo UserInfo) error {
	fileName := fmt.Sprintf(fileNameRaw, userID)

	file, err := json.MarshalIndent(userInfo, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		return err
	}

	return nil
}


