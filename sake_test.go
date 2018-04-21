package main

import (
    "testing"
    "encoding/json"
)

func TestSakeList(t *testing.T) {
    body := Get(t, Host + "/beverages?category=2")

    var data []map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "Number of sake", 1, len(data))

    id1 := Find(data, func(v map[string]interface{}) bool { return v["id"] == 90.0 })
    if id1 == nil {
        t.Fatalf("Id 90 not found in list")
    }

    Assert(t, "id", 90.0, (*id1)["id"])
    Assert(t, "name", "Junmai Shu Matsukura ", (*id1)["name"])
    Assert(t, "brewery", "Akita seishu ", (*id1)["brewery"])
    Assert(t, "percentage", 15.0, (*id1)["percentage"])
    Assert(t, "country", "Japan", (*id1)["country"])
    Assert(t, "style", "Junmai Shu", (*id1)["style"])
    Assert(t, "comment", "Smakar lite som piggelin ", (*id1)["comment"])
    Assert(t, "oscore", 4.0, (*id1)["oscore"])
    Assert(t, "sscore", 4.0, (*id1)["sscore"])
    Assert(t, "category", 2.0, (*id1)["category"])
}

func TestSakeGet(t *testing.T) {
    body := Get(t, Host + "/beverages/90")

    var data map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "id", 90.0, data["id"])
    Assert(t, "name", "Junmai Shu Matsukura ", data["name"])
    Assert(t, "brewery", "Akita seishu ", data["brewery"])
    Assert(t, "percentage", 15.0, data["percentage"])
    Assert(t, "country", "Japan", data["country"])
    Assert(t, "style", "Junmai Shu", data["style"])
    Assert(t, "comment", "Smakar lite som piggelin ", data["comment"])
    Assert(t, "oscore", 4.0, data["oscore"])
    Assert(t, "sscore", 4.0, data["sscore"])
    Assert(t, "category", 2.0, data["category"])
}
