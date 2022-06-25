package websocket

// Connection is a middleman between the websocket Connection and the Hub.

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan message),
		register:   make(chan subscription),
		unregister: make(chan subscription),
		rooms:      make(map[string]map[*Connection]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room.Id]
			if connections == nil {
				connections = make(map[*Connection]bool)
				h.rooms[s.room.Id] = connections
			}
			h.rooms[s.room.Id][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room.Id]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room.Id)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room.Id]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room.Id)
					}
				}
			}
		}
	}
}
