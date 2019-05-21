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

func getJobs(w http.ResponseWriter, r *http.Request) {
	resultsPerPage := r.FormValue("resultsPerPage")
	city := r.FormValue("city")
	state := r.FormValue("state")
	sendResponse(w, "GET", getJobsHelper(resultsPerPage, city, state))
}
func getJobsHelper(resultsPerPage string, city string, state string) interface{} {

	request, urlError := createUSAJobsHTTPGetRequest(createQueryURL(resultsPerPage, city, state))
	if urlError != nil {
		return urlError
	}

	resp, getRequestError := sendUSAJobsHTTPGetRequest(request)
	if getRequestError != nil {
		return getRequestError
	}

	// Convert to Go object
	defer resp.Body.Close()
	jobs, decodeJobsError := decodeJobs(resp)
	if decodeJobsError != nil {
		return decodeJobsError
	}

	if jobs.SearchResult.SearchResultCount == 0 {
		return itemNotFound()
	}

	// Format response
	var successResponse = SuccessResponse{}
	successResponse.Data = transformUSAJobAPIResultToOurOwnAPIResult(jobs, 1)
	return successResponse
}

func sendUSAJobsHTTPGetRequest(request *http.Request) (*http.Response, *ErrorResponse) {
	client := &http.Client{}
	resp, _ := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		var errorResponse = new(ErrorResponse)
		var error = Error{}
		error.Code = http.StatusInternalServerError
		error.Message = "Issue getting jobs from USAJobs server."
		return nil, errorResponse
	}
	return resp, nil
}

func createUSAJobsHTTPGetRequest(url string) (*http.Request, *ErrorResponse) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		var errorResponse = new(ErrorResponse)
		var error = Error{}
		error.Code = http.StatusBadRequest
		error.Message = "Query parameters are invalid."
		return nil, errorResponse
	}
	req.Header.Set("Authorization-Key", "Xv2Ytw9hK/5C6WitRWqUjje7n8/RDqZ8lOSElegLrEo=")
	req.Header.Set("Host", "data.usajobs.gov")
	req.Header.Set("User-Agent", "nhweller37@yahoo.com")
	return req, nil
}

func createQueryURL(resultsPerPage string, city string, state string) string {
	var url = "https://data.usajobs.gov/api/search?" +
		"&LocationName=" + city + ",%20" + state +
		"&ResultsPerPage=" + resultsPerPage +
		"&SortField=PositionTitle"
	return url
}

func transformUSAJobAPIResultToOurOwnAPIResult(usaJobAPIResponse *USAJobResponse, pageNumber int) JobSearchResult {
	var searchResultItems = usaJobAPIResponse.SearchResult.SearchResultItems
	var jobSearchResult JobSearchResult
	var a = make([]string, len(searchResultItems))
	for i := 0; i < len(searchResultItems); i++ {
		a[i] = searchResultItems[i].MatchedObjectDescriptor.PositionTitle
	}
	jobSearchResult.Jobs = a
	jobSearchResult.TotalItemCount = usaJobAPIResponse.SearchResult.SearchResultCountAll
	jobSearchResult.Page = pageNumber
	return jobSearchResult
}

// Conversion of JSON objects to Go structs
func decodeJobs(resp *http.Response) (*USAJobResponse, *ErrorResponse) {
	var errorResponse = new(ErrorResponse)
	var USAJobResponse = new(USAJobResponse)
	var jsonErr = json.NewDecoder(resp.Body).Decode(USAJobResponse)

	if jsonErr != nil {
		var error = Error{}
		error.Code = http.StatusInternalServerError
		error.Message = "getJobsHelper function: UnmarshalListOfMaps error. " + jsonErr.Error()
		return nil, errorResponse
	}
	return USAJobResponse, nil
}
