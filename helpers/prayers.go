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

func GetPrayer(info UserInfo) (text string, isLastPart bool) {
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

func Get5MinPrayerName() string {
	in := []string{me, ah, nb}
	randomIndex := rand.Intn(len(in))
	return in[randomIndex]
}

func Get15MinPrayerName(userInfo UserInfo) string {
	switch userInfo.PrayerCount {
	case 1:
		return nb
	case 2:
		return ah
	default:
		return me
	}
}

func Get30MinPrayerName(userInfo UserInfo) string {
	switch userInfo.PrayerCount {
	case 1:
		return nb
	case 2:
		return ah
	default:
		return af
	}
}

func Get1hPrayerName(userInfo UserInfo) string {
	switch userInfo.PrayerCount {
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