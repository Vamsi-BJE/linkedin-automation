package state

import (
	"encoding/json"
	"os"
)

type State struct {
	Runs int `json:"runs"`
}

func Load() *State {
	data, err := os.ReadFile("state.json")
	if err != nil {
		return &State{}
	}
	var s State
	_ = json.Unmarshal(data, &s)
	return &s
}

func Save(s *State) {
	data, _ := json.MarshalIndent(s, "", "  ")
	_ = os.WriteFile("state.json", data, 0644)
}
