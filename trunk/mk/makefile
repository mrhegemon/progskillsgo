include $(GOROOT)/src/Make.$(GOARCH)

all:		$(O) $(O)/dag $(O)/mk

clean:
	rm -rf [568vq] *.[568vq]
	
$(O):
	mkdir $(O)

$(O)/dag:	$(O) dagmain.$(O) dag.$(O)
	$(O)l -o $(O)/dag -L . dagmain.$(O)
	
dag.$(O):	dag.go dagimp.go
	$(O)g -o dag.$(O) dag.go dagimp.go

dagmain.$(O):	dagmain.go dag.$(O)
	$(O)g -I . dagmain.go

$(O)/mk:	$(O) mkmain.$(O) dag.$(O) mk.$(O)
	$(O)l -o $(O)/mk -L . mkmain.$(O)
	
mk.$(O):
	$(O)g -I . mk.go

mkmain.$(O):	mkmain.go dag.$(O) mk.$(O)
	$(O)g -I . mkmain.go

test:
	$(O)/mk
	$(O)/mk clean
	$(O)/mk dag.$(O)
