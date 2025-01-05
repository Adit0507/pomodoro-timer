package app

import (
	"context"

	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"
)

type widgets struct {
	donTimer       *donut.Donut
	disType        *segmentdisplay.SegmentDisplay
	txtInfo        *text.Text
	updateDonTimer chan []int
	updateTxtInfo  chan string
	updateTxtTimer chan string
	updateTxtType  chan string
}

func (w *widgets) update(timer []int,
	txtType, txtInfo, txtTimer string, //txtType: update the SegmentDisplay, txtInfo and txtTimer to update Text widgets
	redrawCh chan<- bool) { //redrawch of vhannel bool indicates when the app should redraw the screen
	if txtInfo != "" {
		w.updateTxtInfo <- txtInfo
	}

	if txtType != "" {
		w.updateTxtType <- txtType
	}
	if txtTimer != "" {
		w.updateTxtTimer <- txtTimer
	}

	if len(timer) > 0 {
		w.updateDonTimer <- timer
	}

	redrawCh <- true
}
