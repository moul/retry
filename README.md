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
   0.1.0

AUTHOR(S):
   Manfred Touron <https://github.com/moul/retry>

COMMANDS:
GLOBAL OPTIONS:
   --interval value, -n value      	seconds to wait between attempts (default: 1) [$RETRY_INTERVAL]
   --quiet, -q 			            don't print errors [$RETRY_QUIET]
   --clear, -c                      clear screen between each attempts [$RETRY_CLEAR]
   --timeout value, -t value        maximum seconds per attempt (disabled=0) (default: 0) [$RETRY_TIMEOUT]
   --help, -h  			            show help
   --version, -v       		        print the version
```

## Install

#### Using Homebrew

1. `brew tap moul/moul`
2. `brew update`
3. `brew install retry`

#### Using Golang

1. Install and configure Golang
2. Get the sources with `go get github.com/moul/retry/cmd/retry`
3. Compile and install the binary with `go install github.com/moul/retry/cmd/retry`
4. Profit

## License

MIT
