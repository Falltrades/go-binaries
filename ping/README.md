# Usage and thanks

All thanks to https://github.com/go-ping/ping/blob/master/cmd/ping/ping.go !


To be able to run without root privileges, set privileged to true:
```
//privileged := flag.Bool("privileged", false, "")
privileged := flag.Bool("privileged", true, "")
```

and use setcap on the binary to allow it to bind to raw sockets:
```
sudo setcap cap_net_raw+ep /path/to/your/compiled/binary
```
