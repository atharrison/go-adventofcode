package adventofcode

import "fmt"

func RunDay22Game() {

	RunSpellPermutations()
}

type Wizard struct {
	*Character

	ManaPoints      int
	ActiveDurations []*DurationSpell
	ManaSpent       int
}

func (w *Wizard) String() string {
	return fmt.Sprintf("{HP: %v Dmg: %v Def: %v, Mana: %v}", w.HitPoints, w.DamagePoints, w.ArmorPoints, w.ManaPoints)
}

func (w *Wizard) SpellInEffect(s Spell) bool {
	for _, duration := range w.ActiveDurations {
		if duration.SpellCast == s {
			return true
		}
	}
	return false
}

func (w *Wizard) Attack(opponent *Boss, cast Spell) int {
	dealt := 0

	// Assumes the cast spell is valid.

	w.ManaPoints -= cast.ManaCost
	w.ManaSpent += cast.ManaCost

	if w.ManaPoints < 0 {
		return 0 // Ooops...
	}

	//fmt.Printf("Player Casts %v\n", cast.Name)
	if cast.Duration > 0 {
		// Cast Duration spell, then finish
		// Add duration to list of active durations.
		w.ActiveDurations = append(w.ActiveDurations, &DurationSpell{SpellCast: cast, RemainingDuration: cast.Duration})
		return 0
	} else {
		if cast.Damage > 0 {
			opponent.Character.HitPoints -= cast.Damage
			dealt = cast.Damage
		}
		if cast.Heal > 0 {
			w.Character.HitPoints += cast.Heal
		}
	}

	return dealt
}

func (w *Wizard) ApplyDurationSpells(opponent *Boss) {

	remainingDurations := []*DurationSpell{}
	for _, duration := range w.ActiveDurations {
		if duration.SpellCast.Armor > 0 {
			//fmt.Printf("Shield provides %v Armor; its timer is now %v.\n", duration.SpellCast.Armor, duration.RemainingDuration-1)
			w.ArmorPoints = duration.SpellCast.Armor
		}

		//if duration.SpellCast.Heal > 0 {
		//	//fmt.Printf("Drain provides %v HP; its timer is now %v.\n", duration.SpellCast.Heal, duration.RemainingDuration-1)
		//	w.HitPoints += duration.SpellCast.Heal
		//}

		if duration.SpellCast.ManaReplenish > 0 {
			//fmt.Printf("Recharge provides %v mana; its timer is now %v.\n", duration.SpellCast.ManaReplenish, duration.RemainingDuration-1)
			w.ManaPoints += duration.SpellCast.ManaReplenish
		}

		if duration.SpellCast.Damage > 0 {
			//fmt.Printf("Poison inflicts %v to Boss; its timer is now %v.\n", duration.SpellCast.Damage, duration.RemainingDuration-1)
			opponent.HitPoints -= duration.SpellCast.Damage
		}

		duration.RemainingDuration -= 1
		if duration.RemainingDuration > 0 {
			remainingDurations = append(remainingDurations, duration)
		}
	}
	w.ActiveDurations = remainingDurations
}

func (w *Wizard) ClearSpellEffects() {
	w.Character.ArmorPoints = 0 // Shield gone
}

func (b *Boss) AttackWizard(opponent *Wizard) int {
	return b.Character.Attack(opponent.Character)
}

type Spell struct {
	Name          string
	ManaCost      int
	Damage        int
	Heal          int
	Armor         int
	Duration      int
	ManaReplenish int
}

var Spells = []Spell{
	Spell{Name: "Magic Missle", ManaCost: 53, Damage: 4},
	Spell{Name: "Drain", ManaCost: 73, Damage: 2, Heal: 2},
	Spell{Name: "Shield", ManaCost: 113, Armor: 7, Duration: 6},
	Spell{Name: "Poison", ManaCost: 173, Damage: 3, Duration: 6},
	Spell{Name: "Recharge", ManaCost: 229, ManaReplenish: 101, Duration: 5},
}

var MIN_MANA_COST = 53
var MAX_SPELL_LIST_SIZE = 50
var SPELL_SEED_START = int64(1000000000000000000)

type DurationSpell struct {
	SpellCast         Spell
	RemainingDuration int
}

