package structs_services

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
)

type StringJokes struct {
	json string
}

type JokeStorage struct {
	Jokes []string `json:"jokes"`
}

func NewStringJokesService(initialJokes ...string) (stringJokeService *StringJokes) {
	var jokes []string
	for _, joke := range initialJokes {
		jokes = append(jokes, joke)
	}

	service := &StringJokes{
		json: "",
	}
	err := service.SaveJokeStorage(&JokeStorage{
		Jokes: jokes,
	})
	if err != nil {

	}

	return service
}

func (service *StringJokes) GiveMeAJoke() (joke string, err error) {
	jokeCount, err := service.JokeCount()
	if err != nil {
		return "", err
	}

	randomInt := rand.IntN(jokeCount)
	joke, err = service.GiveMeJokeAtIndex(randomInt)
	return joke, err
}

func (service *StringJokes) GiveMeJokeAtIndex(jokeIndex int) (joke string, err error) {
	jokeStorage, err := service.GetJokeStorage()

	if err != nil {
		return "", err
	}
	jokeCount := len(jokeStorage.Jokes)

	if jokeCount == 0 {
		return "", fmt.Errorf("Joke service does not contain a joke")
	}

	if jokeIndex > jokeCount-1 {
		return "", fmt.Errorf("Joke service does not contain joke at that index %v, max is %v", jokeIndex, jokeCount)
	}
	return jokeStorage.Jokes[jokeIndex], nil
}

func (service *StringJokes) RemoveJokeAtIndex(jokeIndex int) (err error) {
	jokeStorage, err := service.GetJokeStorage()

	if err != nil {
		return err
	}
	jokeCount := len(jokeStorage.Jokes)
	if jokeCount == 0 {
		return fmt.Errorf("Joke service does not contain a joke")
	}

	if jokeIndex > jokeCount-1 {
		return fmt.Errorf("Joke service does not contain joke at that index %v, max is %v", jokeIndex, jokeCount)
	}
	jokeStorage.Jokes = removeIndexFromSlice(jokeStorage.Jokes, jokeIndex)

	err = service.SaveJokeStorage(jokeStorage)
	if err != nil {
		return
	}

	return nil
}

func (service *StringJokes) JokeCount() (jokeCount int, err error) {
	jokeStorage, err := service.GetJokeStorage()

	if err != nil {
		return 0, err
	}

	return len(jokeStorage.Jokes), nil

}

func (service *StringJokes) AddJoke(joke string) (err error) {
	jokeStorage, err := service.GetJokeStorage()
	if err != nil {
		return
	}

	jokeStorage.Jokes = append(jokeStorage.Jokes, joke)

	err = service.SaveJokeStorage(jokeStorage)
	if err != nil {
		return
	}

	return
}

func (service *StringJokes) GetJokeStorage() (jokeStorage *JokeStorage, err error) {
	var jokes JokeStorage

	err = json.Unmarshal([]byte(service.json), &jokes)
	if err != nil {
		return nil, err
	}

	return &jokes, nil
}

func (service *StringJokes) SaveJokeStorage(jokeStorage *JokeStorage) (err error) {
	jsonData, err := json.Marshal(jokeStorage)
	if err != nil {
		return
	}

	service.json = string(jsonData)
	return
}
