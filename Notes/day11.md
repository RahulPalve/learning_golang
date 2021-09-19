## Notes 
-  Pointer to pointer is userful in case we want to change pointers value, as you know in go even pointers are passed by value ref to https://stackoverflow.com/a/8769095/7707030
- in json.Marshal private fields are not detected so, json with lowercase keys wont work. In this case either write json marshal method to that type or during type def 
    ```go
    type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}
    ```
- 