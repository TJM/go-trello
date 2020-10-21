# Golang Trello API client

go-trello is a [Go](http://golang.org/) client package for accessing the [Trello](http://www.trello.com/) [API](https://developer.atlassian.com/cloud/trello/guides/rest-api/api-introduction/).

<a href="http://golang.org"><img alt="Go package" src="https://golang.org/doc/gopher/pencil/gopherhat.jpg" width="20%" /></a>
<a href="http://trello.com"><img src="https://d2k1ftgv7pobq7.cloudfront.net/meta/p/res/images/c13d1cd96a2cff30f0460a5e1860c5ea/header-logo-blue.svg" style="height: 80px; margin-bottom: 2em;"></a>

[![GoDoc](https://godoc.org/github.com/TJM/go-trello?status.png)](https://godoc.org/github.com/TJM/go-trello)
[![Travis](https://travis-ci.org/TJM/go-trello.svg?branch=master)](https://travis-ci.org/TJM/go-trello)

## Example

Prerequisites:

* Retrieve your `appKey`: <https://trello.com/app-key> (NOTE: This identifies "you" as the developer of the application)
* Retrieve your (temporary) `token`: (put the space there to prevent the link) https ://trello\.com/1/connect?key=<MYKEYFROMABOVE>&name=<MYAPPNAME>&response_type=token&scope=read,write&expiration=1day

```go
package main

import (
  "fmt"
  "log"

  "github.com/TJM/go-trello"
)

func main() {
  // New Trello Client
  appKey := "application-key"
  token := "token"
  trello, err := trello.NewAuthClient(appKey, &token)
  if err != nil {
    log.Fatal(err)
  }

  // User @trello
  user, err := trello.Member("trello")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(user.FullName)

  // @trello Boards
  boards, err := user.Boards()
  if err != nil {
    log.Fatal(err)
  }

  if len(boards) > 0 {
    board := boards[0]
    fmt.Printf("* %v (%v)\n", board.Name, board.ShortURL)

    // @trello Board Lists
    lists, err := board.Lists()
    if err != nil {
      log.Fatal(err)
    }

    for _, list := range lists {
      fmt.Println("   - ", list.Name)

      // @trello Board List Cards
      cards, _ := list.Cards()
      for _, card := range cards {
        fmt.Println("      + ", card.Name)
      }
    }
  }
}
```

prints

```console
Trello
* Bucket List (https://trello.com/b/Nl2oG77n)
   -  Goals
      +  How To Use This Board
      +  Do volunteer work
   -  Up Next
      +  Solve a Rubik’s Cube!
      +  Visit Japan
   -  Underway
      +  Improve writing skills
   -  Done!
      +  Learn to sail
```

## Acknowledgements

Forked From:

* [VojtechVitek/go-trello](https://github.com/VojtechVitek/go-trello)
* [aQaTL/go-trello](https://github.com/aQaTL/go-trello)

(previously) Influenced by:

* [fsouza/go-dockerclient](https://github.com/fsouza/go-dockerclient)
* [jeremytregunna/ruby-trello](https://github.com/jeremytregunna/ruby-trello)

## License

Licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).
