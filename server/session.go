// global
package server

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"model"
)

type LoginResult struct {
	Success	bool	`json:"success"`
	Message	string	`json:"message"`
	Token	string	`json:"token"`
	Exp	int64	`json:"exp"`
}

func (this Remote) SessionInsertion(args []byte, result *LoginResult) error {
	err := json.Unmarshal(args, result)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if !result.Success {
		err := &sError{1, "login failed"}
		return err
	}

	message := strings.Split(result.Message, " ")

	token := &model.Token {
		message[1][1:len(message[1]) - 1],
		0,
	}
	log.Println(token)
	model.Logined[result.Token] = token

	go token.Count(result.Token, result.Exp - time.Now().Unix())

	return nil
}

func (this Remote) SessionValidation(args map[string][]string, result *bool) error {
	authorization, ok := args["Authorization"]
	if !ok {
		log.Println("Authorization required")
		return nil
	}

	username, ok := args["Username"]
	if !ok {
		log.Println("Username required")
		return nil
	}

	log.Println(authorization[0], username[0])

	log.Println(authorization[0])
	log.Println(username[0])
	log.Println(model.Logined[authorization[0]])
	_, ok = model.Logined[authorization[0]]
	if !ok || model.Logined[authorization[0]].U != username[0] {
		return nil
	}

	*result = true
	return nil
}
