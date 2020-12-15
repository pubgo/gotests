package models

import "time"

type Article struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" xorm:"varchar(25) notnull unique 'name'"`
	Comments  []Comment `json:"comments" xorm:"-"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type Comment struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" xorm:"varchar(25) notnull unique 'name'"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	Article   Article   `json:"article" xorm:"varchar(25) notnull unique 'article'"`
}
