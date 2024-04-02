package utils

import (
	"encoding/json"
	"fmt"
	"go-pokemon-api/model"
	"net/http"
	"time"
)

const URL string = "https://pokeapi.co/api/v2/pokemon"

func FetchData(pokemonId int) (*model.Pokemon, error) {
	client := &http.Client{
		Timeout: 3 * time.Minute,
	}

	url := fmt.Sprintf("%s/%d", URL, pokemonId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data model.Pokemon
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
