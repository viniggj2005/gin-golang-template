package routes

import (
	"log"
	"sync"

	socketio "github.com/doquangtan/socketio/v4"
)

var sockets sync.Map

func NewSocketIOServer() *socketio.Io {
	io := socketio.New()

	io.OnConnection(func(s *socketio.Socket) {
		sockets.Store(s.Id, s)
		log.Println("conectado:", s.Id)

		s.Emit("message", "Bem-vindo")
		s.On("chat message", func(ev *socketio.EventPayload) {
			s.Emit("message", "olá ")
		})

		s.On("disconnect", func(ev *socketio.EventPayload) {
			log.Println("disconnect:", ev.SID)
		})
	})

	return io
}

func EmitSinal(p string) {
	sockets.Range(func(_ interface{}, v interface{}) bool {
		if s, ok := v.(*socketio.Socket); ok {
			s.Emit("MESSAGEMAIN", p)
		}
		return true
	})
}
