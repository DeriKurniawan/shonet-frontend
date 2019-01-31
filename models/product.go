package models

import (
	"database/sql"
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Brand       Brand     `json:"brand"`
	Category    Category  `json:"category"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	SiteURL     string    `json:"site_url"`
	MustHave    uint      `json:"must_have"`
	TopProduct  uint      `json:"top_product"`
	CreatedBy   int64     `json:"created_by"`
	View        int       `json:"view"`
	Click       int       `json:"click"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductNull struct {
	Description sql.NullString
	CreatedAt   mysql.NullTime
	UpdatedAt   mysql.NullTime
	View		sql.NullInt64
	Click		sql.NullInt64
}

type ProductElastic struct {
	ID				uint				`json:"id"`
	Name 			string				`json:"name"`

	Brand 			struct {
		ID 			uint				`json:"id"`
		Name 		string				`json:"name"`
	}									`json:"brand"`

	Categories		ProductCategories	`json:"category"`

	Thumbnail 		string				`json:"thumbnail"`
	Description		string				`json:"description"`
	Price 			float64				`json:"price"`
	SiteURL			string				`json:"site_url"`

	CreatedBy 		struct {
		ID 	        uint				`json:"id"`
		Name 		string				`json:"name"`
		Photo		string				`json:"photo"`
	}									`json:"created_by"`

	CreatedAt 		string				`json:"created_at"`
	UpdatedAt 		string				`json:"updated_at"`
}

type ProductCategories struct {
	ID 			uint			`json:"id"`
	Parent      json.RawMessage	`json:"parent"`
	Title		string			`json:"title"`
}

type DescriptionProduct struct {
	EditorsNote   string `json:"editors_note"`
	ProductDetail string `json:"product_detail"`
}
