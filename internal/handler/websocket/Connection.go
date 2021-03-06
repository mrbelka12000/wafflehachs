package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wafflehacks/entities/response"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func (h *Handler) GetConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomId, ok := vars["id"]
	if !ok {
		response.SendErrorResponse(w, "Не удалось найти айди комнаты", 400)
		h.log.Debug("Не удалось найти айди комнаты")
		return
	}

	if conns, ok := h.Hub.rooms[roomId]; ok {
		if len(conns) == 2 {
			response.SendErrorResponse(w, "Комната занята", http.StatusForbidden)
			h.log.Debug("Комната занята")
			return
		}
	}
	h.serveWs(w, r, roomId)
}

// serveWs handles websocket requests from the peer.
func (h *Handler) serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	ws, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		response.SendErrorResponse(w, "не удалось обновить соединение", 500)
		h.log.Debug(err.Error())
		return
	}

	_, msg, err := ws.ReadMessage()
	if err != nil {
		response.SendErrorResponse(w, "не удалось прочитать сообщение", 400)
		h.log.Debug("Не удалось прочитать сообщение: " + err.Error())
		return
	}

	fmt.Println(string(msg))
	c := &Connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{conn: c}
	h.Hub.register <- s
	go s.writePump()
	go s.readPump(h)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump(h *Handler) {
	c := s.conn
	defer func() {
		h.Hub.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := message{msg, s.room}
		h.Hub.broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}
