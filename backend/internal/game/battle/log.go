package battle

import "jrpg-auto-battler/backend/internal/game/character"

type Result string

const (
	ResultPlayerWon Result = "player_won"
	ResultEnemyWon  Result = "enemy_won"
)

type Request struct {
	BattleID string              `json:"battleId`
	Seed     int64               `json:"seed"`
	Player   character.Character `json:"player"`
	Enemy    character.Character `json:"enemy"`
}

type Log struct {
	BattleID string `json:"battleId"`
	Result   Result `json:"result"`
	Turns    []Turn `json:"turns"`
}

type Turn struct {
	Turn          int    `json:"turn"`
	ActorID       string `json:"actorId"`
	ActorName     string `json:"actorName"`
	TargetID      string `json:"targetId"`
	TargetName    string `json:"targetName"`
	SkillID       string `json:"skillId"`
	SkillName     string `json:"skillName"`
	Damage        int    `json:"damage"`
	Healing       int    `json:"healing"`
	TargetHPAfter int    `json:"targetHpAfter"`
	Message       string `json:"message"`
}
