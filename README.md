Exec Util
===========
`executil` is a Go (golang) convenience utility package for [os/exec](https://golang.org/pkg/os/exec/) that will output stdout and stderr to the terminal.

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

Resulting output: `[echo] test echo`

### SetOutputChan
Optionally you can pass an output channel to send the command stdout and stderr output. If there is an output channel set CmdStart will not print the output but instead send it to the channel.
```
outputChan := make(chan string)
executil.SetOutputChan(outputChan)

err := executil.CmdStart("echo", "test echo")
if err != nil {
  log.Fatal(err)
}
```

### SetOutputPrefix
Optionally you can pass an output prefix string to be included in the stdout and stderr output. By default it will use the command you run.
```
executil.SetOutputPrefix("echoandthebunnymen")
err := executil.CmdStart("echo", "test echo")
if err != nil {
  log.Fatal(err)
}
```

Resulting output: `[echoandthebunnymen] test echo`
