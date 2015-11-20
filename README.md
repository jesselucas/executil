Exec Util
===========
`executil` is a Go (golang) package that wraps [os/exec](https://golang.org/pkg/os/exec/). It extends cmd.Start() to automatically output the cmd's stderr and stdout pipes to the terminal or to a supplied channel.

## Install
`go get -u github.com/jesselucas/executil`

## Usage
### StartAndWait
This will automatically Println the stdout and stderr pipes to the terminal and Wait for command to finish.
```
cmd := executil.Command("echo", "test echo")
err := cmd.StartAndWait()
if err != nil {
  log.Fatal(err)
}
```

Resulting output: `[echo] test echo`

### OutputChan
Optionally you can pass an output channel to send the command stdout and stderr output. If there is an output channel set CmdStart will not print the output but instead send it to the channel.
```
outputChan := make(chan string)

cmd := executil.Command("echo", "test echo")
cmd.OutputChan = outputChan

err := cmd.StartAndWait()

if err != nil {
  log.Fatal(err)
}
```

### OutputPrefix
Optionally you can pass an output prefix string to be included in the stdout and stderr output. By default it will use the command you run.
```
cmd := executil.Command("echo", "test echo")
cmd.OutputPrefix = "echoandthebunnymen"

err := cmd.StartAndWait()

if err != nil {
  log.Fatal(err)
}
```

Resulting output: `[echoandthebunnymen] test echo`

### ShowOutput
By default ShowOutput is true you can set it to false if you need to.

```
cmd := executil.Command("echo", "test echo")
cmd.ShowOutput = false

err := cmd.StartAndWait()

if err != nil {
  log.Fatal(err)
}
```
