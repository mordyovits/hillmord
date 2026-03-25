package game

import (
	"fmt"
	"strconv"
	"strings"
)

var WeaponsForSale = []Weapon{
	{Name: "Rusty Spoon", Class: ClassDagger, MinDmg: 1, MaxDmg: 3, Price: 5, Quip: "It's seen better days. And better meals."},
	{Name: "Sharpened Stick", Class: ClassDagger, MinDmg: 2, MaxDmg: 4, Price: 10, Quip: "Nature's switchblade."},
	{Name: "Iron Short Sword", Class: ClassSword, MinDmg: 3, MaxDmg: 7, Price: 30, Quip: "Short sword, big dreams."},
	{Name: "Bandit's Cleaver", Class: ClassAxe, MinDmg: 4, MaxDmg: 9, Price: 45, Quip: "Fell off the back of a bandit."},
	{Name: "Morningstar of Regret", Class: ClassMace, MinDmg: 5, MaxDmg: 11, Price: 60, Quip: "Every swing makes you question your life choices."},
	{Name: "Longbow of Optimism", Class: ClassBow, MinDmg: 4, MaxDmg: 10, Price: 55, Quip: "You'll definitely hit something. Maybe even the enemy."},
	{Name: "Staff of Moderate Power", Class: ClassStaff, MinDmg: 3, MaxDmg: 12, Price: 70, Quip: "It hums with energy. Or bees. Might be bees."},
	{Name: "Vorpal Cheese Knife", Class: ClassExotic, MinDmg: 6, MaxDmg: 14, Price: 100, Quip: "Cuts through enemies like a hot knife through... well, cheese."},
	{Name: "Doom Blade of Excessive Sharpness", Class: ClassSword, MinDmg: 8, MaxDmg: 18, Price: 200, Quip: "So sharp it cut the price tag. Twice."},
	{Name: "Squid Launcher", Class: ClassExotic, MinDmg: 10, MaxDmg: 20, Price: 350, Quip: "Fires live squid. The squid are NOT happy about this."},
}

var ItemsForSale = []Item{
	{Name: "🧪 Health Potion", Kind: "potion", Price: 15, Heal: 25, Quip: "Tastes like cherry cough syrup and broken dreams."},
	{Name: "🧪 Greater Health Potion", Kind: "potion", Price: 40, Heal: 60, Quip: "Now with 60% more healing and 40% more aftertaste."},
	{Name: "🍖 Suspicious Meat", Kind: "food", Price: 8, Heal: 15, Quip: "Don't ask what animal. You won't like the answer."},
	{Name: "🧀 Wheel of Courage Cheese", Kind: "food", Price: 12, Heal: 20, Quip: "Aged for bravery. Smells like feet."},
	{Name: "🍞 Hero's Sandwich", Kind: "food", Price: 20, Heal: 35, Quip: "Two slices of bread with DESTINY in between."},
	{Name: "📜 Scroll of Mild Healing", Kind: "scroll", Price: 25, Heal: 40, Quip: "The instructions say 'just feel better, lol'."},
	{Name: "📜 Scroll of Decent Healing", Kind: "scroll", Price: 50, Heal: 80, Quip: "Written by a wizard who actually graduated."},
	{Name: "🐔 Rubber Chicken", Kind: "junk", Price: 3, Heal: 0, Quip: "No practical use whatsoever. A must-buy."},
	{Name: "🎺 Tiny Trumpet", Kind: "junk", Price: 5, Heal: 0, Quip: "Play a little fanfare when you enter battle. Enemies hate it."},
}

func (g *Game) RunMarket(input func() string) {
	for {
		fmt.Println("\n💰 ═══ MARKET ═══ 💰")
		fmt.Printf("Your gold: %d 🪙\n\n", g.Player.Gold)
		fmt.Println("[W] Browse weapons  [I] Browse items  [S] Sell items  [L] Leave market")
		fmt.Print("> ")

		cmd := strings.ToLower(strings.TrimSpace(input()))
		switch cmd {
		case "w":
			g.browseWeapons(input)
		case "i":
			g.browseItems(input)
		case "s":
			g.sellItems(input)
		case "l":
			fmt.Println("🚶 You leave the market, wallet slightly lighter.")
			return
		default:
			fmt.Println("🤷 The merchant stares blankly.")
		}
	}
}

