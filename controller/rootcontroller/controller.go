package rootcontroller

import (
	"database/sql"
	"flag"
	"net/http"

	"github.com/L-oris/yabb/logger"
	"github.com/L-oris/yabb/models/tpl"
	"github.com/gorilla/mux"
)

type Config struct {
	DB         *sql.DB
	PathPrefix string
	Tpl        tpl.Template
	ServeFile  serveFile
}

type rootController struct {
	Router *mux.Router
	db     *sql.DB
	serveFile
	tpl tpl.Template
}

// New creates a new controller and registers the routes
func New(config *Config) rootController {
	c := rootController{
		db:        config.DB,
		serveFile: config.ServeFile,
		tpl:       config.Tpl,
	}

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(c.static())

	routes := router.PathPrefix(config.PathPrefix).Subrouter()
	routes.HandleFunc("/", c.home).Methods("GET")
	routes.HandleFunc("/ping", c.ping).Methods("GET")
	routes.HandleFunc("/favicon.ico", c.favicon).Methods("GET")

	c.Router = router
	return c
}

// static serves static files
func (c rootController) static() http.Handler {
	var dir string
	flag.StringVar(&dir, "dir", "public/", "the directory to serve files from /public")
	flag.Parse()

	return http.StripPrefix("/static/", http.FileServer(http.Dir(dir)))
}

// home serves the Home page
func (c rootController) home(w http.ResponseWriter, req *http.Request) {
	c.tpl.Render(w, "home.gohtml", nil)
}

// ping is used for health check
func (c rootController) ping(w http.ResponseWriter, req *http.Request) {
	logger.Log.Debug("ping pong request")
	w.Write([]byte("pong"))
}

func (c rootController) favicon(w http.ResponseWriter, req *http.Request) {
	c.serveFile(w, req, "public/favicon.ico")
}
