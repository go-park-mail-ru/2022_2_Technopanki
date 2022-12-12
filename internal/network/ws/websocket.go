package ws

import (
	"HeadHunter/internal/network/handlers/utils"
	"HeadHunter/internal/usecases"
	"HeadHunter/pkg/errorHandler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Pool struct {
	connections map[uint][]*websocket.Conn
	mutex       *sync.Mutex
	upgrader    websocket.Upgrader
	userUseCase usecases.User
}

func NewWSPool(userUseCase usecases.User) *Pool {
	return &Pool{
		connections: make(map[uint][]*websocket.Conn),
		mutex:       new(sync.Mutex),
		upgrader:    websocket.Upgrader{},
		userUseCase: userUseCase,
	}
}

func (p *Pool) Connect(c *gin.Context) {
	email, contextErr := utils.GetEmailFromContext(c)
	if contextErr != nil {
		_ = c.Error(contextErr)
		return
	}

	user, getId := p.userUseCase.GetUserByEmail(email)
	if getId != nil {
		_ = c.Error(getId)
		return
	}

	conn, upgradeErr := p.upgrader.Upgrade(c.Writer, c.Request, nil)
	if upgradeErr != nil {
		_ = c.Error(upgradeErr)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		err = p.Delete(user.ID, conn)
		if err != nil {
			_ = c.Error(err)
		}

	}(conn)
	p.Add(user.ID, conn)
	var readErr error
	for {
		_, _, readErr = conn.ReadMessage()
		if readErr != nil {
			if !isNormalClosure(readErr) {
				_ = c.Error(readErr)
				return
			}
			break
		}
	}
	c.Status(http.StatusOK)
}

func (p *Pool) Send(id uint, data []byte) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, conn := range p.connections[id] {
		writeErr := conn.WriteMessage(websocket.TextMessage, data)
		if writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func (p *Pool) Add(id uint, conn *websocket.Conn) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.connections[id] = append(p.connections[id], conn)
}

func (p *Pool) Delete(id uint, conn *websocket.Conn) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for i, connection := range p.connections[id] {
		if connection == conn {
			p.connections[id][i] = p.connections[id][len(p.connections[id])-1]
			p.connections[id] = p.connections[id][:len(p.connections[id])-1]
		}
	}
	return errorHandler.ErrConnectionNotFound
}

func isNormalClosure(err error) bool {
	closeErr, ok := err.(*websocket.CloseError)
	if !ok {
		return false
	}

	if closeErr.Code != websocket.CloseNormalClosure {
		return false
	}

	return true
}
