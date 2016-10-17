# deepgram-go
A Go wrapper for the Deepgram API - https://www.deepgram.com

[![GoDoc](https://godoc.org/github.com/agouil/deepgram-go?status.svg)](https://godoc.org/github.com/agouil/deepgram-go)

## Installation
Install this package with:
```
go get github.com/agouil/deepgram-go
```

## Documentation
There is an online reference to the documentation on the [GoDoc site](https://godoc.org/github.com/agouil/deepgram-go).

## Usage
```go
package main

import (
    "fmt"
    "github.com/agouil/deepgram-go"
)

func main() {
    deepgramClient := deepgram.Deepgram{
        ApiKey: <DEEPGRAM_API_KEY>,
    }
    result, err := deepgramClient.CheckBalance()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(
        "Balance:", result.Balance,
        "UserID:", result.UserId,
    )
}
```

## Contributing
Pull requests are welcome! :muscle:

### Testing
Run tests with:
```
go test -v
```

## Issues
To submit any issues, raise an issue through the [Issues Page](https://github.com/agouil/deepgram-go/issues)

## License
[MIT](LICENSE)

