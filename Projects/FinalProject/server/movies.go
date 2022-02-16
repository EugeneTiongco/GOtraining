package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type Movie struct {
	ID     int
	Title  string
	Rating int
	Review string
}

type MovieList struct {
	MovieList []Movie `json:"movielist"`
}

type User struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	MoviesToWatch MovieList `json:"moviestowatch"`
	MoviesWatched MovieList `json:"movieswatched"`
	NextId        int       `json:"nextid"`
}

type JsonDatabase struct {
	Users []User `json:"users"`
}

type Database struct {
	mu         sync.Mutex
	userList   []User
	activeUser string
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

//Initiates the loggers
func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	dbContents := ReadDatabase()
	db := &Database{userList: dbContents.Users}
	http.ListenAndServe(":8080", db.handler())
}

//ReadDatabase reads the contents of the database in the JSON File
func ReadDatabase() JsonDatabase {
	file, err := os.Open("movieDatabase.json")
	if err != nil {
		log.Fatalf("failed to open json file")
	}
	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read json file")
	}
	var result JsonDatabase
	_ = json.Unmarshal([]byte(byteData), &result)

	InfoLogger.Println("Database read.")
	return result
}

//WriteToDatabase updates the database in the JSON file
func WriteToDatabase(ul []User) {
	data := &JsonDatabase{Users: ul}
	file, _ := json.Marshal(data)
	ioutil.WriteFile("movieDatabase.json", file, 0644)
	InfoLogger.Println("Database updated.")
}

//LoadDatabase reads the json Database and assigns the values to the virtual database
func (db *Database) LoadDatabase() {
	moviesDb := ReadDatabase()
	db.userList = moviesDb.Users
	InfoLogger.Println("Database loaded.")
}

//handler reads the URL and calls designated processes
func (db *Database) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		if r.URL.Path == "/moviestracker/towatch" {
			db.ProcessMoviesToWatch(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/moviestracker/towatch/%d", &id); n == 1 {
			db.ProcessMoviesToWatchID(id, w, r)
		} else if r.URL.Path == "/moviestracker/towatch/finish" {
			db.ProcessMoviesToWatchFinished(w, r)
		} else if r.URL.Path == "/moviestracker/watched" {
			db.ProcessMoviesWatched(w, r)
		} else if n, _ := fmt.Sscanf(r.URL.Path, "/moviestracker/watched/%d", &id); n == 1 {
			db.ProcessMoviesWatchedID(id, w, r)
		} else if r.URL.Path == "/moviestracker/login" {
			db.ProcessLogin(w, r)
		} else if r.URL.Path == "/moviestracker/signup" {
			db.ProcessSignUp(w, r)
		} else if r.URL.Path == "/moviestracker/logout" {
			db.ProcessLogout(w, r)
		} else {
			http.Error(w, "\"Error 400\": Incorrect URL format.", http.StatusBadRequest)
		}
	}
}