func (g *Game) browseWeapons(input func() string) {
	fmt.Println("\n⚔️  ─── WEAPONS FOR SALE ─── ⚔️")
	fmt.Printf("  %-4s %-35s %-14s %-10s %s\n", "#", "Name", "Class", "Damage", "Price")
	fmt.Println("  " + strings.Repeat("─", 80))
	for i, w := range WeaponsForSale {
		fmt.Printf("  %-4d %-35s %-14s %2d-%-7d %d 🪙\n", i+1, w.Name, w.Class, w.MinDmg, w.MaxDmg, w.Price)
		fmt.Printf("       💬 %s\n", w.Quip)
	}
	fmt.Printf("\n  Current weapon: %s (%s, %d-%d dmg)\n", g.Player.Weapon.Name, g.Player.Weapon.Class, g.Player.Weapon.MinDmg, g.Player.Weapon.MaxDmg)
	fmt.Print("\nEnter number to buy (or 0 to go back): > ")

	choice, err := strconv.Atoi(strings.TrimSpace(input()))
	if err != nil || choice < 1 || choice > len(WeaponsForSale) {
		return
	}
	w := WeaponsForSale[choice-1]
	if g.Player.Gold < w.Price {
		fmt.Printf("😢 You need %d 🪙 but only have %d. Embarrassing.\n", w.Price, g.Player.Gold)
		return
	}
	g.Player.Gold -= w.Price
	old := g.Player.Weapon
	g.Player.Weapon = w
	fmt.Printf("🎉 You bought %s! Equipped immediately.\n", w.Name)
	fmt.Printf("   (Your old %s clatters to the ground. No refunds.)\n", old.Name)
}

func (g *Game) browseItems(input func() string) {
	fmt.Println("\n🎒 ─── ITEMS FOR SALE ─── 🎒")
	for i, item := range ItemsForSale {
		healStr := ""
		if item.Heal > 0 {
			healStr = fmt.Sprintf(" (heals %d HP)", item.Heal)
		}
		fmt.Printf("  %d. %s — %d 🪙%s\n", i+1, item.Name, item.Price, healStr)
		fmt.Printf("     💬 %s\n", item.Quip)
	}
	fmt.Print("\nEnter number to buy (or 0 to go back): > ")

	choice, err := strconv.Atoi(strings.TrimSpace(input()))
	if err != nil || choice < 1 || choice > len(ItemsForSale) {
		return
	}
	item := ItemsForSale[choice-1]
	if g.Player.Gold < item.Price {
		fmt.Printf("😢 Can't afford %s. You need %d 🪙.\n", item.Name, item.Price)
		return
	}
	g.Player.Gold -= item.Price
	g.Player.Inventory = append(g.Player.Inventory, item)
	fmt.Printf("🎉 Bought %s! It's now rattling around in your bag.\n", item.Name)
}

func (g *Game) sellItems(input func() string) {
	if len(g.Player.Inventory) == 0 {
		fmt.Println("🎒 Your bag is emptier than a goblin's promise.")
		return
	}
	fmt.Println("\n🎒 ─── YOUR INVENTORY ─── 🎒")
	for i, item := range g.Player.Inventory {
		sellPrice := item.Price / 2
		if sellPrice < 1 {
			sellPrice = 1
		}
		fmt.Printf("  %d. %s (sell for %d 🪙)\n", i+1, item.Name, sellPrice)
	}
	fmt.Print("\nEnter number to sell (or 0 to go back): > ")

	choice, err := strconv.Atoi(strings.TrimSpace(input()))
	if err != nil || choice < 1 || choice > len(g.Player.Inventory) {
		return
	}
	item := g.Player.Inventory[choice-1]
	sellPrice := item.Price / 2
	if sellPrice < 1 {
		sellPrice = 1
	}
	g.Player.Gold += sellPrice
	g.Player.Inventory = append(g.Player.Inventory[:choice-1], g.Player.Inventory[choice:]...)
	fmt.Printf("💸 Sold %s for %d 🪙. The merchant grunts approvingly.\n", item.Name, sellPrice)
}
