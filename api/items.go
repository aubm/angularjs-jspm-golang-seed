package api

import (
	"io"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	items := []interface{}{
		item{"An item", "An item description"},
		item{"An other item", "An other item description"},
		item{"This is a third item", "This is a third item description"},
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(toJson(items)[:]))
}

type item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
