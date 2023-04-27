package main

import (
	"ChitChat/route"
	"ChitChat/utils"
	"net/http"
)

func main() {
	utils.LoadConfig()

	//// --- Controller ---
	//// define in controller/users.go
	//mux.HandleFunc("/register", controller.RegisterHandler)
	//mux.HandleFunc("/login", controller.LoginHandler)
	//mux.HandleFunc("/logout", controller.LogoutHandler)
	//
	//// define in controller/forum.go
	//mux.HandleFunc("/forumCreate", controller.ForumCreateHandler)
	//
	//// define in controller/replay.go
	//mux.HandleFunc("/replay", controller.ReplyHandler)
	//
	//// define in controller/user.go
	//mux.HandleFunc("/messageSend", controller.MessageSendHandler)

	// --- Controller End ---

	server := http.Server{
		Addr:    utils.Config.Address,
		Handler: route.Mux,
	}
	server.ListenAndServe()
}
