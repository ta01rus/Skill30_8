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
func (hs *HttpServer) HomeEndPoint(c *gin.Context) {
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

func (hs *HttpServer) AddTaskEndPoint(c *gin.Context) {
	var (
		task = storage.TaskView{}
	)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	h := c.Request.Header.Get("HX-Request")
	log.Println(h)
	if h == "true" {
		err := c.Bind(&task)
		if err != nil {
			log.Println(err)
			c.HTML(403, "index.html", nil)
			return
		}
		err = task.Check()
		if err != nil {
			log.Println(err)
			c.HTML(403, "index.html", nil)
			return
		}

		task, err := hs.Db.AddTasks(ctx, &task)
		if err != nil {
			log.Println(err)
			c.HTML(403, "index.html", nil)
			return
		}

		htmlStr := fmt.Sprintf(`
		<tr>						
		<td>%d</td>
		<td>%s</td>    
		<td>%s</td>    
		<td>%s</td>    
		<td>%s</td>    
		<td>
			<button class="btn btn-danger" hx-delete="/del-task/{{.ID}}">
				X
			</button>
		</td>                           
		</tr>`, task.ID, task.Title, task.AuthorName, task.AssignedName, task.Content)

		tmpl, err := template.New("tr").Parse(htmlStr)
		if err != nil {
			log.Panicln(err)
		}
		tmpl.Execute(c.Writer, task)
	}
}

func (hs *HttpServer) DelTaskEndPoint(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.HTML(403, "index.html", nil)
		return
	}
	err = hs.Db.DelTasks(ctx, taskId)
	if err != nil {
		log.Println(err)
		c.HTML(403, "index.html", nil)
		return
	}

	c.HTML(200, "index.html", nil)
}
