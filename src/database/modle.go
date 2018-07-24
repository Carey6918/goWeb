package database

import "time"

type User struct {
	Id         int
	Uuid       string
	Name       string
	Email      string
	Password   string
	CreatedAt time.Time
}

type Session struct {
	Id         int
	Uuid       string
	Name       string
	Email      string
	UserId	   int
	CreatedAt time.Time
}

func (u *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), u.Email, u.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (s *Session) Check() (valid bool,err error) {
	defer Db.Close()
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}
	return
}

