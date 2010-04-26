include $(GOROOT)/src/Make.$(GOARCH)

all:	$(O) $(O)/rps-stdin

clean:
	rm -rf [568vq] *.[586vq]
	rm *~
	
$(O):
	mkdir $(O)

$(O)/rps-stdin: games.(O) rps-stdin.$(O)
	$(O)l -o $(O)/rps-stdin -L . rps-stdin.$(O)

games.(O): games.go
	$(O)g games.go

rps-stdin.$(O):	games.(O) ttt.go
	$(O)g -o rps-stdin.$(O) -I . ttt.go

