package helpers

type UserInfo struct {
	CurrentPrayer  string // the prayer user is reading currently
	PrayerPart     int    // part of prayer being read
	PrayerCount    int    // num of prayer
	UserState      string //
	PrayersInState int
}

const (
	FiveMins    = "5min"
	FifteenMins = "15min"
	ThirtyMins  = "30min"
	OneHour     = "1h"
)

type UsersMap map[int]UserInfo

func (um *UsersMap) AddUser(userID int) {
	(*um)[userID] = UserInfo{
		CurrentPrayer:  "",
		PrayerPart:     1,
		PrayerCount:    1,
		UserState:      "",
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

func (um *UsersMap) UpdateState(userID int, state string) {
	userInfo := (*um)[userID]
	userInfo.UserState = state
	var prayersInState int
	switch state {
	case FifteenMins, ThirtyMins:
		prayersInState = 3
	case OneHour:
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
		UserState:      "",
		PrayersInState: 1,
	}

	(*um)[userID] = userInfo
}
