package battle

import (
	"fmt"

	"jrpg-auto-battler/backend/internal/game/character"
	"jrpg-auto-battler/backend/internal/game/skill"
)

func Simulate(request Request) Log {
	player := request.Player
	enemy := request.Enemy

	log := Log{
		BattleID: request.BattleID,
		Turns:    []Turn{},
	}

	turnNumber := 1

	for isAlive(player) && isAlive(enemy) {
		if player.Stats.Speed >= enemy.Stats.Speed {
			playerTurn := resolveAction(turnNumber, &player, &enemy)
			log.Turns = append(log.Turns, playerTurn)
			turnNumber++

			if !isAlive(enemy) {
				break
			}

			enemyTurn := resolveAction(turnNumber, &enemy, &player)
			log.Turns = append(log.Turns, enemyTurn)
			turnNumber++
		} else {
			enemyTurn := resolveAction(turnNumber, &enemy, &player)
			log.Turns = append(log.Turns, enemyTurn)
			turnNumber++

			if !isAlive(player) {
				break
			}

			playerTurn := resolveAction(turnNumber, &player, &enemy)
			log.Turns = append(log.Turns, playerTurn)
			turnNumber++
		}
	}

	if isAlive(player) {
		log.Result = ResultPlayerWon
	} else {
		log.Result = ResultEnemyWon
	}

	return log
}

func isAlive(c character.Character) bool {
	return c.Stats.HP > 0
}

func calculateDamage(actor character.Character, target character.Character, selectedSkill skill.Skill) int {
	damage := selectedSkill.Power + actor.Stats.Attack - target.Stats.Defense

	if damage < 1 {
		return 1
	}

	return damage
}

func resolveAction(turnNumber int, actor *character.Character, target *character.Character) Turn {
	selectedSkillIndex := chooseSkillIndex(*actor)
	selectedSkill := actor.Skills[selectedSkillIndex]

	damage := calculateDamage(*actor, *target, selectedSkill)

	target.Stats.HP -= damage
	if target.Stats.HP < 0 {
		target.Stats.HP = 0
	}

	actor.Skills[selectedSkillIndex].CurrentCooldown = selectedSkill.Cooldown

	tickCooldowns(actor)

	message := fmt.Sprintf(
		"%s used %s on %s and dealt %d damage.",
		actor.Name,
		selectedSkill.Name,
		target.Name,
		damage,
	)

	return Turn{
		Turn:          turnNumber,
		ActorID:       actor.ID,
		ActorName:     actor.Name,
		TargetID:      target.ID,
		TargetName:    target.Name,
		SkillID:       selectedSkill.ID,
		SkillName:     selectedSkill.Name,
		Damage:        damage,
		Healing:       0,
		TargetHPAfter: target.Stats.HP,
		Message:       message,
	}
}

func chooseSkillIndex(c character.Character) int {
	for i, currentSkill := range c.Skills {
		if currentSkill.CurrentCooldown == 0 {
			return i
		}
	}

	return 0
}

func tickCooldowns(c *character.Character) {
	for i := range c.Skills {
		if c.Skills[i].CurrentCooldown > 0 {
			c.Skills[i].CurrentCooldown--
		}
	}
}
