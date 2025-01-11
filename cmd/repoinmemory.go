package cmd

import (
	"github.com/Adit0507/pomodoro-timer/pomodoro"
	"github.com/Adit0507/pomodoro-timer/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}