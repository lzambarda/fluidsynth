package fluidsynth

import (
	"fmt"
	"testing"
	"time"
)

func TestSequencer(t *testing.T) {
	settings := NewSettings()
	synth := NewSynth(settings)

	driver := NewAudioDriver(settings, synth)
	defer driver.Delete()

	hasWorked := synth.SFLoad("/Users/iyadassaf/go/src/github.com/form3tech/f3-plugin-drummachine/sounds/808.sf2", true)
	fmt.Println("has worked", hasWorked)


	seq := NewSequencer()
	defer seq.Delete()
	seq.RegisterSynth(synth)

	for {
		seq.SendNote(1, 36, 127)
		time.Sleep(time.Second)
	}
}

