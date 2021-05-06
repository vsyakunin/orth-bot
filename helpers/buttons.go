package helpers

import tb "gopkg.in/tucnak/telebot.v2"

var (
	InlineBtnProceedStart = tb.InlineButton{
		Unique: "proceedStart",
		Text:   "Дальше",
	}

	InlineBtnLetsStart = tb.InlineButton{
		Unique: "letsStart",
		Text:   "Начнём",
	}

	InlineBtn5min = tb.InlineButton{
		Unique: "btn5min",
		Text:   "5 минут",
	}

	InlineBtn15min = tb.InlineButton{
		Unique: "btn15min",
		Text:   "15 минут",
	}

	InlineBtn30min = tb.InlineButton{
		Unique: "btn30min",
		Text:   "30 минут",
	}

	InlineBtn1h = tb.InlineButton{
		Unique: "btn1h",
		Text:   "Час и более",
	}

	InlineBtnNextPart = tb.InlineButton{
		Unique: "nextPart",
		Text:   "Далее",
	}

	InlineBtnNextPrayer = tb.InlineButton{
		Unique: "nextPrayer",
		Text:   "Следующая молитва",
	}

	InlineBtnAmen = tb.InlineButton{
		Unique: "amen",
		Text:   "Аминь",
	}
)

func getKeysForButtons(buttons ...tb.InlineButton) (ret [][]tb.InlineButton) {
	for _, button := range buttons {
		ret = append(ret, []tb.InlineButton{button})
	}
	return
}

func MakeReplyMarkup(buttons ...tb.InlineButton) *tb.ReplyMarkup {
	return &tb.ReplyMarkup{
		InlineKeyboard: getKeysForButtons(buttons...),
	}
}