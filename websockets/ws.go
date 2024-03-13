package websockets

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"project-websocket/controllers"
	"project-websocket/models"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Aceptar todas las solicitudes
		},
	}

	// Guarda las conexiones activas
	activeConnections = make(map[*websocket.Conn]bool)

	// Mutex para manejar el acceso concurrente a activeConnections
	connMutex = &sync.Mutex{}
)

type UserMessage struct {
	Action string       `json:"action"`
	User   *models.User  `json:"user,omitempty"`
	Users  []*models.User `json:"users,omitempty"`
}

func broadcastMessage(message UserMessage) {
	connMutex.Lock()
	defer connMutex.Unlock()

	for conn := range activeConnections {
		err := conn.WriteJSON(message)
		if err != nil {
			log.Println(err)
			delete(activeConnections, conn)
		}
	}
}

func handleUserAction(action string, user *models.User, conn *websocket.Conn) {
	var err error
	var result *models.User

	switch action {
	case "createUser":
		result, err = controllers.CreateUser(user)
	case "updateUser":
		result, err = controllers.UpdateUser(user.Id.Hex(), user)
	case "deleteUser":
		result, err = controllers.DeleteUser(user.Id.Hex())
	}

	if err != nil {
		log.Println(err)
		return
	}

	broadcastMessage(UserMessage{
		Action: action,
		User:   result,
	})
	conn.WriteJSON(result)
}

// WSHandler maneja las conexiones WebSocket
func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Agrega la conexión al mapa de conexiones activas
	connMutex.Lock()
	activeConnections[conn] = true
	connMutex.Unlock()

	defer func() {
		// Cuando la conexión se cierra, elimina la conexión del mapa de conexiones activas
		connMutex.Lock()
		delete(activeConnections, conn)
		connMutex.Unlock()

		conn.Close()
	}()

	fmt.Printf("Client connected. Active connections: %d\n", len(activeConnections))

	for {
		var req UserMessage
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Println(err)
			break
		}

		switch req.Action {
		case "getUsers":
			users, err := controllers.GetAllUsers()
			if err != nil {
				log.Println(err)
				break
			}
			// Enviar un objeto con la propiedad 'action' y la propiedad 'users'
			conn.WriteJSON(UserMessage{
				Action: "getUsers",
				Users:  users,
			})
		case "getUser":
			user, err := controllers.GetUser(req.User.Id.Hex()) // Asume que ID es el identificador del usuario
			if err != nil {
				log.Println(err)
				break
			}
			conn.WriteJSON(user)
		case "createUser", "updateUser", "deleteUser":
			handleUserAction(req.Action, req.User, conn)
		}
	}
}