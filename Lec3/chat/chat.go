package chat

import (
	context "context"
	"log"
)

//Server ...
type Server struct {
}

//SayHello ...
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello from this server!"}, nil
}

//SayGoodBye ...
func (s *Server) SayGoodBye(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Recive message body from client: %s", in.Body)
	return &Message{Body: "GoodBye from this server!"}, nil
}
