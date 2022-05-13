package pipelines

import (
	"app/server/database"
	"app/server/models"
)

func SelectOptionResponse(Message *models.Message) error {
	var Session models.Session
	db := database.GetDatabase()
	db.Last(&Session)
	Template, err := Session.GetActualTemplate(db)
	newMessage := models.Message{Message: Template.GetMessage(),
		ProcessedAt: true}
	db.Create(&newMessage)
	return err
}

func ChainProcess(Message *models.Message) error {
	err := SelectOptionResponse(Message)
	Message.ProcessedAt = true
	db := database.GetDatabase()
	db.Save(&Message)
	return err
}

func WelcomeMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Step == 0 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.Message = "Olá, Boa noite! \n \n Conhece nossas opções? \n \n 1 - Função A \n \n 2 - Função B \n \n 3 - Função C"
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Step = 999
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.Step = 1
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func SecondLayerMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Step == 1 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Step = 999
		switch Message.Message {
		case "1":
			AnswerMessage.Message = "Função A"
		case "2":
			AnswerMessage.Message = "Função B"
		case "3":
			AnswerMessage.Message = "Função C"
		default:
			return DefaultMessage(Message)
		}
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.Step = 2
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func ThirdLayerMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if Message.ProcessedAt && Message.Step == 2 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Step = 999
		AnswerMessage.Message = "O que mais posso estar fazendo por você? \n \n Digite uma das opções abaixo: \n \n 1 - Desejo voltar para o menu inicial. \n \n 2 - Desejo Encerrar o atendimento."
		Message.Step = 3
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func FourthLayerMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Step == 3 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Step = 999
		switch Message.Message {
		case "1":
			Message.Step = 0
			Message.ProcessedAt = false
			return ChainProcess(Message)
		case "2":
			Message.Step = 0
			AnswerMessage.Message = "Obrigado pela preferência!"
		default:
			return DefaultMessage(Message)
		}
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func DefaultMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	AnswerMessage.WidSender = "Numero do Ítalo"
	AnswerMessage.WidReceiver = Message.WidSender
	AnswerMessage.ProcessedAt = true
	AnswerMessage.Step = 999
	if !Message.ProcessedAt && Message.Step == 1 && Message.WidSender != "Numero do italo" {
		AnswerMessage.Message = "Não Entendi :( \n \n Por favor digite uma das opções abaixo: \n \n 1 - Função A. \n \n 2 - Função B. \n \n 3 - Função C."
	} else if !Message.ProcessedAt && Message.Step == 3 {
		AnswerMessage.Message = "Não Entendi :( \n \n Por favor digite uma das opções abaixo: \n \n 1 - Desejo voltar para o menu inicial. \n \n 2 - Desejo Encerrar o atendimento."
	}
	err := db.Save(&AnswerMessage).Error
	if err != nil {
		return err
	}
	Message.ProcessedAt = true
	err = db.Save(&Message).Error
	return err
}
