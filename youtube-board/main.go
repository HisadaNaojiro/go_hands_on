package main

import (
	"./my"
	"github.com/gorilla/sessions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var dbDriver = "postgres"
var dbName = "postgres"
var dsn = "host=localhost dbname=postgres sslmode=disable"

var sesName = "ytboard-session"
var cs = sessions.NewCookieStore([]byte("secret-key-1234"))

func checkLogin(w http.ResponseWriter, rq *http.Request) *my.User {
	ses, _ := cs.Get(rq, sesName)
	if ses.Values["login"] == nil || !ses.Values["login"].(bool) {
		http.Redirect(w, rq, "/login", 302)
	}

	ac := ""
	if ses.Values["account"] != nil {
		ac = ses.Values["account"].(string)
	}

	var user my.User

	v2Db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, _ := v2Db.DB()
	defer db.Close()

	v2Db.Where("account = ?", ac).Find(&user)

	return &user
}

func notemp() *template.Template {
	tmp, _ := template.New("index").Parse("NO PAGE.")
	return tmp
}

func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles(
		"youtube-board/templates/"+fname+".html",
		"youtube-board/templates/head.html",
		"youtube-board/templates/foot.html",
	)
	return tmps
}

func index(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)

	v2Db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, _ := v2Db.DB()
	defer db.Close()

	var pl []my.Post
	v2Db.Where("group_id > 0").Order("created_at desc").Limit(10).Find(&pl)
	var gl []my.Group
	v2Db.Order("created_at desc").Limit(10).Find(&gl)

	item := struct {
		Title   string
		Message string
		Name    string
		Account string
		Plist   []my.Post
		Glist   []my.Group
	}{
		Title:   "index",
		Message: "This is Top page.",
		Account: user.Account,
		Plist:   pl,
		Glist:   gl,
	}

	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func post(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)

	pid := rq.FormValue("pid")
	v2Db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, _ := v2Db.DB()
	defer db.Close()

	if rq.Method == "POST" {
		msg := rq.PostFormValue("message")
		pId, _ := strconv.Atoi(pid)
		cmt := my.Comment{
			UserId:  int(user.Model.ID),
			PostId:  pId,
			Message: msg,
		}
		v2Db.Create(cmt)
	}

	var pst my.Post
	var cmts []my.CommentJoin

	v2Db.Where("id = ?", pid).First(&pst)
	v2Db.Table("comments").Select("comments.*, users.id, users.name").Joins("join users on users.id = comments.user_id").Where("comments.post_id = ?", pid).Order("created_at desc").Find(&cmts)

	item := struct {
		Title   string
		Message string
		Name    string
		Account string
		Post    my.Post
		Clist   []my.CommentJoin
	}{
		Title:   "Post",
		Message: "Post id=" + pid,
		Name:    user.Name,
		Account: user.Account,
		Post:    pst,
		Clist:   cmts,
	}
	er := page("post").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func home(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)

	v2Db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, _ := v2Db.DB()
	defer db.Close()

	if rq.Method == "POST" {
		switch rq.PostFormValue("form") {
		case "post":
			ad := rq.PostFormValue("address")
			ad = strings.TrimSpace(ad)
			if strings.HasPrefix(ad, "https://youtu.be/") {
				ad = strings.TrimPrefix(ad, "https://youtu.be/")
			}

			pt := my.Post{
				UserId:  int(user.Model.ID),
				Address: ad,
				Message: rq.PostFormValue("message"),
			}
			v2Db.Create(&pt)
		case "group":
			gp := my.Group{
				UserId:  int(user.Model.ID),
				Name:    rq.PostFormValue("name"),
				Message: rq.PostFormValue("message"),
			}
			v2Db.Create(&gp)
		}
	}

	var pts []my.Post
	var gps []my.Group
	v2Db.Where("user_id = ?", user.ID).Order("created_at desc").Limit(10).Find(&pts)
	v2Db.Where("user_id = ?", user.ID).Order("created_at desc").Limit(10).Find(&gps)

	itm := struct {
		Title   string
		Message string
		Name    string
		Account string
		Plist   []my.Post
		Glist   []my.Group
	}{
		Title:   "Home",
		Message: "User account=\"" + user.Account + "\".",
		Name:    user.Name,
		Account: user.Account,
		Plist:   pts,
		Glist:   gps,
	}

	er := page("home").Execute(w, itm)
	if er != nil {
		log.Fatal(er)
	}
}

func group(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)

	gid := rq.FormValue("gid")
	v2Db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, _ := v2Db.DB()
	defer db.Close()

	if rq.Method == "POST" {
		ad := rq.PostFormValue("address")
		ad = strings.TrimSpace(ad)
		if strings.HasPrefix(ad, "https://youtu.be/") {
			ad = strings.TrimPrefix(ad, "https://youtu.be/")
		}
		gid, _ := strconv.Atoi(gid)
		pt := my.Post{
			UserId:  int(user.Model.ID),
			Address: ad,
			Message: rq.PostFormValue("message"),
			GroupId: gid,
		}
		v2Db.Create(&pt)
	}

	var grp my.Group
	var pts []my.Post

	v2Db.Where("id=?", gid).Find(&grp)
	v2Db.Order("created_at desc").Model(&grp).Association("Post").Find(&pts)

	itm := struct {
		Title   string
		Message string
		Name    string
		Account string
		Group   my.Group
		Plist   []my.Post
	}{
		Title:   "Group",
		Message: "Group id=" + gid,
		Name:    user.Name,
		Account: user.Account,
		Group:   grp,
		Plist:   pts,
	}
	er := page("group").Execute(w, itm)
	if er != nil {
		log.Fatal(er)
	}
}

func login(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Title   string
		Message string
		Account string
	}{
		Title:   "Login",
		Message: "type your account & password:",
		Account: "",
	}

	if rq.Method == "GET" {
		er := page("login").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
		return
	}

	if rq.Method == "POST" {
		v2Db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db, _ := v2Db.DB()
		defer db.Close()

		usr := rq.PostFormValue("account")
		pass := rq.PostFormValue("pass")
		item.Account = usr

		var re int64
		var user my.User

		v2Db.Where("account = ? and password = ? ", usr, pass).Find(&user).Count(&re)

		if re <= 0 {
			item.Message = "Wrong account or password."
			page("login").Execute(w, item)
			return
		}

		ses, _ := cs.Get(rq, sesName)
		ses.Values["login"] = true
		ses.Values["account"] = usr
		ses.Values["name"] = user.Name
		ses.Save(rq, w)
		http.Redirect(w, rq, "/", 302)

	}
	er := page("login").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func logout(w http.ResponseWriter, rq *http.Request) {
	ses, _ := cs.Get(rq, sesName)
	ses.Values["login"] = nil
	ses.Values["account"] = nil
	ses.Save(rq, w)
	http.Redirect(w, rq, "/login", 302)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, rq *http.Request) {
		home(w, rq)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, rq *http.Request) {
		post(w, rq)
	})

	http.HandleFunc("/group", func(w http.ResponseWriter, rq *http.Request) {
		group(w, rq)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, rq *http.Request) {
		login(w, rq)
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, rq *http.Request) {
		logout(w, rq)
	})

	http.ListenAndServe("", nil)
}
