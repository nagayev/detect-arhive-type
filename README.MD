# detect-arhive-type

## About package

This package provide function for determition file's type using [file signature](https://en.wikipedia.org/wiki/File_format#Magic_number).

### Installation

`go get https://github.com/nagayev/detect-arhive-type`

### Usage

It's very simple.

```go
package main

import(
	"github.com/nagayev/detect-arhive-type"
)

func main(){

	type,err:=dat.getType("test_files/test.zip")
	if err!=nil{
		// Handle your error
	}
}
```

### Supported formats:

- tar.gz
- zip

## Contributing

We are always open for your issues and PRs.

## Autor and license

Author and current mainter: Marat Nagayev  
License: MIT. See License file for details.
