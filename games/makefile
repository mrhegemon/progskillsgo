include $(GOROOT)/src/Make.$(GOARCH)

all:	$(O) $(O)/rps-net $(O)/rps-stdin $(O)/ttt-stdin $(O)/rps-net  $(O)/ttt-net #$(O)/rps-stdin-ai

clean:
	rm -rf [568vq] *.[586vq]
	rm *~

$(O):
	mkdir $(O)

$(O)/rps-stdin:	 rps-main.$(O)
	$(O)l -o $(O)/rps-stdin -L . rps-main.$(O)

$(O)/rps-stdin-ai:	rps-main-ai.$(O)
	$(O)l -o $(O)/rps-stdin-ai -L . rps-main-ai.$(O)

$(O)/ttt-stdin:  ttt-main.$(O)
	$(O)l -o $(O)/ttt-stdin -L . ttt-main.$(O)

$(O)/rps-net:	rps-net.$(O)
	$(O)l -o $(O)/rps-net -L . rps-net.$(O)

$(O)/ttt-net:	ttt-net.$(O)
	$(O)l -o $(O)/ttt-net -L . ttt-net.$(O)

rps-net.$(O):	rps-net.go view.$(O) rps.$(O)
	$(O)g -I . rps-net.go

ttt-net.$(O):	ttt-net.go view.$(O) ttt.$(O)
	$(O)g -I . ttt-net.go

games.$(O): games.go
	$(O)g games.go

rps-main.$(O):	view.$(O) rps.$(O) rps-main.go
	$(O)g -I . rps-main.go

rps-main-ai.$(O):	view.$(O) rps.$(O) ai.$(O) rps-main-ai.go
	$(O)g -I . rps-main-ai.go

ai.$(O):	games.$(O) ai.go
	$(O)g -I . ai.go
	
view.$(O):	games.$(O) sstruct.$(O)
	$(O)g -I . view.go

rps.$(O):	rps.go games.$(O) sstruct.$(O) view.$(O)
	$(O)g -I . rps.go

ttt-main.$(O): view.$(O) ttt.$(O) ttt-main.go
	$(O)g -o ttt-main.$(O) -I . ttt-main.go

ttt.$(O):	ttt.go games.$(O)
	$(O)g -I . ttt.go

sstruct.$(O):	sstruct.go
	$(O)g sstruct.go


test:	all
	$(O)/rps-stdin
	$(O)/ttt-stdin
