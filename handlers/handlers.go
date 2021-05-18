package handlers

import (
	"log"
	"os"
	"time"

	"github.com/vsyakunin/orth-bot/helpers"

	tb "gopkg.in/tucnak/telebot.v2"
)

type MessageHandler struct {
	Bot *tb.Bot
}

func newMessageHandler() *MessageHandler {
	bot, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	return &MessageHandler{
		Bot: bot,
	}
}

func StartHandling() {
	h := newMessageHandler()

	h.Bot.Handle(tb.OnText, h.IntroHandler())
	h.Bot.Handle(&helpers.InlineBtnProceedStart, h.ProceedStartHandler())
	h.Bot.Handle(&helpers.InlineBtnLetsStart, h.LetsStartHandler())
	h.Bot.Handle(&helpers.InlineBtn5min, h.PrayerHandler(helpers.StFiveMins))
	h.Bot.Handle(&helpers.InlineBtn15min, h.PrayerHandler(helpers.StFifteenMins))
	h.Bot.Handle(&helpers.InlineBtn30min, h.PrayerHandler(helpers.StThirtyMins))
	h.Bot.Handle(&helpers.InlineBtn1h, h.PrayerHandler(helpers.StOneHour))
	h.Bot.Handle(&helpers.InlineBtnNextPrayer, h.NextPrayerHandler())
	h.Bot.Handle(&helpers.InlineBtnNextPart, h.NextPartHandler())
	h.Bot.Handle(&helpers.InlineBtnAmen, h.AmenHandler())
	h.Bot.Start()
}

func (h MessageHandler) IntroHandler() func(*tb.Message) {
	return func(m *tb.Message) {
		userID := m.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		msg, err := h.Bot.Send(m.Sender, helpers.GetText(helpers.IntroText), helpers.MakeReplyMarkup(helpers.InlineBtnProceedStart), tb.ModeHTML)
		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo = helpers.InitialUserInfo
		userInfo.LastMsg = msg
		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (h MessageHandler) ProceedStartHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		userID := c.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = h.Bot.EditReplyMarkup(userInfo.LastMsg, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
			return
		}

		msg, err := h.Bot.Send(c.Sender, helpers.GetText(helpers.StartText), helpers.MakeReplyMarkup(helpers.InlineBtnLetsStart), tb.ModeHTML)
		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.LastMsg = msg
		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (h MessageHandler) LetsStartHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		userID := c.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = h.Bot.EditReplyMarkup(userInfo.LastMsg, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
			return
		}

		msg, err := h.Bot.Send(c.Sender, helpers.GetText(helpers.QuestionText), helpers.MakeReplyMarkup(
			helpers.InlineBtn5min, helpers.InlineBtn15min,
			helpers.InlineBtn30min, helpers.InlineBtn1h), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.LastMsg = msg
		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (h MessageHandler) PrayerHandler(handlerState helpers.State) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		userID := c.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = h.Bot.EditReplyMarkup(userInfo.LastMsg, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.UserState = handlerState
		userInfo.PrayersInState = helpers.GetPrayersInState(handlerState)
		userInfo.CurrentPrayer = helpers.GetPrayerName(userInfo)

		text, _ := helpers.GetPrayerPart(userInfo)
		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.LastMsg = msg
		userInfo.PrayerPart++

		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (h MessageHandler) NextPrayerHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		userID := c.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = h.Bot.EditReplyMarkup(userInfo.LastMsg, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.PrayerCount++
		userInfo.PrayerPart = 1
		userInfo.CurrentPrayer = helpers.GetPrayerName(userInfo)

		text, _ := helpers.GetPrayerPart(userInfo)
		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.LastMsg = msg
		userInfo.PrayerPart++

		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (h MessageHandler) NextPartHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		userID := c.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = h.Bot.EditReplyMarkup(userInfo.LastMsg, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
			return
		}

		var msg *tb.Message
		text, isLastPart := helpers.GetPrayerPart(userInfo)
		if !isLastPart {
			msg, err = h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
				helpers.InlineBtnNextPart), tb.ModeHTML)
		} else {
			if userInfo.UserState == helpers.StFiveMins ||
				userInfo.PrayerCount == userInfo.PrayersInState {
				msg, err = h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
					helpers.InlineBtnAmen), tb.ModeHTML)
			} else {
				msg, err = h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
					helpers.InlineBtnNextPrayer), tb.ModeHTML)
			}
		}

		if err != nil {
			log.Println(err.Error())
			return
		}

		userInfo.LastMsg = msg
		userInfo.PrayerPart++

		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (h MessageHandler) AmenHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		userID := c.Sender.ID
		userInfo, err := helpers.GetUserInfo(userID)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = h.Bot.EditReplyMarkup(userInfo.LastMsg, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		userInfo = helpers.InitialUserInfo

		msg, err := h.Bot.Send(c.Sender, helpers.GetText(helpers.FinalText))
		if err != nil {
			log.Println(err.Error())
		}

		userInfo.LastMsg = msg

		err = helpers.UpdateUserInfo(userID, userInfo)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}
