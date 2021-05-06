package helpers

type UserInfo struct {
	ChosenPrayer string
	PrayerPart   int
}

type UsersMap map[int]UserInfo

func (um *UsersMap) AddUser(userID int) {
	(*um)[userID] = UserInfo{
		ChosenPrayer: "",
		PrayerPart: 1,
	}
}

func (um *UsersMap) UpdatePrayer(userID int, prayerName string) {
	userInfo := (*um)[userID]
	userInfo = UserInfo{
		ChosenPrayer: prayerName,
		PrayerPart: 1,
	}
	(*um)[userID] = userInfo
}

func (um *UsersMap) UpdatePrayerPart(userID int) {
	userInfo := (*um)[userID]
	userInfo.PrayerPart++
	(*um)[userID] = userInfo
}

func (um *UsersMap) GetUserInfo(userID int) UserInfo {
	return (*um)[userID]
}
