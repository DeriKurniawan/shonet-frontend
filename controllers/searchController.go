package controllers

import (
	"encoding/json"
	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"github.com/jacky-htg/shonet-frontend/models"
	"github.com/jacky-htg/shonet-frontend/repositories"
	"github.com/pkg/errors"
	"html/template"
	"io/ioutil"
	"log"
	_ "log"
	"net/http"
	"net/url"
	_ "reflect"
	"strings"
)

type DataSearchPage struct {
	Menu 		  Menu
	Auth		  interface{}
	GoogleID 	  string
	OnesignalID   string
	PixelID		  string
	Word		  string
	Catex		  string
	Typex         string
	BaseURL		  template.URL
	ApiUrl		  interface{}
	Articles    []interface{}
	Products    []interface{}
	Users       []interface{}
	Banner        interface{}
	Categories  []models.Category
}

var errx     error
var users    []interface{}
var products []interface{}
var articles []interface{}

func SearchIndexHandler(w http.ResponseWriter, r *http.Request) {

	authUser, err := GetUserClient(r)
	if libraries.CheckError(err) {
		return
	}

	menu, err := GetMenu()
	if libraries.CheckError(err) {
		return
	}

	var api map[string]interface{}
	var word  string
	var catx  string
	var typex string

	api = map[string]interface{}{
		"apiKey": config.GetString("elasticsearch.apiKey"),
		"url"   : config.GetString("elasticsearch.url"),
	}

	err = r.ParseForm()
	if libraries.CheckError(err) {return}

	word   = r.FormValue("search")
	catx   = r.FormValue("category")
	typex  = r.FormValue("type")


	GetUsersDataSearch(r)
	if libraries.CheckError(errx) {return}

	GetProductsDataSearch(r)
	if libraries.CheckError(errx) {return}

	GetArticlesDataSearch(r)
	if libraries.CheckError(errx) {return}

	banner, err := repositories.GetBannerForFront()
	if libraries.CheckError(err) {return}

	categories, err := repositories.GetCategoryForMenu([]string{"group", "J"})
	if libraries.CheckError(err) {return}

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
			BaseURL:		template.URL(config.GetString("app.url")),
			Word:			word,
			Catex:			catx,
			Typex:          typex,
			ApiUrl:			api,
			Users:          users,
			Products:       products,
			Articles:		articles,
			Banner:         banner,
			Categories:     categories,
		},
	}

	tmpl := template.Must(template.New("main.html").Funcs(template.FuncMap{
		"strtoupper": func(words string) string {
			return strings.ToUpper(words)
		},
		"urlencode": func(words string) string {
			result, err := url.Parse(words)
			if err!=nil {log.Printf("ERROR: <urlencode> - ", err.Error()); return words}

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

func GetUsersDataSearch(r *http.Request) {
	err := r.ParseForm()
	if err!=nil {
		err = errors.New(" Error @searchController:GetUsersDataSearch #ParseForm :: " + err.Error())
		errx = err
		return
	}

	search, err := url.Parse(r.FormValue("search"))
	typex, err  := url.Parse(r.FormValue("type"))
	catx, err   := url.Parse(r.FormValue("category"))
	if err!=nil {
		err = errors.New(" Error @searchController:GetUsersDataSearch #url.Parse#r.FormValue :: " + err.Error())
		errx = err
		return
	}

	urlx := config.GetString("elasticsearch.url") +
		    "/elastic/search/users?word=" + search.EscapedPath() +
		    "&type=" + typex.EscapedPath() +
		    "&category=" + catx.EscapedPath()

	request, err := http.NewRequest("GET", urlx, nil)
	if err!=nil {
		err = errors.New(" Error @searchController:GetUsersDataSearch #request#http.NewRequest :: " + err.Error())
		errx = err
		return
	}

	request.Header.Set("X-Api-Key", config.GetString("elasticsearch.apiKey"))

	client := http.Client{}
	response, err := client.Do(request)
	if err!=nil {
		err = errors.New(" Error @searchController:GetUsersDataSearch #response#client.Do(request) :: " + err.Error())
		errx = err
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err!=nil {
		err = errors.New(" Error @searchController:GetUsersDataSearch #ioutil.ReadAll#response.Body :: " + err.Error())
		errx = err
		return
	}

	err = json.Unmarshal(body, &users)
	if err!=nil {
		err = errors.New(" Error @searchController:GetUsersDataSearch #json.Unmarshal#&user :: " + err.Error())
		errx = err
		return
	}
}

func GetProductsDataSearch(r *http.Request) {
	err := r.ParseForm()
	if err!=nil {
		err = errors.New(" Error @searchController:GetProductsDataSearch #r.ParseForm :: " + err.Error())
		errx = err
		return
	}

	search, err := url.Parse(r.FormValue("search"))
	typex, err  := url.Parse(r.FormValue("type"))
	catx, err   := url.Parse(r.FormValue("category"))
	if err!=nil {
		err = errors.New(" Error @searchController:GetProductsDataSearch #url.Parse#r.FormValue :: " + err.Error())
		errx = err
		return
	}

	urlx := config.GetString("elasticsearch.url") +
		    "/elastic/search/products?word=" + search.EscapedPath() +
			"&type=" + typex.EscapedPath() +
			"&category=" + catx.EscapedPath()

	request, err := http.NewRequest("GET", urlx, nil)
	if err!=nil {
		err = errors.New(" Error @searchController:GetProductsDataSearch #url.Parse#r.FormValue :: " + err.Error())
		errx = err
		return
	}

	request.Header.Set("X-Api-Key", config.GetString("elasticsearch.apiKey"))

	client := http.Client{}
	response, err := client.Do(request)
	if err!=nil {
		err = errors.New(" Error @searchController:GetProductsDataSearch #response#client.Do(request) :: " + err.Error())
		errx = err
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err!=nil {
		err = errors.New(" Error @searchController:GetProductsDataSearch #ioutil.ReadAll#response.Body :: " + err.Error())
		errx = err
		return
	}

	err = json.Unmarshal(body, &products)
	if err!=nil {
		err = errors.New(" Error @searchController:GetProductsDataSearch #json.Unmarshal#&user :: " + err.Error())
		errx = err
		return
	}
}

func GetArticlesDataSearch(r *http.Request) {
	err := r.ParseForm()
	if err!=nil {
		err = errors.New(" Error @searchController:GetArticlesDataSearch #r.ParseForm :: " + err.Error())
		errx = err
		return
	}

	search, err := url.Parse(r.FormValue("search"))
	typex, err  := url.Parse(r.FormValue("type"))
	catx, err   := url.Parse(r.FormValue("category"))
	if err!=nil {
		err = errors.New(" Error @searchController:GetArticlesDataSearch #url.Parse#r.FormValue :: " + err.Error())
		errx = err
		return
	}

	urlx := config.GetString("elasticsearch.url") +
		    "/elastic/search/articles?word=" + search.EscapedPath() +
			"&type=" + typex.EscapedPath() +
			"&category=" + catx.EscapedPath()

	request, err := http.NewRequest("GET", urlx, nil)
	if err!=nil {
		err = errors.New(" Error @searchController:GetArticlesDataSearch #url.Parse#r.FormValue :: " + err.Error())
		errx = err
		return
	}

	request.Header.Set("X-Api-Key", config.GetString("elasticsearch.apiKey"))

	client := http.Client{}
	response, err := client.Do(request)
	if err!=nil {
		err = errors.New(" Error @searchController:GetArticlesDataSearch #response#client.Do(request) :: " + err.Error())
		errx = err
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err!=nil {
		err = errors.New(" Error @searchController:GetArticlesDataSearch #ioutil.ReadAll#response.Body :: " + err.Error())
		errx = err
		return
	}

	err = json.Unmarshal(body, &articles)
	if err!=nil {
		err = errors.New(" Error @searchController:GetArticlesDataSearch #json.Unmarshal#&user :: " + err.Error())
		errx = err
		return
	}
}