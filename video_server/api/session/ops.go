package session

import ("sync"
	"../dbops"
	"../defs"
	"../utils"
	"time"
)


// sync.Map在读上可满足于10w+的并发量，但是在于写上因为加了全局锁，所以效率很低
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool{
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func nowInmilli() int64{
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ct := nowInmilli()
	ttl := ct + 30 * 60 * 1000

	ss := &defs.SimpleSession{Username: uname, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, uname)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInmilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
