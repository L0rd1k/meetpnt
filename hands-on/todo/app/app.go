package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/L0rd1k/meetpnt/app/config"
	"github.com/L0rd1k/meetpnt/app/handler"
	"github.com/L0rd1k/meetpnt/app/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	config config.Config
	dbase  *gorm.DB
	router *mux.Router
}

func NewApp() *App {
	app := &App{config: *config.NewConfig("settings")}
	return app
}

func (app *App) Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		app.config.DataBase.Username,
		app.config.DataBase.Password,
		app.config.DataBase.Host,
		app.config.DataBase.Port,
		app.config.DataBase.Name,
		app.config.DataBase.Charset)

	db, err := gorm.Open(app.config.DataBase.Dialect, dsn)
	if err != nil {
		log.Fatal("Error: No connection with DB:", err)
	}

	app.dbase = model.Migrate(db)
	app.router = mux.NewRouter()
	fmt.Println("Set routers...")
	app.setRouters()
}

func (app *App) setRequest(handler func(db *gorm.DB, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.dbase, w, r)
	}
}

func (app *App) setRouters() {
	// Project struct
	app.Get("/projects", app.setRequest(handler.GetProjects))
	app.Get("/projects/{title}", app.setRequest(handler.GetProject))
	app.Post("/projects", app.setRequest(handler.CreateProject))
	app.Put("/projects/{title}", app.setRequest(handler.UpdateProject))
	app.Delete("/projects/{title}", app.setRequest(handler.DeleteProject))
	// Task struct
	// app.Get("/projects/{title}/tasks", app.setRequest(handler.GetTasks))
	// app.Post("/projects/{title}/tasks", app.setRequest(handler.CreateTask))
	// app.Get("/projects/{title}/tasks/{id:[0-9]+}", app.setRequest(handler.GetTask))
	// app.Delete("/projects/{title}/tasks/{id:[0-9]+}", app.setRequest(handler.DeleteTask))
	// app.Put("/projects/{title}/tasks/{id:[0-9]+}", app.setRequest(handler.UpdateTask))

}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("PUT")
}

func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.router.HandleFunc(path, f).Methods("DELETE")
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.router))
}
