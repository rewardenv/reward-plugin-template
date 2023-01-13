# Reward Plugin Template

## Structure

```
.
├── cmd/             # CLI commands
│   ├── greeter/     # The Application's entrypoint. This contains the main.go file.
│   ├── root/        # root command (the top level command)
│   ├── greet/       # greet command
│   └── helpers.go   # Helper functions for the CLI commands
├── go.mod
├── go.sum
└── internal/
    ├── config/      # The app configuration
    └── logic/       # The logic of various commands. 
                     # Keeping them in a separate package makes it easier to call one command's logic from another 
                     # and vice-versa.
```

## Running

```console
$ go run cmd/greeter/main.go greet Foo Bar

Hello Foo Bar!
```

## Running in Debug mode

```console
$ DEBUG=true go run cmd/greeter/main.go greet Foo Bar

DEBUG[2023-01-13T18:13:48+01:00] greet.go:12 reward-greeter/internal/logic.(*Client).RunCmdGreet() Creating a greeting...                       
Hello Foo Bar!
DEBUG[2023-01-13T18:13:48+01:00] greet.go:25 reward-greeter/internal/logic.(*Client).RunCmdGreet() ...greeting created.
```

## Running with Command Line Flags

```console
$ go run cmd/greeter/main.go greet Foo Bar --add-cakes

Hello Foo Bar, here are some cakes 🎂!
```

# Using with Reward

```console
$ go build -o reward-greeter cmd/greeter/main.go
$ mkdir -p ~/.reward/plugins.d
$ mv reward-greeter ~/.reward/plugins.d/

$ reward greeter greet Foo Bar --add-cakes

Hello Foo Bar, here are some cakes 🎂!
```
