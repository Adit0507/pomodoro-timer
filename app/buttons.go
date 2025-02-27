package app

import (
	"context"
	"fmt"

	"github.com/Adit0507/pomodoro-timer/pomodoro"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
)

type buttonSet struct {
	btStart *button.Button
	btPause *button.Button
}

func newButtonSet(ctx context.Context, config *pomodoro.IntervalConfig, w *widgets, redrawCh chan<- bool, errorCh chan<- error) (*buttonSet, error) {
	startInterval := func() {
		i, err := pomodoro.GetInterval(config)
		errorCh <- err

		start := func(i pomodoro.Interval) {
			msg := "Take a break"
			if i.Category == pomodoro.CategoryPomodoro {
				msg = "Focus on your task"
			}
			w.update([]int{}, i.Category, msg, "", redrawCh)
		}

		end := func(pomodoro.Interval) {
			w.update([]int{}, "", "Nothing running...", "", redrawCh)
		}
		periodic := func(i pomodoro.Interval) {
			w.update([]int{int(i.ActualDuration), int(i.PlannedDuration)},
				"", "",
				fmt.Sprint(i.PlannedDuration-i.ActualDuration),
				redrawCh,
			)
		}
		errorCh <- i.Start(ctx, config, start, periodic, end)
	}

	pauseInterval := func() {
		i, err := pomodoro.GetInterval(config)
		if err != nil {
			errorCh <- err
			return
		}

		if err := i.Pause(config); err != nil {
			if err == pomodoro.ErrIntervalNotRunning {
				return
			}

			errorCh <- err
			return
		}
		w.update([]int{}, "", "Paused... pres start to continue", "", redrawCh)
	}

	btStart, err := button.New("(s)tart", func() error {
		go startInterval()
		return nil
	},
		button.GlobalKey('s'),
		button.WidthFor("(p)ause"),
		button.Height(2),
	)
	if err != nil {
		return nil, err
	}

	btPause, err := button.New("(p)ause", func() error {
		go pauseInterval()
		return nil
	},
		button.FillColor(cell.ColorNumber(220)),
		button.GlobalKey('p'),
		button.Height(2),
	)
	if err != nil {
		return nil, err
	}

	return &buttonSet{btStart, btPause}, nil
}
