package structs_services

import (
	"fmt"
	"math/rand/v2"
)

type InMemoryJokes struct {
	jokes []string
}

func NewInMemoryJokeService(initialJokes ...string) (inmemoryJokes JokeService) {
	var jokes []string
	for _, joke := range initialJokes {
		jokes = append(jokes, joke)
	}

	service := &InMemoryJokes{
		jokes: jokes,
	}

	return service
}

func (inmemoryJoke *InMemoryJokes) GiveMeAJoke() (joke string, err error) {
	randomInt := rand.IntN(len(inmemoryJoke.jokes))
	joke, err = inmemoryJoke.GiveMeJokeAtIndex(randomInt)
	return joke, err
}

func (inmemoryJoke *InMemoryJokes) GiveMeJokeAtIndex(jokeIndex int) (joke string, err error) {
	jokeCount := len(inmemoryJoke.jokes)

	if jokeCount == 0 {
		return "", fmt.Errorf("Joke service does not contain a joke")
	}

	if jokeIndex > jokeCount-1 {
		return "", fmt.Errorf("Joke service does not contain joke at that index %v, max is %v", jokeIndex, jokeCount)
	}
	return inmemoryJoke.jokes[jokeIndex], nil
}

func (inmemoryJoke *InMemoryJokes) RemoveJokeAtIndex(jokeIndex int) (err error) {
	jokeCount := len(inmemoryJoke.jokes)

	if jokeCount == 0 {
		return fmt.Errorf("Joke service does not contain a joke")
	}

	if jokeIndex > jokeCount-1 {
		return fmt.Errorf("Joke service does not contain joke at that index %v, max is %v", jokeIndex, jokeCount)
	}
	inmemoryJoke.jokes = removeIndexFromSlice(inmemoryJoke.jokes, jokeIndex)
	return nil
}

func (inmemoryJoke *InMemoryJokes) JokeCount() (jokeCount int, err error) {

	return len(inmemoryJoke.jokes), nil
}

func (inmemoryJoke *InMemoryJokes) AddJoke(joke string) (err error) {
	inmemoryJoke.jokes = append(inmemoryJoke.jokes, joke)
	return nil
}

func removeIndexFromSlice(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
