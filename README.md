**Note**: There's now a Wasm based version of this thing available in [the photocopy-wasm repository](https://github.com/fuglede/photocopy-wasm/) -- in particular, you can try it out on https://photocopy.fuglede.dk/

# Photocopy

Have you ever had to send in a signature, been annoyed about the process of printing a piece of paper to write your signature on it, only to immediately throw away the paper after scanning it, when you could just as well have written your signature in paint and sent the result, only it looks just a little bit too clean and forged that way? Then you've come to the right place!

Based on [this wonderfully silly StackExchange post](https://tex.stackexchange.com/a/94541/88992), this web app allows users to upload images and get back as a result the simulated scan of their input.

![Example](tmpl/Example.png "Example output")

# Usage

Depending on your setup, the server can be run in a few different ways. In each case, complete the instructions below to have the application available at `http://localhost:8000`.

## From source

Ensure that you have a reasonably recent version of Go and ImageMagick installed; on Debian-like systems

```
$ sudo apt install golang imagemagick
```

should suffice. Then, to launch the service on port 8000, clone this repository, and run

```
$ go run main.go
```

## From Docker Hub

If you have Docker available, the application is available from Docker Hub through

```
$ docker run -d -p 8000:8000 fuglede/photocopy
```
