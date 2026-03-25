package game

import (
	"fmt"
	"math/rand"
	"strings"
)

func (g *Game) RunCombat(enemy Enemy, input func() string) bool {
	e := enemy // copy so we don't mutate the template
	g.printf("\n⚔️ ═══════════════════════════════════════════════ ⚔️\n")
	g.printf("   %s  %s appears!\n", e.Emoji, e.Name)
	g.printf("   💬 %s\n", e.Quip)
	g.printf("   ❤️  HP: %d  |  ⚔️  ATK: %d  |  🛡️  DEF: %d\n", e.Stats.HP, e.Stats.Attack, e.Stats.Defense)
	g.printf("⚔️ ═══════════════════════════════════════════════ ⚔️\n")

	round := 1
	for e.Stats.HP > 0 && g.Player.Stats.HP > 0 {
		g.printf("\n── Round %d ──\n", round)
		g.printf("  🧍 %s: ❤️  %d/%d  |  %s %s equipped\n",
			g.Player.Name, g.Player.Stats.HP, g.Player.Stats.MaxHP,
			g.Player.Weapon.Class, g.Player.Weapon.Name)
		g.printf("  %s %s: ❤️  %d/%d\n", e.Emoji, e.Name, e.Stats.HP, e.Stats.MaxHP)
		g.println("\n  [A] Attack  [U] Use item  [F] Flee")
		g.print("  > ")

		cmd := strings.ToLower(strings.TrimSpace(input()))
		switch cmd {
		case "a":
			g.playerAttack(&e)
			if e.Stats.HP > 0 {
				g.enemyAttack(&e)
			}
		case "u":
			used := g.useItemInCombat(input)
			if !used {
				continue // don't advance round
			}
			if e.Stats.HP > 0 {
				g.enemyAttack(&e)
			}
		case "f":
			if rand.Intn(100) < 40+g.Player.Stats.Speed*3 {
				g.println("  🏃 You flee like the wind! A cowardly, sensible wind.")
				return true // alive
			}
			g.println("  🏃 You try to flee but trip over your own ego!")
			g.enemyAttack(&e)
		default:
			g.println("  🤷 Confusion isn't a combat move! (Try A, U, or F)")
			continue
		}

		if g.Player.Stats.HP <= 0 {
			g.println("\n  💀 You have been slain! Your adventure ends here.")
			g.printf("  💀 %s laughs over your crumpled form.\n", e.Name)
			return false
		}
		round++
	}

	// Victory!
	g.printf("\n  🎉🎉🎉 VICTORY! You defeated %s %s! 🎉🎉🎉\n", e.Emoji, e.Name)
	g.printf("  💰 Loot: %d 🪙  |  ✨ XP: %d\n", e.Gold, e.XP)
	g.Player.Gold += e.Gold
	g.Player.XP += e.XP
	g.checkLevelUp()
	return true
}

func (g *Game) playerAttack(e *Enemy) {
	baseDmg := g.Player.Weapon.RollDamage() + g.Player.Stats.Attack
	def := e.Stats.Defense
	dmg := baseDmg - def
	// Critical hit chance
	crit := false
	if rand.Intn(100) < 15 {
		dmg = baseDmg // ignore defense on crit
		crit = true
	}
	if dmg < 1 {
		dmg = 1
	}
	e.Stats.HP -= dmg
	if e.Stats.HP < 0 {
		e.Stats.HP = 0
	}

	if crit {
		g.printf("  💥 CRITICAL HIT! You smash %s for %d damage!\n", e.Name, dmg)
		g.println("     " + critQuip())
	} else {
		g.printf("  🗡️  You hit %s for %d damage with your %s!\n", e.Name, dmg, g.Player.Weapon.Name)
	}
}

func (g *Game) enemyAttack(e *Enemy) {
	baseDmg := e.Weapon.RollDamage() + e.Stats.Attack
	def := g.Player.Stats.Defense
	dmg := baseDmg - def

	crit := false
	if rand.Intn(100) < 10 {
		dmg = baseDmg
		crit = true
	}
	if dmg < 1 {
		dmg = 1
	}

	g.Player.Stats.HP -= dmg
	if g.Player.Stats.HP < 0 {
		g.Player.Stats.HP = 0
	}

	if crit {
		g.printf("  💥 %s lands a CRITICAL HIT on you for %d damage! Ouch!\n", e.Name, dmg)
	} else {
		g.printf("  🩸 %s hits you for %d damage with %s!\n", e.Name, dmg, e.Weapon.Name)
	}
}

func (g *Game) useItemInCombat(input func() string) bool {
	healItems := []int{}
	for i, item := range g.Player.Inventory {
		if item.Heal > 0 {
			healItems = append(healItems, i)
		}
	}
	if len(healItems) == 0 {
		g.println("  🎒 No usable items! Maybe buy some potions next time, genius.")
		return false
	}
	g.println("  🎒 Usable items:")
	for j, idx := range healItems {
		item := g.Player.Inventory[idx]
		g.printf("    %d. %s (heals %d HP)\n", j+1, item.Name, item.Heal)
	}
	g.print("  Choose (or 0 to cancel): > ")

	choice := 0
	fmt.Sscanf(strings.TrimSpace(input()), "%d", &choice)
	if choice < 1 || choice > len(healItems) {
		return false
	}
	idx := healItems[choice-1]
	item := g.Player.Inventory[idx]
	g.Player.Stats.HP += item.Heal
	if g.Player.Stats.HP > g.Player.Stats.MaxHP {
		g.Player.Stats.HP = g.Player.Stats.MaxHP
	}
	g.printf("  ✨ You use %s and recover %d HP! (now %d/%d)\n",
		item.Name, item.Heal, g.Player.Stats.HP, g.Player.Stats.MaxHP)
	// Remove item
	g.Player.Inventory = append(g.Player.Inventory[:idx], g.Player.Inventory[idx+1:]...)
	return true
}

func (g *Game) checkLevelUp() {
	needed := g.Player.Level * 50
	for g.Player.XP >= needed {
		g.Player.XP -= needed
		g.Player.Level++
		g.Player.Stats.MaxHP += 10
		g.Player.Stats.HP += 10
		if g.Player.Stats.HP > g.Player.Stats.MaxHP {
			g.Player.Stats.HP = g.Player.Stats.MaxHP
		}
		g.Player.Stats.Attack += 2
		g.Player.Stats.Defense += 1
		g.Player.Stats.Speed += 1
		g.printf("\n  🌟 LEVEL UP! You are now level %d! 🌟\n", g.Player.Level)
		g.printf("  ❤️  MaxHP +10  |  ⚔️  ATK +2  |  🛡️  DEF +1  |  💨 SPD +1\n")
		needed = g.Player.Level * 50
	}
}

func critQuip() string {
	quips := []string{
		"Right in the dignity!",
		"That's gonna leave a mark... and a story.",
		"Even the narrator winced!",
		"Somewhere, a bard just wrote that down.",
		"The enemy briefly reconsiders its life choices.",
		"*chef's kiss* Beautiful violence.",
	}
	return quips[rand.Intn(len(quips))]
}
