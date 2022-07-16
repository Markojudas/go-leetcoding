package main

import "sort"

func topKFrequentWord(words []string, k int) []string {

	//mapping the frequency of each word
	freqMap := map[string]int{}

	// holding the unique words
	set := []string{}

	//both setting the frequency of each word
	//and adding the unique words to "set"
	for _, word := range words {
		freqMap[word] += 1
		if freqMap[word] == 1 {
			set = append(set, word)
		}
	}

	//sorting the "set" in order of frequency
	//if the same frequency we order in lexicographic order
	sort.Slice(set, func(i, j int) bool {
		if freqMap[set[i]] == freqMap[set[j]] {
			return set[i] < set[j]
		} else {
			return freqMap[set[i]] > freqMap[set[j]]
		}
	})

	//we return the slice of unique words of the specified frequency
	return set[:k]
}
