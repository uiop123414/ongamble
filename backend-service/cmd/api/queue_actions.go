package main

import (
	"encoding/json"
	"ongambl/event"
)

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *application) logEventViaRabbit(payload LogPayload) error {
	return app.pushToQueue(payload.Name, payload.Data)
}

func (app *application) pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return err
	}

	payload := LogPayload{
		Name: name, 
		Data: msg,
	}

	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}

	return nil
}