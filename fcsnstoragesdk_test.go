package fcsnstoragesdk

import (
	"context"
	"testing"
	"time"
)

func SetChart(client *Client, uid, chart string) (time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.SetChart(ctx, uid, chart)
}

func GetChart(client *Client, uid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.GetChart(ctx, uid)
}

func Subscribe(client *Client, source_uid, target_uid string) (time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Subscribe(ctx, source_uid, target_uid)
}

func Unsubscribe(client *Client, source_uid, target_uid string) (time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Unsubscribe(ctx, source_uid, target_uid)
}

func TestSetChart(t *testing.T) {
	conn, cli, err := NewConnectionAndClient("localhost:9090", true, false)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer conn.Close()

	uid := "337525026058144928"
	chart := `[ { "id": "0", "rels": { "father": "49b952a3-f42b-4807-b216-bb01364ee444", "mother": "57c348b3-c987-4cce-b48b-e48e7738acb1", "spouses": [ "4a8a329e-65c2-4567-a8e9-4c55ff3a4367" ], "children": [ "2582cfcb-248c-4fc0-8e67-d819b7dc3e6b" ] }, "data": { "firstname": "Мераб", "lastname": "Мамардашвили", "birthday": "", "deathday": "", "avatar": "", "gender": "M", "birthlocation": "Гори" } }, { "id": "49b952a3-f42b-4807-b216-bb01364ee444", "data": { "gender": "M", "firstname": "Константин", "lastname": "Мамардашвили", "birthday": null, "deathday": null, "birthlocation": null }, "rels": { "children": [ "0" ], "spouses": [ "57c348b3-c987-4cce-b48b-e48e7738acb1" ] } }, { "id": "57c348b3-c987-4cce-b48b-e48e7738acb1", "data": { "gender": "F", "firstname": "Ксения", "lastname": "Гарсеванишвили", "birthday": null, "deathday": null, "birthlocation": null }, "rels": { "children": [ "0" ], "spouses": [ "49b952a3-f42b-4807-b216-bb01364ee444" ] } }, { "id": "4a8a329e-65c2-4567-a8e9-4c55ff3a4367", "data": { "gender": "F", "firstname": "Нина", "lastname": "Мордасова", "birthday": null, "deathday": null, "birthlocation": "Москва" }, "rels": { "spouses": [ "0" ], "children": [ "2582cfcb-248c-4fc0-8e67-d819b7dc3e6b" ] } }, { "id": "2582cfcb-248c-4fc0-8e67-d819b7dc3e6b", "data": { "gender": "F", "firstname": "Елена", "lastname": "Мамардашвили", "birthday": null, "deathday": null, "birthlocation": "Москва" }, "rels": { "father": "0", "mother": "4a8a329e-65c2-4567-a8e9-4c55ff3a4367" } } ]`

	r, err := SetChart(cli, uid, chart)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("chart has been updated: %s.\n", r.Local())
}

func TestGetChart(t *testing.T) {
	conn, cli, err := NewConnectionAndClient("localhost:9090", true, false)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer conn.Close()

	uid := "337525026058144928"

	chart, err := GetChart(cli, uid)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("chart is: %s.\n", chart)
}

/*
func TestSubscribe(t *testing.T) {
	conn, cli, err := makeConnection("localhost:9090", true)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer conn.Close()

	uid1 := "337525026058144928"
	uid2 := "337525026058144929"

	timestamp, err := Subscribe(cli, uid1, uid2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("subscribed at: %s.\n", timestamp.Local())
}

func TestUnsubscribe(t *testing.T) {
	conn, cli, err := makeConnection("localhost:9090", true)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer conn.Close()

	uid1 := "337525026058144928"
	uid2 := "337525026058144929"

	timestamp, err := Unsubscribe(cli, uid1, uid2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("unsubscribed at: %s.\n", timestamp.Local())
}
*/
