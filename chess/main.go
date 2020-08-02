package main

import "godemo/chess/landlords"

func main() {
	landlords.Play()
	// cases := []struct {
	// 	id       int
	// 	st       landlords.Stauts
	// 	expected landlords.Stauts
	// }{
	// 	{
	// 		id: 0,
	// 		st: landlords.Stauts{
	// 			Plcrds: [landlords.PLAYERSNUM][]string{
	// 				{"1", "2", "3"},
	// 				{"4", "5", "6"},
	// 				{"7", "8", "9"},
	// 			},
	// 			Plts:  [landlords.PLAYERSNUM]landlords.PLTYPE{0, 0, 1},
	// 			Cards: []string{},
	// 			Win:   [landlords.PLAYERSNUM]bool{false, false, false},
	// 		},
	// 		expected: landlords.Stauts{},
	// 	},
	// }

	// for _, c := range cases {
	// 	st := landlords.NextCards(c.id, c.st)
	// 	fmt.Println(st)
	// }
}
