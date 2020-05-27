package main

import "fmt"

/*
	我有一个梦想 每个字符出现次数 rune => int, count => []rune a = 2 b = 3 c = 2 2 => ['a', 'c'] 3 => ['b']
*/

func main() {
	article := `Good afternoon! Today I would like to talk about the importance of keeping optimistic. 
	When we encounter difficulties in life, we notice that some of us choose to bury their heads in the sand. 
	Unfortunately, however, this attitude will do you no good, because if you will have no courage even to face them, 
	how can you conquer them? Thus, be optimistic, ladies and gentlemen, as it can give you confidence and help you see yourself 
	through the hard times, just as Winston Churchill once said, “An optimist sees an opportunity in every calamity; 
	a pessimist sees a calamity in every opportunity.” 　　

	Ladies and Gentlemen, keeping optimistic, you will be able to realize, in spite of some hardship, 
	there’s always hope waiting for you, which will lead you to the ultimate success. Historically as well as currently, 
	there are too many optimists of this kind to enumerate. You see, Thomas Edison is optimistic; if not, 
	the light of hope in his heart could not illuminate the whole world. Alfred Nobel is optimistic; 
	if not, the explosives and the prestigious Nobel Prize would not have come into being. 
	And Lance Armstrong is also optimistic; if not, the devil of cancer would have devoured his life and the world would 
	not see a 5-time winner of the Tour De France.`
	// 定义word_count
	word_count := map[rune]int{}

	// 统计每个字母出现的次数
	for _, word := range article {
		if word >= 'A' && word <= 'Z' || word >= 'a' && word <= 'z' {
			word_count[word]++
		}
	}
	// fmt.Println(word_count)
	countStats := map[int][]rune{}
	for ch, count := range word_count {
		// 	if _, ok := countStats[count]; ok {
		// 		countStats[count] = append(countStats[count], ch)
		// 	} else {
		// 		countStats[count] = []rune{ch}
		// 	}
		countStats[count] = append(countStats[count], ch)
	}
	// fmt.Println(countStats)

	for k, v := range countStats {
		fmt.Printf("%d : %s\n", k, string(v))
	}

}
