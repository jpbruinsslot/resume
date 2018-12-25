Resume
======

Create an online resume that can be accessed by a `curl` command.

Example:

![Screenshot](/screenshot.png?raw=true)


Usage
-----

#### Docker

```bash
$ docker run -p 80:80 erroneousboat/resume
$ curl http://localhost/resume
```

#### Command line

You can also run this application from the commandline, and I provided a
binary for linux amd64 based systems [here](https://github.com/erroneosboat/resume).

Configuration
-------------

You'll be able to configure some settings, mainly:

| Var    | Explanation                                          |
|--------|------------------------------------------------------|
| `PORT`   | set the port on which the application will be served |
| `USER`   | set username for basic authentication                |
| `PASS`   | set password for basic authentication                |

A full example:

```bash
$ docker run -p 8080:8080 \
    --env PORT=8080 \
    --env USER=show \
    --env PASS=me \
    -v path-to-resume.tmpl:/resume.tmpl \
    erroneousboat/resume

$ curl http://localhost:8080/resume -u show:me
```

Template
--------

The [`resume.tmpl`](./resume.tmpl) has an example of how you could structure your resume. To
see what colors you'll be able to use, inspect the [`colors.go`](./colors.go) file.

For more color styles, you can use the function
``{{.ColorCode "your-style-here"}}`` in the template. See https://github.com/mgutz/ansi#style-format
for the style format.
