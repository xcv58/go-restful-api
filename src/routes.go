package main

import "net/http"

// Route struct describes a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains some routes
type Routes []Route

var routes = Routes{
	Route{
		"Usage",
		"GET",
		"/api-v1/",
		Usage,
	},
	Route{
		"SubmitSubmission",
		"POST",
		"/api-v1/submissions",
		SubmissionCreate,
	},
	Route{
		"ListSubmissions",
		"GET",
		"/api-v1/submissions/{submissionID}",
		SubmissionShow,
	},
}
