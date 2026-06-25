package job

import (
	"jrpg-auto-battler/backend/internal/game/character"
	"jrpg-auto-battler/backend/internal/game/skill"
)

type Job struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	BaseStats   character.Stats `json:"baseStats"`
	Skills      []skill.Skill   `json:"skills"`
}
