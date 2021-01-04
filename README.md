# go mod required as per https://github.com/labstack/echo/issues/1374
# had to install echo this way

Figured it out. As initially mentioned by @ShrewdSpirit, your code also needs to be in a module.
As a relative go newbie I was following https://golang.org/doc/code.html which talks about workspaces but now it seems there is an entirely different way to organize your code via modules.

You can read about it here:
https://blog.golang.org/using-go-modules

Note:

    "As of Go 1.11, the go command enables the use of modules when the current directory or any parent directory has a go.mod, provided the directory is outside $GOPATH/src. (Inside $GOPATH/src, for compatibility, the go command still runs in the old GOPATH mode, even if a go.mod is found. See the go command documentation for details.) Starting in Go 1.13, module mode will be the default for all development."

So, you can still follow the echo README example code:

    But make sure your project folder is OUTSIDE of the $GOPATH/src directory structure
    Make sure your GOPATH environment variable is set
    still use "github.com/labstack/echo/v4" in your import
    in your project directory, run go mod init example.com/main (example.com can be whatever namespace you want and main is your module name which can be anything as well)
    this will simply create a go.mod file which is just a very simple text file.
    run go build this will download echo and its dependencies to your workspace folder. Also it will automatically update your go.mod file to include the echo lab dependency


# URLS
# /articles
# /article/12

With go 1.15 and later, just clone and run.

Set PORT to 6660 macos export PORT="6660"
Set DATABASE_URL macos export DATABASE_URL="postgressurl"
go run server.go 

