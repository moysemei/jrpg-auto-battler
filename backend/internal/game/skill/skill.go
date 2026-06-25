package skill

type SkillType string

const (
	SkillTypeAttack SkillType = "attack"
	SkillTypeHeal   SkillType = "heal"
	SkillTypeBuff   SkillType = "buff"
)

type Skill struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Type            SkillType `json:"type"`
	Power           int       `json:"power"`
	Cooldown        int       `json:"cooldown"`
	CurrentCooldown int       `json:"currentCooldown"`
	Description     string    `json:"description"`
}
