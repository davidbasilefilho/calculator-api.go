# `calculator-api.go`

A simple calculator API, made for learning Go's `net/http` package from the standard library.
With this project I learned how to make API endpoints, middleware, error handling, and more.


## Dependencies

- Docker for portable deployment


## Build

```bash
docker build -t calculator-api . # Build using docker
docker run -p 8080:8080 calculator-api # Run using docker
```


## Example API calls

```bash
# Addition (2 + 3)
curl "http://localhost:8080/add?x=2&y=3"

# Subtraction (5 - 2)
curl "http://localhost:8080/subtract?x=5&y=2"

# Multiplication (4 * 3)
curl "http://localhost:8080/multiply?x=4&y=3"

# Division (10 / 2)
curl "http://localhost:8080/divide?x=10&y=2"

# Power (2^3 = 8)
curl "http://localhost:8080/power?x=2&y=3"

# Root (root of 8 with index 3)
curl "http://localhost:8080/root?x=8&y=3"

# Test division by zero (should return error)
curl "http://localhost:8080/divide?x=5&y=0"

# Test root with zero index (should return error)
curl "http://localhost:8080/root?x=8&y=0"
```
