package structs_services

type JokeService interface {
	GiveMeAJoke() (joke string, err error)
	GiveMeJokeAtIndex(jokeIndex int) (joke string, err error)
	RemoveJokeAtIndex(jokeIndex int) (err error)
	JokeCount() (jokeCount int, err error)
	AddJoke(joke string) (err error)
}
