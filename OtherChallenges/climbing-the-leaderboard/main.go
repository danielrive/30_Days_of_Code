package main

import (
	"fmt"
	"sort"
)

var ranking = make(map[int32]int32)

func main() {

	ranked := []int32{100, 400, 200, 300, 8050, 700, 5, 800, 5500, 100, 1040, 4040}
	player := []int32{5, 0, 400, 20000, 3000, 90, 700, 3, 2, 4, 3, 23, 4, 2, 3, 2}

	for j := range player {
		ranked = append(ranked, player[j])
		calculateRanking(ranked, player[j])
	}
	fmt.Println(ranking)
}
func calculateRanking(ranked []int32, valueToAdd int32) int32 {
	var ranking = make(map[int32]int32)
	//sorting ranking
	sort.Slice(ranked, func(i, j int) bool { return ranked[i] > ranked[j] })
	rankNumber := int32(0)
	for i := range ranked {
		if ranking[ranked[i]] == 0 {
			rankNumber = rankNumber + 1
		}
		ranking[ranked[i]] = rankNumber
	}
	return ranking[valueToAdd]

}
