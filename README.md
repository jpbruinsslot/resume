# 📄 Resume

Create an online resume that can be accessed by a `curl` command.

## 📸 Screenshot

![Screenshot](/screenshot.png?raw=true)

## 📦 Installation

### Binary installation

[Download](https://github.com/jpbruinsslot/resume/releases) a
compatible binary for your system. For convenience, place `weather` in a
directory where you can access it from the command line.

### Via Go

```text
$ go install github.com/jpbruinsslot/resume/cmd/resume@latest
```

### Via Docker

```
$ docker build  -t resume .
```

## 💻 Usage

### Command line

```
$ resume \
    -port 8080 \
    -path /resume \
    -user show \
    -pass me
$ curl http://localhost:8080/resume -u show:me
```

### Docker

```
$ docker run -p 8080:8080 \
    --env PORT=8080 \
    --env PATH=/resume \
    --env USER=show \
    --env PASS=me \
    -v path-to-resume.tmpl:/resume.tmpl \
    resume
$ curl http://localhost:8080/resume -u show:me
```

## 🔧 Configuration

## Arguments

You'll be able to pass some arguments to the application

| Argument   | Description                                    |
| ---------- | ---------------------------------------------- |
| `port`     | port to serve on (default: 80)                 |
| `path`     | optional url path (default: /)                 |
| `template` | filename of template (default: resume.tmpl)    |
| `user`     | username used for authentication (default: "") |
| `pass`     | password used for authentication (default: "") |

## Templating

The [`resume.tmpl`](./resume.tmpl) has an example of how you could structure your resume. To
see what colors you'll be able to use, inspect the [`colors.go`](./colors.go) file.

For more color styles, you can use the function
`{{.ColorCode "your-style-here"}}` in the template. See https://github.com/mgutz/ansi#style-format
for the style format.
