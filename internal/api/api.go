package api

import (
	// "context"
	// "encoding/json"
	// "errors"
	// "log/slog"
	"net/http"
	// "sync"

	"github.com/bsmedina/tech-go-react-server/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}


func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w,r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)


  r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300,
  }))


	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)
			
			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateRoomMessages)
				r.Get("/", a.handleGetMessages)
				
				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemoveReactFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleCreateRoomMessages(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleGetMessages(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleGetMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleReactToMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request){}
func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request){}