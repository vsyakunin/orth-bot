package helpers

import tb "gopkg.in/tucnak/telebot.v2"

type UserInfo struct {
	CurrentPrayer  string // the prayer user is reading currently
	PrayerPart     int    // part of prayer being read
	PrayerCount    int    // num of prayer
	UserState      State
	PrayersInState int
	LastMsg        *tb.Message
}

type State string

const (
	StEmpty       State = "empty"
	StFiveMins    State = "5min"
	StFifteenMins State = "15min"
	StThirtyMins  State = "30min"
	StOneHour     State = "1h"
)

type UsersMap map[int]UserInfo

func (um *UsersMap) AddUser(userID int) {
	(*um)[userID] = UserInfo{
		CurrentPrayer:  "",
		PrayerPart:     1,
		PrayerCount:    1,
		UserState:      StEmpty,
		PrayersInState: 1,
	}
}

func (um *UsersMap) UpdatePrayer(userID int) {
	userInfo := (*um)[userID]
	userInfo.CurrentPrayer = getPrayerName(userInfo)
	(*um)[userID] = userInfo
}

func (um *UsersMap) UpdatePrayerPart(userID int) {
	userInfo := (*um)[userID]
	userInfo.PrayerPart++
	(*um)[userID] = userInfo
}

func (um *UsersMap) UpdatePrayerCount(userID int) {
	userInfo := (*um)[userID]
	userInfo.PrayerCount++
	userInfo.PrayerPart = 1
	(*um)[userID] = userInfo
}

func (um *UsersMap) GetUserInfo(userID int) UserInfo {
	return (*um)[userID]
}

func (um *UsersMap) UpdateState(userID int, state State) {
	userInfo := (*um)[userID]
	userInfo.UserState = state
	var prayersInState int
	switch state {
	case StFifteenMins, StThirtyMins:
		prayersInState = 3
	case StOneHour:
		prayersInState = 4
	default:
		prayersInState = 1
	}
	userInfo.PrayersInState = prayersInState
	(*um)[userID] = userInfo
}

func (um *UsersMap) FlushUserInfo(userID int) {
	userInfo := (*um)[userID]
	userInfo = UserInfo{
		CurrentPrayer:  "",
		PrayerPart:     1,
		PrayerCount:    1,
		UserState:      StEmpty,
		PrayersInState: 1,
	}

	(*um)[userID] = userInfo
}

func (um *UsersMap) UpdateLastMsg(userID int, msg *tb.Message) {
	userInfo := (*um)[userID]
	userInfo.LastMsg = msg
	(*um)[userID] = userInfo
}

func (um *UsersMap) GetLastMsg(userID int) *tb.Message {
	userInfo := (*um)[userID]
	return userInfo.LastMsg
}
