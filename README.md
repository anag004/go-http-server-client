# HTTP server and client in go

To run the HTTP server (listens on port 80)

```
go server.go -port <port to serve on> -home <path of directory to serve> 
```

```
go client.go -nreqs <number of requests> -url <url to GET from> -print <should response body be printed>
```