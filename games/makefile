include $(GOROOT)/src/Make.$(GOARCH)

all:	$(O) $(O)/rps-stdin $(O)/ttt-stdin

clean:
	rm -rf [568vq] *.[586vq]
	rm *~

$(O):
	mkdir $(O)

$(O)/rps-stdin: games.$(O) view.$(O) rps-stdin.$(O)
	$(O)l -o $(O)/rps-stdin -L . rps-stdin.$(O)

$(O)/ttt-stdin: games.$(O) view.$(O) ttt-stdin.$(O)
	$(O)l -o $(O)/ttt-stdin -L . ttt-stdin.$(O)

games.$(O): games.go
	$(O)g games.go
	
view.$(O):	games.$(O)
	$(O)g -I . view.go

rps-stdin.$(O):	games.$(O) view.$(O) rps.go
	$(O)g -o rps-stdin.$(O) -I . rps.go

ttt-stdin.$(O):	games.$(O) view.$(O) ttt.go
	$(O)g -o ttt-stdin.$(O) -I . ttt.go

test:	all
	$(O)/rps-stdin
	$(O)/ttt-stdin
