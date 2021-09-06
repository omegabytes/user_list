package service

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

type ServiceConfig struct{}

// Service implements the service.
type Service struct{}

// User represents a user record
// todo: unique UUID
type User struct {
	Email     string
	FirstName string
	LastName  string
	ZipCode   int
}

// NewService instatiates a service with given config optiions
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

// hello handles default or unexpected url paths
func hello() http.Handler {
	return http.HandlerFunc(ServeHTTP)
}

// AddUserRecord inserts a single user record into the data store.
func (s *Service) AddUserRecord() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo
	})
}

// ListUserRecord fetches a specific user record.
func (s *Service) ListUserRecord() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo
	})
}

// UpdateUserRecord updates a specific user record already present in the data store.
func (s *Service) UpdateUserRecord() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo
	})
}

// ListUserRecords lists all user records contained in the data store.
func (s *Service) ListUserRecords() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo
	})
}

// UploadCSV ingests a CSV of one or more user records.
// We assume that all CSVs have Email,FirstName,LastName,ZipCode header rows.
func (s *Service) UploadCSV() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reader := csv.NewReader(r.Body)

		lines, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		records := make([]User, 0)

		for i, line := range lines {
			fmt.Println("line: ", line)
			// skip header
			if i == 0 {
				continue
			}

			zipCode, err := strconv.Atoi(line[3])
			if err != nil {
				panic(err)
			}

			record := User{
				Email:     line[0],
				FirstName: line[1],
				LastName:  line[2],
				ZipCode:   zipCode,
			}

			records = append(records, record)
		}

		// todo: dbService.InsertRecords(records)
	})
}
