package helpers

const (
	ErrorText    = "error_text"
	IntroText    = "intro_text"
	StartText    = "start_text"
	QuestionText = "question_text"
	FinalText    = "final_text"
	StatsText    = "stats_text"
)

func GetText(textName string) (text string) {
	prayer, err := getPrayerByName(textName)
	if err != nil {
		return ""
	}

	prayerPart, _, err := prayer.getPart(1)
	if err != nil {
		return ""
	}

	return prayerPart
}
