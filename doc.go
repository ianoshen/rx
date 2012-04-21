/*
The rx command is a dependency and version management system for Go projects.
It is built on top of the go tool and utilizes the $GOPATH convention.

General Usage

The rx command is composed of numerous sub-commands.
Sub-commands can be abbreviated to any unique prefix on the command-line.
The general usage is:

  rx [rx args] [command] [command args]

Options:
  --rescan = false        Force a rescan of repositories
  --rxdir  = $HOME/.rx    Directory in which to save state

See below for a description of the various sub-commands understood by rx.

Help Command

Help on the rx command and subcommands.

Usage:
    rx help [command]

Options:
  --godoc = false    Dump the godoc output for the command(s)


List Command

List recognized repositories.

Usage:
    rx list 

Options:
  -f                List output format
  --long = false    Use long output format

The list command scans all available packages and collects information about
their repositories.  By default, each repository is listed along with its
dependencies and contained packages.

The -f option takes a template as a format.  The data passed into the
template invocation is an (rx/repo) RepoMap, and the default format is:

  {{range .}}{{.Path}}:{{range .Packages}} {{.Name}}{{end}}
  {{end}}

If you specify --long, the format will be:

  {{range .}}Repository ({{.VCS}}) {{printf "%q" .Path}}:
      Dependencies:{{range .RepoDeps}}
          {{.}}{{end}}
      Packages:{{range .Packages}}
          {{.ImportPath}}{{end}}
  
  {{end}}
Tags Command

List known repository tags.

Usage:
    rx tags repo

Options:
  --down = false    Only show downgrades
  -f                tags output format
  --long = false    Use long output format
  --up   = false    Only show updates (overrides --down)

The tags command scans the specified repository and lists
information about its tags.  The [repo] can be the suffix of the repository
root path, as long as it is unique.

The -f option takes a template as a format.  The data passed into the
template invocation is an (rx/repo) TagList, and the default format is:

  {{range .}}{{.Rev}} {{.Name}}
  {{end}}
Prescribe Command

Update the repository to the given tag/rev.

Usage:
    rx prescribe repo tag

Options:
  --build   = true    build all updated packages
  --install = true    install all updated packages
  --test    = true    test all updated packages

The prescribe command updates the repository to the named tag or
revision.  The [repo] can be the suffix of the repository root path,
as long as it is unique.  The [tag] is anything understood by the
underlying version control system as a commit, usually a tag, branch,
or commit.

After updating, prescribe will test, build, and the install each package
in the updated repository.  These steps can be disabled via flags such as
"rx prescribe --test=false repo tag".  If a step is disabled, the next
steps will be disabled as well.

*/
package main