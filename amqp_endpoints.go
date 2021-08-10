package lesson20acl

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	common_lib21 "github.com/kirigaikabuto/common-lib21"
)

type RoleAmqpEndpoints struct {
	ch common_lib21.CommandHandler
}

func (r *RoleAmqpEndpoints) MakeCreateRoleAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateRoleCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			panic(err)
			return nil
		}
		response, err := r.ch.ExecCommand(cmd)
		if err != nil {
			panic(err)
			return nil
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			panic(err)
			return nil
		}
		return &amqp.Message{Body: jsonData}
	}
}
