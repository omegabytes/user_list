# constituent user list

## Requirements

- app can ingest CSVs of constituents
  - all CSVs have header rows Email,FirstName,LastName,ZipCode
- users should be able to view a list of all constituents
- users should be able to create individual constituent record
- users should be able to update or delete individual constituent record

## Use

Clone the project and setting up your go environment. Then, run the following from the project root:

```
go build ./...
go run cmd/main.go
```

You can send a test CSV to be uploaded to the `users/add` path from the project root via:

```
curl -X POST -H 'Content-Type: text/csv' -d @fixtures/records.csv http://localhost:3000/users/add
```

Halt the program with cmd/ctrl+C.

As submitted, this parses the entire CSV into a single line. Next steps are to instead parse each record into individual lines, perform validation on each field, and store it into a `UserRecord` struct with a generated uuid. The array of structs is used as the input to the (unimplemented) `InsertRecord` method, which is responsible for inserting user records into the data store.

## Next steps

- implement remaining CRUD handlers
- Redis datastore with mutex locks
- fully containerized application with setup and useage info
- graceful, context-aware shutdown
- unit tests, field validations, proper error handling
- metrics + logs with log levels
