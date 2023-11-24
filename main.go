package main

import (
	"fmt"
	"net/http"
)

//GLOBAL VARIABLES###########################
var userslice = []string

//ENDPOINT FUNCTIONS#########################

//Basic endpoint saying gday
func helloHandler(w http.ResponseWriter, r *http.Request) 
{
	fmt.Fprintf(w, "Hello, world!\n")
}

//Endpoint taking user input "userid", appending it to slice "userslice" and saying g'day
func dyn-userHandler(w http.ResponseWriter, r *http.Request)
{
	userid := r.URL.Query().Get("userid")

	if user != ""
	{
		fmt.Fprintf(w, "Hello user %s",userid)
		userslice = append(userslice,user)
	}
	else
	{
		fmt.Fprintf(w, "Hello Stranger!")
	}
}

//Endpoint checking if user is in the slice, and if not - sending a 403
func check-user-idHandler(w http.ResponseWriter, r *http.Request)
{
	
}

//MAIN############################

func main() {

	http.HandleFunc("/check-userid",dyn-userHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/dyn-user", dyn-userHandler)

	port := ":8081"

	fmt.Printf("Server listening on http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

}

// Graceful exit needed that shuts down the listener/kills the process
