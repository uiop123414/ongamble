package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"ongambl/internal/models"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/xeipuuv/gojsonschema"
)

const RequestSizeLimit = 2 * 1024 * 1024 // 2 kBs

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, loader gojsonschema.JSONLoader, data interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(RequestSizeLimit))
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = app.validateJSON(loader, bodyBytes)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}

func (app *application) errorJSONWithMSG(w http.ResponseWriter, err error, errors map[string]string, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()
	payload.Data = errors

	return app.writeJSON(w, statusCode, payload)
}

func (app *application) GetAuthToken(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		return "", errors.New("invalid token")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid token")
	}

	token := headerParts[1]

	return token, nil
}

func (app *application) SendMail(msg EmailPayload) error {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	mailServiceURL := "http://mail-service/send"

	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}

func (app *application) connectToRabbit() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ is not yet ready...")
			counts++
		} else {
			fmt.Println("Connected to RabbitMQ!")
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		app.logger.PrintInfo("backing off...", map[string]string{})
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}

func (app *application) validateJSON(loader gojsonschema.JSONLoader, data []byte) error {
	payloadLoader := gojsonschema.NewBytesLoader(data)

	result, err := gojsonschema.Validate(loader, payloadLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return models.ErrJSONNotValid
	}

	return nil
}

func (app *application) sendLoggedInMailViaRabbitMQ() {
	var msg EmailPayload

	msg.From = "from@email.com"
	msg.To = "to@email.com"
	msg.Body = fmt.Sprintf("You're were logged in at %v", time.Now())
	msg.Subject = "You're logged in"

	app.sendEmailViaRabbit(msg)
}
