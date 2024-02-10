package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	ID     int       `json:"id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
	UserID int       `json:"userId"`
}

type CalendarAPI struct{}

func (api *CalendarAPI) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var event Event
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Парсинг и валидация параметров
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		http.Error(w, "Invalid date", http.StatusBadRequest)
		return
	}

	event = Event{
		Title:  r.FormValue("title"),
		Date:   date,
		UserID: userID,
	}

	// Здесь должна быть  логика создания события в календаре

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{"result": "Event created successfully"})
}

// Другие обработчики для /update_event, /delete_event и т.д.

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Processing request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	calendarAPI := &CalendarAPI{}

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", calendarAPI.CreateEventHandler)
	// Добавьте другие маршруты здесь

	server := &http.Server{
		Addr:    ":8080", // Порт из конфига
		Handler: loggingMiddleware(mux),
	}

	log.Printf("Server starting on port %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}
}
