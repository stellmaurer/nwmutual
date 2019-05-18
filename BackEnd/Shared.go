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
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
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

/*
{
    "LanguageCode": "EN",
    "SearchParameters": {},
    "SearchResult": {
        "SearchResultCount": 10,
        "SearchResultCountAll": 64,
        "SearchResultItems": [
            {
                "MatchedObjectId": "510569800",
                "MatchedObjectDescriptor": {
                    "PositionID": "ML-10292043-18-AB",
                    "PositionTitle": "Physician - Psychiatrist",
                    "PositionURI": "https://www.usajobs.gov:443/GetJob/ViewDetails/510569800",
                    "ApplyURI": [
                        "https://www.usajobs.gov:443/GetJob/ViewDetails/510569800?PostingChannelID=RESTAPI"
                    ],
                    "PositionLocationDisplay": "Milwaukee, Wisconsin",
                    "PositionLocation": [
                        {
                            "LocationName": "Milwaukee, Wisconsin",
                            "CountryCode": "United States",
                            "CountrySubDivisionCode": "Wisconsin",
                            "CityName": "Milwaukee, Wisconsin",
                            "Longitude": -87.9068451,
                            "Latitude": 43.04181
                        }
                    ],
                    "OrganizationName": "Veterans Affairs, Veterans Health Administration",
                    "DepartmentName": "Department of Veterans Affairs",
                    "JobCategory": [
                        {
                            "Name": "Medical Officer",
                            "Code": "0602"
                        }
                    ],
                    "JobGrade": [
                        {
                            "Code": "VM"
                        }
                    ],
                    "PositionSchedule": [
                        {
                            "Name": "Full-Time",
                            "Code": "1"
                        }
                    ],
                    "PositionOfferingType": [
                        {
                            "Name": "Permanent",
                            "Code": "15317"
                        }
                    ],
                    "QualificationSummary": "Basic Requirements: United States Citizenship: Non-citizens may only be appointed when it is not possible to recruit qualified citizens in accordance with VA Policy. Degree of doctor of medicine or an equivalent degree resulting from a course of education in medicine or osteopathic medicine. The degree must have been obtained from one of the schools approved by the Department of Veterans Affairs for the year in which the course of study was completed. Current, full and unrestricted license to practice medicine or surgery in a State, Territory, or Commonwealth of the United States, or in the District of Columbia. Completion of residency training, or its equivalent, approved by the Secretary of Veterans Affairs in an accredited core specialty training program leading to eligibility for board certification. Proficiency in spoken and written English. Physical requirements outlined below. Preferred Experience: 3 years of experience Reference: VA Regulations, specifically VA Handbook 5005, Part II, Appendix G-2 Physician Qualification Standard. This can be found in the local Human Resources Office. Physical Requirements: Walking, bending, stooping, and some lifting in the direct care of patients. There may also be extended periods of sitting.",
                    "PositionRemuneration": [
                        {
                            "MinimumRange": "110000.00",
                            "MaximumRange": "264000.00",
                            "RateIntervalCode": "Per Year"
                        }
                    ],
                    "PositionStartDate": "2018-09-11",
                    "PositionEndDate": "2019-09-11",
                    "PublicationStartDate": "2018-09-11",
                    "ApplicationCloseDate": "2019-09-11",
                    "PositionFormattedDescription": [
                        {
                            "Label": "Dynamic Teaser",
                            "LabelDescription": "Hit highlighting for keyword searches."
                        }
                    ],
                    "UserArea": {
                        "Details": {
                            "JobSummary": "We are looking for a board-certified or board-eligible Psychiatrist to join our psychiatric staff at the Clement J. Zablocki VA Medical Center in Milwaukee WI.",
                            "WhoMayApply": {
                                "Name": "",
                                "Code": ""
                            },
                            "LowGrade": "00",
                            "HighGrade": "00",
                            "PromotionPotential": "00",
                            "HiringPath": [
                                "public"
                            ],
                            "AgencyMarketingStatement": "OUR MISSION: To fulfill President Lincoln's promise - \"To care for him who shall have borne the battle and for his widow, and his orphan\" - by serving and honoring the men and women who are America's Veterans. How would you like to become a part of a team providing compassionate care to Veterans?",
                            "TravelCode": "0",
                            "ApplyOnlineUrl": "https://apply.usastaffing.gov/Application/Apply",
                            "DetailStatusUrl": "https://apply.usastaffing.gov/Application/ApplicationStatus"
                        },
                        "IsRadialSearch": false
                    }
                },
                "RelevanceRank": 0
            },*/
