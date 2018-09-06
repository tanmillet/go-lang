package IPC

import (
	"sync"
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Server = &CenterServer{} // 确认实现了Server接口

type Message struct {
	From    string
	To      string
	Content string
}
type CenterServer struct {
	servers map[string]Server
	players []*Player
	//rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]Server)
	players := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}
func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	// 偷懒了，没做重复登录检查
	server.players = append(server.players, player)
	return nil
}
func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
				//} elseif i == len(server.players) - 1 {
				//	server.players = server.players[:i - 1]
				//} elseif i == 0 {
				//	server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[:i +
					1]...)
			}
			return nil
		}
	}
	return errors.New("Player not found.")
}
func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.RLock()
	defer server.mutex.RUnlock()
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("No player online.")
	}
	return
}

func (server *CenterServer) broadcast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("No player online.")
	}

	return err
}
func (server *CenterServer) Handle(method, params string) *Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &Response{Code: err.Error()}
		}
		return &Response{"200", players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &Response{Code: err.Error()}
		}
		return &Response{Code: "200"}
	default:
		return &Response{Code: "404", Body: method + ":" + params}
	}
	return &Response{Code: "200"}
}
func (server *CenterServer) Name() string {
	return "CenterServer"
}
