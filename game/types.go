package game

import "math/rand"

// ---------- Items & Weapons ----------

type WeaponClass int

const (
	ClassDagger WeaponClass = iota
	ClassSword
	ClassAxe
	ClassMace
	ClassBow
	ClassStaff
	ClassExotic
)

func (c WeaponClass) String() string {
	return [...]string{"🗡️ Dagger", "⚔️ Sword", "🪓 Axe", "🔨 Mace", "🏹 Bow", "🪄 Staff", "🦑 Exotic"}[c]
}

type Weapon struct {
	Name    string
	Class   WeaponClass
	MinDmg  int
	MaxDmg  int
	Price   int
	Quip    string // flavour text
}

func (w Weapon) RollDamage() int {
	if w.MaxDmg <= w.MinDmg {
		return w.MinDmg
	}
	return w.MinDmg + rand.Intn(w.MaxDmg-w.MinDmg+1)
}

type Item struct {
	Name  string
	Kind  string // "potion", "food", "scroll", "junk"
	Price int
	Heal  int
	Quip  string
}

// ---------- Characters ----------

type Stats struct {
	MaxHP   int
	HP      int
	Attack  int
	Defense int
	Speed   int
}

type Player struct {
	Name      string
	Stats     Stats
	Gold      int
	XP        int
	Level     int
	Weapon    Weapon
	Inventory []Item
	Location  string
}

type Enemy struct {
	Name    string
	Emoji   string
	Stats   Stats
	Gold    int
	XP      int
	Quip    string
	Weapon  Weapon
}

// ---------- World ----------

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
)

type Location struct {
	Name        string
	Emoji       string
	Description string
	Connections map[Direction]string // direction -> location name
	HasMarket   bool
	Enemies     []Enemy
	NPCLines    []string
}
