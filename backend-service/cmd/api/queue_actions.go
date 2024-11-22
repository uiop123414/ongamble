package main

import (
	"encoding/json"
	"ongambl/event"
)

type Payload struct {
	Name string       `json:"name"`
	Data string       `json:"data,omitempty"`
	Mail EmailPayload `json:"mail,omitempty"`
}

type EmailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (app *application) logEventViaRabbit(payload Payload) error {
	return app.pushToQueue(payload, "log")
}

func (app *application) sendEmailViaRabbit(e EmailPayload) error {
	var payload Payload
	payload.Name = "email.task"
	payload.Mail = e
	return app.pushToQueue(payload, "email.task")
}

func (app *application) pushToQueue(payload interface{}, severity string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return err
	}

	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = emitter.Push(string(j), severity)
	if err != nil {
		return err
	}

	return nil
}
