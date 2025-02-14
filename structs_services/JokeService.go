package structs_services

type JokeService interface {
	GiveMeAJoke() (joke string, err error)
	GiveMeJokeAtIndex(jokeIndex int) (joke string, err error)
	RemoveJokeAtIndex(jokeIndex int) (err error)
	JokeCount() (jokeCount int, err error)
	AddJoke(joke string) (err error)
}

type Joker struct {
	Jokes JokeService
}

func NewJoker(service JokeService) *Joker {
	return &Joker{
		Jokes: service,
	}
}

func (j *Joker) AddJoke(joke string) (err error) {
	return j.Jokes.AddJoke(joke)
}

func (j *Joker) JokeCount(joke string) (count int, err error) {
	return j.Jokes.JokeCount()
}
