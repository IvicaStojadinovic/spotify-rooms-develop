package controller

import (
	"github.com/timfuhrmann/spotify-rooms/backend/conn"
	"github.com/timfuhrmann/spotify-rooms/backend/entity"
)

func UpdateRooms(ws *conn.WebSocket, event *entity.Event) {
	ws.Out <- (event).Raw()
}