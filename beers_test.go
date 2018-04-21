package main

import (
    "testing"
    "encoding/json"
)

func TestBeersList(t *testing.T) {
    body := Get(t, Host + "/beverages?category=0")

    var data []map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "Number of beers", 82, len(data))

    id1 := Find(data, func(v map[string]interface{}) bool { return v["id"] == 1.0 })
    if id1 == nil {
        t.Fatalf("Id 1 not found in list")
    }

    Assert(t, "id", float64(1), (*id1)["id"])
    Assert(t, "name", "Ale", (*id1)["name"])
    Assert(t, "brewery", "Flying Brewery", (*id1)["brewery"])
    Assert(t, "percentage", 7.2, (*id1)["percentage"])
    Assert(t, "country", "Sweden", (*id1)["country"])
    Assert(t, "style", "Ale", (*id1)["style"])
    Assert(t, "comment", nil, (*id1)["comment"])
    Assert(t, "oscore", float64(3), (*id1)["oscore"])
    Assert(t, "sscore", float64(4), (*id1)["sscore"])
    Assert(t, "category", float64(0), (*id1)["category"])
}

func TestBeersGet(t *testing.T) {
    body := Get(t, Host + "/beverages/1")

    var data map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "id", float64(1), data["id"])
    Assert(t, "name", "Ale", data["name"])
    Assert(t, "brewery", "Flying Brewery", data["brewery"])
    Assert(t, "percentage", 7.2, data["percentage"])
    Assert(t, "country", "Sweden", data["country"])
    Assert(t, "style", "Ale", data["style"])
    Assert(t, "comment", nil, data["comment"])
    Assert(t, "oscore", float64(3), data["oscore"])
    Assert(t, "sscore", float64(4), data["sscore"])
    Assert(t, "category", float64(0), data["category"])
}
