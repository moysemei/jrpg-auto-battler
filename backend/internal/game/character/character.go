package character

import "jrpg-auto-battler/backend/internal/game/skill"

type Side string

const (
	SidePlayer Side = "player"
	SideEnemy  Side = "enemy"
)

type Character struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Side   Side          `json:"side"`
	Stats  Stats         `json:"stats"`
	Skills []skill.Skill `json:"skills"`
}
