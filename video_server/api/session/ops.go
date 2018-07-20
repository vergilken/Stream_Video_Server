package session

import "sync"


// sync.Map在读上可满足于10w+的并发量，但是在于写上因为加了全局锁，所以效率很低
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {

}

func GenerateNewSessionId(un string) string {

}

func IsSessionExpired(sid string) (string, bool) {

}
