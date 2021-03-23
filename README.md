# portscanner

A portscanner written in Go based on https://sj14.gitlab.io/post/2015/03-01-port-scanner/

### Examples

```bash
$ ./portscan -host=google.com
open 80
scan finished
```

```bash
$ ./portscan -start=80 -end=600 -timeout=250ms
open 80
open 443
```

### Arguments

```bash
$ ./portscan -h
        -closed
                list closed ports (true/false)
        -end int
                the upper end to scan (default -1)
        -host string
                the host to scan (default "localhost")
        -pause string
                pause after each scanned port (e.g. 5ms) (default "1ms")
        -start int
                the lower end to scan (default 80)
        -timeout string
                timeout (e.g. 50ms or 1s) (default "1000ms")
```
