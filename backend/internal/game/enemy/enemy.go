package enemy

import (
	"jrpg-auto-battler/backend/internal/game/character"
	"jrpg-auto-battler/backend/internal/game/skill"
)

type Enemy struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Stats       character.Stats `json:"stats"`
	Skill       []skill.Skill   `json:"skills"`
}