func RunSpellPermutations() {

	leastManaSpent := 99999
	var bestSpellList []int

	var spellSeed int64
	spellSeed = SPELL_SEED_START
	for {
		// Just iterate, starting at spellSeed,
		// converting the seed to a number in base len(Spells),
		// treated as an array of ints.
		// This array represents your 'list of spells' that can be tried.
		// Try it, and see if you best any previous attempt.
		spellSeed++
		spellList := ConvertToBase(spellSeed, int64(len(Spells)))
		//spellList := []int{4, 2, 1, 3, 0}
		boss := &Boss{
			//&Character{
			//	HitPoints:    14,
			//	DamagePoints: 8,
			//},
			&Character{
				HitPoints:    51,
				DamagePoints: 9,
			},
		}

		wizard := &Wizard{
			//Character: &Character{
			//	HitPoints: 10,
			//},
			//ManaPoints:      250,
			Character: &Character{
				HitPoints: 50,
			},
			ManaPoints:      500,
			ManaSpent:       0,
			ActiveDurations: []*DurationSpell{},
		}

		//fmt.Printf("Running SpellList %v...\n", spellList)
		result := RunWizardBattle(wizard, boss, spellList, leastManaSpent)
		//fmt.Printf("Battle Won? %v Mana Spent: %v\n", result, wizard.ManaSpent)
		if result == true && leastManaSpent >= wizard.ManaSpent {
			leastManaSpent = wizard.ManaSpent
			copy(spellList, bestSpellList)
			fmt.Printf("Battle Won? %v Mana Spent: %v, Spells: %v\n", result, wizard.ManaSpent, spellList)
		}

		//if spellSeed%500 == 0 {
		//	fmt.Printf("Best Option spends %v Mana: %v\n", leastManaSpent, spellList)
		//}
	}

}

func RunWizardBattle(wizard *Wizard, boss *Boss, spellList []int, maxManaToSpend int) bool {
	round := 0

	for idx := 0; idx < len(spellList); idx++ {

		if wizard.ManaSpent > maxManaToSpend {
			//fmt.Println("Bailing, overspent on Mana")
			return false
		}

		round = round + 1

		cast := Spells[spellList[idx]]

		//Player Turn
		if wizard.SpellInEffect(cast) {
			//Ignore, Find another spell...
			continue
		}

		_ = ExecutePlayerTurn(wizard, boss, cast, round)
		if boss.HitPoints <= 0 {
			//fmt.Printf("Player won after %v rounds!\n", round)
			return true
		}
		if wizard.HitPoints <= 0 {
			//fmt.Printf("Boss won after %v rounds!\n", round)
			return false
		}

		// Boss Turn
		_ = ExecuteBossTurn(wizard, boss, round)
		if boss.HitPoints <= 0 {
			//fmt.Printf("Player won after %v rounds!\n", round)
			return true
		}
		if wizard.HitPoints <= 0 {
			//fmt.Printf("Boss won after %v rounds!\n", round)
			return false
		}
		if wizard.ManaPoints < MIN_MANA_COST {
			//fmt.Printf("Boss won after %v rounds! (Player OOM)\n", round)
			return false
		}

	}
	fmt.Printf("!!!!!!!Battle took too long!!!!!!! %v\n", spellList)
	return false
}

func ExecutePlayerTurn(wizard *Wizard, boss *Boss, cast Spell, round int) int {
	//fmt.Printf("--- Player Turn (%v) ---\n", round)
	wizard.ClearSpellEffects()

	// Part 2
	wizard.HitPoints -= 1
	if wizard.HitPoints <= 0 {
		return 0 //Ooops, dead.
	}

	//fmt.Printf("\tBoss: %v\n\tPlayer: %v\n", boss, wizard)
	wizard.ApplyDurationSpells(boss)
	dealt := wizard.Attack(boss, cast)
	//fmt.Printf("\t\tPlayer dealt %v Damage\n", dealt)

	return dealt
}

func ExecuteBossTurn(wizard *Wizard, boss *Boss, round int) int {
	//fmt.Printf("--- Boss Turn   (%v) ---\n", round)
	//fmt.Printf("\tBoss: %v\n\tPlayer: %v\n", boss, wizard)
	wizard.ClearSpellEffects()
	wizard.ApplyDurationSpells(boss)

	dealt := boss.AttackWizard(wizard)
	//fmt.Printf("\t\tBoss dealt %v Damage\n", dealt)

	return dealt
}

func ConvertToBase(num int64, base int64) []int {
	var digit []int
	for num != 0 {
		remainder := num % base               // assume base > 1
		num = num / base                      // integer division
		digit = append(digit, int(remainder)) // We're going to make the array backwards on purpose...
	}
	return digit
}
