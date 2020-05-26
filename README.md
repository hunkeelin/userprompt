# UserPrompt
[![CircleCI](https://circleci.com/gh/hunkeelin/userprompt.svg?style=shield)](https://circleci.com/gh/hunkeelin/userprompt)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunkeelin/userprompt)](https://goreportcard.com/report/github.com/hunkeelin/userprompt)
[![GoDoc](https://godoc.org/github.com/hunkeelin/userprompt?status.svg)](https://godoc.org/github.com/hunkeelin/userprompt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hunkeelin/userprompt/master/LICENSE)


Examples:
```
package main

import (
    "fmt"
    "github.com/hunkeelin/userprompt"
)

func main() {
    userName, err := userprompt.UserPromptWithDefault("Enter username (default: foo)", "foo", false)
    if err != nil {
        panic(err)
    }
    fmt.Println(userName)
    pass, err := userprompt.UserPrompt("Enter Okta Password", true)
    if err != nil {
        panic(err)
    }
    fmt.Println(pass)
}
```
