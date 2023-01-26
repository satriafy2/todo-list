## Todo List Project

This is a mini project API to create to-do list and sub list.


### Installation

1. Make sure you have `go 1.18.10 or earlier` and `PostgreSQL 13.9 or earlier` installed in your local machine

2. Install required dependencies using command below
```
go mod tidy
```
3. This project use auto migration, so make sure you create database beforehand to avoid migration failure. Then rename `.env.example` to `.env` and modify values inside the file to configure database connection
4. Run project using the following command
```
go run main.go
```

To check the API specification via swagger you can open `{base_url}/swagger/index.html` in your browser.