package domain

import "time"

// 领域对象user,是DDD中的聚合根，entity
type User struct {
	Id       int64
	Email    string
	Password string
	Nickname string
	Phone    string
	Birthday string
	Profile  string
	Ctime    time.Time
	Utime    time.Time
}

//
//type Address struct{}
