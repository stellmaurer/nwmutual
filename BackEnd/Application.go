/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/

package main

import "net/http"
import "github.com/gorilla/mux"

func main() {

	r := mux.NewRouter().StrictSlash(true)

	go r.HandleFunc("/jobs", getJobs).Methods("OPTIONS")
	go r.HandleFunc("/jobs", getJobs).Methods("GET")

	http.ListenAndServe(":5000", r)
}
