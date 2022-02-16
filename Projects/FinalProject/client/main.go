package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const toWatchURL = "http://localhost:8080/moviestracker/towatch"
const finishMovieURL = "http://localhost:8080/moviestracker/towatch/finish"
const watchedURL = "http://localhost:8080/moviestracker/watched"
const loginURL = "http://localhost:8080/moviestracker/login"
const signupURL = "http://localhost:8080/moviestracker/signup"
const logoutURL = "http://localhost:8080/moviestracker/logout"

func main() {
	cmd := flag.String("cmd", "", "add, viewToWatch, viewWatched, update, delete, finish, login, logout, signup")
	user := flag.String("user", "", "username input")
	pass := flag.String("pass", "", "password input")
	title := flag.String("title", "", "movie title")
	rating := flag.Int("rating", 0, "movie score from 0 to 5")
	review := flag.String("review", "", "movie review")
	id := flag.Int("id", -1, "ID of movie to process")
	flag.Parse()
	switch *cmd {
	case "add":
		addMovie(*title)
	case "viewToWatch":
		listToWatch(*id)
	case "viewWatched":
		listWatched(*id)
	case "update":
		update(*id, *title)
	case "delete":
		deleteMovie(*id)
	case "finish":
		finishedMovie(*title, *rating, *review)
	case "login":
		login(*user, *pass)
	case "logout":
		logout(*user)
	case "signup":
		signup(*user, *pass)
	}
}

//addMovie sends a request to add a movie to the user's list
func addMovie(title string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"Title\":\"%s\"}", title)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(toWatchURL, "application/json", outData)

	if err != nil {
		log.Fatal(err)
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	defer resp.Body.Close()
}

//listToWatch sends a request to display the list of movies the in the user's watchlist
func listToWatch(id int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	var url string
	if id != -1 {
		url = fmt.Sprintf("%s/%d", toWatchURL, id)
	} else {
		url = toWatchURL
	}

	resp, err := c.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	defer resp.Body.Close()
}

//listWatched sends a request to display the movies the user has finished watching
func listWatched(id int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	var url string
	if id != -1 {
		url = fmt.Sprintf("%s/%d", watchedURL, id)
	} else {
		url = watchedURL
	}

	resp, err := c.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	defer resp.Body.Close()
}

//update sends a request to edit the title of a movie in the user's watchlist
func update(id int, title string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/%d", toWatchURL, id)
	jsonRec := fmt.Sprintf("{\"Title\":\"%s\"}", title)
	outData := bytes.NewBuffer([]byte(jsonRec))
	req, err := http.NewRequest("PUT", url, outData)
	if err != nil {
		log.Fatal(err)
	}

	resp, err2 := c.Do(req)
	if err != nil {
		log.Fatal(err2)
	}
	data, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

//deleteMovie sends a request to delete a movie
func deleteMovie(id int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/%d", toWatchURL, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err2 := c.Do(req)
	if err2 != nil {
		log.Fatal(err)
	}

	data, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

//finishedMovie sends a request to mark a movie as finished
func finishedMovie(title string, rating int, review string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"Title\":\"%s\", \"Rating\":%d, \"Review\":\"%s\"}", title, rating, review)
	outData := bytes.NewBuffer([]byte(jsonRec))
	req, err := http.NewRequest("PUT", finishMovieURL, outData)
	if err != nil {
		log.Fatal(err)
	}

	resp, err2 := c.Do(req)
	if err != nil {
		log.Fatal(err2)
	}
	data, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

//login sends a request for user login
func login(user string, pass string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"Username\":\"%s\", \"Password\":\"%s\"}", user, pass)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(loginURL, "application/json", outData)

	if err != nil {
		log.Fatal(err)
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	defer resp.Body.Close()
}

//logout sends a request for user logout
func logout(user string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"Username\":\"%s\"}", user)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(logoutURL, "application/json", outData)

	if err != nil {
		log.Fatal(err)
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	defer resp.Body.Close()
}

//signup sends a request for account registration
func signup(user string, pass string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"Username\":\"%s\", \"Password\":\"%s\"}", user, pass)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(signupURL, "application/json", outData)

	if err != nil {
		log.Fatal(err)
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	defer resp.Body.Close()
}
