package adventofcode

import "fmt"

type Character struct {
	HitPoints    int
	DamagePoints int
	ArmorPoints  int
}

func (c Character) String() string {
	return fmt.Sprintf("{HP: %v Dmg: %v Def: %v}", c.HitPoints, c.DamagePoints, c.ArmorPoints)
}

func (c *Character) Attack(opponent *Character) int {

	damage := c.DamagePoints - opponent.ArmorPoints
	if damage < 1 {
		damage = 1
	}
	opponent.HitPoints = opponent.HitPoints - damage
	return damage
}

type Boss struct {
	*Character
}

func (b *Boss) Attack(opponent *Player) int {
	return b.Character.Attack(opponent.Character)
}

type Player struct {
	*Character
	Equipment Rack
}

func (p *Player) EquipWeapon(item Item) {
	p.Equipment.Weapon = item
}

func (p *Player) EquipArmor(item Item) {
	p.Equipment.Armor = item
}

func (p *Player) EquipRing(item Item) {
	if p.Equipment.Rings == nil {
		p.Equipment.Rings = make([]Item, 2)
	}
	p.Equipment.Rings = append(p.Equipment.Rings, item)
}

func (p *Player) Attack(opponent *Boss) int {
	return p.Character.Attack(opponent.Character)
}

func (p *Player) AdjustStats() {
	p.HitPoints = 100 // Always 100
	p.DamagePoints = 0
	p.ArmorPoints = 0
	for _, ring := range p.Equipment.Rings {
		p.DamagePoints = p.DamagePoints + ring.Damage
		p.ArmorPoints = p.ArmorPoints + ring.Armor
	}
	p.DamagePoints = p.DamagePoints + p.Equipment.Weapon.Damage
	p.ArmorPoints = p.ArmorPoints + p.Equipment.Armor.Armor
}

func (p *Player) EquipmentCost() int {
	return p.Equipment.Cost()
}

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

type Rack struct {
	Weapon Item
	Armor  Item
	Rings  []Item
}

func (r *Rack) Cost() int {
	cost := 0
	for _, item := range r.Rings {
		cost = cost + item.Cost
	}
	cost = cost + r.Armor.Cost + r.Weapon.Cost
	return cost
}

var Weapons = []Item{
	/*
		Weapons:    Cost  Damage  Armor
		Dagger        8     4       0
		Shortsword   10     5       0
		Warhammer    25     6       0
		Longsword    40     7       0
		Greataxe     74     8       0
	*/
	Item{Name: "Dagger", Cost: 8, Damage: 4, Armor: 0},
	Item{Name: "Shortsword", Cost: 10, Damage: 5, Armor: 0},
	Item{Name: "Warhammer", Cost: 25, Damage: 6, Armor: 0},
	Item{Name: "Longsword", Cost: 40, Damage: 7, Armor: 0},
	Item{Name: "Greataxe", Cost: 74, Damage: 8, Armor: 0},
}

var Armor = []Item{
	/*
		Armor:      Cost  Damage  Armor
		Leather      13     0       1
		Chainmail    31     0       2
		Splintmail   53     0       3
		Bandedmail   75     0       4
		Platemail   102     0       5
	*/
	Item{Name: "Leather", Cost: 13, Damage: 0, Armor: 1},
	Item{Name: "Chainmail", Cost: 31, Damage: 0, Armor: 2},
	Item{Name: "Splintmail", Cost: 53, Damage: 0, Armor: 3},
	Item{Name: "Bandedmail", Cost: 75, Damage: 0, Armor: 4},
	Item{Name: "Platemail", Cost: 102, Damage: 0, Armor: 5},
}

var Rings = []Item{
	/*
		Rings:      Cost  Damage  Armor
		Damage +1    25     1       0
		Damage +2    50     2       0
		Damage +3   100     3       0
		Defense +1   20     0       1
		Defense +2   40     0       2
		Defense +3   80     0       3
	*/
	Item{Name: "No Ring1", Cost: 0, Damage: 0, Armor: 0},
	Item{Name: "No Ring2", Cost: 0, Damage: 0, Armor: 0},
	Item{Name: "Damage +1", Cost: 25, Damage: 1, Armor: 0},
	Item{Name: "Damage +2", Cost: 50, Damage: 2, Armor: 0},
	Item{Name: "Damage +3", Cost: 100, Damage: 3, Armor: 0},
	Item{Name: "Defense +1", Cost: 20, Damage: 0, Armor: 1},
	Item{Name: "Defense +2", Cost: 40, Damage: 0, Armor: 2},
	Item{Name: "Defense +3", Cost: 80, Damage: 0, Armor: 3},
}

func RunGamePart1() {

	RunEquipmentPermutations()
}

func RunEquipmentPermutations() {

	battleCount := 0
	bestCost := 999999
	var bestRack Rack
	for armorIdx := 0; armorIdx < len(Armor); armorIdx++ {
		for weaponIdx := 0; weaponIdx < len(Weapons); weaponIdx++ {
			for ring1Idx := 0; ring1Idx < len(Rings); ring1Idx++ {
				for ring2Idx := 0; ring2Idx < len(Rings); ring2Idx++ {
					if ring1Idx == ring2Idx {
						continue //Don't equip the same ring twice
					}

					armor := Armor[armorIdx]
					weapon := Weapons[weaponIdx]
					ring1 := Rings[ring1Idx]
					ring2 := Rings[ring2Idx]

					player := EquipPlayer(armor, weapon, ring1, ring2)

					player.AdjustStats()

					boss := &Boss{&Character{
						HitPoints:    109,
						DamagePoints: 8,
						ArmorPoints:  2,
					},
					}

					fmt.Printf("Battle %v\n", battleCount)
					if RunBattle(player, boss) {
						lastCost := player.EquipmentCost()
						if lastCost < bestCost {
							bestCost = lastCost
							var bestRings []Item
							copy(bestRings, player.Equipment.Rings)
							bestRack = Rack{
								Armor:  player.Equipment.Armor,
								Weapon: player.Equipment.Weapon,
								Rings:  bestRings,
							}
						}
					}
					battleCount++
				}
			}
		}
	}
	fmt.Printf("Best:\n\tCost: %v\n\tRack: %v\n", bestCost, bestRack)
}

func RunBattle(player *Player, boss *Boss) bool {
	rounds := 0
	for {
		rounds = rounds + 1
		//fmt.Printf("Round %v\n \tBoss: %v\n\tPlayer: %v\n", rounds, boss, player)
		dealt := player.Attack(boss)
		//fmt.Printf("\t\tPlayer dealt %v Damage\n", dealt)
		if boss.HitPoints <= 0 {
			fmt.Printf("Player won after %v rounds!\n", rounds)
			return true
		}
		dealt = boss.Attack(player)
		fmt.Printf("\t\tBoss dealt %v Damage\n", dealt)
		if player.HitPoints <= 0 {
			fmt.Printf("Boss won after %v rounds!\n", rounds)
			return false
		}
	}
}

func EquipPlayer(armor Item, weapon Item, ring1 Item, ring2 Item) *Player {

	player := &Player{
		Character: &Character{},
		Equipment: Rack{},
	}

	player.EquipArmor(armor)
	player.EquipWeapon(weapon)
	player.EquipRing(ring1)
	player.EquipRing(ring2)

	return player
}
