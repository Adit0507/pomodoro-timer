package pomodoro_test

import (
	"testing"

	"github.com/Adit0507/pomodoro-timer/pomodoro"
	"github.com/Adit0507/pomodoro-timer/repository"
)

// returns repository instance and cleanup func.
func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}