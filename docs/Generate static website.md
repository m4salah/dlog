dlog CLI allow for generating static website from source directory. this is how this website is generated.

To generate a static website using dlog use the `--build` flag with a path as destination for example:

```shell
dlog --build /path/to/output
```

dlog will build all markdown files to HTML and extract all static files from inside the binary executable file to that destination directory. Then it will terminate.

Building process creates a dlog server instance and request all pages and save it to desk. That allow dlog extensions to define a new handler that renders a page. the page will work in both usecases: local server, static site generation. extensions has to also register the path for build using [`RegisteBuildPage`](https://pkg.go.dev/github.com/m4salah/dlog#RegisterBuildPage) function

While building static site dlog turns on **READONLY** mode. so specifying `--build` flag is equal to `--build --readonly`.

dlog builds `/docs` directory every commit to update this website. it uses Github workflow to do that. https://github.com/m4salah/dlog/blob/master/.github/workflows/dlog.yml