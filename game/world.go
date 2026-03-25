package game

func BuildWorld() map[string]*Location {
	w := map[string]*Location{
		"Bumbleford": {
			Name:  "Bumbleford",
			Emoji: "🏘️",
			Description: "A sleepy hamlet where the chickens are braver than the guards. " +
				"Cobblestones wobble underfoot and the air smells faintly of cheese.",
			Connections: map[Direction]string{
				North: "Gloomhollow Forest",
				East:  "Cragmaw Pass",
				South: "Soggy Flats",
			},
			HasMarket: true,
			Enemies:   nil,
			NPCLines: []string{
				"👨‍🌾 Farmer Giles: \"Watch out for the hills. They mord, you know.\"",
				"👵 Old Nana: \"Back in my day, we fought dragons uphill both ways!\"",
				"🧒 Urchin: \"Give us a coin, mister! Or I'll tell everyone you smell.\"",
			},
		},
		"Gloomhollow Forest": {
			Name:  "Gloomhollow Forest",
			Emoji: "🌲",
			Description: "Ancient trees loom overhead, whispering rude things about your haircut. " +
				"Mushrooms glow in colours that probably aren't safe.",
			Connections: map[Direction]string{
				South: "Bumbleford",
				North: "Moonpeak Summit",
				East:  "Witch's Knuckle",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Angry Squirrel", Emoji: "🐿️", Stats: Stats{MaxHP: 12, HP: 12, Attack: 3, Defense: 1, Speed: 8}, Gold: 2, XP: 5, Quip: "It chatters furiously and throws an acorn at your head.", Weapon: Weapon{Name: "Tiny Claws", Class: ClassDagger, MinDmg: 1, MaxDmg: 3}},
				{Name: "Mossy Bandit", Emoji: "🧟", Stats: Stats{MaxHP: 25, HP: 25, Attack: 6, Defense: 3, Speed: 4}, Gold: 12, XP: 15, Quip: "\"Your money or your... wait, what was the other thing?\"", Weapon: Weapon{Name: "Rusty Cleaver", Class: ClassAxe, MinDmg: 3, MaxDmg: 7}},
				{Name: "Offended Treant", Emoji: "🌳", Stats: Stats{MaxHP: 40, HP: 40, Attack: 9, Defense: 6, Speed: 2}, Gold: 20, XP: 30, Quip: "\"You STEPPED on my COUSIN.\"", Weapon: Weapon{Name: "Branch Arm", Class: ClassMace, MinDmg: 5, MaxDmg: 10}},
			},
			NPCLines: []string{
				"🧝 Elf Hermit: \"I moved here for the peace and quiet. Then YOU showed up.\"",
				"🍄 Talking Mushroom: \"Pssst. Eat me. I dare you.\"",
			},
		},
		"Cragmaw Pass": {
			Name:  "Cragmaw Pass",
			Emoji: "⛰️",
			Description: "A narrow mountain pass littered with the bones of those who " +
				"thought they were 'pretty good at climbing, actually'.",
			Connections: map[Direction]string{
				West:  "Bumbleford",
				North: "Witch's Knuckle",
				East:  "Dreadmire Swamp",
				South: "Rattlesnake Bazaar",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Rock Goblin", Emoji: "👺", Stats: Stats{MaxHP: 18, HP: 18, Attack: 5, Defense: 2, Speed: 6}, Gold: 8, XP: 10, Quip: "It licks a rock menacingly.", Weapon: Weapon{Name: "Pointy Rock", Class: ClassDagger, MinDmg: 2, MaxDmg: 5}},
				{Name: "Mountain Goat of Doom", Emoji: "🐐", Stats: Stats{MaxHP: 22, HP: 22, Attack: 7, Defense: 4, Speed: 7}, Gold: 5, XP: 12, Quip: "Those horns are NOT for decoration.", Weapon: Weapon{Name: "Horns of Fury", Class: ClassExotic, MinDmg: 4, MaxDmg: 8}},
				{Name: "Cliff Harpy", Emoji: "🦅", Stats: Stats{MaxHP: 30, HP: 30, Attack: 8, Defense: 3, Speed: 9}, Gold: 18, XP: 25, Quip: "She screeches a song that's technically copyrighted.", Weapon: Weapon{Name: "Razor Talons", Class: ClassDagger, MinDmg: 4, MaxDmg: 9}},
			},
			NPCLines: []string{
				"🧗 Stranded Climber: \"I've been up here for three days. Got any snacks?\"",
			},
		},
		"Soggy Flats": {
			Name:  "Soggy Flats",
			Emoji: "🌾",
			Description: "Flat, damp, and deeply uninspiring. A sign reads: " +
				"'WELCOME TO SOGGY FLATS — IT'S NOT MUCH BUT IT'S HONEST'.",
			Connections: map[Direction]string{
				North: "Bumbleford",
				East:  "Rattlesnake Bazaar",
				South: "Dreadmire Swamp",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Mud Crab", Emoji: "🦀", Stats: Stats{MaxHP: 15, HP: 15, Attack: 4, Defense: 5, Speed: 3}, Gold: 3, XP: 8, Quip: "It snaps passive-aggressively.", Weapon: Weapon{Name: "Pinchy Claws", Class: ClassExotic, MinDmg: 2, MaxDmg: 5}},
				{Name: "Depressed Scarecrow", Emoji: "🧣", Stats: Stats{MaxHP: 20, HP: 20, Attack: 5, Defense: 2, Speed: 3}, Gold: 7, XP: 10, Quip: "\"I was meant for greater things... like scaring crows.\"", Weapon: Weapon{Name: "Floppy Arm", Class: ClassStaff, MinDmg: 2, MaxDmg: 6}},
			},
			NPCLines: []string{
				"🐸 Philosophical Frog: \"Ribbit. But, like, why ribbit? You know?\"",
				"👨‍🌾 Soggy Pete: \"Rain again. Always rain. I love it here. (He does not love it here.)\"",
			},
		},
		"Rattlesnake Bazaar": {
			Name:  "Rattlesnake Bazaar",
			Emoji: "🐍",
			Description: "A chaotic open-air market where everything is for sale and nothing is refundable. " +
				"A snake in a tiny hat runs the information booth.",
			Connections: map[Direction]string{
				North: "Cragmaw Pass",
				West:  "Soggy Flats",
				East:  "Dreadmire Swamp",
				South: "Skull Plateau",
			},
			HasMarket: true,
			Enemies:   nil,
			NPCLines: []string{
				"🐍 Sssylvia the Snake: \"Welcomesss to the bazzzaar. No refundsss.\"",
				"🎪 Juggler: \"Watch this!\" *drops everything* \"...I'm still learning.\"",
				"🧙 Shady Wizard: \"Psst. Want to buy a spell? Mostly works. Sometimes explodes.\"",
			},
		},
		"Witch's Knuckle": {
			Name:  "Witch's Knuckle",
			Emoji: "🧙‍♀️",
			Description: "A craggy hilltop shaped like a fist. Locals say a witch cursed it. " +
				"The witch says locals are being dramatic.",
			Connections: map[Direction]string{
				South: "Cragmaw Pass",
				West:  "Gloomhollow Forest",
				North: "Moonpeak Summit",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Hex Bat", Emoji: "🦇", Stats: Stats{MaxHP: 14, HP: 14, Attack: 5, Defense: 2, Speed: 10}, Gold: 6, XP: 8, Quip: "It squeaks a tiny curse at you. Your left shoe feels tighter.", Weapon: Weapon{Name: "Sonic Screech", Class: ClassStaff, MinDmg: 2, MaxDmg: 6}},
				{Name: "Cursed Knight", Emoji: "🛡️", Stats: Stats{MaxHP: 35, HP: 35, Attack: 9, Defense: 7, Speed: 3}, Gold: 25, XP: 30, Quip: "\"I've been cursed to fight anyone who makes eye contact. So thanks for THAT.\"", Weapon: Weapon{Name: "Hexblade", Class: ClassSword, MinDmg: 5, MaxDmg: 11}},
				{Name: "Warty Familiar", Emoji: "🐸", Stats: Stats{MaxHP: 10, HP: 10, Attack: 3, Defense: 1, Speed: 5}, Gold: 4, XP: 5, Quip: "It's a toad. It's angry. It knows spells. Good luck.", Weapon: Weapon{Name: "Tongue Lash", Class: ClassExotic, MinDmg: 1, MaxDmg: 4}},
			},
			NPCLines: []string{
				"🧙‍♀️ Griselda the Witch: \"I'm NOT evil. I just have resting hex face.\"",
			},
		},
		"Dreadmire Swamp": {
			Name:  "Dreadmire Swamp",
			Emoji: "🐊",
			Description: "Bubbling mud, ominous fog, and the distinct aroma of regret. " +
				"A crocodile watches you with unsettling intelligence.",
			Connections: map[Direction]string{
				West:  "Rattlesnake Bazaar",
				North: "Cragmaw Pass",  // loop back
				South: "Skull Plateau",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Bog Zombie", Emoji: "🧟‍♂️", Stats: Stats{MaxHP: 28, HP: 28, Attack: 7, Defense: 3, Speed: 2}, Gold: 10, XP: 15, Quip: "\"Braaains... or whatever. I'm not picky.\"", Weapon: Weapon{Name: "Muddy Fist", Class: ClassMace, MinDmg: 3, MaxDmg: 8}},
				{Name: "Swamp Witch", Emoji: "🧹", Stats: Stats{MaxHP: 22, HP: 22, Attack: 10, Defense: 2, Speed: 6}, Gold: 20, XP: 22, Quip: "She offers you tea. The tea is alive.", Weapon: Weapon{Name: "Hex Bolt", Class: ClassStaff, MinDmg: 5, MaxDmg: 12}},
				{Name: "Enormous Leech", Emoji: "🪱", Stats: Stats{MaxHP: 16, HP: 16, Attack: 6, Defense: 1, Speed: 4}, Gold: 3, XP: 8, Quip: "It wiggles toward you with alarming enthusiasm.", Weapon: Weapon{Name: "Sucker Bite", Class: ClassExotic, MinDmg: 3, MaxDmg: 7}},
			},
			NPCLines: []string{
				"🐊 Old Chomps: *stares* *blinks once* *keeps staring*",
				"👻 Swamp Ghost: \"OooOOoo! ...sorry, that's all I've got.\"",
			},
		},
		"Moonpeak Summit": {
			Name:  "Moonpeak Summit",
			Emoji: "🌙",
			Description: "The highest point in the land. The view is breathtaking, " +
				"which is good because the altitude already took your breath.",
			Connections: map[Direction]string{
				South: "Gloomhollow Forest",
				East:  "Witch's Knuckle",
				West:  "Dragon's Pantry",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Frost Wraith", Emoji: "👻", Stats: Stats{MaxHP: 32, HP: 32, Attack: 10, Defense: 4, Speed: 7}, Gold: 22, XP: 28, Quip: "\"I'm cold. YOU'RE cold. Let's fight about it.\"", Weapon: Weapon{Name: "Icy Touch", Class: ClassStaff, MinDmg: 5, MaxDmg: 12}},
				{Name: "Sky Serpent", Emoji: "🐉", Stats: Stats{MaxHP: 45, HP: 45, Attack: 12, Defense: 5, Speed: 8}, Gold: 35, XP: 40, Quip: "It coils through the clouds like a very angry rainbow.", Weapon: Weapon{Name: "Lightning Fang", Class: ClassExotic, MinDmg: 7, MaxDmg: 14}},
			},
			NPCLines: []string{
				"🧘 Meditating Monk: \"I came here to find myself. Found frostbite instead.\"",
			},
		},
		"Dragon's Pantry": {
			Name:  "Dragon's Pantry",
			Emoji: "🐲",
			Description: "A massive cave filled with gold, bones, and inexplicably, a fully stocked spice rack. " +
				"The dragon is a foodie.",
			Connections: map[Direction]string{
				East: "Moonpeak Summit",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Gourmet Dragon", Emoji: "🐲", Stats: Stats{MaxHP: 80, HP: 80, Attack: 18, Defense: 10, Speed: 6}, Gold: 100, XP: 100, Quip: "\"You look like you'd go well with a nice chianti and some fava beans.\"", Weapon: Weapon{Name: "Flame Breath", Class: ClassExotic, MinDmg: 12, MaxDmg: 22}},
			},
			NPCLines: []string{
				"🍖 A talking leg of ham: \"Don't eat me! I have a FAMILY!\"",
			},
		},
		"Skull Plateau": {
			Name:  "Skull Plateau",
			Emoji: "💀",
			Description: "A windswept plateau shaped like a skull when viewed from above. " +
				"Nobody knows who did the landscaping but they had commitment issues with subtlety.",
			Connections: map[Direction]string{
				North: "Rattlesnake Bazaar",
				West:  "Dreadmire Swamp",
				South: "The Underbelly",
			},
			HasMarket: true,
			Enemies: []Enemy{
				{Name: "Skeleton Merchant", Emoji: "💀", Stats: Stats{MaxHP: 20, HP: 20, Attack: 6, Defense: 3, Speed: 5}, Gold: 30, XP: 15, Quip: "\"I used to sell potions. Now I AM a potion... of sadness.\"", Weapon: Weapon{Name: "Bone Club", Class: ClassMace, MinDmg: 3, MaxDmg: 7}},
				{Name: "Vulture King", Emoji: "🦅", Stats: Stats{MaxHP: 38, HP: 38, Attack: 11, Defense: 5, Speed: 7}, Gold: 25, XP: 35, Quip: "He wears a tiny crown. It suits him.", Weapon: Weapon{Name: "Regal Talons", Class: ClassDagger, MinDmg: 6, MaxDmg: 12}},
			},
			NPCLines: []string{
				"💀 Skull on a pike: \"Lovely weather we're having.\"",
				"🦴 Skeleton busker: *plays a xylophone made of ribs*",
			},
		},
		"The Underbelly": {
			Name:  "The Underbelly",
			Emoji: "🕳️",
			Description: "Beneath the plateau lies a vast underground city lit by stolen starlight. " +
				"It smells like ambition and old socks.",
			Connections: map[Direction]string{
				North: "Skull Plateau",
				South: "The Throne of Mild Inconvenience",
			},
			HasMarket: true,
			Enemies: []Enemy{
				{Name: "Sewer Rat King", Emoji: "🐀", Stats: Stats{MaxHP: 50, HP: 50, Attack: 13, Defense: 6, Speed: 6}, Gold: 40, XP: 45, Quip: "It's literally twelve rats in a trench coat wearing a crown.", Weapon: Weapon{Name: "Plague Teeth", Class: ClassExotic, MinDmg: 7, MaxDmg: 15}},
				{Name: "Underground Tax Collector", Emoji: "🧛", Stats: Stats{MaxHP: 35, HP: 35, Attack: 8, Defense: 8, Speed: 5}, Gold: 50, XP: 30, Quip: "\"You owe back taxes. On EXISTING.\"", Weapon: Weapon{Name: "Stamp of Authority", Class: ClassMace, MinDmg: 4, MaxDmg: 9}},
			},
			NPCLines: []string{
				"🐀 Ratfolk Bartender: \"What'll it be? We got ale, and slightly different ale.\"",
				"🕯️ Candlemaker: \"Business is lit. Get it? LIT? I hate it here.\"",
			},
		},
		"The Throne of Mild Inconvenience": {
			Name:  "The Throne of Mild Inconvenience",
			Emoji: "👑",
			Description: "The final chamber. A ridiculous throne made of slightly uncomfortable chairs " +
				"stacked to the ceiling. On it sits the Dark Lord of Hillmord... " +
				"who is mostly just annoying.",
			Connections: map[Direction]string{
				North: "The Underbelly",
			},
			HasMarket: false,
			Enemies: []Enemy{
				{Name: "Lord Hillmord", Emoji: "👑", Stats: Stats{MaxHP: 120, HP: 120, Attack: 22, Defense: 12, Speed: 7}, Gold: 500, XP: 200, Quip: "\"Ah, a hero! How TEDIOUSLY predictable. I was in the middle of a sudoku.\"", Weapon: Weapon{Name: "Sceptre of Annoyance", Class: ClassStaff, MinDmg: 14, MaxDmg: 26}},
			},
			NPCLines: []string{
				"🪑 A sentient chair: \"He sits on us ALL DAY. Send help.\"",
			},
		},
	}
	return w
}
