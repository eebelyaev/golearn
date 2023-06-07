// Выводит таблицу - результат поиска в GitHub.
package main

import (
	"fmt"
	"golearn/chapter04/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	periods := getPeriods(result.Items)
	printPeriods(periods)
}

func getPeriods(items []*github.Issue) [][]*github.Issue {
	periods := make([][]*github.Issue, 3)
	for i := 0; i < 3; i++ {
		periods[i] = make([]*github.Issue, 0)
	}

	t := time.Now()
	for _, item := range items {
		h := t.Sub(item.CreatedAt).Hours()
		switch {
		case h < 24*30:
			periods[0] = append(periods[0], item)
		case h < 24*365:
			periods[1] = append(periods[1], item)
		default:
			periods[2] = append(periods[2], item)
		}
	}
	return periods
}

func printPeriods(periods [][]*github.Issue) {
	for i, p := range periods {
		fmt.Printf("\n%v\n", i)
		for _, item := range p {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
