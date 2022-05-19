# DisGo
[DisGo](https://pkg.go.dev/github.com/denkylabs/disgo) is a powerful Golang library for interacting with Discord.

---

⚠️ This library is not ready for production use.

---

## Installation
### Stable version:
```sh-session
go get -u github.com/denkylabs/disgo
```

### Development version:
```sh-session
$ cd $GOPATH
$ mkdir -p src/github.com/denkylabs
$ cd src/github.com/denkylabs
$ git clone https://github.com/denkylabs/disgo.git
$ cd disgo
$ go install
```

## Example
```go
package main

import "fmt"
import "github.com/denkylabs/disgo"

func main() {
    session, err := disgo.New("Bot token")
    session.Identify.Intents = disgo.CalcIntents("Guilds", "GuildMessages")

    if err != nil {
        fmt.Println(err)
        return
    }

    err = session.Connect()

    if err != nil {
        fmt.Println(err)
        return
    }
}
```