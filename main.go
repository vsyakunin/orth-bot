package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	inlineBtnProceedStart = tb.InlineButton{
		Unique: "proceedStart",
		Text:   "Дальше",
	}

	inlineBtnLetsStart = tb.InlineButton{
		Unique: "letsStart",
		Text:   "Начнём",
	}

	inlineBtn5min = tb.InlineButton{
		Unique: "btn5min",
		Text:   "5 минут",
	}

	inlineBtn15min = tb.InlineButton{
		Unique: "btn15min",
		Text:   "15 минут",
	}

	inlineBtn30min = tb.InlineButton{
		Unique: "btn30min",
		Text:   "30 минут",
	}

	inlineBtn1h = tb.InlineButton{
		Unique: "btn1h",
		Text:   "Час и более",
	}
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

	inlineKeys := [][]tb.InlineButton{
		{inlineBtnProceedStart},
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, `Господи, прости!

Ты столкнулся с проблемой зависимости от порнографии? Твоё сердце разъедает чувство вины за совершённое блудодеяние, гложет совесть? 

Этот чат-бот подскажет тебе, как правильно обратиться к Господу, чтобы он помог отпустить этот грех.`,
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys, OneTimeKeyboard: true})
	})

	inlineKeysLetsStart := [][]tb.InlineButton{
		{inlineBtnLetsStart},
	}

	b.Handle(&inlineBtnProceedStart, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{ShowAlert: false})

		b.Send(c.Sender, `Конечно, лучше всего отправиться на исповедь в церковь. Но, если такой возможности нет, или ты пока не решаешься, пусть этот бот будет твоим помощником. 

Проект разработан SMIT.Studio совместно с православным иереем Константином Мальцевым. Отец Константин лично отобрал подходящие молитвы и благословил проект.`,
			&tb.ReplyMarkup{InlineKeyboard: inlineKeysLetsStart, OneTimeKeyboard: true})
	})

	timeKeys := [][]tb.InlineButton{
		{inlineBtn5min},
		{inlineBtn15min},
		{inlineBtn30min},
		{inlineBtn1h},
	}

	b.Handle(&inlineBtnLetsStart, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{ShowAlert: false})

		b.Send(c.Sender, `Итак, ты поддался искушению и потратил своё время на просмотр порнографии. Возможно, ты чувствуешь вину, не знаешь, как облегчить свою душу и очистить разум? 

Мы предлагаем тебе уделить для молитвы столько же времени, сколько ты потратил на просмотр порнографии.

Просто ответь на этот простой вопрос

Сколько времени ты потратил сейчас на порно? 
`,
			&tb.ReplyMarkup{InlineKeyboard: timeKeys, OneTimeKeyboard: true})
	})

	b.Start()
}
