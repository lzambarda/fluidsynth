package fluidsynth

/*
#cgo pkg-config: fluidsynth
#include <fluidsynth.h>
#include <stdlib.h>

typedef void (*closure)();
void go_sequencer_callback(unsigned int time, fluid_event_t *event, fluid_sequencer_t *seq, void *data);
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

//export go_sequencer_callback
func go_sequencer_callback(t C.uint, evt *C.fluid_event_t, seq *C.fluid_sequencer_t, data unsafe.Pointer) {
}

type Sequencer struct {
	ptr       *C.fluid_sequencer_t
	synthPtr  C.fluid_seq_id_t
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
	s.clientPtr = C.fluid_sequencer_register_client(s.ptr, C.CString("sequencer"), C.closure(C.go_sequencer_callback), C.NULL)
}

func (s *Sequencer) SendNoteNow(ch, note, velocity uint8) {
	evt := C.new_fluid_event()
	C.fluid_event_set_source(evt, s.clientPtr)
	C.fluid_event_set_dest(evt, s.synthPtr)
	C.fluid_event_timer(evt, C.NULL)
	C.fluid_event_noteon(evt, C.int(ch), C.short(note), C.short(velocity))
	C.fluid_sequencer_send_now(s.ptr, evt)
	C.fluid_sequencer_unregister_client(s.ptr, s.clientPtr)
	C.fluid_event_unregistering(evt)
	C.delete_fluid_event(evt)
}

func (s *Sequencer) ScheduleSendNote(ch, note, velocity uint8, t time.Duration) error {
	evt := C.new_fluid_event()
	C.fluid_event_set_source(evt, s.clientPtr)
	C.fluid_event_set_dest(evt, s.synthPtr)
	C.fluid_event_noteon(evt, C.int(ch), C.short(note), C.short(velocity))
	C.fluid_event_timer(evt, C.NULL)

	if C.fluid_sequencer_send_at(s.ptr, evt, C.uint(t.Milliseconds()), 0) != C.FLUID_OK {
		return fmt.Errorf("failed to schedule time")
	}

	C.fluid_event_unregistering(evt)
	C.delete_fluid_event(evt)
	return nil
}
