# RestTester
simple way to test rest apis

# Usage
## json file:
add your tests to a json file of the form:
```jsonc
[
    {
        "url":"",
        "status":"", // expected http status
        "body":"" // expected json body ( must be valid json)
    },
    {...}
]
```

## cli args
use -url, -status and -body cli arguments to run a single quick test against the url

## http request

run : `restTester -server <<port>>` to run a local http server then send to a test post request to the server
for example with httpie:
```bash
http :8080/test url=https://jsonplaceholder.typicode.com/todos/1 status=200 body='{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
'
``` 
