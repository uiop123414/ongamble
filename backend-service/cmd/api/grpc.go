package main

import (
	"context"
	"ongambl/logs"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *application) LogViaGRPC(Name, Data string) error {
	conn, err := grpc.NewClient("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: Name,
			Data: Data,
		},
	})
	if err != nil {
		return err
	}
	app.logger.PrintInfo("Exit Log ViaGRPC", map[string]string{})
	return nil	
}