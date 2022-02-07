package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Contact struct {
	ID       int
	Last     string
	First    string
	Company  string
	Address  string
	Country  string
	Position string
}

type Database struct {
	nextID int
	mu     sync.Mutex
	conts  []Contact
}

func main() {
	db := &Database{conts: []Contact{}}
	http.ListenAndServe(":8080", db.handler())
}

//handler reads the URL and calls designated processes
func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/contacts" {
			db.process(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/contacts/%d", &id); n == 1 {
			db.processID(id, w, r)
		} else {
			http.Error(w, "\"Error 400\": Incorrect URL format.", http.StatusBadRequest)
		}
	}
}

//process handles requests with no ID
func (db *Database) process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var contact Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if db.isDuplicateContact(contact) {
			http.Error(w, "\"Error 409\": Contact already exists.", http.StatusConflict)
			return
		} else {
			db.mu.Lock()
			contact.ID = db.nextID
			db.nextID++
			db.conts = append(db.conts, contact)
			db.mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "{\"success\": Contact with ID#%v saved to database}\n", contact.ID)
		}
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(db.conts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
}

//proceessID handles requests with given ID
func (db *Database) processID(id int, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		exists, contactIndex := db.isIdExists(id)
		if exists {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(db.conts[contactIndex]); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "\"Error 404\": Contact not found in database.", http.StatusNotFound)
			return
		}
	case "PUT":
		exists, contactIndex := db.isIdExists(id)
		if exists {
			var newContact Contact
			if err := json.NewDecoder(r.Body).Decode(&newContact); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			db.mu.Lock()
			newContact.ID = db.conts[contactIndex].ID
			db.conts[contactIndex] = newContact
			db.mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "{\"success\": Contact with ID#%v updated.}\n", newContact.ID)

		} else {
			http.Error(w, "\"Error 404\": Contact not found in database.", http.StatusNotFound)
			return
		}
	case "DELETE":
		exists := false
		db.mu.Lock()
		for j, item := range db.conts {
			if id == item.ID {
				db.conts = append(db.conts[:j], db.conts[j+1:]...)
				exists = true
				break
			}
		}
		db.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if exists {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "{\"success\": Contact with ID#%v deleted.}\n", id)
		} else {
			http.Error(w, "\"Error 404\": Contact not found in database.", http.StatusNotFound)
			return
		}
	default:
		http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
}

//isDuplicateContact checks if the new contact already exists in the database and returns a boolean
func (db *Database) isDuplicateContact(contact Contact) bool {
	isDuplicate := false
	for _, item := range db.conts {
		if item.First == contact.First && item.Last == contact.Last {
			isDuplicate = true
			break
		}
	}

	return isDuplicate
}

//isIdExists finds the ID in the database and returns a boolean and the index of that ID in the database
func (db *Database) isIdExists(id int) (bool, int) {
	isIdExists := false
	var contactIndex int
	for i, item := range db.conts {
		if item.ID == id {
			isIdExists = true
			contactIndex = i
			break
		}
	}

	return isIdExists, contactIndex
}
