package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item

}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

type Queue struct {
	list *list.List
}

func (q *Queue) Enqueue(chalenge string) {
	q.list.PushBack(chalenge)
}

func (q *Queue) Dequeue() string {
	if q.list.Len() == 0 {
		return ""
	}
	element := q.list.Front()
	q.list.Remove(element)
	return element.Value.(string)
}

var challenges = []string{
	"Open the treasure chest",
	"Find a hidden treasure",
	"Cure a disease",
	"Build a shelter",
	"Keep yourself warm",
	"Battle a dragon",
	"Fight a monster",
	"Survive a storm",
	"Escape a maze",
	"Darkness in the forest",
	"Dive into the ocean",
	"Defend a castle",
	"Cross a river",
	"Climb a wall",
}

var tools = map[string]string{
	"Open the treasure chest": "Key",
	"Find a hidden treasure":  "Shovel",
	"Cure a disease":          "Vaccine",
	"Build a shelter":         "Axe",
	"Keep yourself warm":      "Fire",
	"Battle a dragon":         "Shield",
	"Fight a monster":         "Sword",
	"Survive a storm":         "Raincoat",
	"Escape a maze":           "Map",
	"Darkness in the forest":  "Lantern",
	"Dive into the ocean":     "Scuba gear",
	"Defend a castle":         "Bow and arrow",
	"Cross a river":           "Rope",
	"Climb a wall":            "Ladder",
}

func showChallenges(challengeQueue *Queue) {
	fmt.Println("Memorize the challenges and their order!")
	for e := challengeQueue.list.Front(); e != nil; e = e.Next() {
		fmt.Println("Challenge:", e.Value)
		time.Sleep(3 * time.Second)
	}
	for i := 0; i < 20; i++ {
		fmt.Println()
	}
	fmt.Println("\nTime's up! Now stack the correct tools.")
}

func getPlayerStack(numChallenges int) *Stack {
	stack := &Stack{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Available tools to choose from:")
	for tool := range tools {
		fmt.Println("-", tools[tool])
	}
	fmt.Println()

	for i := 1; i <= numChallenges; i++ {
		fmt.Printf("Enter tool #%d: ", i)
		tool, _ := reader.ReadString('\n')
		stack.Push(strings.TrimSpace(tool))
	}

	return stack
}

func main() {
	fmt.Println("Welcome to the Challenge Game!")
	fmt.Println("In each level, you will see a list of challenges in the order you need to complete them.")
	fmt.Println("Your goal is to stack the correct tools in reverse order. Remember, the last tool you put in the stack is the first one you’ll use!")
	fmt.Println("Good luck!\n")

	level := 1

	for {
		fmt.Printf("\n--- Level %d ---\n", level)
		challengeQueue := NewQueue()

		for i := 0; i < level+2; i++ {
			challenge := challenges[rand.Intn(len(challenges))]
			challengeQueue.Enqueue(challenge)
		}

		showChallenges(challengeQueue)

		playerStack := getPlayerStack(challengeQueue.list.Len())
		isCorrect := true

		for e := challengeQueue.list.Front(); e != nil && isCorrect; e = e.Next() {
			expectedTool := strings.ToLower(tools[e.Value.(string)])
			playerTool := playerStack.Pop()

			if playerTool != expectedTool {
				fmt.Printf("Incorrect! You needed %s but selected %s. Try again.\n", expectedTool, playerTool)
				isCorrect = false
			}
		}

		if isCorrect {
			fmt.Println("Congratulations! You completed the level!")
			level++
		} else {
			fmt.Println("Game over! Restarting the level.")
		}
	}
}
