# notabar
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build
	scdoc < notabar.1.scd | sed "s/VERSION/$(VERSION)/g" > notabar.1
	scdoc < notabar.5.scd | sed "s/VERSION/$(VERSION)/g" > notabar.5

clean:
	rm -f notabar
	rm -f notabar.1
	rm -f notabar.5

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f notabar $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/notabar
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f notabar.1 $(DESTDIR)$(MANPREFIX)/man1/notabar.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/notabar.1
	cp -f notabar.5 $(DESTDIR)$(MANPREFIX)/man5/notabar.5
	chmod 644 $(DESTDIR)$(MANPREFIX)/man5/notabar.5

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/notabar
	rm -f $(DESTDIR)$(MANPREFIX)/man1/notabar.1
	rm -f $(DESTDIR)$(MANPREFIX)/man5/notabar.5

.PHONY: all build clean install uninstall
