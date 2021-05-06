package helpers

type UserInfo struct {
	CurrentPrayer string // the prayer user is reading currently
	PrayerPart    int    // part of prayer being read
	PrayerCount   int    // num of prayer
	UserState     string //
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
		CurrentPrayer: "",
		PrayerPart:    1,
		PrayerCount:   1,
		UserState: "",
	}
}

func (um *UsersMap) UpdatePrayer(userID int, prayerName string) {
	userInfo := (*um)[userID]
	userInfo.CurrentPrayer = prayerName
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
	(*um)[userID] = userInfo
}

func (um *UsersMap) FlushUserInfo(userID int) {
	userInfo := (*um)[userID]
	userInfo = UserInfo{
		CurrentPrayer: "",
		PrayerPart:    1,
		PrayerCount:   1,
		UserState: "",
	}

	(*um)[userID] = userInfo
}