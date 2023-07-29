package service

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	storage "github.com/ta01rus/Skill30_8/pkg/storage"
)

// стартовая страница отабражаает поль
func (hs *HttpServer) HomeEndPoooint(c *gin.Context) {
	var (
		err              error
		page             = 1
		id, athID, asgID = 0, 0, 0
		offset, limit    = 0, 5

		args = struct {
			Error string
			Tasks []*storage.TaskView
		}{}
	)

	id, _ = strconv.Atoi(c.Query("id"))
	asgID, _ = strconv.Atoi(c.Query("assigned"))
	athID, _ = strconv.Atoi(c.Query("author"))

	page, _ = strconv.Atoi(c.Query("page"))

	offset = (page * offset) - offset

	ctx, cancel := context.WithTimeout(c.Request.Context(), 100*time.Second)
	defer cancel()

	args.Tasks, err = hs.Db.Tasks(ctx, id, athID, asgID, offset, limit)

	if err != nil {
		log.Println(err)
		c.HTML(403, "index.html", nil)
		return
	}

	c.HTML(200, "index.html", args)

}

func (hs *HttpServer) AddUserEndPoint(c *gin.Context) {
	var (
		user = storage.Users{}
	)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h := c.Request.Header.Get("HX-Request")
	log.Println(h)
	if h == "true" {
		err := c.Bind(&user)
		if err != nil {
			log.Println(err)
			c.HTML(403, "index.html", nil)
			return
		}
		if user.Valid() {
			user, err := hs.Db.AddUsers(ctx, &user)
			htmlStr := fmt.Sprintf(`
			<tr>						
			<td>%d</td>
			<td>%s</td>    
			<td>
				<button class="btn btn-danger" hx-delete="/del-user/{{.ID}}">
					X
				</button>
			</td>                           
		  	</tr>`, user.ID, user.Name)

			tmpl, err := template.New("tr").Parse(htmlStr)
			if err != nil {
				log.Panicln(err)
			}
			tmpl.Execute(c.Writer, user)

		}
	}
}

func (hs *HttpServer) DelUserEndPoint(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.HTML(403, "index.html", nil)
		return
	}
	err = hs.Db.DelUsers(ctx, userId)
	if err != nil {
		log.Println(err)
		c.HTML(403, "index.html", nil)
		return
	}

	c.HTML(200, "index.html", nil)
}
