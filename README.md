# ⚔️ HILLMORD

**A Dubiously Heroic Text-Based RPG**

> *"They said the hills had eyes. Turns out they also have swords."*

Hillmord is a terminal RPG written in Go where you wander a large, populated world, fight increasingly absurd monsters, haggle in markets, and ultimately face Lord Hillmord on the Throne of Mild Inconvenience. The tone is not serious. The danger is.

## 🚀 Quick Start

```bash
go run .
```

Or build a binary:

```bash
go build -o hillmord
./hillmord
```

## 🎮 How to Play

You navigate the world by typing commands at the `🎮 >` prompt:

| Command | Short | What it does |
|---|---|---|
| `north` | `n` | Travel north |
| `south` | `s` | Travel south |
| `east` | `e` | Travel east |
| `west` | `w` | Travel west |
| `look` | `l` | Describe your current location |
| `talk` | `t` | Chat with a local NPC |
| `fight` | `f` | Pick a fight with something nearby |
| `market` | `m` | Enter the local market (if there is one) |
| `inventory` | `i` | Check your bag and equipped weapon |
| `stats` | `st` | View your character sheet |
| `use` | `u` | Use a healing item outside of combat |
| `rest` | `r` | Take a nap and recover some HP |
| `map` | | Show the world map |
| `help` | `h` | List all commands |
| `quit` | `q` | Abandon your quest |

## ⚔️ Combat

Encounters are turn-based. When you meet an enemy (by exploring or seeking a fight), you can:

- **[A]ttack** — Swing your weapon. Damage is based on your weapon roll + ATK vs. enemy DEF. There's a 15% chance of a critical hit that ignores armor.
- **[U]se item** — Chug a potion or eat some Suspicious Meat mid-fight.
- **[F]lee** — Run away. Success depends on your Speed stat. Failure means you eat a free hit.

Defeating enemies earns 💰 gold and ✨ XP. Accumulate enough XP and you level up, gaining MaxHP, Attack, Defense, and Speed.

## 🗡️ Weapon Classes

Weapons come in seven classes, each with different damage ranges and price points:

| Class | Emoji | Example |
|---|---|---|
| Dagger | 🗡️ | Rusty Spoon, Sharpened Stick |
| Sword | ⚔️ | Iron Short Sword, Doom Blade of Excessive Sharpness |
| Axe | 🪓 | Bandit's Cleaver |
| Mace | 🔨 | Morningstar of Regret |
| Bow | 🏹 | Longbow of Optimism |
| Staff | 🪄 | Staff of Moderate Power |
| Exotic | 🦑 | Vorpal Cheese Knife, Squid Launcher |

## 💰 Markets

Four locations have markets where you can buy weapons and items or sell your loot at half price:

- 🏘️ **Bumbleford** — Starter town, friendly prices
- 🐍 **Rattlesnake Bazaar** — No refunds
- 💀 **Skull Plateau** — Skeleton-run, surprisingly well stocked
- 🕳️ **The Underbelly** — Underground black market vibes

Items for sale include health potions, dubious food, healing scrolls, and utterly useless junk (the Rubber Chicken is a must-buy).

## 🗺️ World Map

```
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
                     👑 Throne of Mild Inconvenience
```

You start in **Bumbleford** and must journey south and underground to reach **Lord Hillmord** on his ridiculous throne. The path there gets progressively harder — stock up on gear and potions before venturing deep.

## 👹 Notable Residents

| Enemy | Location | Flavour |
|---|---|---|
| 🐿️ Angry Squirrel | Gloomhollow Forest | Throws acorns at your head |
| 🐐 Mountain Goat of Doom | Cragmaw Pass | Those horns are NOT for decoration |
| 🧟‍♂️ Bog Zombie | Dreadmire Swamp | "Braaains... or whatever. I'm not picky." |
| 🐉 Sky Serpent | Moonpeak Summit | Coils through the clouds like a very angry rainbow |
| 🐲 Gourmet Dragon | Dragon's Pantry | Pairs adventurers with a nice chianti |
| 🧛 Underground Tax Collector | The Underbelly | "You owe back taxes. On EXISTING." |
| 👑 Lord Hillmord | Throne of Mild Inconvenience | The final boss. Mostly just annoying. |

## 📋 Requirements

- Go 1.21 or later
- A terminal with Unicode/emoji support

## 📁 Project Structure

```
hillmord/
├── main.go            Entry point
├── go.mod             Module definition
└── game/
    ├── types.go       Core types: weapons, items, stats, locations
    ├── title.go       Title screen
    ├── world.go       12 locations and their inhabitants
    ├── game.go        Main loop, movement, exploration
    ├── combat.go      Turn-based combat engine
    └── market.go      Buy/sell system
```

## 📜 License

Do whatever you want with it. Lord Hillmord doesn't care about licensing.
