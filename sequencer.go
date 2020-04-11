package fluidsynth

// #cgo pkg-config: fluidsynth
// #include <fluidsynth.h>
// #include <stdlib.h>
// #include <clibrary.h>
/*


*/
import "C"
import (
	"fmt"
	"time"
)

//export go_sequencer_callback
func go_sequencer_callback(t C.int, evt *C.fluid_event_t, seq *C.fluid_sequencer_t, data interface{}) {

}

type Sequencer struct {
	ptr      *C.fluid_sequencer_t
	synthPtr C.fluid_seq_id_t
	clientPtr C.fluid_seq_id_t
}

func NewSequencer() *Sequencer {
	return &Sequencer{ptr: C.new_fluid_sequencer2(0)}
}

func (s *Sequencer) Delete() {
	C.delete_fluid_sequencer(s.ptr)
}

func (s *Sequencer) RegisterSynth(synth Synth) {
	s.synthPtr = C.fluid_sequencer_register_fluidsynth(s.ptr, synth.ptr)

	//sequencer_callback(unsigned int time, fluid_event_t *event, fluid_sequencer_t *seq, void *data)
	//cb := func(t C.int, evt *C.fluid_event_t, seq *C.fluid_sequencer_t, data interface{}) {
	//
	//}

	s.clientPtr = C.fluid_sequencer_register_client(s.ptr,
		C.CString("fluidsynth_metronome"), C.closure(C.CallMyFunction), C.NULL);

}

func (s *Sequencer) SendNoteNow(ch, note, velocity uint8) {
	evt := C.new_fluid_event()
	C.fluid_event_set_source(evt, s.clientPtr)
	C.fluid_event_set_dest(evt, s.synthPtr)
	C.fluid_event_noteon(evt, C.int(ch), C.short(note), C.short(velocity))
	C.fluid_sequencer_send_now(s.ptr, evt)
	C.fluid_event_unregistering(evt)
	C.delete_fluid_event(evt)
}

func (s *Sequencer) ScheduleSendNote(ch, note, velocity uint8, t time.Duration) error {
	evt := C.new_fluid_event()
	C.fluid_event_set_source(evt, s.clientPtr)
	C.fluid_event_set_dest(evt, s.synthPtr)
	C.fluid_event_noteon(evt, C.int(ch), C.short(note), C.short(velocity))

	if C.fluid_sequencer_send_at(s.ptr, evt, C.uint(t.Milliseconds()), 0) != C.FLUID_OK {
		return fmt.Errorf("failed to schedule time")
	}

	C.fluid_event_unregistering(evt)
	C.delete_fluid_event(evt)
	return nil
}
