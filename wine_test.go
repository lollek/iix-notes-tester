package main

import (
    "testing"
    "encoding/json"
)

func TestWineList(t *testing.T) {
    body := Get(t, Host + "/beverages?category=1")

    var data []map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "Number of wine", 1, len(data))

    id1 := Find(data, func(v map[string]interface{}) bool { return v["id"] == 99.0 })
    if id1 == nil {
        t.Fatalf("Id 99 not found in list")
    }

    Assert(t, "id", 99.0, (*id1)["id"])
    Assert(t, "name", "Chateau Ste Michelle", (*id1)["name"])
    Assert(t, "brewery", "Ste Michelle", (*id1)["brewery"])
    Assert(t, "percentage", 14.0, (*id1)["percentage"])
    Assert(t, "country", "USA", (*id1)["country"])
    Assert(t, "style", "Chardonnay", (*id1)["style"])
    Assert(t, "comment", nil, (*id1)["comment"])
    Assert(t, "oscore", 4.0, (*id1)["oscore"])
    Assert(t, "sscore", 4.0, (*id1)["sscore"])
    Assert(t, "category", 1.0, (*id1)["category"])
}

func TestWineGet(t *testing.T) {
    body := Get(t, Host + "/beverages/99")

    var data map[string]interface{}
    err := json.Unmarshal(body, &data)
    if err != nil {
        t.Fatalf("Failed to unmarshal body: %s", err.Error())
    }

    Assert(t, "id", 99.0, data["id"])
    Assert(t, "name", "Chateau Ste Michelle", data["name"])
    Assert(t, "brewery", "Ste Michelle", data["brewery"])
    Assert(t, "percentage", 14.0, data["percentage"])
    Assert(t, "country", "USA", data["country"])
    Assert(t, "style", "Chardonnay", data["style"])
    Assert(t, "comment", nil, data["comment"])
    Assert(t, "oscore", 4.0, data["oscore"])
    Assert(t, "sscore", 4.0, data["sscore"])
    Assert(t, "category", 1.0, data["category"])
}
