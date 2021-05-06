package main

import (
	"log"
	"time"

	"github.com/vsyakunin/orth-bot/helpers"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	usersMap := make(helpers.UsersMap)

	b.Handle(tb.OnText, func(m *tb.Message) {
		if _, ok := usersMap[m.Sender.ID]; !ok {
			usersMap.AddUser(m.Sender.ID)
			log.Printf("added new user with ID %v", m.Sender.ID)
		}
		usersMap.FlushUserInfo(m.Sender.ID)
		_, err = b.Send(m.Sender, helpers.IntroText, helpers.MakeReplyMarkup(helpers.InlineBtnProceedStart), tb.ModeHTML)
		if err != nil {
			log.Println(err.Error())
		}
	})

	b.Handle(&helpers.InlineBtnProceedStart, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		_, err = b.Send(c.Sender, helpers.StartText, helpers.MakeReplyMarkup(helpers.InlineBtnLetsStart), tb.ModeHTML)
		if err != nil {
			log.Println(err.Error())
		}
	})

	b.Handle(&helpers.InlineBtnLetsStart, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		_, err = b.Send(c.Sender, helpers.QuestionText, helpers.MakeReplyMarkup(
			helpers.InlineBtn5min, helpers.InlineBtn15min,
			helpers.InlineBtn30min, helpers.InlineBtn1h), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}
	})

	b.Handle(&helpers.InlineBtn5min, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		usersMap.UpdateState(c.Sender.ID, helpers.FiveMins)
		usersMap.UpdatePrayer(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(usersMap.GetUserInfo(c.Sender.ID))

		_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}
		usersMap.UpdatePrayerPart(c.Sender.ID)
	})

	b.Handle(&helpers.InlineBtn15min, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		usersMap.UpdateState(c.Sender.ID, helpers.FifteenMins)
		usersMap.UpdatePrayer(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(usersMap.GetUserInfo(c.Sender.ID))

		_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}
		usersMap.UpdatePrayerPart(c.Sender.ID)
	})

	b.Handle(&helpers.InlineBtn30min, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		usersMap.UpdateState(c.Sender.ID, helpers.ThirtyMins)
		usersMap.UpdatePrayer(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(usersMap.GetUserInfo(c.Sender.ID))

		_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}
		usersMap.UpdatePrayerPart(c.Sender.ID)
	})

	b.Handle(&helpers.InlineBtn1h, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		usersMap.UpdateState(c.Sender.ID, helpers.OneHour)
		usersMap.UpdatePrayer(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(usersMap.GetUserInfo(c.Sender.ID))

		_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}
		usersMap.UpdatePrayerPart(c.Sender.ID)
	})

	b.Handle(&helpers.InlineBtnNextPrayer, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		usersMap.UpdatePrayerCount(c.Sender.ID)
		usersMap.UpdatePrayer(c.Sender.ID)

		text, _ := helpers.GetPrayerPart(usersMap.GetUserInfo(c.Sender.ID))

		_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}
		usersMap.UpdatePrayerPart(c.Sender.ID)
	})

	b.Handle(&helpers.InlineBtnNextPart, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		text, isLastPart := helpers.GetPrayerPart(usersMap.GetUserInfo(c.Sender.ID))
		if !isLastPart {
			_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
				helpers.InlineBtnNextPart), tb.ModeHTML)
		} else {
			userInfo := usersMap.GetUserInfo(c.Sender.ID)
			if userInfo.UserState == helpers.FiveMins ||
				userInfo.PrayerCount == userInfo.PrayersInState {
				_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
					helpers.InlineBtnAmen), tb.ModeHTML)
			} else {
				_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
					helpers.InlineBtnNextPrayer), tb.ModeHTML)
			}
		}

		if err != nil {
			log.Println(err.Error())
		}
		usersMap.UpdatePrayerPart(c.Sender.ID)
	})

	b.Handle(&helpers.InlineBtnAmen, func(c *tb.Callback) {
		err = b.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		usersMap.FlushUserInfo(c.Sender.ID)

		_, err = b.Send(c.Sender, helpers.FinalText)
		if err != nil {
			log.Println(err.Error())
		}
	})

	b.Start()
}
