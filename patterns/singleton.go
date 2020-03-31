package patterns

import (
	"sync"
	"gopkg.in/mgo.v2"
)

type SingletonMongoDBSession struct{
	session *mgo.Session
	errDial error
}

var instance *SingletonMongoDBSession
var once sync.Once

func GetSessErrMongoDBSession(dialInfo string) (*mgo.Session, error) {
	var instance *SingletonMongoDBSession
	if dialInfo == "Dial" {
		instance = NewMongoDBSesssion()
	}	else if dialInfo == "DialWithInfo" {
		instance = NewMongoDBSessionInfo()
	}
	return instance.session, instance.errDial
}

func NewMongoDBSessionInfo() *SingletonMongoDBSession {
	once.Do(func() {
		sess, err := mgo.DialWithInfo(mongoDBDialInfo)
		instance = &SingletonMongoDBSession{session: sess, errDial: err}
	})
	return instance
}

func NewMongoDBSesssion() *SingletonMongoDBSession {
	once.Do(func() {
		sess, err := mgo.Dial(urldb)
		instance = &SingletonMongoDBSession{session: sess, errDial: err}
	})
	return instance
}

func NewMongoDBSessionClassic() *SingletonMongoDBSession {
	if instance == nil {
		instance = &SingletonMongoDBSession{}
	}
	return instance
}