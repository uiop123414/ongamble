package handlers

import (
	"chatgpt-article/internal/models"
	"chatgpt-article/internal/schemas"
	"chatgpt-article/internal/utils"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleMessage(message []byte, offset kafka.Offset) error {
	caap := models.CreateAiArticlePayload

	err := utils.ReadJSON(message, schemas.CreateAiArticleLoader, &caap)
	if err != nil {
		logrus.Error(err)
		return err
	}
	
	logrus.Info("Message caap = ", caap, ",message = ", string(message))

	switch caap.Type{

	case "create-chatgpt-article":

	}

	
	return nil
}

func (h *Handler) createChatgptArticle(caap interface{}) {

}