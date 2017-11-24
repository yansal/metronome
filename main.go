package main

import (
	"strconv"
	"time"

	"honnef.co/go/js/dom"
)

func main() {
	var (
		taps         []time.Time
		interval     int
		window       = dom.GetWindow()
		document     = window.Document()
		clickAudio   = document.GetElementByID("click").(*dom.HTMLAudioElement)
		tempoDisplay = document.GetElementByID("tempo")
	)

	document.AddEventListener("touchstart", true, func(event dom.Event) {
		taps = append(taps, time.Now())
		if len(taps) < 2 {
			return
		}
		if len(taps) > 4 {
			taps = taps[1:]
		}
		var diffs []time.Duration
		for i := 0; i < len(taps)-1; i++ {
			diffs = append(diffs, taps[i+1].Sub(taps[i]))
		}
		avgMs := int(avg(diffs) / time.Millisecond)
		tempo := strconv.Itoa(60000 / avgMs)

		window.ClearInterval(interval)
		interval = window.SetInterval(func() {
			clickAudio.Play()
		}, avgMs)
		clickAudio.Play()
		tempoDisplay.SetTextContent(tempo)
	})
}

func avg(durations []time.Duration) time.Duration {
	var sum time.Duration
	for _, i := range durations {
		sum += i
	}
	return sum / time.Duration(len(durations))
}
