# go-binaries

Some of the most infamous network tools code written in golang so you can build them into executables binaries.

Imagine you need to troubleshoot network issues and the command are somehow missing on the system. What if you are not allowed to install packages because of lack of permission ? Or maybe you are just lazy to properly configure proxies and mirrors. Think also when you need to troubleshoot inside a container...  
Fear not, because you may be able to workaround by using those binaries.

How come those binaries could work ? As long as glibc libraries are compatible. This means you may need to build it on a similar system to your target.

Official golang images can be found there : https://hub.docker.com/_/golang

Use script/build.sh to build your binary. Example:
```
cd script
./build.sh golang:1.23rc1-alpine3.20 curl
./build.sh golang:1.23rc1-bullseye netcat
```
