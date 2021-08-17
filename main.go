package main

import "github.com/han-joker/moo-layout/moo"

func main() {
	//启动 websocket服务
	moo.Server("websocket").Start()
}