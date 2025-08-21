package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	cpb "datenote/gunk/v1/category"
	epb "datenote/gunk/v1/event"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
)

type Event struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Date     string `db:"date"`
	Info     string `db:"info"`
	Category string `db:"category"`
}

func (event *Event) Validate() error {
	return validation.ValidateStruct(event,
		validation.Field(&event.Name, validation.Required.Error("Name field can not be empty."), validation.Length(3, 50).Error("Name field should have atleast 3 characters and atmost 50 characters")),
		validation.Field(&event.Date, validation.Required.Error("Date is required")),
	)
}

func (h *Handler) listEvent(w http.ResponseWriter, r *http.Request) {
	events, err := h.ec.GetAllEvents(r.Context(), &epb.GetAllEventsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.templates.ExecuteTemplate(w, "home.html", events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	int_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	event, err := h.ec.GetEvent(r.Context(), &epb.GetEventRequest{
		ID: int64(int_id),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if event.Event.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	var e Event
	e.ID = int64(int_id)
	e.Name = event.Event.Name
	e.Date = event.Event.Date
	e.Info = event.Event.Info
	e.Category = event.Event.Category

	err = h.templates.ExecuteTemplate(w, "event.html", e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	ErrorValue := map[string]string{}
	categories, err := h.cc.GetAllCategories(r.Context(), &cpb.GetAllCategoriesRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category := []Category{}
	for _, v := range categories.Categories {
		category = append(category, Category{
			Title: v.Title,
		})
	}

	event := Event{}
	h.loadCreateForm(w, event, category, ErrorValue)
}

func (h *Handler) storeEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var event Event
	err = h.decoder.Decode(&event, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := h.cc.GetAllCategories(r.Context(), &cpb.GetAllCategoriesRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category := []Category{}
	for _, v := range categories.Categories {
		category = append(category, Category{
			Title: v.Title,
		})
	}

	err = event.Validate()
	if err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			ErrorValue := make(map[string]string)
			for key, value := range vErrors {
				ErrorValue[strings.Title(key)] = value.Error()
			}
			h.loadCreateForm(w, event, category, ErrorValue)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.ec.CreateEvent(r.Context(), &epb.CreateEventRequest{
		Event: &epb.Event{
			Name:     event.Name,
			Date:     event.Date,
			Info:     event.Info,
			Category: event.Category,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) editEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	int_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	event, err := h.ec.GetEvent(r.Context(), &epb.GetEventRequest{
		ID: int64(int_id),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if event.Event.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	var e Event
	e.ID = int64(int_id)
	e.Name = event.Event.Name
	e.Date = event.Event.Date
	e.Info = event.Event.Info
	e.Category = event.Event.Category

	categories, err := h.cc.GetAllCategories(r.Context(), &cpb.GetAllCategoriesRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category := []Category{}
	for _, v := range categories.Categories {
		category = append(category, Category{
			Title: v.Title,
		})
	}

	h.loadEditForm(w, e, category, map[string]string{})
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	int_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	var e Event
	e.ID = int64(int_id)

	if e.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.decoder.Decode(&e, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := h.cc.GetAllCategories(r.Context(), &cpb.GetAllCategoriesRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category := []Category{}
	for _, v := range categories.Categories {
		category = append(category, Category{
			Title: v.Title,
		})
	}

	err = e.Validate()
	if err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			ErrorValue := make(map[string]string)
			for key, value := range vErrors {
				ErrorValue[strings.Title(key)] = value.Error()
			}
			h.loadEditForm(w, e, category, ErrorValue)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.ec.UpdateEvent(r.Context(), &epb.UpdateEventRequest{
		Event: &epb.Event{
			ID:       e.ID,
			Name:     e.Name,
			Date:     e.Date,
			Info:     e.Info,
			Category: e.Category,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	int_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	_, err = h.ec.DeleteEvent(r.Context(), &epb.DeleteEventRequest{
		ID: int64(int_id),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

type SearchedFormData struct {
	SearchResult []Event
	SearchQuery  string
}

func (h *Handler) searchEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sq := r.FormValue("SearchPost")

	posts, err := h.ec.SearchEvent(context.Background(), &epb.SearchEventRequest{SearchEventQuery: sq})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var sResult []Event
	for _, event := range posts.SearchEventResult {
		sResult = append(sResult, Event{
			ID:       event.ID,
			Name:     event.Name,
			Date:     event.Date,
			Info:     event.Info,
			Category: event.Category,
		})
	}
	slt := SearchedFormData{
		SearchResult: sResult,
		SearchQuery:  sq,
	}
	if len(sResult) == 0 {
		err = h.templates.ExecuteTemplate(w, "no-search-result.html", slt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err = h.templates.ExecuteTemplate(w, "search-result-event.html", slt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type FormData struct {
	Event    Event
	Category []Category
	Errors   map[string]string
}

func (h *Handler) loadCreateForm(w http.ResponseWriter, event Event, category []Category, myErrors map[string]string) {
	form := FormData{
		Event:    event,
		Category: category,
		Errors:   myErrors,
	}

	err := h.templates.ExecuteTemplate(w, "create-event.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) loadEditForm(w http.ResponseWriter, event Event, category []Category, myErrors map[string]string) {
	form := FormData{
		Event:    event,
		Category: category,
		Errors:   myErrors,
	}

	err := h.templates.ExecuteTemplate(w, "edit-event.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
