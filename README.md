# retry
:shell: retry shell commands

## Examples

```console
$ retry sh -c 'exit "$(($RANDOM % 3))"'
2016/08/27 19:41:11 run 1: command finished with error: exit status 1
2016/08/27 19:41:12 run 2: command finished with error: exit status 2
2016/08/27 19:41:13 run 3: command finished with error: exit status 2
2016/08/27 19:41:14 run 4: command finished with error: exit status 2
2016/08/27 19:41:15 run 5: command finished with error: exit status 2
2016/08/27 19:41:16 run 6: command finished with error: exit status 1
2016/08/27 19:41:17 Command succeeded on attempt 7 with a total duration of 6 seconds
$
```

```console
$ retry -q sh -c 'exit "$(($RANDOM % 3))"'
$
```

```console
$ retry false
2016/08/27 19:38:25 run 1: command finished with error: exit status 1
2016/08/27 19:38:26 run 2: command finished with error: exit status 1
2016/08/27 19:38:27 run 3: command finished with error: exit status 1
2016/08/27 19:38:28 run 4: command finished with error: exit status 1
2016/08/27 19:38:29 run 5: command finished with error: exit status 1
^C
$
```

```console
$ retry true
2016/08/27 19:38:10 Command succeeded on attempt 1 with a total duration of 0 second
$
```

## Usage

```console
$ retry -h
NAME:
   retry - retry

USAGE:
   retry [global options] command [command options] [arguments...]

VERSION:
   0.4.0

AUTHOR(S):
   Manfred Touron <https://moul.io/retry>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --interval value, -n value      seconds to wait between attempts (default: 1) [$RETRY_INTERVAL]
   --quiet, -q                     don't print errors [$RETRY_QUIET]
   --clear, -c                     clear screen between each attempts [$RETRY_CLEAR]
   --timeout value, -t value       maximum seconds per attempt (disabled=0) (default: 0) [$RETRY_TIMEOUT]
   --max-attempts value, -m value  quit after NUM attempts (default: 0) [$RETRY_MAX_ATTEMPTS]
   --reverse-behavior, -r          inverse behavior, stop on first fail [$RETRY_REVERSE_BEHAVIOR]
   --help, -h                      show help
   --version, -v                   print the version
```

## Install

#### With Golang

1. Install and configure Golang
2. `go get moul.io/retry`

#### With Homebrew

    brew install moul/moul/retry

## License

MIT
