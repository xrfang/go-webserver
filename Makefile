GOMOD=webserver
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
HASH=$(shell git log -n1 --pretty=format:%h)
REVS=$(shell git log --oneline|wc -l)
native: release
arm: export GOOS=linux
arm: export GOARCH=arm
arm: export GOARM=7
arm: release
win: export GOOS=windows
win: export GOARCH=amd64
win: conv release
	#restore charset encoding after compiling
	git checkout resources/*
conv:
	#change line-ending and charset encoding for Windows
	find resources/init/ -type f |grep -v gfwlist|xargs -L1 -I% unix2dos -n % %
	find resources/init/ -type f |grep -v gfwlist|xargs -L1 -I% iconv -futf-8 -tgb18030 -o% %
debug: setver geneh compdbg pack
release: setver geneh comprel pack
geneh: #generate error handler
	@for tpl in `find . -type f |grep errors.tpl`; do \
	    target=`echo $$tpl|sed 's/\.tpl/\.go/'`; \
	    pkg=`basename $$(dirname $$tpl)`; \
		sed "s/package main/package $$pkg/" src/errors.go > $$target; \
    done
setver:
	cp src/verinfo.tpl src/version.go
	sed -i 's/{_BRANCH}/$(BRANCH)/' src/version.go
	sed -i 's/{_G_HASH}/$(HASH)/' src/version.go
	sed -i 's/{_G_REVS}/$(REVS)/' src/version.go
comprel:
	mkdir -p bin && cd src && go build -ldflags="-s -w" . && mv $(GOMOD)* ../bin
compdbg:
	mkdir -p bin && cd src && go build -race -gcflags=all=-d=checkptr=0 . && mv $(GOMOD)* ../bin
pack: export GOOS=
pack: export GOARCH=
pack: export GOARM=
pack:
	cd utils && go build . && ./pack && rm pack
clean:
	rm -fr bin src/version.go src/*/errors.go
	git checkout resources/*
