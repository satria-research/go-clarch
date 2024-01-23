package ws_v2

import (
	"errors"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// hold all client registered
type Manager struct {
	sync.RWMutex
	clients   ClientList
	functions map[string]FunctionHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients:   make(ClientList),
		functions: make(map[string]FunctionHandler),
	}

	m.setupEventHandlers()

	return m
}

// register other function here...
func (m *Manager) setupEventHandlers() {
	m.functions[EventTopTipper] = SendMessageHandler
}

func (m *Manager) addClient(client *Client, wg *sync.WaitGroup) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true

	wg.Done()
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()

		delete(m.clients, client)
	}
}

func (m *Manager) routeFunction(fn Function, c *Client) error {
	if fnHandler, ok := m.functions[fn.Type]; ok {
		if err := fnHandler(fn, c); err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("error: wrong function")
	}
}

func (m *Manager) ServeWS(app *fiber.App) {

	log.Println("New connection")

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/websocket", websocket.New(func(c *websocket.Conn) {
		wg := sync.WaitGroup{}
		wg.Add(3)

		client := NewClient(c, m)
		go m.addClient(client, &wg)

		go client.readMessages()
		go client.writeMessages()

		wg.Wait()

	}))
}
