# detect-arhive-type

## About package

This package provide function for determition file's type using [file signature](https://en.wikipedia.org/wiki/File_format#Magic_number).  
Current version: 1.1

### Installation

`go get https://github.com/nagayev/detect-arhive-type`

### Usage

It's very simple.

```go
package main

import(
    "fmt"
	"github.com/nagayev/detect-arhive-type"
)

func main(){

	type,err:=dat.getTypeString("test_files/test.zip")
	if err!=nil{
		// Handle your error
    }
    fmt.Println(type) //zip
}
```

### Supported formats:

- tar.gz
- zip
- 7z

## Contributing

We are always open for your issues and PRs.

## Autor and license

Author and current mainter: Marat Nagayev  
Thanks to: @gernest for test_files and golang test template  
License: MIT. See License file for details.
