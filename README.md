# captive portal checker

`captive` is a tiny cmdline tool to check if the system it is running on has working Internet access. Useful to detect captive portals on public WiFis and corporate networks. Checks http://captive.apple.com (which powers the Captive Portal Assistant on Apple devices).

Intended to be used in shell scripts: Single binary/no dependencies, cross-platform and returns the online status as exit code.

## install
download executable from [release page](https://github.com/christian-korneck/captive/releases) and copy it to a PATH location (i.e. `/usr/local/bin` or `c:\windows\system32`)

## usage

Run `captive` (there are no parameters). It will return with exit code `0` when online or exit code `1` when offline.

Example:

```
$ captive
online
$ echo $?    #check exit code
0
$ ifconfig eth0 down    #break Internet connection
$ captive
offline
$ echo $?    #check exit code
1
```

## build

with Go 1.13+ installed:

```
go get github.com/christian-korneck/captive/...
```

## run with Docker

```
docker run --rm chko/captive
```
