# Notes


## Structure

- package should have a folder named as it is
- module can have multiple packages module > package
- `go build` to build cache in GOPATH for created packages
- mutiple files with same package name can be made
- basically structure is 
- `go install . OR go install <module_name>` compile entire project.
- In order to functions to be exported it first letter should be Capital.


```
module
-- main.go
-- /custom_package
---- package_file1.go (this files should have `package custom_package`
---- package_file2.go
```
