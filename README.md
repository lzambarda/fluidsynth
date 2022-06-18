# FluidSynth go bindings

Please refer to [github.com/jonathanslenders/fluidsynth](github.com/jonathanslenders/fluidsynth) to see the underlying C types.


## Getting started

This repo depends on fluidsynth.

```bash
# Set up dependencies on macOS
brew install portmidi fluidsynth pkg-config readline
# or (currently only supporting macOS)
make dependencies

# Testing
source ./config/dev.env
go test ./...
# or
make unit_test
```

Check out `config/dev.env` to see how `PKG_CONFIG_PATH` is correctly set for the `readline` package.
