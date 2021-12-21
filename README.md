# Smart Reboot

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/BrenekH/smartreboot)
[![GitHub](https://img.shields.io/github/license/BrenekH/smartreboot)](https://github.com/BrenekH/smartreboot/blob/master/LICENSE)

## What?

Smart Reboot is a daemon(`smartrebootd`) and cli(`smartreboot`) that allow for a unix machine to be rebooted without interrupting any work that may be happening at the current time.
This is accomplished by using user-defined scripts in `/etc/smartreboot/rebootchecks` and `/etc/smartreboot/blockchecks` to determine when a reboot is required and if it can accomplished at the current time.

## Why?

Smart Reboot was designed as a way to confidently let a server handle its own rebooting while not compromising the work being done.

**For example:** A home media aficionado doesn't want their Jellyfin server to restart itself while a stream is currently playing, but updates to the kernel still require that the machine be rebooted once in a while.
Instead of requiring the manual intervention of the sysadmin, Smart Reboot can be configured to reboot when the system needs to, but only when the Jellyfin instance is idle.

## Installing

### Debian (and derivatives)

At this time, Smart Reboot must be installed on Debian using the [manual instructions](#manual-installation).

We are working on packaging a `.deb` for distribution.

### Arch Linux (and derivatives)

At this time, Smart Reboot must be installed on Arch Linux using the [manual instructions](#manual-installation).

We are working on submitting a `PKGBUILD` to the AUR for distribution.

### Manual installation

You will need a [Go compiler](https://go.dev)(>=1.17) and the Make command(preferably [GNU Make](https://www.gnu.org/software/make/)) to begin.

First, clone `BrenekH/smartreboot` like so: `git clone https://github.com/BrenekH/smartreboot.git`.

Then, change into the directory containing the source files: `cd smartreboot`.

Next, build using Make: `make`.

Lastly, install Smart Reboot as follows: `sudo make install`.

If everything worked, you should now be able to use Smart Reboot!

## Usage
