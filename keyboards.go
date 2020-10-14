package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

//  use to do
func MainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Про ОСББ"),
			tgbotapi.NewKeyboardButton("Правила"),
			tgbotapi.NewKeyboardButton("Інструкції"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Зареєструвати запит"),
		),
	)
}

// ruleKeyboard use to do
func RuleKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Територія двору"),
			tgbotapi.NewKeyboardButton("Кондиціонери"),
			tgbotapi.NewKeyboardButton("Платіжна дисципліна"),
		),
	)
}

// instructionKeyboard use to do
func InstructionKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Затоплення"),
			tgbotapi.NewKeyboardButton("Витоку газу"),
			tgbotapi.NewKeyboardButton("Каналізація"),
		),
	)
}

//contactKeyboard use to do
func RequestKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Електрика"),
			tgbotapi.NewKeyboardButton("Водопостачання"),
			tgbotapi.NewKeyboardButton("Загальні питання"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Прибирання під'їздів"),
			tgbotapi.NewKeyboardButton("Чистота двору"),
		),
	)
}
