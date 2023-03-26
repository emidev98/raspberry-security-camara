# Raspberry PI Security Camara

<p align="center">
    <img href="./icon.png">
</p>

Monorepo to run a security camara using the RaspberryPi. It contains the [api](./api/) written in GoLang and the [frontend](./frontend/) that uses the driver [libcamera-vid](https://www.raspberrypi.com/documentation/computers/camera_software.html) to record.

## Development of API

Move into the api folder and execute `go install` to install the dependencies and build the project into the go binaries path. If you want to have a build of the project directly in the api folder you can execute `go build main.go` which will generate a binary for your platform directly in the api folder.

> Compile the project directly into your Rasbperry Pi, that way you avoid possible architecture or OS incompatibilities.