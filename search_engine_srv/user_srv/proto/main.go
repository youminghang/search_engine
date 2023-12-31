package main

import (
	user "github/youminghang/search_engine/search_engine_srv/user_srv/proto/kitex_gen/user/user"
	"log"
)

func main() {
	svr := user.NewServer(new(UserImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
