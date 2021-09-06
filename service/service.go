package service

import (
	"fmt"
	"net/http"
)

type ServiceConfig struct{}

// Service implements the service.
type Service struct{}

func NewService(config ServiceConfig) Service {
	return Service{}
}

func (s *Service) NewServiceHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", hello())
	mux.Handle("/user/add", s.AddUserRecord())
	mux.Handle("/user/list", s.ListUserRecord())
	mux.Handle("/user/update", s.UpdateUserRecord())
	mux.Handle("/users/add", s.UploadCSV())
	mux.Handle("/users/list", s.ListUserRecord())
	return mux
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!\n", r.URL.Path[1:])
}

func hello() http.Handler {
	return http.HandlerFunc(ServeHTTP)
}

// AddUserRecord inserts a single user record into the data store.
func (s *Service) AddUserRecord() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AddUserRecord")
	})
}

// ListUserRecord fetches a specific user record.
func (s *Service) ListUserRecord() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ListUserRecord")
	})
}

// UpdateUserRecord updates a specific user record already present in the data store.
func (s *Service) UpdateUserRecord() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UpdateUserRecord")
	})
}

// ListUserRecords lists all user records contained in the data store.
func (s *Service) ListUserRecords() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ListUserRecords")
	})
}

// UploadCSV ingests a CSV of one or more user records.
// We assume that all CSVs have Email,FirstName,LastName,ZipCode header rows.
func (s *Service) UploadCSV() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UploadCSV")
	})
}