//ProcessMoviesToWatch handles requests for the user's to watch list
func (db *Database) ProcessMoviesToWatch(w http.ResponseWriter, r *http.Request) {
	//get data from json database
	db.LoadDatabase()

	if db.activeUser != "" {
		userIndex := db.GetUserDatabaseIndex()
		movieList := db.userList[userIndex].MoviesToWatch.MovieList
		switch r.Method {
		case "POST":
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				ErrorLogger.Println(err.Error())
				return
			}
			if db.CheckDuplicateMovie(movie, movieList) {
				http.Error(w, "\"Error 409\": Movie already listed.", http.StatusConflict)
				ErrorLogger.Println("\"Error 409\": Movie already listed.")
				return
			} else {

				db.mu.Lock()
				movie.ID = db.userList[userIndex].NextId
				db.userList[userIndex].NextId++
				movieList = append(movieList, movie)
				db.userList[userIndex].MoviesToWatch.MovieList = movieList
				db.mu.Unlock()

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintf(w, "\"success\": Movie with ID#%v saved to database.\n", movie.ID)
				InfoLogger.Println(fmt.Sprintf("%v added %v to their to-watch list.", db.activeUser, movie.Title))

			}
		case "GET":
			if len(movieList) == 0 {
				http.Error(w, "\"Error 404\": No movies found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": No movies found in database.")
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				InfoLogger.Println(fmt.Sprintf("%v viewed their to-watch movie list.", db.activeUser))
				// if err := json.NewEncoder(w).Encode(movieList); err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// 	ErrorLogger.Println(err.Error())
				// 	return
				// }
				DisplayMovieList(movieList, w)
			}

		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Requires user log in.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Requires user log in.")
	}

	//save to json database
	WriteToDatabase(db.userList)

}

//ProcessMoviesToWatchID handles requests for the user's to watch list with ID
func (db *Database) ProcessMoviesToWatchID(id int, w http.ResponseWriter, r *http.Request) {
	//get data from json database
	db.LoadDatabase()

	if db.activeUser != "" {
		userIndex := db.GetUserDatabaseIndex()
		movieList := db.userList[userIndex].MoviesToWatch.MovieList
		switch r.Method {
		case "GET":
			exists, movieIndex := db.CheckIdExists(id, movieList)
			if exists {
				w.Header().Set("Content-Type", "application/json")
				InfoLogger.Println(fmt.Sprintf("%v viewed Movie#%v info.", db.activeUser, id))
				// if err := json.NewEncoder(w).Encode(movieList[movieIndex]); err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// 	ErrorLogger.Println(err.Error())
				// 	return
				// }
				DisplayMovie(movieList[movieIndex], w)
			} else {
				http.Error(w, "\"Error 404\": Movie not found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": Movie not found in database.")
				return
			}
		case "PUT":
			exists, movieIndex := db.CheckIdExists(id, movieList)
			if exists {
				var newMovie Movie
				if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					ErrorLogger.Println(err.Error())
					return
				}

				db.mu.Lock()
				newMovie.ID = movieList[movieIndex].ID
				movieList[movieIndex] = newMovie
				db.userList[userIndex].MoviesToWatch.MovieList = movieList
				db.mu.Unlock()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "{\"success\": Movie with ID#%v updated.}\n", newMovie.ID)
				InfoLogger.Println(fmt.Sprintf("%v updated Movie#%v info.", db.activeUser, id))

			} else {
				http.Error(w, "\"Error 404\": Movie not found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": Movie not found in database.")
				return
			}
		case "DELETE":
			exists := false
			db.mu.Lock()
			for j, item := range movieList {
				if id == item.ID {
					movieList = append(movieList[:j], movieList[j+1:]...)
					db.userList[userIndex].MoviesToWatch.MovieList = movieList
					exists = true
					break
				}
			}
			db.mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			if exists {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "{\"success\": Movie with ID#%v deleted.}\n", id)
				InfoLogger.Println(fmt.Sprintf("%v deleted Movie#%v from database.", db.activeUser, id))
			} else {
				http.Error(w, "\"Error 404\": Movie not found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": Movie not found in database.")
				return
			}
		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Requires user log in.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Requires user log in.")
	}
	//save to json database
	WriteToDatabase(db.userList)
}

//ProcessMoviesToWatchFinished handles requests for marking movies as finished
func (db *Database) ProcessMoviesToWatchFinished(w http.ResponseWriter, r *http.Request) {
	db.LoadDatabase()

	if db.activeUser != "" {
		userIndex := db.GetUserDatabaseIndex()
		movieList := db.userList[userIndex].MoviesToWatch.MovieList
		watchedMovieList := db.userList[userIndex].MoviesWatched.MovieList
		switch r.Method {
		case "PUT":
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				ErrorLogger.Println(err.Error())
				return
			}
			exists, movieIndex := db.CheckMovieExists(movie.Title, movieList)
			if exists {
				db.mu.Lock()
				movie.ID = movieList[movieIndex].ID
				for j, item := range movieList {
					if movie.Title == item.Title {
						movieList = append(movieList[:j], movieList[j+1:]...)
						db.userList[userIndex].MoviesToWatch.MovieList = movieList
						break
					}
				}
				watchedMovieList = append(watchedMovieList, movie)
				db.userList[userIndex].MoviesWatched.MovieList = watchedMovieList
				db.mu.Unlock()

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "\"success\": Congratulations, you've finished watching %v.\n", movie.Title)
				InfoLogger.Println(fmt.Sprintf("%v finished watching %v.", db.activeUser, movie.Title))
			} else {
				http.Error(w, "\"Error 404\": Movie not found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": Movie not found in database.")
				return
			}

		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Requires user log in.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Requires user log in.")
	}
	//save to json database
	WriteToDatabase(db.userList)
}

//ProcessMoviesWatched handles requests for user's list of watched movies
func (db *Database) ProcessMoviesWatched(w http.ResponseWriter, r *http.Request) {
	//get data from json database
	db.LoadDatabase()

	if db.activeUser != "" {
		userIndex := db.GetUserDatabaseIndex()
		movieList := db.userList[userIndex].MoviesWatched.MovieList
		switch r.Method {
		case "GET":
			if len(movieList) == 0 {
				http.Error(w, "\"Error 404\": No movies found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": No movies found in database.")
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				InfoLogger.Println(fmt.Sprintf("%v viewed the list of movies they have watched.", db.activeUser))
				// if err := json.NewEncoder(w).Encode(movieList); err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// 	ErrorLogger.Println(err.Error())
				// 	return
				// }
				DisplayMovieWatchedList(movieList, w)
			}
		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Requires user log in.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Requires user log in.")
	}
	//save to json database
	WriteToDatabase(db.userList)
}

//ProcessMoviesWatched handles requests for user's list of watched movies with ID
func (db *Database) ProcessMoviesWatchedID(id int, w http.ResponseWriter, r *http.Request) {
	//get data from json database
	db.LoadDatabase()

	if db.activeUser != "" {
		userIndex := db.GetUserDatabaseIndex()
		movieList := db.userList[userIndex].MoviesWatched.MovieList
		switch r.Method {
		case "GET":
			exists, movieIndex := db.CheckIdExists(id, movieList)
			if exists {
				w.Header().Set("Content-Type", "application/json")
				InfoLogger.Println(fmt.Sprintf("%v viewed Movie#%v info.", db.activeUser, id))
				// if err := json.NewEncoder(w).Encode(movieList[movieIndex]); err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// 	ErrorLogger.Println(err.Error())
				// 	return
				// }
				DisplayMovie(movieList[movieIndex], w)
			} else {
				http.Error(w, "\"Error 404\": Movie not found in database.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": Movie not found in database.")
				return
			}

		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Requires user log in.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Requires user log in.")
	}
	//save to json database
	WriteToDatabase(db.userList)
}

//ProcessLogin handles requests regarding user login
func (db *Database) ProcessLogin(w http.ResponseWriter, r *http.Request) {
	db.LoadDatabase()

	if db.activeUser == "" {
		switch r.Method {
		case "POST":
			var user User
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ErrorLogger.Println(err.Error())
				return
			}

			if db.CheckUserExists(user.Username) {
				if db.VerifyLogin(user) {
					db.mu.Lock()
					db.activeUser = user.Username
					db.mu.Unlock()
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, "\"success\": Log in successfull.\n")
					InfoLogger.Println(fmt.Sprintf("%v logged in successfully.", user.Username))
				} else {
					http.Error(w, "\"Error 403\": The username and password do not match our database.", http.StatusForbidden)
					ErrorLogger.Println("\"Error 403\": The username and password do not match our database.")
					return
				}
			} else {
				http.Error(w, "\"Error 404\": User not found.", http.StatusNotFound)
				ErrorLogger.Println("\"Error 404\": User not found.")
				return
			}

		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Active user already found.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Active user already found.")
	}

	WriteToDatabase(db.userList)
}

//ProcessSignUp handles requests regarding account creation
func (db *Database) ProcessSignUp(w http.ResponseWriter, r *http.Request) {
	db.LoadDatabase()

	if db.activeUser == "" {
		switch r.Method {
		case "POST":
			var user User
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ErrorLogger.Println(err.Error())
				return
			}

			if !db.CheckUserExists(user.Username) {
				db.mu.Lock()
				db.userList = append(db.userList, user)
				db.mu.Unlock()
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "\"success\": Sign up succesfull.\n")
				InfoLogger.Print(fmt.Sprintf("New account created with username: %v", user.Username))
			} else {
				http.Error(w, "\"Error 404\": Username already exists.", http.StatusConflict)
				ErrorLogger.Println("\"Error 404\": Username already exists.")
				return
			}

		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 401\": Active user found, kindly log out first.", http.StatusUnauthorized)
		ErrorLogger.Println("\"Error 401\": Active user already found, log out first.")
	}

	WriteToDatabase(db.userList)
}

//ProcessLogout handles requests for logging out
func (db *Database) ProcessLogout(w http.ResponseWriter, r *http.Request) {
	if db.activeUser != "" {
		switch r.Method {
		case "POST":
			var user User
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				ErrorLogger.Println(err.Error())
				return
			}
			if user.Username == db.activeUser {
				InfoLogger.Println(fmt.Sprintf("%v logged out successfully.", db.activeUser))
				db.activeUser = ""
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "\"success\": Logged out successfully.\n")
			} else {
				http.Error(w, "\"Error 409\": Incorrect username passed", http.StatusConflict)
			}

		default:
			http.Error(w, "\"Error 405\": Method not allowed.", http.StatusMethodNotAllowed)
			ErrorLogger.Println("\"Error 405\": Method not allowed.")
			return
		}
	} else {
		http.Error(w, "\"Error 409\": User already logged out.", http.StatusConflict)
		ErrorLogger.Println("\"Error 409\": User already logged out.")
	}
}

//CheckDuplicateMovie checks if the new movie already exists in the database and returns a boolean
func (db *Database) CheckDuplicateMovie(movie Movie, movieList []Movie) bool {
	for _, item := range movieList {
		if item.Title == movie.Title {
			return true
		}
	}
	return false
}

//CheckIdExists finds the ID in the database and returns a boolean and the index of that ID in the database
func (db *Database) CheckIdExists(id int, movieList []Movie) (bool, int) {

	var contactIndex int
	for i, item := range movieList {
		if item.ID == id {
			contactIndex = i
			return true, contactIndex
		}
	}

	return false, contactIndex
}

//CheckMovieExists finds the movie in the database and returns a boolean and the index of that movie in the database
func (db *Database) CheckMovieExists(title string, movieList []Movie) (bool, int) {

	var contactIndex int
	for i, item := range movieList {
		if item.Title == title {
			contactIndex = i
			return true, contactIndex
		}
	}

	return false, contactIndex
}

//GetUserDatabaseIndex searches the database for the username and returns its index in the User slice
func (db *Database) GetUserDatabaseIndex() int {
	var index int
	for i, item := range db.userList {
		if db.activeUser == item.Username {
			index = i
		}
	}
	return index
}

//CheckUserExists checks if the username exists in the database
func (db *Database) CheckUserExists(username string) bool {
	for _, item := range db.userList {
		if item.Username == username {
			return true
		}
	}
	return false
}

//VerifyLogin checks if the inputted username and password matches the database entries
func (db *Database) VerifyLogin(user User) bool {
	for _, item := range db.userList {
		if item.Username == user.Username && item.Password == user.Password {
			return true
		}
	}
	return false
}

//DisplayMovieWatchedList displays a list of finished movies to the user
func DisplayMovieWatchedList(movieList []Movie, w http.ResponseWriter) {

	fmt.Fprintln(w, "\n---------------------------------------")
	fmt.Fprintln(w, "\n++++++++ Movies You've Watched ++++++++")
	fmt.Fprintln(w, "")

	for _, item := range movieList {
		fmt.Fprintf(w, "| #%v\n", item.ID)
		fmt.Fprintf(w, "| %v\n", item.Title)
		fmt.Fprintf(w, "| %v/5\n", item.Rating)
		fmt.Fprintf(w, "| %v\n", item.Review)

		fmt.Fprintln(w, "__________________________________________")
	}
}

//DisplayMovieList displays a list of movies to the user
func DisplayMovieList(movieList []Movie, w http.ResponseWriter) {

	fmt.Fprintln(w, "\n---------------------------------------")
	fmt.Fprintln(w, "\n++++++++ Movies to Watch ++++++++")
	fmt.Fprintln(w, "")

	for _, item := range movieList {
		fmt.Fprintf(w, "| #%v\n", item.ID)
		fmt.Fprintf(w, "| %v\n", item.Title)
		fmt.Fprintf(w, "| %v/5\n", item.Rating)
		fmt.Fprintf(w, "| %v\n", item.Review)

		fmt.Fprintln(w, "__________________________________________")
	}
}

//DisplayMovie displays a movie to the user
func DisplayMovie(movie Movie, w http.ResponseWriter) {
	fmt.Fprintln(w, "__________________________________________")
	fmt.Fprintf(w, "| #%v\n", movie.ID)
	fmt.Fprintf(w, "| %v\n", movie.Title)
	fmt.Fprintf(w, "| %v/5\n", movie.Rating)
	fmt.Fprintf(w, "| %v\n", movie.Review)
	fmt.Fprintln(w, "__________________________________________")
}
