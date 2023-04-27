package route

import (
	"ChitChat/data"
	"log"
	"net/http"
	"strconv"
)

func userLogin(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Println(err)
	}
	if user.Password == request.PostFormValue("password") {
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    strconv.Itoa(user.Id),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login/?密码错误", 302)
	}
}
func userRegister(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println("Cannot parse form" + err.Error())
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		log.Println("Cannot create user" + err.Error())
	}
	http.Redirect(writer, request, "/login", 302)
}
func userLogout(writer http.ResponseWriter, request *http.Request) {

	// !!! Important !!!
	http.SetCookie(writer, &http.Cookie{
		Name:   "_cookie",
		MaxAge: -1,
		Path:   "/",
	})
	http.Redirect(writer, request, "/", 302)
}
