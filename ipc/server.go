package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "methon"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(methon, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			requset := <- c
			if requset == "CLOSE" || requset == "close" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(requset), &req)
			if err != nil {
				fmt.Println("Invalid request format:", requset)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)

			c <- string(b)
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}

