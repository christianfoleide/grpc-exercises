### Playing with grpc :man_technologist:

The API exposes a single endpoint, creating a gRPC client which dials the gRPC server, hopefully producing a result


**From top-level directory:**

```bash
go run main.go
```

**then, in a new terminal**

```bash
cd server && go run main.go
```
From the terminal where you ran the httpserver, results from the server-side stream will be printed.

**You can also use curl:**

```bash
curl 'http://localhost:8080/currencies?base=<base>&target=<target>'
```
which performs a unary request.
