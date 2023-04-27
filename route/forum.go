package route

import (
	"ChitChat/data"
	"log"
	"net/http"
	"strconv"
)

func forumNew(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println("Cannot parse form: ", err)
	}
	cookie, _ := request.Cookie("_cookie")
	userId, _ := strconv.Atoi(cookie.Value)
	user := data.User{
		Id: userId,
	}
	title := request.PostFormValue("title")
	content := request.PostFormValue("content")
	if _, err := user.CreateForum(title, content); err != nil {
		log.Println("Cannot create forum:", err)
	}
	http.Redirect(writer, request, "/", 302)
}
