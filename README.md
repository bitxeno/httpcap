# httpcap
A simple network analyzer that captures http network traffic.

* support Windows/MacOS/Linux/OpenWrt(x64)
* https only capture `clienthello`
* colorful output

![screenshot](https://cdn.jsdelivr.net/gh/kozalak-robot/assets@main/img/1634055826518iShot2021-10-13%2000.23.26.jpg)

## Usage

```
NAME:
   A simple network analyzer that capture http network traffic. run - run capture

USAGE:
   A simple network analyzer that capture http network traffic. run [command options] [arguments...]

OPTIONS:
   --interface value, -i value  interface to listen on (e.g. eth0, en1, or 192.168.1.1, 127.0.0.1 etc.) (default: "any")
   --port value, -p value       port to listen on (default listen on all port) (default: 0)
   --ip value                   capture traffic to and from ip
   --keyword value, -k value    filte output match the keyword
   --body                       print body content (only support text content body) (default: false)
   --raw                        print raw request / response (default: false)
   --debug, --vv                print debug message (default: false)
   --verbose, --vvv             print more debug message (default: false)
   --help, -h                   show help (default: false)

```

## How to use in openwrt

Openwrt maybe lack of `glibc` or `pcap`，Please try `static` or `nocgo` build


## Example

```
httpcap run                        # capture all interface (not support Windows)
httpcap run -i eth0 -p 80          # capture [eth0] interface with port 80
httpcap run -i 192.168.1.1 -p 80   # capture [192.168.1.1] interface with port 80
```
