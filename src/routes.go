package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
		"/api-v1/submissions/{submissionId}",
		SubmissionShow,
	},
}
