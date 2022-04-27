package pipelines

import (
	"app/server/database"
	"app/server/models"
)

func ChainProcess(Message *models.Message) error {
	err := WelcomeMessage(Message)
	return err
}

func WelcomeMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Step == 1 && Message.WidSender != "Numero do italo" {
		Message.ProcessedAt = true
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.Message = "Olá, Boa noite! \n \n Conhece nossas opções?"
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Step = 999
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.Step = 2
	}
	Message.ProcessedAt = true
	err := db.Save(&Message).Error
	return err
}
