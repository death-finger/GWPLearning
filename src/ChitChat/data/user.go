package data

import "time"

type User struct {
	Id       int
	Uuid     int
	Name     string
	Email    string
	Password string
	CreateAt time.Time
}

type Session struct {
	Id        int
	Uuid      int
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (sess *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", sess.Uuid).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if sess.Id != 0 {
		valid = true
	}
	return
}
