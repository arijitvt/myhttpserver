package scheduler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobData struct {
	JobId string
}

type Scheduler struct {
	Id string
}

func AddJob(c *gin.Context, scheduler *Scheduler, job JobData) {
	c.HTML(http.StatusOK, "alarm.tmpl", gin.H{})

}

func RemoveJobStr(c *gin.Context, sch *Scheduler, jobId string) {

}

func RemoveJob(c *gin.Context, scheduler *Scheduler, job JobData) {

}
