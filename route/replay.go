package route

import (
	"ChitChat/data"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func replayNew(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println("Cannot parse form:", err)
	}
	content := request.PostFormValue("content")
	forumId := request.PostFormValue("forum_id")
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	userId, _ := strconv.Atoi(cookie.Value)
	user := data.User{
		Id: userId,
	}
	user.ReflectDetail()
	forum, err := data.ForumById(forumId)
	if err != nil {
		log.Println("Cannot read forum: ", err)
	}
	if _, err := user.CreatePost(forum, content); err != nil {
		log.Println("Cannot create post: ", err)
	}
	url := fmt.Sprint("/forumDetail?forum_id=", forumId)
	http.Redirect(writer, request, url, 302)
}
