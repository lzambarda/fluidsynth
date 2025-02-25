package fluidsynth

/*
#cgo pkg-config: fluidsynth
#include <fluidsynth.h>
#include <stdlib.h>

typedef void (*closure)();
*/
import "C"
import "unsafe"

type AudioDriver struct {
	ptr *C.fluid_audio_driver_t
}

func NewAudioDriver(settings Settings, synth Synth) AudioDriver {
	return AudioDriver{C.new_fluid_audio_driver(settings.ptr, synth.ptr)}
}

func (d *AudioDriver) Delete() {
	C.delete_fluid_audio_driver(d.ptr)
}

type FileRenderer struct {
	ptr *C.fluid_file_renderer_t
}

func NewFileRenderer(synth Synth) FileRenderer {
	return FileRenderer{C.new_fluid_file_renderer(synth.ptr)}
}

func (r *FileRenderer) Delete() {
	C.delete_fluid_file_renderer(r.ptr)
}

func (r *FileRenderer) ProcessBlock() bool {
	return C.fluid_file_renderer_process_block(r.ptr) == C.FLUID_OK
}

type MIDIDriver struct {
	ptr *C.fluid_midi_driver_t
}

func NewMIDIDriver(settings Settings, synth Synth) MIDIDriver {
	return MIDIDriver{C.new_fluid_midi_driver(settings.ptr, C.closure(C.fluid_synth_handle_midi_event), unsafe.Pointer(synth.ptr))}
}

func (d *MIDIDriver) Delete() {
	C.delete_fluid_midi_driver(d.ptr)
}
