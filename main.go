package main

import (
	"log"
	"time"

	"orth-bot/helpers"

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

		usersMap.UpdatePrayer(c.Sender.ID, helpers.Get5MinPrayerName())
		text, _ := helpers.GetPrayer(usersMap.GetUserInfo(c.Sender.ID))

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

		text, isLastPart := helpers.GetPrayer(usersMap.GetUserInfo(c.Sender.ID))
		if !isLastPart {
			_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
				helpers.InlineBtnNextPart), tb.ModeHTML)
		} else {
			_, err = b.Send(c.Sender, text, helpers.MakeReplyMarkup(
				helpers.InlineBtnAmen), tb.ModeHTML)
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

		_, err = b.Send(c.Sender, helpers.IntroText, helpers.MakeReplyMarkup(helpers.InlineBtnProceedStart))
		if err != nil {
			log.Println(err.Error())
		}
	})

	b.Start()
}
