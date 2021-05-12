package handlers

import (
	"log"
	"os"
	"time"

	"github.com/vsyakunin/orth-bot/helpers"

	tb "gopkg.in/tucnak/telebot.v2"
)

type MessageHandler struct {
	Bot      *tb.Bot
	UsersMap helpers.UsersMap
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
		Bot:      bot,
		UsersMap: make(helpers.UsersMap),
	}
}

func StartHandling() {
	h := newMessageHandler()

	h.Bot.Handle(tb.OnText, h.IntroHandler())
	h.Bot.Handle(&helpers.InlineBtnProceedStart, h.ProceedStartHandler())
	h.Bot.Handle(&helpers.InlineBtnLetsStart, h.LetsStartHandler())
	h.Bot.Handle(&helpers.InlineBtn5min, h.FiveMinHandler())
	h.Bot.Handle(&helpers.InlineBtn15min, h.FifteenMinHandler())
	h.Bot.Handle(&helpers.InlineBtn30min, h.ThirtyMinHandler())
	h.Bot.Handle(&helpers.InlineBtn1h, h.OneHourHandler())
	h.Bot.Handle(&helpers.InlineBtnNextPrayer, h.NextPrayerHandler())
	h.Bot.Handle(&helpers.InlineBtnNextPart, h.NextPartHandler())
	h.Bot.Handle(&helpers.InlineBtnAmen, h.AmenHandler())
	h.Bot.Start()
}

func (h MessageHandler) IntroHandler() func(*tb.Message) {
	return func(m *tb.Message) {
		if _, ok := h.UsersMap[m.Sender.ID]; !ok {
			h.UsersMap.AddUser(m.Sender.ID)
			log.Printf("added new user with ID %v", m.Sender.ID)
		}
		h.UsersMap.FlushUserInfo(m.Sender.ID)

		msg, err := h.Bot.Send(m.Sender, helpers.GetText(helpers.IntroText), helpers.MakeReplyMarkup(helpers.InlineBtnProceedStart), tb.ModeHTML)
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(m.Sender.ID, msg)
	}
}

func (h MessageHandler) ProceedStartHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		msg, err := h.Bot.Send(c.Sender, helpers.GetText(helpers.StartText), helpers.MakeReplyMarkup(helpers.InlineBtnLetsStart), tb.ModeHTML)
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
	}
}

func (h MessageHandler) LetsStartHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		msg, err := h.Bot.Send(c.Sender, helpers.GetText(helpers.QuestionText), helpers.MakeReplyMarkup(
			helpers.InlineBtn5min, helpers.InlineBtn15min,
			helpers.InlineBtn30min, helpers.InlineBtn1h), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
	}
}

func (h MessageHandler) FiveMinHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateState(c.Sender.ID, helpers.StFiveMins)
		h.UsersMap.UpdatePrayer(c.Sender.ID)

		userInfo := h.UsersMap.GetUserInfo(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(userInfo)

		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
		h.UsersMap.UpdatePrayerPart(c.Sender.ID)
	}
}

func (h MessageHandler) FifteenMinHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateState(c.Sender.ID, helpers.StFifteenMins)
		h.UsersMap.UpdatePrayer(c.Sender.ID)

		userInfo := h.UsersMap.GetUserInfo(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(userInfo)

		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
		h.UsersMap.UpdatePrayerPart(c.Sender.ID)
	}
}

func (h MessageHandler) ThirtyMinHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateState(c.Sender.ID, helpers.StThirtyMins)
		h.UsersMap.UpdatePrayer(c.Sender.ID)

		userInfo := h.UsersMap.GetUserInfo(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(userInfo)

		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
		h.UsersMap.UpdatePrayerPart(c.Sender.ID)
	}
}

func (h MessageHandler) OneHourHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateState(c.Sender.ID, helpers.StOneHour)
		h.UsersMap.UpdatePrayer(c.Sender.ID)

		userInfo := h.UsersMap.GetUserInfo(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(userInfo)

		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
		h.UsersMap.UpdatePrayerPart(c.Sender.ID)
	}
}

func (h MessageHandler) NextPrayerHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdatePrayerCount(c.Sender.ID)
		h.UsersMap.UpdatePrayer(c.Sender.ID)

		userInfo := h.UsersMap.GetUserInfo(c.Sender.ID)
		text, _ := helpers.GetPrayerPart(userInfo)

		icon := helpers.GetIconForPrayer(userInfo.CurrentPrayer)
		if icon != nil {
			_, err = h.Bot.Send(c.Sender, icon)
		}

		msg, err := h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
			helpers.InlineBtnNextPart), tb.ModeHTML)

		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
		h.UsersMap.UpdatePrayerPart(c.Sender.ID)
	}
}

func (h MessageHandler) NextPartHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		var msg *tb.Message
		text, isLastPart := helpers.GetPrayerPart(h.UsersMap.GetUserInfo(c.Sender.ID))
		if !isLastPart {
			msg, err = h.Bot.Send(c.Sender, text, helpers.MakeReplyMarkup(
				helpers.InlineBtnNextPart), tb.ModeHTML)
		} else {
			userInfo := h.UsersMap.GetUserInfo(c.Sender.ID)
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
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
		h.UsersMap.UpdatePrayerPart(c.Sender.ID)
	}
}

func (h MessageHandler) AmenHandler() func(c *tb.Callback) {
	return func(c *tb.Callback) {
		_, err := h.Bot.EditReplyMarkup(h.UsersMap.GetLastMsg(c.Sender.ID), nil)
		if err != nil {
			log.Println(err.Error())
		}

		err = h.Bot.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.FlushUserInfo(c.Sender.ID)

		msg, err := h.Bot.Send(c.Sender, helpers.GetText(helpers.FinalText))
		if err != nil {
			log.Println(err.Error())
		}

		h.UsersMap.UpdateLastMsg(c.Sender.ID, msg)
	}
}

