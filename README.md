

## Fetching deps

`gb vendor fetch github.com/heroku/slog`

## Compilation

`gb build`

Creates:

* `pkg/` where package .a files are created.
* `bin/` where main packages are compiled to. In this case we get `bin/gb-sample`, since that is our only main.

Both of these directories are ignored in .gitignore
