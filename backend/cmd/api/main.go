package main

import (
	"encoding/json"
	"fmt"
	"log"

	"jrpg-auto-battler/backend/internal/game/battle"
	"jrpg-auto-battler/backend/internal/game/character"
	"jrpg-auto-battler/backend/internal/game/skill"
)

func main() {
	player := character.Character{
		ID:   "player-1",
		Name: "Void Mage",
		Side: character.SidePlayer,
		Stats: character.Stats{
			MaxHP:   30,
			HP:      30,
			Attack:  6,
			Defense: 2,
			Speed:   7,
		},
		Skills: []skill.Skill{
			{
				ID:              "firebolt",
				Name:            "Firebolt",
				Type:            skill.SkillTypeAttack,
				Power:           8,
				Cooldown:        2,
				CurrentCooldown: 0,
				Description:     "A small burst of fire damage.",
			},
			{
				ID:              "staff-hit",
				Name:            "Staff Hit",
				Type:            skill.SkillTypeAttack,
				Power:           3,
				Cooldown:        0,
				CurrentCooldown: 0,
				Description:     "A basic staff attack.",
			},
		},
	}

	enemy := character.Character{
		ID:   "enemy-1",
		Name: "Goblin",
		Side: character.SideEnemy,
		Stats: character.Stats{
			MaxHP:   24,
			HP:      24,
			Attack:  4,
			Defense: 1,
			Speed:   5,
		},
		Skills: []skill.Skill{
			{
				ID:              "slash",
				Name:            "Slash",
				Type:            skill.SkillTypeAttack,
				Power:           4,
				Cooldown:        1,
				CurrentCooldown: 0,
				Description:     "A basic melee slash.",
			},
		},
	}

	request := battle.Request{
		BattleID: "battle-1",
		Seed:     123,
		Player:   player,
		Enemy:    enemy,
	}

	result := battle.Simulate(request)

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
