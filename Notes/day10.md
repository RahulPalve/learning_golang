# net/http package
 - net package contains all necessary things for network IO
 - http is one of them
 

## http
### Serving
- First understand what is handler, any type which implements below interface is **Handler**.
```go
type Handler Interface{
    ServeHTTP(http.ResponseWriter, *http.Request)
}
```
- for simplicity you can consider handle is similar to the **view** part of MVC framework.
- there is a method 
```go
http.ListenAndServe(pattern string, handler Handler) //any type implementing Handler interface
```
- this take urlpattern (including port no) and serve to the Handler with Request type     
 There are multiple ways of url routing to handle which are

1) 
```go
http.ListenAndServe("http://localhost:8000/endpoint", handleEndpoint)
```
2)
```go
http.Handle("/endpoint", handleEndpoint) 
//here handleEndpoint is type implementing Handler interface ServeHttp method of that type will be called.
http.ListenAndServe("http://localhost:8000", nil)

```
3)
 ```go
func handleEndpointFunc(resp http.ResponseWriter, req *http.Request){
    // do some processing
}
http.HandleFunc("/endpoint", handleEndpointFunc) //here handleEndpoint is type implementing Handler interface
http.ListenAndServe("http://localhost:8000", nil)
```     

Understanding ResponseWriter and Request
- ResponseWriter (https://pkg.go.dev/net/http#ResponseWriter) is Interface. which have Write, Header and HeaderWrite methods. ResponseWriter.write([]byte) is used to write to the usser. but dont have to reteurn this cause `ServeHttp` returns nothing. In **ListenAndServe** actually **http.response** implements ResponseWriter.
 - Request is of type struct contains, method, header, body and other things that a http request have. https://golang.org/src/net/http/response.go here you can see the actual implementation of Write([]byte) method

-  Writing to user
### Other
- mux is HTTP request multiplexer, which helps in handling url part.
- []byte("Here is a string....") converts string to byte array.
- https://stackoverflow.com/questions/13255907/in-go-http-handlers-why-is-the-responsewriter-a-value-but-the-request-a-pointer
- https://cs.opensource.google/go/go/+/refs/tags/go1.17.1:src/net/http/server.go;l=1515
- https://stackoverflow.com/questions/69232941/what-original-type-is-passed-in-servehttps-http-responsewriter-interface/69233010#69233010

// The Life Of A Write is like this:
//
// Handler starts. No header has been sent. The handler can either
// write a header, or just start writing. Writing before sending a header
// sends an implicitly empty 200 OK header.
//
// If the handler didn't declare a Content-Length up front, we either
// go into chunking mode or, if the handler finishes running before
// the chunking buffer size, we compute a Content-Length and send that
// in the header instead.
//
// Likewise, if the handler didn't set a Content-Type, we sniff that
// from the initial chunk of output.
//