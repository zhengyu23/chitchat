package route

import (
	data "ChitChat/data"
	utils "ChitChat/utils"
	"log"
	"net/http"
	"strconv"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()
	// HTML
	Mux.HandleFunc("/", index)
	Mux.HandleFunc("/login", login)
	Mux.HandleFunc("/register", register)
	Mux.HandleFunc("/userCenter", userCenter)
	Mux.HandleFunc("/messageCenter", messageCenter)
	Mux.HandleFunc("/forumDetail", forumDetail)
	Mux.HandleFunc("/newForum", newForum)

	// FUNCTION
	Mux.HandleFunc("/user/login", userLogin)
	Mux.HandleFunc("/user/register", userRegister)
	Mux.HandleFunc("/user/logout", userLogout)
	Mux.HandleFunc("/forum/new", forumNew)
	Mux.HandleFunc("/replay/new", replayNew)
	Mux.HandleFunc("/message/send", messageSend)

	// FUNCTION-GET
	// 获取用户信息 - user.Detail(userId)(User)
	// 获取帖子列表 - ForumList()([]Forum)
	// 获取帖子信息 - forum.forumDetail()(Forum)
	// 获取回帖列表 - forum.ReplayList()([]Replay)
	// 获取私信列表 - user.messageList()([]Message)

}

func index(writer http.ResponseWriter, request *http.Request) {
	// index包括：导航栏+获取帖子列表
	// 获取帖子列表
	forums, err := data.Forums()
	if err != nil {
		log.Println("Cannot get forumList:", err)
	}
	// 先判断登录状态
	_, err = request.Cookie("_cookie")
	if err == nil { // err为空，说明有user_id在cookie
		utils.GenerateHTML(writer, forums, "layout", "auth/private/nav", "auth/private/forum")
	} else {
		utils.GenerateHTML(writer, forums, "layout", "auth/public/nav", "auth/public/forum")
	}
}
func login(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, nil, "layout", "login")
}
func register(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, nil, "layout", "register")
}

// auth-user
func userCenter(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	userId, _ := strconv.Atoi(cookie.Value)
	user := data.User{
		Id: userId,
	}
	err = user.ReflectDetail()
	if err != nil {
		log.Println(err)
	}
	utils.GenerateHTML(writer, &user, "layout", "auth/private/nav", "auth/private/user")
}

// auth-user
func messageCenter(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	userId, _ := strconv.Atoi(cookie.Value)
	user := data.User{
		Id: userId,
	}
	messages, err := user.Messages()
	if err != nil {
		log.Println("Cannot get messageList:", err)
	}
	// 先判断登录状态
	_, err = request.Cookie("_cookie")

	if err == nil { // err为空，说明有user_id在cookie
		if messages == nil {
			utils.GenerateHTML(writer, nil, "layout", "auth/private/nav", "auth/private/message")
		} else {
			utils.GenerateHTML(writer, &messages, "layout", "auth/private/nav", "auth/private/message")
		}
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

func forumDetail(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	forumId := values.Get("forum_id")
	forum, err := data.ForumById(forumId)
	if err != nil {
		log.Println("Cannot read forum: " + err.Error())
	} else {
		_, err := request.Cookie("_cookie")
		if err != nil {
			utils.GenerateHTML(writer, &forum, "layout", "/auth/public/nav", "/auth/public/replay")
		} else {
			utils.GenerateHTML(writer, &forum, "layout", "/auth/private/nav", "/auth/private/replay")
		}
	}
}

func newForum(writer http.ResponseWriter, request *http.Request) {
	// 先判断登录状态
	_, err := request.Cookie("_cookie")
	if err == nil { // err为空，说明有user_id在cookie
		utils.GenerateHTML(writer, nil, "layout", "auth/private/nav", "auth/private/forumNew")
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
