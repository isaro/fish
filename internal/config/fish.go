package config

import (
	"github.com/isaro/fish/internal/types"
	"github.com/isaro/fish/internal/types/fish_archtypes"
)

var ()

type FishConfig struct {
	AvailableFish []types.Fish
}

func initFish(fishConfig FishConfig) {
	fishConfig.AvailableFish = []types.Fish{
		&fish_archtypes.IRLFish{Name: "Minnow", Type: "Small Fish"},
	}
}
