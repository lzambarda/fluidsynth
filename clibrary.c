#include <fluidsynth.h>


void CallMyFunction(unsigned int time, fluid_event_t *event, fluid_sequencer_t *seq, void *data) {
  printf("call my function");
}