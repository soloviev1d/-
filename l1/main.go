package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

type RuleMap map[rune][]string

// Правила порождения
var g = RuleMap{
	'S': {"aF"},
	'F': {"Sb", "Fb", "b"},
}

// Случайное слово из языка грамматики G
func genWord(nonterm rune) string {
	if rules, ok := g[nonterm]; ok {
		r := rules[rand.Intn(len(rules))]
		word := ""
		for _, c := range r {
			word += genWord(c)
		}
		return word
	}
	return string(nonterm)
}

func (rm *RuleMap) toStringUnfold() string {
	s := ""
	for k, _ := range *rm {
		for _, v := range (*rm)[k] {
			s += fmt.Sprintf("%c -> %s\n", k, v)
		}
	}
	return s
}
func main() {
	var (
		s   = 'S' // Начальный нетерминал
		scn = bufio.NewReader(os.Stdin)
	)

	fmt.Print("Ruleset:\n", g.toStringUnfold())
	fmt.Println("Press return for next word")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		for {
			_, _, err := scn.ReadRune()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("word: %s", genWord(s))
		}

	}()
	<-ctx.Done()

}
