package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Blog struct {
	ID           int
	Name         string
	Title        string
	Startdate    time.Time
	Enddate      time.Time
	Description  string
	Technologies []string
	Image        string
	Content      string
	Check1       string
	Check2       string
	Check3       string
	Check4       string
}

var Blogs = []Blog{
	{
		Title:     "test",
		ID : 0,
		Name : "Dion",
		Startdate: time.Now(),
		Enddate:   time.Now(),
		Description : "test",
		Technologies : []string{"NodeJS", "ReactJS", "PHP", "Javascript"},
		Image : "public/img/dummy1",
		// Startdate: "2022-11-24",
		// Enddate:   "2022-11-25",
		Content: "Test",
		Check1:  "checked",
		Check2:  "checked",
		Check3:  "checked",
		Check4:  "checked",
	},
}

func main() {
	route := mux.NewRouter()

	// connection to databse
	connection.DatabaseConnect()

	//route untuk menginisialisai folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/formblog", formblog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/addblog", addblog).Methods("POST")
	route.HandleFunc("/delete-blog/{index}", deleteBlog).Methods("GET")
	route.HandleFunc("/update-blog/{index}", updateBlog).Methods("POST")
	route.HandleFunc("/update-blog/{index}", getUpdateBlog).Methods("GET")

	fmt.Println("Server berjalan pada port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	dataBlog, errQuery := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tbl_blog")
	if errQuery != nil {
		fmt.Println("Message : " + errQuery.Error())
		return
	}

	var result []Blog

	for dataBlog.Next() {
		var each = Blog{}

		err := dataBlog.Scan(&each.ID, &each.Name, &each.Startdate, &each.Enddate, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println("Message : " + err.Error())
			return
		}

		result = append(result, each)
	}

	fmt.Print(result)
	resData := map[string]interface{}{
		"Blogs": result,
	}

	tmpt.Execute(w, resData)

	// if err != nil {
	// 	w.Write([]byte("Message : " + err.Error()))
	// 	return
	// }
	// respData := map[string]interface{}{
	// 	"Blogs": Blogs,
	// }
	// tmpt.Execute(w, respData)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, tmpt)
}

func formblog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/addblog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, tmpt)
}

func addblog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	// startdate := r.PostForm.Get("std")
	// enddate := r.PostForm.Get("etd")
	check1 := r.PostForm.Get("check1")
	check2 := r.PostForm.Get("check2")
	check3 := r.PostForm.Get("check3")
	check4 := r.PostForm.Get("check4")

	var newBlog = Blog{
		Title:   title,
		Content: content,
		// Startdate: startdate,
		// Enddate:   enddate,
		Check1: check1,
		Check2: check2,
		Check3: check3,
		Check4: check4,
	}

	Blogs = append(Blogs, newBlog)

	fmt.Println(check1)
	fmt.Println(check2)
	fmt.Println(check3)
	fmt.Println(check4)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// get update blog[index]
func getUpdateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/update-blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["index"])

	var BlogDetail = Blog{}

	for index, data := range Blogs {
		if index == id {
			BlogDetail = Blog{
				Title:     data.Title,
				Startdate: data.Startdate,
				Enddate:   data.Enddate,
				Content:   data.Content,
				Check1:    data.Check1,
				Check2:    data.Check2,
				Check3:    data.Check3,
				Check4:    data.Check4,
			}
		}
	}

	Detail := map[string]interface{}{
		"Blogs": BlogDetail,
	}

	tmpt.Execute(w, Detail)
}

// update blog berdasarkan id

func updateBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(mux.Vars(r)["index"])

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	// startdate := r.PostForm.Get("std")
	// enddate := r.PostForm.Get("etd")
	check1 := r.PostForm.Get("check1")
	check2 := r.PostForm.Get("check2")
	check3 := r.PostForm.Get("check3")
	check4 := r.PostForm.Get("check4")

	var newBlog = Blog{
		Title:   title,
		Content: content,
		// Startdate: startdate,
		// Enddate:   enddate,
		Check1: check1,
		Check2: check2,
		Check3: check3,
		Check4: check4,
	}

	Blogs[id] = newBlog

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var BlogDetail = Blog{}

	for index, data := range Blogs {
		if index == id {
			BlogDetail = Blog{
				Title:     data.Title,
				Startdate: data.Startdate,
				Enddate:   data.Enddate,
				Content:   data.Content,
			}
		}
	}

	Detail := map[string]interface{}{
		"Blogs": BlogDetail,
	}

	tmpt.Execute(w, Detail)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	Blogs = append(Blogs[:index], Blogs[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}
