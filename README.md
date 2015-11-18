Exec Util
===========
`executil` is a go (golang) convenience utility package for [os/exec](https://golang.org/pkg/os/exec/).

## Install
`go get -u github.com/jesselucas/executil`

## Usage
### CmdStart
Creates a new exec.Cmd and calls the cmd Start() function. This will automatically Println the stdout and stderr.
```
err := executil.CmdStart("echo", "test echo")
if err != nil {
  log.Fatal(err)
}
```

### SetOutputChan
Optionally you can pass an output channel to send the command stdout and stderr output. If there is an output channel set CmdStart will not print the output but instead send it to the channel.
```
outputChan := make(chan string)
SetOutputChan(outputChan)

err := executil.CmdStart("echo", "test echo")
if err != nil {
  log.Fatal(err)
}
```
