package main

import (
    "fmt"
    "os"
    "testing"
    "io/ioutil"
    "net/http"
)

func Get(t *testing.T, uri string) []byte {
    res, err := http.Get(uri)
    if err != nil {
        t.Fatalf("Could not connect: %s", err.Error())
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Fatalf("Error reading body: %s", err.Error())
    }
    return body
}

func Find(data []map[string]interface{}, f func(map[string]interface{}) bool) *map[string]interface{} {
    for _, val := range data {
        if f(val) {
            return &val
        }
    }
    return nil
}

func Assert(t *testing.T, what string, expected interface{}, was interface{}) {
    if expected != was {
        t.Fatalf("Wrong %s! Expected `%s` but was `%s`", what, expected, was)
    }
}

var Host = fmt.Sprintf("%s/api", os.Getenv("test_target"))
var Jwt = os.Getenv("test_jwt")

func TestSetup(t *testing.T) {
    if Host == "" {
        t.Fatalf("test_target not set")
    }
    if Jwt == "" {
        t.Fatalf("test_jwt not set")
    }
}

func main(){}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}
