package game

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

type Game struct {
	Player  *Player
	World   map[string]*Location
	inputFn func() string
	output  io.Writer
}

func New(playerName string) *Game {
	p := &Player{
		Name:  playerName,
		Stats: Stats{MaxHP: 50, HP: 50, Attack: 3, Defense: 2, Speed: 5},
		Gold:  25,
		XP:    0,
		Level: 1,
		Weapon: Weapon{
			Name:   "Soggy Twig",
			Class:  ClassDagger,
			MinDmg: 1,
			MaxDmg: 3,
			Quip:   "It's barely a weapon. More of a strong opinion.",
		},
		Inventory: []Item{
			{Name: "🧪 Health Potion", Kind: "potion", Price: 15, Heal: 25, Quip: "Starter freebie!"},
			{Name: "🍖 Suspicious Meat", Kind: "food", Price: 8, Heal: 15, Quip: "Was in your pocket when you woke up."},
		},
		Location: "Bumbleford",
	}

	return &Game{
		Player: p,
		World:  BuildWorld(),
		output: os.Stdout,
	}
}

// RunWithWriter runs the game with a custom output writer (e.g., for SSH sessions)
func (g *Game) RunWithWriter(reader *bufio.Reader, writer io.Writer) {
	g.output = writer
	g.Run(reader)
}

// Helper methods for printing to the game's output writer
func (g *Game) printf(format string, a ...interface{}) {
	fmt.Fprintf(g.output, format, a...)
}

func (g *Game) println(a ...interface{}) {
	fmt.Fprintln(g.output, a...)
}

func (g *Game) print(a ...interface{}) {
	fmt.Fprint(g.output, a...)
}

func (g *Game) Run(reader *bufio.Reader) {
	inputFn := func() string {
		line, _ := reader.ReadString('\n')
		return line
	}
	g.inputFn = inputFn

	g.printf( "\n🌅 Welcome, %s! You awaken in Bumbleford with a headache,\n", g.Player.Name)
	g.println( "   a soggy twig, and absolutely no idea how you got here.")
	g.println( "   Legend says Lord Hillmord sits on the Throne of Mild Inconvenience")
	g.println( "   deep beneath Skull Plateau. Someone should probably do something about that.")
	g.println( "   That someone, unfortunately, is you.")
	g.println()

	g.describeLocation()

	for {
		g.print("\n🎮 > ")
		cmd := strings.ToLower(strings.TrimSpace(inputFn()))

		switch {
		case cmd == "help" || cmd == "h" || cmd == "?":
			g.showHelp()
		case cmd == "look" || cmd == "l":
			g.describeLocation()
		case cmd == "north" || cmd == "n":
			g.move(North)
		case cmd == "south" || cmd == "s":
			g.move(South)
		case cmd == "east" || cmd == "e":
			g.move(East)
		case cmd == "west" || cmd == "w":
			g.move(West)
		case cmd == "fight" || cmd == "f":
			g.seekFight(inputFn)
		case cmd == "market" || cmd == "m":
			g.enterMarket(inputFn)
		case cmd == "talk" || cmd == "t":
			g.talkToNPC()
		case cmd == "inventory" || cmd == "i":
			g.showInventory()
		case cmd == "stats" || cmd == "st":
			g.showStats()
		case cmd == "use" || cmd == "u":
			g.useItemOutOfCombat(inputFn)
		case cmd == "rest" || cmd == "r":
			g.rest()
		case cmd == "map":
			g.showMap()
		case cmd == "quit" || cmd == "q":
			g.println("\n👋 You wander off into the sunset, never to be seen again.")
			g.println("   Thanks for playing HILLMORD! 🏔️⚔️")
			return
		default:
			g.println("🤷 Unknown command. Type 'help' for a list of actions.")
		}

		if g.Player.Stats.HP <= 0 {
			g.println("\n💀 ═══════════════════════════════════════")
			g.println("   G A M E   O V E R")
			g.printf("   %s reached Level %d with %d 🪙\n", g.Player.Name, g.Player.Level, g.Player.Gold)
			g.println("   Better luck next reincarnation! 👻")
			g.println("💀 ═══════════════════════════════════════")
			return
		}
	}
}

