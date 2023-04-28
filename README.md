# perf-tests

## Notes

* Runs on port 8090
* Contains a single GET endpoint: add

## "Add" Endpoint

Takes in two numeric parameters: **val1** and **val2**

## Running

```
go run webapp.go
```

From a browser: `http://localhost:8090/add?val1=1&val2=2`

Command line: ` curl "http://localhost:8090/add?val1=1&val2=2"`
