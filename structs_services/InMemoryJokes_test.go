package structs_services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InMemoryJokes_InitialJokeAdded(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewInMemoryJokeService(initialJokes...)

	actualLen, err := service.JokeCount()
	assert.Nil(t, err)
	assert.Equal(t, len(initialJokes), actualLen)

}

func Test_InMemoryJokes_JokesAddedSuccessfully(t *testing.T) {
	service := NewInMemoryJokeService()

	joke := "Im funny"

	err := service.AddJoke(joke)
	assert.Nil(t, err)
	actualJoke, err := service.GiveMeJokeAtIndex(0)
	assert.Nil(t, err)

	assert.Equal(t, joke, actualJoke)
}

func Test_InMemoryJokes_GiveMeAJokeWorks(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewInMemoryJokeService(initialJokes...)

	actualJoke, err := service.GiveMeAJoke()
	assert.Nil(t, err)

	assert.NotEqualf(t, actualJoke, "", "The joke service failed to give a joke %v", err)
}

func Test_InMemoryJokes_GiveMeAJokeFailsWhenNoJokes(t *testing.T) {
	service := NewInMemoryJokeService()

	jokeIndex := 1
	_, err := service.GiveMeJokeAtIndex(jokeIndex)
	assert.NotNil(t, err)

}

func Test_InMemoryJokes_GiveMeNJokeWorks(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewInMemoryJokeService(initialJokes...)

	jokeIndex := 1
	actualJoke, err := service.GiveMeJokeAtIndex(jokeIndex)
	assert.Nil(t, err)

	assert.Equal(t, initialJokes[jokeIndex], actualJoke)
}

func Test_InMemoryJokes_GiveMeNJokeFailsWhenNoJokes(t *testing.T) {
	service := NewInMemoryJokeService()
	actualJoke, err := service.GiveMeJokeAtIndex(0)
	assert.NotNil(t, err)
	assert.Equal(t, "", actualJoke)

}

func Test_InMemoryJokes_GiveMeNJokeFailsWHenIndexOutofBounds(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewInMemoryJokeService(initialJokes...)

	jokeIndex := 2
	actualJoke, err := service.GiveMeJokeAtIndex(jokeIndex)
	assert.NotNil(t, err)
	assert.Equal(t, "", actualJoke)
}

func Test_InMemoryJokes_JokeCountReturnsCountOfJokes(t *testing.T) {
	initialJokes := []string{"Funny", "Joke Two"}
	service := NewInMemoryJokeService(initialJokes...)

	jokeCount, err := service.JokeCount()
	assert.Nil(t, err)
	assert.Equal(t, len(initialJokes), jokeCount)
}

func Test_InMemoryJokes_RemoveJokeAtIndexWorks(t *testing.T) {
	initialJokes := []string{"Funny One", "Joke Two", "Hahah three"}
	service := NewInMemoryJokeService(initialJokes...)

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
