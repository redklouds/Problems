package main

import "strings"

func longestCommonPrefix(strs []string) string {
    //condiiton if we find a word that index is END this is the current longgest prefix
    //longest prefix is when the next word, is at the END of itself,
    
    //if the word  next idx IS NOT there, then return "" no common prefix
    ///assuming that each word will build to the next common prefix means they are common for all words
   var sb strings.Builder
   //enqueue the first
    for i:=0; i < len(strs); i++{
        //start with the first character in this word add it to the prefix, and check the next word index
        
        wordArr := []byte(word[i])
		//get the current ith character here
        if wordArr[idx] != nextByte
        sb.
    }
	return sb.String()
}