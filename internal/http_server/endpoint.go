package httpserver

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/ta01rus/Skill30_8/internal/storage"
)

func (hs *HttpServer) Home(w http.ResponseWriter, _ *http.Request) {
	var (
		connDB = hs.db.Conn()
		args   = struct {
			Users []storage.Users
		}{}
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := connDB.QueryContext(ctx, `
		select id , name  from users
	`)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		user := storage.Users{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Println(err)
		}

		args.Users = append(args.Users, user)
	}
	tmpl := template.Must(template.ParseFiles("./template/index.html"))
	tmpl.Execute(w, args)
}

func (hs *HttpServer) AddUser(w http.ResponseWriter, r *http.Request) {
	var (
		user   = storage.Users{}
		connDB = hs.db.Conn()
	)
	// ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	// defer cancel()

	h := r.Header.Get("HX-Request")
	log.Println(h)
	if h == "true" {
		// r.ParseMultipartForm(10 << 20)
		r.ParseForm()
		if r.PostForm.Has("user") {
			user.Name = r.PostFormValue("user")
		}
		if user.Valid() {
			tx, err := connDB.Begin()
			if err != nil {
				log.Println(err)
			}

			stmt, err := tx.Prepare("INSERT INTO users (name) VALUES ($1) RETURNING id")

			if err != nil {
				tx.Rollback()
				log.Println(err)
				return
			}
			defer stmt.Close()

			err = stmt.QueryRow(user.Name).Scan(&user.ID)
			if err != nil {
				log.Println(err)
				return
			}
			tx.Commit()

			htmlStr := fmt.Sprintf(`
			<tr>			
			<td>%d</td>
			<td>%s</td>                              
		  	</tr>`, user.ID, user.Name)

			tmpl, err := template.New("tr").Parse(htmlStr)
			if err != nil {
				log.Panicln(err)
			}
			tmpl.Execute(w, tmpl)
		}
	}
}

func (hs *HttpServer) DelUser(w http.ResponseWriter, r *http.Request) {
	var (
	// user   = storage.Users{}
	// connDB = hs.db.Conn()
	)
	// ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	// defer cancel()

	h := r.Header.Get("HX-Request")
	log.Println(h)
	if h == "true" {
		fmt.Println(r.URL.Fragment)

	}
}
