package structs_services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StringJokesService_InitialJokeAdded(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewStringJokesService(initialJokes...)

	actualLen, err := service.JokeCount()
	assert.Nil(t, err)
	assert.Equal(t, len(initialJokes), actualLen)

}

func Test_StringJokesService_JokesAddedSuccessfully(t *testing.T) {
	service := NewStringJokesService()

	joke := "Im funny"

	err := service.AddJoke(joke)
	assert.Nil(t, err)
	actualJoke, err := service.GiveMeJokeAtIndex(0)
	assert.Nil(t, err)

	assert.Equal(t, joke, actualJoke)
}

func Test_StringJokesService_GiveMeAJokeWorks(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewStringJokesService(initialJokes...)

	actualJoke, err := service.GiveMeAJoke()
	assert.Nil(t, err)

	assert.NotEqualf(t, actualJoke, "", "The joke service failed to give a joke %v", err)
}

func Test_StringJokesService_GiveMeAJokeFailsWhenNoJokes(t *testing.T) {
	service := NewStringJokesService()

	jokeIndex := 1
	_, err := service.GiveMeJokeAtIndex(jokeIndex)
	assert.NotNil(t, err)

}

func Test_StringJokesService_GiveMeNJokeWorks(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewStringJokesService(initialJokes...)

	jokeIndex := 1
	actualJoke, err := service.GiveMeJokeAtIndex(jokeIndex)
	assert.Nil(t, err)

	assert.Equal(t, initialJokes[jokeIndex], actualJoke)
}

func Test_StringJokesService_GiveMeNJokeFailsWhenNoJokes(t *testing.T) {
	service := NewStringJokesService()
	actualJoke, err := service.GiveMeJokeAtIndex(0)
	assert.NotNil(t, err)
	assert.Equal(t, "", actualJoke)

}

func Test_StringJokesService_GiveMeNJokeFailsWHenIndexOutofBounds(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewStringJokesService(initialJokes...)

	jokeIndex := 2
	actualJoke, err := service.GiveMeJokeAtIndex(jokeIndex)
	assert.NotNil(t, err)
	assert.Equal(t, "", actualJoke)
}

func Test_StringJokesService_JokeCountReturnsCountOfJokes(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewStringJokesService(initialJokes...)

	jokeCount, err := service.JokeCount()
	assert.Nil(t, err)
	assert.Equal(t, len(initialJokes), jokeCount)
}

func Test_StringJokesService_RemoveJokeAtIndexWorks(t *testing.T) {
	initialJokes := []string{"Funny One", "Joke Two", "Hahah three"}
	service := NewStringJokesService(initialJokes...)

	err := service.RemoveJokeAtIndex(1)
	assert.Nil(t, err)

	newJokeCount, err := service.JokeCount()
	assert.Nil(t, err)
	assert.Equal(t, len(initialJokes)-1, newJokeCount)

	jokeOne, err := service.GiveMeJokeAtIndex(0)
	assert.Nil(t, err)
	assert.Equal(t, initialJokes[0], jokeOne)
	jokeTwo, err := service.GiveMeJokeAtIndex(1)
	assert.Equal(t, initialJokes[2], jokeTwo)

}
