package ratsit

import (
	"net/http"
	"net/url"
	"strconv"
)

func generatePersonLookupURL(apiURL string, ssn string) (q string) {
	v := url.Values{}
	v.Add("SSN", ssn)
	q = apiURL + "/personinformation?" + v.Encode()
	return
}

func generatePersonSearchURL(apiURL string, name string, location string, limit int, recordFrom int) (q string) {
	v := url.Values{}
	v.Add("who", name)
	v.Add("where", location)
	v.Add("maxNrRecords", strconv.Itoa(limit))
	v.Add("recordFrom", strconv.Itoa(recordFrom))
	q = apiURL + "/personsearch?" + v.Encode()
	return
}

func generateCompanySearchURL(apiURL string, name string, location string, limit int, recordFrom int) (q string) {
	v := url.Values{}
	v.Add("who", name)
	v.Add("where", location)
	v.Add("maxNrRecords", strconv.Itoa(limit))
	v.Add("recordFrom", strconv.Itoa(recordFrom))
	q = apiURL + "/companysearch?" + v.Encode()
	return
}

func authorizeRequest(r *http.Request, apiKey string, pkg string) {
	r.Header.Add("Authorization", apiKey)
	r.Header.Add("package", pkg)
}

func handleResponseError(r *http.Response) (err error) {
	if r.StatusCode == http.StatusInternalServerError {
		err = ErrInternalServer
	} else if r.StatusCode == http.StatusBadRequest {
		err = ErrInvalidInput
	} else if r.StatusCode == http.StatusUnauthorized {
		err = ErrInvalidCredentials
	}
	return
}
