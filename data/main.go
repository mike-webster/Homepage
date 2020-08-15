package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mike-webster/homepage/keys"
)

// GetNext5Events will retrieve the next 5 events for the team with the matching id
func GetNext5Events(ctx context.Context, id string) interface{} {
	key := ctx.Value(keys.SportsDB)
	if key == nil {
		return errors.New("not api key")
	}

	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/1/eventsnext.php?id=%v", id)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("non-200")
	}

	return errors.New("TODO: Finish func")
}

func GetTeamByName(ctx context.Context, name string) interface{} {
	key := ctx.Value(keys.SportsDB)
	if key == nil {
		return errors.New("not api key")
	}

	url := fmt.Sprintf("https://www.thesportsdb.com/api/v1/json/1/searchteams.php?t=%v", name)
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("non-200")
	}

	type resp struct {
		Teams []struct {
			ID   string `json:"idTeam"`
			Name string `json:"strName"`
		} `json:"teams"`
	}

	ret := resp{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return err
	}

	if len(ret.Teams) < 1 {
		return nil
	}

	return &ret.Teams[0]
}
