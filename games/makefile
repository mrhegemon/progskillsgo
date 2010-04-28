include $(GOROOT)/src/Make.$(GOARCH)

all:	$(O) $(O)/rps-stdin

clean:
	rm -rf [568vq] *.[586vq]
	rm *~

$(O):
	mkdir $(O)

$(O)/rps-stdin: games.$(O) view.$(O) rps-stdin.$(O)
	$(O)l -o $(O)/rps-stdin -L . rps-stdin.$(O)

games.$(O): games.go
	$(O)g games.go
	
view.$(O):	games.$(O)
	$(O)g -I . view.go

rps-stdin.$(O):	games.$(O) view.$(O) rps.go
	$(O)g -o rps-stdin.$(O) -I . rps.go

