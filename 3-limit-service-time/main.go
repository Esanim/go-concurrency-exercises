//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	lock      sync.Mutex
}

func validate(u *User) bool {
	return u.IsPremium || u.TimeUsed < 10
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if !validate(u) {
		return false
	}

	var gotResult bool

	go func() {
		process()
		gotResult = true
	}()

	t := time.Tick(time.Second * 1)
	for {
		<-t
		if gotResult {
			return true
		}
		u.TimeUsed++
		if !validate(u) {
			return false
		}
	}
}

func main() {
	RunMockServer()
}
