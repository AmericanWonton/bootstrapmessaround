package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Name            string `json:"Name"`
	Age             int    `json:"Age"`
	FavoriteNumbers []int  `json:"FavoriteNumbers"`
}

type AllData struct {
	ThePersons    []Person `json:"ThePerson"`
	SpecialString string   `json:"SpecialString"`
}

var allTheData AllData //All the data passed to index

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

//var templateDelims = []string{"{{%", "%}}"}

func init() {
	fillTestValues() //Fill the special values within
	/* Need to parse the glob with the 'delims' function to change our comments */
	//template1 = template.Must(template1.New("ourtemp").Delims("<left>", "<right>").ParseGlob("./static/templates/*"))
	//template1.Delims("<left>", "<right>")
	//template1.Delims(templateDelims[0], templateDelims[1])
	//template1.Delims("[[", "]]")
	template1 = template.Must(template.ParseGlob("./static/templates/*"))
	//tmpl, err := template.New("index.html").Delims("[[", "]]").ParseFiles("static/templates/index.html", "templates/standTemp.gohtml")
	/*
		indexTmpl := template.New("index.html").Delims("[[", "]]")
		indexTmpl, thErr := indexTmpl.ParseFiles("./static/templates/index.html")
		if thErr != nil {
			panic(thErr)
		}
		template1 = indexTmpl
	*/
	/*
		tmpl, err := template.New("index.html").Delims("[[", "]]").ParseGlob("./static/templates/*")
		template1 = tmpl
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	*/
}

//Serves our index page
func index(w http.ResponseWriter, r *http.Request) {
	passAllData := allTheData
	/* Execute template, handle error */
	//err2 := template1.Delims("[[", "]]").ExecuteTemplate(w, "index.html", nil)
	err1 := template1.ExecuteTemplate(w, "index.html", passAllData)
	HandleError(w, err1)
}

//Serves a test get
func testGet(w http.ResponseWriter, r *http.Request) {
	//Send a response back to Ajax after session is made
	theSuccMessage := "Hello from the test func you called to"

	fmt.Fprint(w, theSuccMessage)
}

//Serves ANOTHER test get
func testGetTheSecond(w http.ResponseWriter, r *http.Request) {
	type ReturnMessage struct {
		TheErr     string  `json:"TheErr"`
		ResultMsg  string  `json:"ResultMsg"`
		SuccOrFail int     `json:"SuccOrFail"`
		AllData    AllData `json:"AllData"`
	}
	theReturnMessage := ReturnMessage{
		TheErr:     "No Error",
		ResultMsg:  "Good Result",
		SuccOrFail: 0,
		AllData:    allTheData,
	}

	/* Send the response back to Ajax */
	theJSONMessage, err := json.Marshal(theReturnMessage)
	//Send the response back
	if err != nil {
		errIs := "Error formatting JSON for return in createUser: " + err.Error()
		fmt.Println(errIs)
	}
	fmt.Fprint(w, string(theJSONMessage))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	http.Handle("/favicon.ico", http.NotFoundHandler()) //For missing FavIcon
	myRouter.HandleFunc("/", index)
	//Handles our test api calls
	myRouter.HandleFunc("/testGet", testGet).Methods("GET")                   //Handles test api call
	myRouter.HandleFunc("/testGetTheSecond", testGetTheSecond).Methods("GET") //Handles test api call
	//Serve our static files
	myRouter.Handle("/", http.FileServer(http.Dir("./static")))
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	fmt.Printf("Hey, we're starting this example\n")
	//Handle Requests
	handleRequests()
}

// Handle Errors passing templates
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

/* This function simulates pulling data from a database we can pass to our
Angular JS to fill */
func fillTestValues() {
	person1 := Person{
		Name:            "Test Person 1",
		Age:             100,
		FavoriteNumbers: []int{1, 2, 3},
	}
	person2 := Person{
		Name:            "Test Person 2",
		Age:             35,
		FavoriteNumbers: []int{4, 5, 6},
	}
	person3 := Person{
		Name:            "Test Person 3",
		Age:             1000,
		FavoriteNumbers: []int{7, 8, 9},
	}
	person4 := Person{
		Name:            "Test Person 4",
		Age:             5,
		FavoriteNumbers: []int{10, 11, 12},
	}

	allTheData = AllData{
		ThePersons:    []Person{person1, person2, person3, person4},
		SpecialString: "This is the special string",
	}
}
