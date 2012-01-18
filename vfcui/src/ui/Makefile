include $(GOROOT)/src/Make.inc

TARG=vfc/ui

GOFILES=\
	uidef.go \
	event.go \
	ui.go \
	
CGOFILES = \
	cdrv.go \

#CGO_CFLAGS=
CGO_LDFLAGS+= -L../lib -lqtdrv

# set CGO_CFLAGS or CGO_LDFLAGS.  For example, to use an
# alternate installation of the library:
#	CGO_CFLAGS=-I/home/example/include
#	CGO_LDFLAGS+=-L/home/example/lib
# Note the += on the second line.

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
