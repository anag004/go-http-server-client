# HTTP server and client in go

To run the HTTP server (listens on port 80)

```
go server.go 
```

```
go client.go -nreqs <number of requests> -url <url to GET from> -print <should response body be printed>
```