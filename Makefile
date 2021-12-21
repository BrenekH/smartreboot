.PHONY: compile clean docs install

compile:
	go build -o build/smartrebootd cmd/daemon/main.go
	go build -o build/smartreboot cmd/cli/main.go

run-daemon:
	go run cmd/daemon/main.go

run-cli:
	go run cmd/cli/main.go

clean:
	rm -rf build/

docs:
	godoc -http=:6060

install:
# Binaries
	install -Dsm755 build/smartrebootd -T $(DESTDIR)/usr/sbin/smartrebootd
	install -Dsm755 build/smartreboot -T $(DESTDIR)/usr/bin/smartreboot

# Default scripts
	install -Dm755 resources/debian-reboot-required -T $(DESTDIR)/etc/smartreboot/rebootchecks/00-debian-reboot-required
	mkdir -p $(DESTDIR)/etc/smartreboot/blockchecks

# Other files
	install -Dm644 resources/conf-template.conf -T $(DESTDIR)/etc/smartreboot/smartreboot.conf
	install -Dm644 resources/systemd.service -T $(DESTDIR)/usr/lib/systemd/system/smartrebootd.service

	@echo "Smart Reboot is now installed"
