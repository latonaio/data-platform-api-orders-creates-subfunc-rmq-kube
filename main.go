package main

import (
	"context"
	api_input_reader "data-platform-api-orders-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-creates-subfunc-rmq-kube/config"
	"data-platform-api-orders-creates-subfunc-rmq-kube/database"
	"data-platform-api-orders-creates-subfunc-rmq-kube/subfunction"
	"encoding/json"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	c := config.NewConf()
	db, err := database.NewMySQL(c.DB)
	if err != nil {
		l.Error(err)
		return
	}

	rmq, err := rabbitmq.NewRabbitmqClient(c.RMQ.URL(), c.RMQ.QueueFrom(), "", c.RMQ.QueueTo(), -1)
	if err != nil {
		l.Fatal(err.Error())
	}
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()
	for msg := range iter {
		err = callProcess(ctx, db, msg, c)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
	}
}

func getSessionID(data map[string]interface{}) string {
	id := fmt.Sprintf("%v", data["runtime_session_id"])
	return id
}

func callProcess(ctx context.Context, db *database.Mysql, msg rabbitmq.RabbitmqMessage, c *config.Conf) (err error) {
	l := logger.NewLogger()
	l.AddHeaderInfo(map[string]interface{}{"runtime_session_id": getSessionID(msg.Data())})
	defer func(msg rabbitmq.RabbitmqMessage) {
		if err != nil {
			msg.Respond(map[string]interface{}{"result": "error"})
		}
	}(msg)

	subfunc := subfunction.NewSubFunction(ctx, db, l)
	sdc := api_input_reader.ConvertToSDC(msg.Raw())

	err = subfunc.Controller(&sdc)
	if err != nil {
		return err
	}
	l.Info(sdc)

	body, err := json.Marshal(sdc)
	if err != nil {
		return err
	}
	var mapData map[string]interface{}
	json.Unmarshal(body, &mapData)

	// l.Info(mapData)
	err = msg.Respond(mapData)
	if err != nil {
		return err
	}
	return nil
}
