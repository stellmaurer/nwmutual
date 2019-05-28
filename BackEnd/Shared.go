/*************************************************************
 *
 * Code written by: Stephen Ellmaurer <stellmaurer@gmail.com>
 *
 ************************************************************/

package main

import (
	"encoding/json"
	"net/http"
)

// Send response to CORS OPTIONS pre-flight request
func options(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin")
	w.WriteHeader(http.StatusOK)
}

// Send response back to client using HTTP
func sendResponse(w http.ResponseWriter, httpMethod string, response interface{}) {
	var httpStatus = http.StatusNotImplemented
	error, responseIsErrorResponse := response.(*ErrorResponse)
	if responseIsErrorResponse {
		httpStatus = error.Error.Code
	} else {
		switch httpMethod {
		case "GET":
			httpStatus = http.StatusOK
		case "POST":
			httpStatus = http.StatusCreated
		case "PUT":
			httpStatus = http.StatusOK
		case "DELETE":
			httpStatus = http.StatusOK
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(response)
}

func itemNotFound() *ErrorResponse {
	var errorResponse = new(ErrorResponse)
	var error = Error{}
	error.Code = http.StatusNotFound
	error.Message = "Item not found"
	errorResponse.Error = error
	return errorResponse
}

// SuccessResponse : A response of an HTTP request
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// Success : A response indicating a successful request
type Success struct {
	Message string `json:"message"`
}

// ErrorResponse : A response of a HTTP request
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Error : An error response from an HTTP Request
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JobSearchResult : Our own job search API response format
type JobSearchResult struct {
	Jobs           []string `json:"jobs"`
	Page           int      `json:"page"`
	TotalItemCount int      `json:"totalItemCount"`
}

// USAJobResponse : A job search result from the USAJobs API
type USAJobResponse struct {
	LanguageCode     string           `json:"LanguageCode"`
	SearchParameters SearchParameters `json:"SearchParameters"`
	SearchResult     SearchResult     `json:"SearchResult"`
}

// SearchParameters : Query parameters used in search request.
type SearchParameters struct {
}

// SearchResult : Jobs returned from the search
type SearchResult struct {
	SearchResultCount    int                `json:"SearchResultCount"`
	SearchResultCountAll int                `json:"SearchResultCountAll"`
	SearchResultItems    []SearchResultItem `json:"SearchResultItems"`
}

// SearchResultItem : An individual job object returned in the search
type SearchResultItem struct {
	MatchedObjectDescriptor MatchedObjectDescriptor `json:"MatchedObjectDescriptor"`
}

// MatchedObjectDescriptor : More info on the job
type MatchedObjectDescriptor struct {
	PositionTitle string `json:"PositionTitle"`
}
