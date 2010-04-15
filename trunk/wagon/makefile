include $(GOROOT)/src/Make.$(GOARCH)

all:		$(O) $(O)/wagon

clean:
	rm -rf [568vq] *.[568vq]
	
$(O):
	mkdir $(O)

$(O)/wagon:	$(O) list.a wagon.a wagonmain.$(O)
	$(O)l -o $(O)/wagon -L . wagonmain.$(O)

list.a:		list.go
	$(O)g -o list.a list.go

wagon.a:	list.a wagon.go
	$(O)g -o wagon.a -I . wagon.go

wagonmain.$(O)	wagonmain.go list.a wagon.a
	$(O)g -I . wagonmain.go
test:
	
