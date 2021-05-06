package helpers

import (
	"math/rand"
)

const (
	me = "MariyaEgipetskaya"
	nb = "NikitaBesogon"
	ah = "AngelHranitel"
	kp = "KanonPokayannyi"
	af = "Akafist"
)

func GetPrayerPart(info UserInfo) (text string, isLastPart bool) {
	prayer, err := getPrayerByName(info.CurrentPrayer)
	if err != nil {
		return "", true
	}

	prayerPart, isLastPart, err := prayer.getPart(info.PrayerPart)
	if err != nil {
		return "", true
	}

	return prayerPart, isLastPart
}

func get5MinPrayerName() string {
	in := []string{me, ah, nb}
	randomIndex := rand.Intn(len(in))
	return in[randomIndex]
}

func get15MinPrayerName(prayerCount int) string {
	switch prayerCount {
	case 1:
		return nb
	case 2:
		return ah
	default:
		return me
	}
}

func get30MinPrayerName(prayerCount int) string {
	switch prayerCount {
	case 1:
		return nb
	case 2:
		return ah
	default:
		return af
	}
}

func get1hPrayerName(prayerCount int) string {
	switch prayerCount {
	case 1:
		return nb
	case 2:
		return ah
	case 3:
		return af
	default:
		return kp
	}
}

func getPrayerName(userInfo UserInfo) string {
	switch userInfo.UserState {
	case FiveMins:
		return get5MinPrayerName()
	case FifteenMins:
		return get15MinPrayerName(userInfo.PrayerCount)
	case ThirtyMins:
		return get30MinPrayerName(userInfo.PrayerCount)
	case OneHour:
		return get1hPrayerName(userInfo.PrayerCount)
	default:
		return ""
	}
}