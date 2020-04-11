#include <fluidsynth.h>
#include "clibrary.h"

void CallMyFunction(unsigned int time, fluid_event_t *event, fluid_sequencer_t *seq, void *data) {
  go_sequencer_callback(time, event, seq, data);
}