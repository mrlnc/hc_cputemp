# hkcputemp

This is a sample project of a HomeKit temperature sensor using [HomeControl](https://github.com/brutella/hc).

# Installation

## Go

Make sure that you have a working [Go installation](http://golang.org/doc/install).

## Checkout

- Clone project `git clone https://github.com/mrlnc/hkcputemp && cd hkcputemp`
- Install dependencies `go get`
- install `lm-sensors`

## Build

Build with `go build hkcputempd.go`.

## Run

Execute the executable `./hkcputemp`. Parameters:
```
USAGE:   ./hkcputemp -pin <accessory PIN> -chip <chip name> -sensor <sensor>
```

You can find valid chip and sensor names by running `sensors`:
```
$ sensors
...
thinkpad-isa-0000
Adapter: ISA adapter
fan1:        3076 RPM
temp1:        +41.0°C  
...
```
Here, chip would be `thinkpad-isa-0000` and sensor is `temp1`

# License

hkcputemp is based on hklight from Matthias Hochgatterer, which is available under a non-commercial license. See the LICENSE file for more info.