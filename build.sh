

set -e

gcc -c clibrary.c
ar cru libclibrary.a clibrary.o

go test ./... -v