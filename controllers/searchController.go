package controllers

import (
	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"html/template"
	_ "log"
	"net/http"
	"net/url"
	_ "reflect"
	"strings"
)

type DataSearchPage struct {
	Menu 		Menu
	Auth		interface{}
	GoogleID 	string
	OnesignalID string
	PixelID		string
}

func SearchIndexHandler(w http.ResponseWriter, r *http.Request) {

	authUser, err := GetUserClient(r)
	if libraries.CheckError(err) {
		return
	}

	menu, err := GetMenu()
	if libraries.CheckError(err) {
		return
	}

	data := Page{
		Title:			"THE SHONET | search ",
		Description:	"searching all products, articles and branch ",
		Uri:			config.GetString("app.url"),
		Routing:		"search.index",
		Data:	DataSearchPage{
			Menu: 			menu,
			Auth: 			authUser,
			GoogleID:		config.GetString("services.google.analytics.id"),
			OnesignalID:	config.GetString("services.onesignal.app.id"),
			PixelID:		config.GetString("services.facebook.pixel.id"),
		},
	}

	tmpl := template.Must(template.New("main.html").Funcs(template.FuncMap{
		"strtoupper": func(words string) string {
			return strings.ToUpper(words)
		},
		"urlencode": func(words string) string {
			result, _ := url.Parse(words)
			return result.EscapedPath()
		},
		"number_format": func(number float64, decimals uint, point, separator string) string {
			return libraries.NumberFormat(number, decimals, point, separator)
		},
	}).ParseFiles(
		 "views/layout/main.html",
					"views/element/header.html",
					"views/element/menu.html",
					"views/element/offcanvas.html",
					"views/element/sidebar.html",
					"views/element/timeline-share.html",
					"views/element/title.html",
					"views/search/search.html",
	))

	err = tmpl.Execute(w, data)
	if libraries.CheckError(err) {return}
}
