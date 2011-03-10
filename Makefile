include  $(GOROOT)/src/Make.inc

TARG=dp

CGOFILES=\
	fann.go\

CGO_LDFLAGS=-lfann -lm

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
