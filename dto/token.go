package dto

import "time"

var t token

type token struct {
	time time.Time
	sqlInfo DBInfo
	token string
}

func NewToken() *token {
	return &t
}

func (t *token)GetToken() string {
	return t.token
}

func (t *token)GetTime() time.Time {
	return t.time
}

func (t *token)Set(token string, time time.Time, sqlInfo *DBInfo) {
	t.token = token
	t.time = time
	if sqlInfo != nil {
		t.sqlInfo = *sqlInfo
	}
}
