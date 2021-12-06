# fizzbuzz-lbc

## installation

This project has been *Dockerised* with Docker Compose. To build the project use:

```
docker-compose up
```

## Usage
A version has been deploy via Heroku, feel free to test it:
`https://fizzbuzz-lbc.herokuapp.com`

### Fizzbuzz endpoint

#### POST `/`
```
curl -X "POST" "localhost:8080" \
     -H 'Content-Type: application/json' \
     -d '{"str2": "world", "int2": 5, "str1": "hello", "int1": 2, "limit": 10}'
```

### Stats endpoint
#### GET `/stats`
```
curl "http://localhost:8080/stats"
```
