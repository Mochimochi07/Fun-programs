package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    opponents := []string{"Freddy"}
    characters := []string{"Chica", "Bonnie", "Foxy","Goldy","Plushtrap"}
    summoned := []string{}
    for {
        fmt.Println("DeeDee: Uh oh! How unfortunate. Uh oh! How unfortunate. I know how much you like to fight.")
        time.Sleep(1 * time.Second) 
        fmt.Println("DeeDee: So I add a new contender to the night.")

        newOpponent := characters[rand.Intn(len(characters))]
        if contains(summoned,newOpponent) {
            continue
        }
        summoned = append(summoned, newOpponent)
        opponents = append(opponents, newOpponent)
        fmt.Println("A new challenger has been summoned!")
        fmt.Println("Current opponents: ", opponents)
    }
    fmt.Println("DeeDee: *Leaves*")
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

		// This is just a warm up for march because ive been deciding on prioritizing tons of ideas for my  programs. 
//Well i do want to make golang apps but i might need to focus more on using frameworks and other ideas.
//so this will be the last program for two months. ill just go get some time for working out.
