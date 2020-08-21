package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"homepage/keys"
)

// GetNext5Events will retrieve the next 5 events for the team with the matching id
func GetNext5Events(ctx context.Context, id string) (*[]Event, error) {
	key := ctx.Value(keys.SportsDB)
	if key == nil {
		return nil, errors.New("not api key")
	}

	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/1/eventsnext.php?id=%v", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("non-200")
	}

	ret := respEvents{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	if len(ret.Events) < 1 {
		return nil, errors.New("no events returned")
	}

	return &ret.Events, nil
}

func GetTeamByName(ctx context.Context, name string) (*Team, error) {
	key := ctx.Value(keys.SportsDB)
	if key == nil {
		return nil, errors.New("not api key")
	}

	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/1/searchteams.php?t=%v", name)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("non-200")
	}

	ret := resp{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	if len(ret.Teams) < 1 {
		return nil, errors.New("no teams returned")
	}

	return &ret.Teams[0], nil
}
