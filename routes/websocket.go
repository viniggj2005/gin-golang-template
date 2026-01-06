package routes

import (
	"fmt"
	"html"

	"github.com/olahol/melody"
)

func newWebsocketServer() *melody.Melody {
	websocket := melody.New()
	websocket.Config.MaxMessageSize = 1024 * 1024
	return websocket
}

func handleWebsocketconnection(websocket *melody.Melody) {
	// 1. Ao conectar, lemos os parâmetros da URL (sala e tópico) e salvamos na sessão
	websocket.HandleConnect(func(session *melody.Session) {
		room := session.Request.URL.Query().Get("room")
		topic := session.Request.URL.Query().Get("topic")

		// Armazena na sessão do usuário atual
		session.Set("room", room)
		session.Set("topic", topic)

		fmt.Printf("Cliente conectado na Sala: %s | Tópico: %s\n", room, topic)
	})

	// 2. Ao receber mensagem, fazemos o Broadcast com Filtro
	websocket.HandleMessage(func(session *melody.Session, msg []byte) {
		// Recupera onde o remetente está
		myRoom, _ := session.Get("room")
		myTopic, _ := session.Get("topic")

		// Filtro: A mensagem será enviada para 'q' (destinatário) SE retornar true
		securityMessage := html.EscapeString(string(msg))
		websocket.BroadcastFilter([]byte(securityMessage), func(query *melody.Session) bool {
			qRoom, _ := query.Get("room")
			qTopic, _ := query.Get("topic")

			// A lógica hierárquica:
			// O destinatário deve estar na MESMA sala E no MESMO tópico
			return qRoom == myRoom && qTopic == myTopic //cenário com echo
			// cenário sem echo
			// return qRoom == myRoom && qTopic == myTopic && query != session cenário sem echo
		})
	})
}
