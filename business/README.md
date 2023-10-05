
# Ecommerce Go-API

This is an API Rest for an ecommerce, developed by Giovanni Borgogno, Felipe Cañas, Maximo Lucero Ruiz, Santiago Quesada.

### Steps:

1) Create a .env file using the .env.example as an example.

2) Create a utils/initializers/test.env using the test.env.example to use in testing.

3) Install dependencies:
```bash
go mod tidy
```
4) Execute API:
```bash
go run main.go
```

### Testing

Test API:
```bash
go test ./... -v
```