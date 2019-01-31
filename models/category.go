package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	ID        uint       `json:"id"`
	ParentID  uint       `json:"parent_id"`
	Group     string     `json:"group"`
	Title     string     `json:"title"`
	Childs    []Category `json:"childs"`
	Image     string     `json:"image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Icon	  string	 `json:"icon"`
}

type CategoryNull struct {
	Id		  sql.NullInt64
	Title	  sql.NullString
	ParentID  sql.NullInt64
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
	Icon	  sql.NullString
}
