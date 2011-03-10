include  $(GOROOT)/src/Make.inc

TARG=fann

CGOFILES=\
	fann.go\

CGO_LDFLAGS=-lfann -lm

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
