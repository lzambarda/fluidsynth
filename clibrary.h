#include <fluidsynth.h>
#include <stdlib.h>

typedef void (*closure)();

void CallMyFunction(unsigned int time, fluid_event_t *event, fluid_sequencer_t *seq, void *data);

//void go_sequencer_callback(unsigned int time, fluid_event_t *event, fluid_sequencer_t *seq, void *data);