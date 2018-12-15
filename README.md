Resume
======

Create an online resume that can be accessed by a `curl` command.

Example:

![Screenshot](/screenshot.png?raw=true)


Usage
-----

To run the application we'll be using Docker, you can also run this application
on the commandline and I provided a binary for linux amd64 based systems in the
`bin/` folder.

```bash
$ git clone git@github.com:erroneousboat/resume.git
$ cd resume
$ docker build -t resume . && docker run -p 80:80 -it resume
$ curl http://localhost/resume
```

Configuration
-------------

You'll be able to configure some settings, mainly:

| Var    | Explanation                                          |
|--------|------------------------------------------------------|
| `PORT`   | set the port on which the application will be served |
| `USER`   | set username for basic authentication                |
| `PASS`   | set password for basic authentication                |

In conjunction with docker:

```bash
$ docker run -p 8080:8080 --env PORT=8080 --env USER=show --env PASS=me -it resume
$ curl http://localhost:8080/resume -u show:me
```

Template
--------

The [`resume.tmpl`](./resume.tmpl) has an example of how you could structure your resume. To
see what colors you'll be able to use, inspect the [`colors.go`](./colors.go) file.

For more color styles, you can use the function
``{{.ColorCode "your-style-here"}}`` in the template. See https://github.com/mgutz/ansi#style-format
for the style format.
