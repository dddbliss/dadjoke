package cmd

import(
	"fmt"
	"os"
	"strings"
	"bufio"
	_ "embed"
	"math/rand"

	"github.com/spf13/cobra"
)

//go:embed dad_jokes.txt
var jokes string

func readJokes() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(jokes))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

var rootCmd = &cobra.Command{
	Use: "dadjoke",
	Short: "dadjoke is a very fun utility to print out a random dad joke",
	Long: "dadjoke is the joke teller your dad never was. Written in go with love by dddbliss and friends",
	Run: func(cmd *cobra.Command, args[]string) {
		lines, err := readJokes()
		if err != nil {
			panic(err)
		}

		joke := lines[rand.Intn(len(lines))]
		fmt.Println(joke)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}