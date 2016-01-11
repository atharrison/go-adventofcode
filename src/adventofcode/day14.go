package adventofcode

import "fmt"

/*
Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.
Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.
Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.
Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.

Race: Fly for 2503 seconds.

*/

const RACE_DURATION = 2503

type Reindeer struct {
	Name     string
	Speed    int
	Duration int
	Rest     int

	DistanceCovered int
	Resting         bool
	CurrentDuration int
	CurrentRest     int
}

func (r *Reindeer) String() string {

	return fmt.Sprintf("%s can fly %d km/s for %d seconds, rest %d seconds.", r.Name, r.Speed, r.Duration, r.Rest)
}

func (r *Reindeer) Tic() {
	if r.Resting {
		r.CurrentRest++
		if r.CurrentRest == r.Rest {
			r.Resting = false
			// Clear rest
			r.CurrentRest = 0
		}
	} else { // Flying
		// Fly one second
		r.DistanceCovered += r.Speed
		r.CurrentDuration++
		if r.CurrentDuration == r.Duration {
			// Switch to resting
			r.Resting = true
			// Clear flying duration
			r.CurrentDuration = 0
		}
	}
}

func ExecuteDay14() {

	participants := []*Reindeer{

		&Reindeer{
			//		Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
			Name:     "Vixen",
			Speed:    19,
			Duration: 7,
			Rest:     124,
		},
		&Reindeer{
			//		Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
			Name:     "Rudolph",
			Speed:    3,
			Duration: 15,
			Rest:     28,
		},
		&Reindeer{
			//		Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
			Name:     "Donner",
			Speed:    19,
			Duration: 9,
			Rest:     164,
		},
		&Reindeer{
			//		Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.
			Name:     "Blitzen",
			Speed:    19,
			Duration: 9,
			Rest:     158,
		},
		&Reindeer{
			//		Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.
			Name:     "Comet",
			Speed:    13,
			Duration: 7,
			Rest:     82,
		},
		&Reindeer{
			//		Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
			Name:     "Cupid",
			Speed:    25,
			Duration: 6,
			Rest:     145,
		},
		&Reindeer{
			//		Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
			Name:     "Dasher",
			Speed:    14,
			Duration: 3,
			Rest:     38,
		},
		&Reindeer{
			//		Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.
			Name:     "Dancer",
			Speed:    3,
			Duration: 16,
			Rest:     37,
		},
		&Reindeer{
			//		Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.
			Name:     "Prancer",
			Speed:    25,
			Duration: 6,
			Rest:     143,
		},
	}
	fmt.Println(participants)

	for tics := 0; tics < RACE_DURATION; tics++ {
		for _, deer := range participants {
			deer.Tic()
		}
	}

	var winner *Reindeer
	for idx, deer := range participants {
		fmt.Printf("%s Flew %d km\n", deer.Name, deer.DistanceCovered)
		if idx == 0 || winner.DistanceCovered < deer.DistanceCovered {
			winner = deer
		}
	}

	fmt.Printf("Winner!!! -> %s flew %d km!\n", winner.Name, winner.DistanceCovered)
}
