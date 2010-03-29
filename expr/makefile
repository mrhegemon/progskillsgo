cmd = expr
include $(GOROOT)/src/Make.$(GOARCH)

all: $O/$(cmd)
clean:  ; rm -rf [$(OS)] *.[$(OS)]
$O:     ; mkdir $@
$O/%: $O %.go; $(GC) $(GCFLAGS) -o $*.$O $*.go && \
	$(LD) $(LDFLAGS) -o $* $*.$O
test: all ; $0/$(cmd)