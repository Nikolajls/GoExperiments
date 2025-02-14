package structs_services

import (
	"fmt"
	"math/rand/v2"
)

type InMemoryJokes struct {
	jokes []string
}

func NewInMemoryJokeService(initialJokes ...string) (jokeService *InMemoryJokes) {
	var jokes []string
	for _, joke := range initialJokes {
		jokes = append(jokes, joke)
	}

	service := &InMemoryJokes{
		jokes: jokes,
	}

	return service
}

func (inMemoryJoke *InMemoryJokes) GiveMeAJoke() (joke string, err error) {
	randomInt := rand.IntN(len(inMemoryJoke.jokes))
	joke, err = inMemoryJoke.GiveMeJokeAtIndex(randomInt)
	return joke, err
}

func (inMemoryJoke *InMemoryJokes) GiveMeJokeAtIndex(jokeIndex int) (joke string, err error) {
	jokeCount := len(inMemoryJoke.jokes)

	if jokeCount == 0 {
		return "", fmt.Errorf("joke service does not contain a joke")
	}

	if jokeIndex > jokeCount-1 {
		return "", fmt.Errorf("joke service does not contain joke at that index %v, max is %v", jokeIndex, jokeCount)
	}
	return inMemoryJoke.jokes[jokeIndex], nil
}

func (inMemoryJoke *InMemoryJokes) RemoveJokeAtIndex(jokeIndex int) (err error) {
	jokeCount := len(inMemoryJoke.jokes)

	if jokeCount == 0 {
		return fmt.Errorf("joke service does not contain a joke")
	}

	if jokeIndex > jokeCount-1 {
		return fmt.Errorf("joke service does not contain joke at that index %v, max is %v", jokeIndex, jokeCount)
	}
	inMemoryJoke.jokes = removeIndexFromSlice(inMemoryJoke.jokes, jokeIndex)
	return nil
}

func (inMemoryJoke *InMemoryJokes) JokeCount() (jokeCount int, err error) {

	return len(inMemoryJoke.jokes), nil
}

func (inMemoryJoke *InMemoryJokes) AddJoke(joke string) (err error) {
	inMemoryJoke.jokes = append(inMemoryJoke.jokes, joke)
	return nil
}

func removeIndexFromSlice(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
