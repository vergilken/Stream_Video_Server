package dbops

import (
	"strconv"
	"../defs"
	"database/sql"
	"sync"
	"log"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)` )
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name from sessions WHERE session_id = ?" )
	if err != nil {
		return nil, err
	}

	var ttl, uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 64, 10); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else  {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare(`SELECT * FROM sessions`)
	if err != nil {
		log.Printf("%s\n", err)
		return m, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s\n", err)
		return m, err
	}

	for rows.Next() {
		var id, ttlstr, login_name string
		if err1 := rows.Scan(&id, &ttlstr, &login_name); err1 != nil {
			log.Printf("%s\n", err1)
			break
		}

		if ttl, err2 := strconv.ParseInt(ttlstr, 10, 64); err2 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl,}
			m.Store(id, ss)
			log.Printf("session id : %s, ttl : %d\n", id, ss.TTL)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare(`DELETE FROM sessions WHERE session_id = ?`)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	if _, err := stmtOut.Exec(sid); err != nil {
		return err
	}

	defer stmtOut.Close()
	return nil
}

