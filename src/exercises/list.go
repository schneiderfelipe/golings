package exercises

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var ErrExerciseNotFound = errors.New("exercise not found")

type Info struct {
	Exercises []Exercise
}

func List() ([]Exercise, error) {
	var info Info

	data, err := os.ReadFile("info.toml")
	if err != nil {
		return info.Exercises, err
	}

	if err := toml.Unmarshal(data, &info); err != nil {
		return info.Exercises, err
	}

	return info.Exercises, nil
}

func Find(exercise string) (Exercise, error) {
	exs, err := List()
	if err != nil {
		return Exercise{}, err
	}

	for _, ex := range exs {
		if ex.Name == exercise {
			return ex, nil
		}
	}

	return Exercise{}, ErrExerciseNotFound
}
