package helpers

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

func GetIconForPrayer(prayerName string) *tb.Photo {
	var fileName string
	fileName = fmt.Sprintf("icons/%s.png", prayerName)
	return &tb.Photo{File: tb.FromDisk(fileName)}
}