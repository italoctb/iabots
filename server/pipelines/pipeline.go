package pipelines

import (
	"app/server/database"
	"app/server/models"
)

func ChainProcess(Message *models.Message) error {
	err := WelcomeMessage(Message)
	if err != nil {
		return err
	}
	err = FirstLayerMessage(Message)
	if err != nil {
		return err
	}
	err = FirstLayerAnswerMessage(Message)
	if err != nil {
		return err
	}
	err = GoodbyeMessage(Message)
	if err != nil {
		return err
	}
	err = GoodbyeAnswerMessage(Message)
	return err
}

func WelcomeMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Path == 0 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.Message = "Olá, Boa noite! \n \n Conhece nossas opções? \n \n 1 - Função A \n \n 2 - Função B \n \n 3 - Função C"
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Path = 999
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.Path = 1
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func FirstLayerMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Path == 0 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.Message = "Olá, Boa noite! \n \n Conhece nossas opções? \n \n 1 - Função A \n \n 2 - Função B \n \n 3 - Função C"
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Path = 999
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.Path = 1
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func FirstLayerAnswerMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Path == 1 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Path = 999
		switch Message.Message {
		case "1":
			AnswerMessage.Message = "Função A"
		case "2":
			AnswerMessage.Message = "Função B"
		case "3":
			AnswerMessage.Message = "Função C"
		default:
			return FirstLayerDefaultMessage(Message)
		}
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.Path = 2
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func GoodbyeMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if Message.ProcessedAt && Message.Path == 2 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Path = 999
		AnswerMessage.Message = "O que mais posso estar fazendo por você? \n \n Digite uma das opções abaixo: \n \n 1 - Desejo voltar para o menu inicial. \n \n 2 - Desejo Encerrar o atendimento."
		Message.Path = 3
		err := db.Save(&AnswerMessage).Error
		if err != nil {
			return err
		}
		Message.ProcessedAt = true
	}
	err := db.Save(&Message).Error
	return err
}

func GoodbyeAnswerMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	if !Message.ProcessedAt && Message.Path == 3 && Message.WidSender != "Numero do italo" {
		AnswerMessage.WidSender = "Numero do Ítalo"
		AnswerMessage.WidReceiver = Message.WidSender
		AnswerMessage.ProcessedAt = true
		AnswerMessage.Path = 999
		switch Message.Message {
		case "1":
			Message.Path = 0
			Message.ProcessedAt = false
			return ChainProcess(Message)
		case "2":
			Message.Path = 0
			AnswerMessage.Message = "Obrigado pela preferência!"
		default:
			return DefaultGoodbyeMessage(Message)
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

func FirstLayerDefaultMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	AnswerMessage.WidSender = "Numero do Ítalo"
	AnswerMessage.WidReceiver = Message.WidSender
	AnswerMessage.ProcessedAt = true
	AnswerMessage.Path = 999
	if !Message.ProcessedAt && Message.Path == 1 && Message.WidSender != "Numero do italo" {
		AnswerMessage.Message = "Não Entendi :( \n \n Por favor digite uma das opções abaixo: \n \n 1 - Função A. \n \n 2 - Função B. \n \n 3 - Função C."
	}
	err := db.Save(&AnswerMessage).Error
	if err != nil {
		return err
	}
	Message.ProcessedAt = true
	err = db.Save(&Message).Error
	return err
}

func DefaultGoodbyeMessage(Message *models.Message) error {
	var AnswerMessage models.Message
	db := database.GetDatabase()
	AnswerMessage.WidSender = "Numero do Ítalo"
	AnswerMessage.WidReceiver = Message.WidSender
	AnswerMessage.ProcessedAt = true
	AnswerMessage.Path = 999
	if !Message.ProcessedAt && Message.Path == 3 {
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
