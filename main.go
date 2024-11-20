package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("-------------------")
	fmt.Println("Um número aleatório será sorteado. tente acertar. O número é um inteiro entre 1 e 100.")

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}
	x := rand.Int64N(101)

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O valor informado não é um número inteiro.")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Seu chute foi menor que o número sorteado.", chuteInt)
		case chuteInt > x:
			fmt.Println("Seu chute foi maior que o número sorteado.", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns! Você acertou o número sorteado que era %d\n"+
					"Você acertou em %d tentativas\n"+
					"Essas foram as suas tentativas: %v\n",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Infelizmente, você não acertou o número sorteado, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram as suas tentativas: %v\n",
		x, chutes,
	)

}
