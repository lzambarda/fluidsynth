package fluidsynth

import (
	"fmt"
	"testing"
	"time"
)

func TestSequencerSendNow(t *testing.T) {
	settings := NewSettings()
	synth := NewSynth(settings)

	driver := NewAudioDriver(settings, synth)
	defer driver.Delete()

	hasWorked := synth.SFLoad("/Users/iyadassaf/Desktop/808.sf2", true)
	fmt.Println("has worked", hasWorked)

	seq := NewSequencer()
	defer seq.Delete()
	seq.RegisterSynth(synth)

	c := 0
	for {
		c++
		seq.SendNoteNow(1, 36, 127)
		time.Sleep(time.Second)
		if c > 10 {
			break
		}
	}
}

func TestSequencerSchedule(t *testing.T) {
	settings := NewSettings()
	synth := NewSynth(settings)

	driver := NewAudioDriver(settings, synth)
	defer driver.Delete()

	hasWorked := synth.SFLoad("/Users/iyadassaf/Desktop/808.sf2", true)
	fmt.Println("has worked", hasWorked)

	seq := NewSequencer()
	defer seq.Delete()
	seq.RegisterSynth(synth)

	var tm time.Duration
	for i := 0; i<10; i++ {
		tm += time.Second
		fmt.Println("Sending note at", tm)
		if err := seq.ScheduleSendNote(1, 36, 127, tm); err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(time.Second * 10)
}

