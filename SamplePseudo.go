package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Sessions []Session

type Session struct {
	unqiueIdentifer string // ALSO PROB UUID4 BUT never changes between sessions unlike seseision id, it is used to identify the usert unqiuely within the databasee
	name            string
	password        string    // THE PASSWORD WILL BE SHA256 for sequity
	session_id      string    // UUID4 (UUID is a unique id generation format used as convention in most coding practices)
	time            time.Time // represnts a specific time
}

type PracticeStruct struct {
	correct   int64
	incorrect int64
	negatives int64
}

// handles login
func handleLogin() {
	// will check database for creditnals, if they exist it will instantiate an instance within the sessions variable
}

func handleRequestPracticeStruct(sessionID string) {
	//  this will use session id to see if you're logged in by checking Sessions, it will grab unqiue id and check databse for this information and format it as a
	// PracticeStruct and return it within the HTTP
}

func addincorrect()

func main() {

	go removeSessions()
	r.GET("/login", handleLoginBetter)
	r.POST()
}

func removeSessions() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// REMOVE EXPIRED SESSION
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func handleLoginBetter(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