func (g *Game) showHelp() {
	g.println(`
📖 ═══ COMMANDS ═══ 📖
  Movement:     [N]orth  [S]outh  [E]ast  [W]est
  Explore:      [L]ook  [T]alk  [F]ight  [M]arket  [Map]
  Character:    [I]nventory  [St]ats  [U]se item  [R]est
  System:       [H]elp  [Q]uit
`)
}

func (g *Game) describeLocation() {
	loc := g.World[g.Player.Location]
	g.printf("\n%s ═══ %s ═══ %s\n", loc.Emoji, loc.Name, loc.Emoji)
	g.printf("  %s\n", loc.Description)

	g.print("  Exits: ")
	exits := []string{}
	for dir, dest := range loc.Connections {
		exits = append(exits, fmt.Sprintf("%s → %s", dir, dest))
	}
	g.println(strings.Join(exits, "  |  "))

	if loc.HasMarket {
		g.println("  💰 There is a MARKET here.")
	}
	if len(loc.Enemies) > 0 {
		g.println("  ⚠️  Dangerous creatures lurk nearby...")
	}
}

func (g *Game) move(dir Direction) {
	loc := g.World[g.Player.Location]
	dest, ok := loc.Connections[dir]
	if !ok {
		quips := []string{
			"You walk face-first into an invisible wall. Classic.",
			"There's nothing that way except existential dread.",
			"A sign says: 'NOPE'. You respect the sign.",
			"You try, but your legs refuse. Smart legs.",
		}
		g.printf("🚫 %s\n", quips[rand.Intn(len(quips))])
		return
	}
	g.Player.Location = dest
	g.printf("🚶 You travel %s to %s...\n", dir, dest)
	g.describeLocation()

	// Random encounter chance (30%)
	newLoc := g.World[dest]
	if len(newLoc.Enemies) > 0 && rand.Intn(100) < 30 {
		g.println("\n⚠️  Something stirs in the shadows!")
		enemy := newLoc.Enemies[rand.Intn(len(newLoc.Enemies))]
		// Reset enemy HP
		enemy.Stats.HP = enemy.Stats.MaxHP
		g.RunCombat(enemy, g.inputFn)
	}
}

func (g *Game) seekFight(input func() string) {
	loc := g.World[g.Player.Location]
	if len(loc.Enemies) == 0 {
		g.println("😌 This place is peaceful. No one to fight. How boring.")
		return
	}
	enemy := loc.Enemies[rand.Intn(len(loc.Enemies))]
	enemy.Stats.HP = enemy.Stats.MaxHP
	g.RunCombat(enemy, input)
}

func (g *Game) enterMarket(input func() string) {
	loc := g.World[g.Player.Location]
	if !loc.HasMarket {
		g.println("🏪 There's no market here. Just dirt and disappointment.")
		return
	}
	g.RunMarket(input)
}

func (g *Game) talkToNPC() {
	loc := g.World[g.Player.Location]
	if len(loc.NPCLines) == 0 {
		g.println("🤐 There's no one here to talk to. You talk to yourself. It's fine. It's normal.")
		return
	}
	line := loc.NPCLines[rand.Intn(len(loc.NPCLines))]
	g.printf("\n  %s\n", line)
}

func (g *Game) showInventory() {
	g.println("\n🎒 ═══ INVENTORY ═══ 🎒")
	g.printf("  ⚔️  Weapon: %s (%s, %d-%d dmg)\n",
		g.Player.Weapon.Name, g.Player.Weapon.Class, g.Player.Weapon.MinDmg, g.Player.Weapon.MaxDmg)
	g.printf("  💰 Gold: %d 🪙\n", g.Player.Gold)
	if len(g.Player.Inventory) == 0 {
		g.println("  🎒 Bag: Empty. Like your future.")
	} else {
		g.println("  🎒 Bag:")
		for i, item := range g.Player.Inventory {
			healStr := ""
			if item.Heal > 0 {
				healStr = fmt.Sprintf(" (heals %d HP)", item.Heal)
			}
			g.printf("    %d. %s [%s]%s\n", i+1, item.Name, item.Kind, healStr)
		}
	}
}

