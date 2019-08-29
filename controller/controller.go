package controller

import "github.com/gin-contrib/sessions"

const (
	sessionUserIdKey = iota
)

func SessionUserId(sess sessions.Session) (uint64, bool) {
	id := sess.Get(sessionUserIdKey)
	if nil == id {
		return 0, false
	}
	return id.(uint64), true
}
func SessionUserIdKey() int {
	return sessionUserIdKey
}
