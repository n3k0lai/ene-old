package Trivia

import (
	"math/rand"
	"time"

	Plugins "github.com/n3k0lai/ene/internal/plugins"
)

type Trivia struct {
	*Plugins.Plugin
	Questions      []Question
	ActiveQuestion Question
}

func (t *Trivia) Test(query string) bool {
	if t.ActiveQuestion.Solved {
		if query == "!trivia" {
			// ask question
			t.ActiveQuestion = t.GetQuestion()
			t.ActiveQuestion.Ask()
			return true
		} else {
			return false
		}
	}

	return t.ActiveQuestion.AddQuery(query)
}

func NewTrivia() *Trivia {
	rand.Seed(time.Now().UnixNano())
	return &Trivia{
		Plugin:    &Plugins.Plugin{},
		Questions: []Question{},
	}
}

func (t *Trivia) GetQuestion() Question {
	return t.Questions[rand.Intn(len(t.Questions))]
}
func (t *Trivia) Reset() {
	t.ActiveQuestion.Reset()
}
