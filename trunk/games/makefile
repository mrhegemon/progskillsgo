include $(GOROOT)/src/Make.$(GOARCH)

all:	$(O) $(O)/rps-stdin $(O)/ttt-stdin $(O)/rps-net

clean:
	rm -rf [568vq] *.[586vq]
	rm *~

$(O):
	mkdir $(O)

$(O)/rps-stdin:	 rps-main.$(O)
	$(O)l -o $(O)/rps-stdin -L . rps-main.$(O)

$(O)/ttt-stdin:  ttt-stdin.$(O)
	$(O)l -o $(O)/ttt-stdin -L . ttt-stdin.$(O)

$(O)/rps-net:	net.$(O)
	$(O)l -o $(O)/rps-net -L . net.$(O)

net.$(O):	net.go view.$(O) rps.$(O)
	$(O)g -I . net.go

games.$(O): games.go
	$(O)g games.go

rps-main.$(O):	view.$(O) rps.$(O) rps-main.go
	$(O)g -I . rps-main.go
	
view.$(O):	games.$(O)
	$(O)g -I . view.go

rps.$(O):	rps.go games.$(O)
	$(O)g -I . rps.go

ttt-stdin.$(O):	games.$(O) view.$(O) ttt.go
	$(O)g -o ttt-stdin.$(O) -I . ttt.go


test:	all
	$(O)/rps-stdin
	$(O)/ttt-stdin
