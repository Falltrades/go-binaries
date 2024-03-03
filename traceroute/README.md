# Usage and thanks

All thanks to https://github.com/pixelbender/go-traceroute/tree/master ! 


Use setcap on the binary to allow it to bind to raw sockets:
```
sudo setcap cap_net_raw+ep /path/to/your/compiled/binary
```
