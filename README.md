# Smart Reboot

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/BrenekH/smartreboot)
[![GitHub License](https://img.shields.io/github/license/BrenekH/smartreboot)](https://github.com/BrenekH/smartreboot/blob/master/LICENSE)

## What?

Smart Reboot is a daemon(`smartrebootd`) and CLI(`smartreboot`) that allow for a Linux machine to be rebooted without interrupting any work that may be happening at the current time.
This is accomplished by using user-defined scripts to determine when a reboot is required and if it can accomplished at the current time.

## Why?

Smart Reboot was designed as a way to confidently let a server handle its own rebooting while not compromising the work being done.

**For example:** A Jellyfin server shouldn't restart itself while a stream is currently playing, but updates to the kernel and other services still require that the machine be rebooted once in a while.
Instead of requiring the manual intervention of the sysadmin, Smart Reboot can be configured to reboot when the system needs to, but only when the Jellyfin instance is idle.

## Installing

### Debian (and derivatives)

Download and install the `.deb` file for your architecture from the [latest release](https://github.com/BrenekH/smartreboot/releases/latest).

### Arch Linux (and derivatives)

SmartReboot is available on the [Arch User Repository](https://aur.archlinux.org) as the [smartreboot package](https://aur.archlinux.org/packages/smartreboot).

### Manual installation

You will need a [Go compiler](https://go.dev)(>=1.17) and the Make command(preferably [GNU Make](https://www.gnu.org/software/make/)) to begin.

First, clone `BrenekH/smartreboot` like so: `git clone https://github.com/BrenekH/smartreboot.git`.

Then, change into the directory containing the source files: `cd smartreboot`.

Next, build using Make: `make`.

Lastly, install Smart Reboot as follows: `sudo make install`.

If everything worked, you should now be able to use Smart Reboot!

## Usage

### Basics

In order to determine whether or not to reboot, Smart Reboot runs the scripts located in `/etc/smartreboot/rebootchecks` and `/etc/smartreboot/blockchecks` and evaluates the exit codes.
If any of the reboot checks return a 0 exit code, the machine is considered to need a reboot.
However, the reboot can be blocked if any one of the block checks return an exit code that is not 0, which indicates that the system is still working on something and cannot be rebooted at this time.

Each of the scripts must be a valid "executable" which includes, but isn't limited to, the following:

- The script must have the executable bit set(using `chmod +x`), ideally for all users as to not cause issues when switching from the daemon to the cli.

- If the script is a text file (Bash, Python, Perl, etc.), it must have a proper [shebang](<https://en.wikipedia.org/wiki/Shebang_(Unix)>) at the very beginning of the file.

- If the script is a compiled binary, it must be an ELF file or equivalent runnable file for your system.

### Automatic mode (via the daemon)

Smart Reboot provides a daemon called `smartrebootd` who's startup and shutdown should be managed by your distro's init system(most likely SystemD).
Since the SystemD service should already be installed, you can run `sudo systemctl enable --now smartrebootd` to start the Smart Reboot Daemon.
If you use a different init system, we welcome any contributions to add support.

By default, `smartrebootd` only checks for reboots every 15 minutes.
This value can be changed by modifying `/etc/smartreboot/smartreboot.conf` and restarting the daemon(`sudo systemctl restart smartrebootd`).

### Manual mode (via the CLI)

Smart Reboot also provides a cli which can be invoked using the `smartreboot` command.
`smartreboot` can be used to manually run through the same process that the daemon does using the `reboot` command like so `smartreboot reboot`.

If you want to skip the reboot checks, you can use the `--force, -f` flag which will only check the block scripts before issuing a reboot. **Example:** `smartreboot reboot --force`.

### Other CLI tricks

In addition to the `reboot` command, `smartreboot` can show the current status of the executable scripts that you have setup, which is very useful for setting up the script directories.

Simply run `smartreboot check` to view each script and its exit code.

## License

Smart Reboot is licensed under the GNU General Public License Version 3, a copy of which can be found in [LICENSE](https://github.com/BrenekH/smartreboot/blob/master/LICENSE).

## Future plans

- Allow community scripts to be uploaded so that others don't need to rewrite common scripts.
  - Add some sort of download command to `smartreboot` which can automatically download and enable these community scripts.
