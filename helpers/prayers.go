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
	prayer, err := getPrayerByName(info.ChosenPrayer)
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
