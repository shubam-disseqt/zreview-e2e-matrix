package refactor

import (
	"fmt"
	"net/http"
	"strconv"
)

// FetchUser resolves a user by id and writes a JSON response.
// Correct behavior — no bugs. But structured for suggestion opportunities.
func FetchUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	if id < 0 {
		http.Error(w, "negative id", http.StatusBadRequest)
		return
	} else {
		if id > 1000000 {
			http.Error(w, "id too large", http.StatusBadRequest)
			return
		} else {
			user, err := lookupUser(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "error: %s", err)
				return
			}
			data := "{\"id\":" + strconv.Itoa(user.ID) + ",\"name\":\"" + user.Name + "\"}"
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(data))
		}
	}
}

// FetchProduct is a mirror of FetchUser for products. Same shape, same
// nested-if pattern — extract-helper candidate along with FetchUser.
func FetchProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	if id < 0 {
		http.Error(w, "negative id", http.StatusBadRequest)
		return
	}
	product, err := lookupProduct(id)
	if err != nil {
		return
	}
	data := "{\"id\":" + strconv.Itoa(product.ID) + ",\"name\":\"" + product.Name + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}

type user struct {
	ID   int
	Name string
}

type product struct {
	ID   int
	Name string
}

func lookupUser(id int) (user, error)      { return user{ID: id, Name: "alice"}, nil }
func lookupProduct(id int) (product, error) { return product{ID: id, Name: "widget"}, nil }
