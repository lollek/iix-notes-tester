package main

import (
    "testing"
    "encoding/json"
)

func TestNotesList(t *testing.T) {
    body := Get(t, Host + "/notes")

    var data []map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    if len(data) != 73 {
        t.Fatalf("Wrong amount in result, expected 73, was %d", len(data))
    }

    id1 := Find(data, func(v map[string]interface{}) bool { return v["id"] == 1.0 })
    if id1 == nil {
        t.Fatalf("Id 1 not found in list")
    }

    Assert(t, "id", float64(1), (*id1)["id"])
    Assert(t, "date", "2013-04-11", (*id1)["date"])
    Assert(t, "title", "Installing Spotify on Linux", (*id1)["title"])
}

func TestNotesGet(t *testing.T) {
    body := Get(t, Host + "/notes/1")

    var data map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }
}