func (g *Game) showStats() {
	g.println("\n📊 ═══ CHARACTER STATS ═══ 📊")
	g.printf("  🧍 %s  |  Level %d\n", g.Player.Name, g.Player.Level)
	g.printf("  ❤️  HP: %d/%d\n", g.Player.Stats.HP, g.Player.Stats.MaxHP)
	g.printf("  ⚔️  Attack: %d  |  🛡️  Defense: %d  |  💨 Speed: %d\n",
		g.Player.Stats.Attack, g.Player.Stats.Defense, g.Player.Stats.Speed)
	g.printf("  ✨ XP: %d / %d (to next level)\n", g.Player.XP, g.Player.Level*50)
	g.printf("  💰 Gold: %d 🪙\n", g.Player.Gold)
	g.printf("  🗡️  Weapon: %s (%d-%d dmg)\n", g.Player.Weapon.Name, g.Player.Weapon.MinDmg, g.Player.Weapon.MaxDmg)
}

func (g *Game) useItemOutOfCombat(input func() string) {
	healItems := []int{}
	for i, item := range g.Player.Inventory {
		if item.Heal > 0 {
			healItems = append(healItems, i)
		}
	}
	if len(healItems) == 0 {
		g.println("🎒 Nothing usable. Try buying potions or food at a market.")
		return
	}
	if g.Player.Stats.HP >= g.Player.Stats.MaxHP {
		g.println("❤️  You're already at full health! No need to waste supplies.")
		return
	}
	g.println("\n🎒 Usable items:")
	for j, idx := range healItems {
		item := g.Player.Inventory[idx]
		g.printf("  %d. %s (heals %d HP)\n", j+1, item.Name, item.Heal)
	}
	g.print("Choose (or 0 to cancel): > ")
	choice := 0
	fmt.Sscanf(strings.TrimSpace(input()), "%d", &choice)
	if choice < 1 || choice > len(healItems) {
		return
	}
	idx := healItems[choice-1]
	item := g.Player.Inventory[idx]
	g.Player.Stats.HP += item.Heal
	if g.Player.Stats.HP > g.Player.Stats.MaxHP {
		g.Player.Stats.HP = g.Player.Stats.MaxHP
	}
	g.printf("✨ You use %s and feel %d HP better! (%d/%d)\n",
		item.Name, item.Heal, g.Player.Stats.HP, g.Player.Stats.MaxHP)
	g.Player.Inventory = append(g.Player.Inventory[:idx], g.Player.Inventory[idx+1:]...)
}

func (g *Game) rest() {
	loc := g.World[g.Player.Location]
	if len(loc.Enemies) > 0 && rand.Intn(100) < 25 {
		g.println("😴 You try to rest but something interrupts your nap!")
		enemy := loc.Enemies[rand.Intn(len(loc.Enemies))]
		enemy.Stats.HP = enemy.Stats.MaxHP
		g.RunCombat(enemy, g.inputFn)
		return
	}
	heal := 10 + g.Player.Level*3
	g.Player.Stats.HP += heal
	if g.Player.Stats.HP > g.Player.Stats.MaxHP {
		g.Player.Stats.HP = g.Player.Stats.MaxHP
	}
	quips := []string{
		"You find a nice rock and have a little sit-down.",
		"You close your eyes for what feels like hours but was probably four minutes.",
		"A passing bird poops on you. But hey, you feel rested!",
		"You lean against a tree and immediately fall asleep. Nobody robs you, amazingly.",
	}
	g.printf("😴 %s Recovered %d HP. (%d/%d)\n",
		quips[rand.Intn(len(quips))], heal, g.Player.Stats.HP, g.Player.Stats.MaxHP)
}

func (g *Game) showMap() {
	g.println(`
🗺️ ═══ MAP OF HILLMORD ═══ 🗺️

                    🐲 Dragon's Pantry
                         |
         🌲 Gloomhollow ── 🌙 Moonpeak Summit
              Forest    |
               |    🧙‍♀️ Witch's
               |     Knuckle
          🏘️ Bumbleford ─── ⛰️ Cragmaw ─── 🐊 Dreadmire
               |              Pass     |      Swamp
          🌾 Soggy ───── 🐍 Rattlesnake  ── /
              Flats          Bazaar
                               |
                          💀 Skull Plateau
                               |
                          🕳️ The Underbelly
                               |
                     👑 Throne of Mild Inconvenience`)

	g.printf("\n  📍 You are at: %s\n", g.Player.Location)
}
