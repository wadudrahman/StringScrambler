/**
 * Aoccdrnig to a rscheearch at Cmabrigde Uinervtisy, it deosn't mttaer in waht oredr the ltteers in a wrod are, the olny iprmoetnt tihng is taht the frist and lsat ltteer be at the rghit pclae. The rset can be a toatl mses and you can sitll raed it wouthit porbelm. Tihs is bcuseae the huamn mnid deos not raed ervey lteter by istlef, but the wrod as a wlohe.
 * Hasin Hayder <me at hasin dot me>
 */
package main

import (
	"fmt"
	"time"
	"math/rand"
	"strings"
	"bufio"
	"os"
)

func main() {
	var shuffledWord string;
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input some words you want to scramble, and hit enter: ")
	text, _ := reader.ReadString('\n')
	texttokens := strings.Split(strings.TrimRight(text, "\n"), " ");
	fmt.Println()
	for i := range (texttokens) {
		if (len(texttokens[i]) <= 3) {
			//no need to change words like i, am, you
			fmt.Print(texttokens[i] + " ")
		} else if (len(texttokens[i]) == 4 && texttokens[i][1] == texttokens[i][2]) {
			//no need to change 4 letter palindromes like book, boob, seen, deed
			fmt.Print(texttokens[i] + " ")
		} else {
			deadlockbuster := 0
			for {
				deadlockbuster++
				//this block makes sure that word is actually scrambled after shuffling
				shuffledWord = ShuffleWord(texttokens[i])
				if (texttokens[i] != shuffledWord ) {
					break
				}
				if (deadlockbuster > 5) {
					break // just for the safety to not fall into deadlock when people enter words like haaaah or fuuuuuk
				}
			}
			fmt.Print(shuffledWord + " ")
		}
	}
	fmt.Print("\n\nThank You!")
}

func ShuffleWord(word string) string {
	if (len(word) <= 3) {
		//no need to change words like i, am, you
		return word
	}
	if (len(word) == 4 && word[2] == word[3]) {
		//no need to change 4 letter palindromes like book, boob, seen, deed
		return word
	}

	letters := []rune(word)

	punctuationMarkFound := false;
	punctuations := "^+%@*$#-<>:,.';?"
	if (strings.ContainsAny(word, punctuations)) {
		punctuationMarkFound = true
	}

	if (punctuationMarkFound) {
		letters = []rune(word[:len(word)-1])
	}

	insiders := letters[1:len(letters)-1]

	rand.Seed(time.Now().UnixNano())

	for i := range insiders {
		j := rand.Intn(i + 1)
		insiders[i], insiders[j] = insiders[j], insiders[i]
	}

	if (!punctuationMarkFound) {
		return string(letters)
	} else {
		return string(letters) + string(word[len(word)-1])
	}
}
