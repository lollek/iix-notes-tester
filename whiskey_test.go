package main

import (
    "testing"
    "encoding/json"
)

func TestWhiskeyList(t *testing.T) {
    body := Get(t, Host + "/beverages?category=3")

    var data []map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "Number of whiskey", 3, len(data))

    id1 := Find(data, func(v map[string]interface{}) bool { return v["id"] == 87.0 })
    if id1 == nil {
        t.Fatalf("Id 87 not found in list")
    }

    Assert(t, "id", 87.0, (*id1)["id"])
    Assert(t, "name", "16 Years", (*id1)["name"])
    Assert(t, "brewery", "Lagavulin", (*id1)["brewery"])
    Assert(t, "percentage", 43.0, (*id1)["percentage"])
    Assert(t, "country", "Scotland", (*id1)["country"])
    Assert(t, "style", nil, (*id1)["style"])
    Assert(t, "comment", nil, (*id1)["comment"])
    Assert(t, "oscore", 4.0, (*id1)["oscore"])
    Assert(t, "sscore", nil, (*id1)["sscore"])
    Assert(t, "category", 3.0, (*id1)["category"])
}

func TestWhiskeyGet(t *testing.T) {
    body := Get(t, Host + "/beverages/87")

    var data map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "id", 87.0, data["id"])
    Assert(t, "name", "16 Years", data["name"])
    Assert(t, "brewery", "Lagavulin", data["brewery"])
    Assert(t, "percentage", 43.0, data["percentage"])
    Assert(t, "country", "Scotland", data["country"])
    Assert(t, "style", nil, data["style"])
    Assert(t, "comment", nil, data["comment"])
    Assert(t, "oscore", 4.0, data["oscore"])
    Assert(t, "sscore", nil, data["sscore"])
    Assert(t, "category", 3.0, data["category"])
}
