package route

import (
	"ChitChat/data"
	"log"
	"net/http"
	"strconv"
)

func messageSend(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	email := request.PostFormValue("email")
	content := request.PostFormValue("content")
	recipient, err := data.UserByEmail(email)
	if err != nil {
		log.Println("Cannot get recipientId by email, or email worn:", err)
		http.Redirect(writer, request, "/messageCenter?email=不存在", 302)
	}
	cookie, _ := request.Cookie("_cookie")
	userId, _ := strconv.Atoi(cookie.Value)
	user := data.User{
		Id: userId,
	}
	_, err = user.CreateMessage(content, user.Id, recipient.Id)
	if err != nil {
		log.Println("Cannot create message:", err)
	}
	http.Redirect(writer, request, "/messageCenter", 302)
}
