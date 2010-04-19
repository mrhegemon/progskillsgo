include $(GOROOT)/src/Make.$(GOARCH)

all:		$(O) $(O)/wagon $(O)/list_tester

clean:
	rm -rf [568vq] *.[568vq]
	rm *~
	
$(O):
	mkdir $(O)

$(O)/wagon:	$(O) list.$(O) wagon.$(O) wagonmain.$(O)
	$(O)l -o $(O)/wagon -L . wagonmain.$(O)

list.$(O):		list.go
	$(O)g -o list.$(O) list.go

wagon.$(O):	list.$(O) wagon.go
	$(O)g -o wagon.$(O) -I . wagon.go

wagonmain.$(O):	wagonmain.go list.$(O) wagon.$(O)
	$(O)g -I . wagonmain.go

$(O)/list_tester:   $(O) list.$(O) listTester.$(O)
	$(O)l -o $(O)/list_tester -L . listTester.$(O)

listTester.$(O):	listTester.go
	$(O)g -I . listTester.go

test:	$(O)/list_tester
	#run linked list unit tests.
	#only prints output on error
	$(O)/list_tester

verbose_test: $(O)/list_tester
	#runs tests in verbose mode
	$(O)/list_tester -v
	
