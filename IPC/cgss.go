package IPC

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

var centerClient *CenterClient

func startCenterService() error {
	server := NewIpcServer(&CenterServer{})
	client := NewIpcClient(server)
	centerClient = &CenterClient{client}
	return nil
}

func Help(args []string) int {
	fmt.Println(`
 Commands:
 login <username><level><exp>
 logout <username>
 send <message>
 listplayer
 quit(q)
 help(h)
 `)
	return 0
}
func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
	return 0
}
func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username><level><exp>")
		return 0
	}
	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid Parameter: <level> should be an integer.")
		return 0
	}
	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Invalid Parameter: <exp> should be an integer.")
		return 0
	}
	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp
	err = centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("Failed adding player", err)
	}
	return 0
}

func ListPlayer(args []string) int {
	ps, err := centerClient.ListPlayer("")
	if err != nil {
		fmt.Println("Failed. ", err)
	} else {
		for i, v := range ps {
			fmt.Println(i+1, ":", v)
		}
	}
	return 0
}
func Send(args []string) int {
	message := strings.Join(args[1:], " ")
	err := centerClient.Broadcast(message)
	if err != nil {
		fmt.Println("Failed.", err)
	}
	return 0
}

// 将命令和处理函数对应
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":       Help,
		"h":          Help,
		"quit":       Quit,
		"q":          Quit,
		"login":      Login,
		"logout":     Logout,
		"listplayer": ListPlayer,
		"send":       Send,
	}
}

func main() {
	fmt.Println("Casual Game Server Solution")
	startCenterService()
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for { // 循环读取用户输入
		fmt.Print("Command> ")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command:", tokens[0])
		}
	}
}

//$ go run cgss.go
//Casual Game Server Solution
//A new session has been created successfully.
//Commands:
//login <username><level><exp>
//logout <username>
//send <message>
//listplayer
//quit(q)
//help(h)
//Command> login Tom 1 101
//Command> login Jerry 2 321
//Command> listplayer
//1 : &{Tom 1 101 0 <nil>}
//2 : &{Jerry 2 321 0 <nil>}
//Command> send Hello everybody.
//Tom received message: Hello everybody.
//Jerry received message: Hello everybody.
//Command> logout Tom
//Command> listplayer
//1 : &{Jerry 2 321 0 <nil>}
//Command> send Hello the people online.
//Jerry received message: Hello the people online.
//Command> logout Jerry
//Command> listplayer
//Failed. No player online.
//Command> q
//$
