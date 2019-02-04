// Package stringutil contains utility functions for working with strings.
package stringutil

import "strings"

//WordCount is My implementation of wordcount function - go tour
func WordCount(s string) map[string]int {
    
	words := strings.Fields(s)
    wordCountMap := make(map[string]int)
    
    for _,word := range words{
        wordCountMap[word]++    
    }
    
    return wordCountMap
}
