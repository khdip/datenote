package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	cpb "datenote/gunk/v1/category"
	epb "datenote/gunk/v1/event"
)

// const sessionName = "cms-session"

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	session   *sessions.CookieStore
	ec        epb.EventServiceClient
	cc        cpb.CategoryServiceClient
}

func GetHandler(decoder *schema.Decoder, session *sessions.CookieStore, ec epb.EventServiceClient, cc cpb.CategoryServiceClient) *mux.Router {
	hand := &Handler{
		decoder: decoder,
		session: session,
		ec:      ec,
		cc:      cc,
	}
	hand.GetTemplate()

	r := mux.NewRouter()
	r.HandleFunc("/", hand.listEvent)
	r.HandleFunc("/categories", hand.listCategory)
	r.HandleFunc("/event/create", hand.createEvent)
	r.HandleFunc("/event/store", hand.storeEvent)
	r.HandleFunc("/category/create", hand.createCategory)
	r.HandleFunc("/category/store", hand.storeCategory)
	r.HandleFunc("/event/{id:[0-9]+}/view", hand.getEvent)
	r.HandleFunc("/event/{id:[0-9]+}/edit", hand.editEvent)
	r.HandleFunc("/event/{id:[0-9]+}/update", hand.updateEvent)
	r.HandleFunc("/event/{id:[0-9]+}/delete", hand.deleteEvent)
	r.HandleFunc("/category/{id:[0-9]+}/edit", hand.editCategory)
	r.HandleFunc("/category/{id:[0-9]+}/update", hand.updateCategory)
	r.HandleFunc("/category/{id:[0-9]+}/delete", hand.deleteCategory)
	r.HandleFunc("/posts/q", hand.searchEvent)
	r.HandleFunc("/categories/q", hand.searchCategory)

	// s.Use(hand.authMiddleware)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("cms/assets/images"))))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := hand.templates.ExecuteTemplate(w, "404.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return r
}

func (h *Handler) GetTemplate() {
	h.templates = template.Must(template.ParseFiles(
		"cms/assets/templates/home.html",
		"cms/assets/templates/event.html",
		"cms/assets/templates/category-list.html",
		"cms/assets/templates/create-event.html",
		"cms/assets/templates/create-category.html",
		"cms/assets/templates/edit-event.html",
		"cms/assets/templates/edit-category.html",
		"cms/assets/templates/404.html",
		"cms/assets/templates/search-result-event.html",
		"cms/assets/templates/search-result-category.html",
		"cms/assets/templates/no-search-result.html",
	))
}
