package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Banner struct {
	ID				uint		`json:"id"`
	Banner 			string		`json:"banner"`
	Title			string		`json:"title"`
	Page			string		`json:"page"`
	Url				string		`json:"url"`
	StartDate		time.Time	`json:"start_date"`
	EndDate			time.Time	`json:"end_date"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt       time.Time	`json:"updated_at"`
}

type BannerNull struct {
	Url             sql.NullString
	EndDate			mysql.NullTime
	CreatedAt		mysql.NullTime
	UpdatedAt       mysql.NullTime
}