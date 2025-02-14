package structs_services

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockJokeService struct {
	GiveMeAJokeResult string
	AddedJoke         string
	CountOfJokes      int
}

func (service *mockJokeService) GiveMeAJoke() (joke string, err error) {
	return service.GiveMeAJokeResult, nil
}

func (service *mockJokeService) GiveMeJokeAtIndex(jokeIndex int) (joke string, err error) {
	return fmt.Sprintf("JOKE:%v", jokeIndex), nil
}

func (service *mockJokeService) RemoveJokeAtIndex(jokeIndex int) (err error) {
	return nil
}

func (service *mockJokeService) JokeCount() (jokeCount int, err error) {
	return service.CountOfJokes, nil
}

func (service *mockJokeService) AddJoke(joke string) (err error) {
	if joke == "" {
		service.AddedJoke = ""
		return errors.New("joke is empty")
	}
	service.AddedJoke = joke
	return nil
}

func TestJoker_AddJoke(t *testing.T) {
	type fields struct {
		Jokes *mockJokeService
	}

	type args struct {
		joke string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		//wantErr  assert.ErrorAssertionFunc
		wantJoke string
	}{
		{
			name: "success",
			fields: fields{
				Jokes: &mockJokeService{},
			},
			args:     args{"JOKE:JOKE"},
			wantErr:  false,
			wantJoke: "JOKE:JOKE",
		},
		{
			name: "Fails due to empty joke",
			fields: fields{
				Jokes: &mockJokeService{},
			},
			args:     args{},
			wantErr:  true,
			wantJoke: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockService JokeService = tt.fields.Jokes
			j := &Joker{
				Jokes: mockService,
			}

			err := j.AddJoke(tt.args.joke)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddJoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//tt.wantErr(t, , fmt.Sprintf("AddJoke(%v)", tt.args.joke))

			assert.Equal(t, tt.args.joke, tt.fields.Jokes.AddedJoke)
		})
	}
}

//TODO:
/*
func TestJoker_JokeCount(t *testing.T) {
	type fields struct {
		Jokes JokeService
	}
	type args struct {
		joke string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int
		wantErr   assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Joker{
				Jokes: tt.fields.Jokes,
			}
			gotCount, err := j.JokeCount(tt.args.joke)
			if !tt.wantErr(t, err, fmt.Sprintf("JokeCount(%v)", tt.args.joke)) {
				return
			}
			assert.Equalf(t, tt.wantCount, gotCount, "JokeCount(%v)", tt.args.joke)
		})
	}
}

func TestNewJoker(t *testing.T) {
	type args struct {
		service JokeService
	}
	tests := []struct {
		name string
		args args
		want *Joker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewJoker(tt.args.service), "NewJoker(%v)", tt.args.service)
		})
	}
}
*/
