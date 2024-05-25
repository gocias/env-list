# EnVar List
Print environment variables: all or by conditions.

## Install
Add a last module version
```shell
go get github.com/gocias/env-list
```
or a specific version
```shell
go get github.com/gocias/env-list@v1.0.1
```

## Usage
```go
package main

import (
	"fmt"
	"log"
	"os"

	el "github.com/gocias/env-list"
	"github.com/joho/godotenv"
)

func main() {
	err := os.Setenv("CONF_FIRST", "conf_first")
	if err != nil {
		return
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	er := os.Setenv("CONF_LAST", "conf_last")
	if er != nil {
		return
	}
	var ose el.EnvList = os.Environ()
	fmt.Println("=== PrintItems ============================================")
	ose.PrintItems()
	fmt.Println("=== get VarPrefix 'CONF' ==================================")
	ose.VarPrefix("CONF").PrintItems()
	fmt.Println("=== get VarContains 'ICE' =================================")
	ose.VarContains("ICE").PrintItems()
	fmt.Println("=== get VarKeyContains 'SESS' =============================")
	ose.VarKeyContains("SESS").PrintItems()
	fmt.Println("=== get VarValueContains '=' ==============================")
	ose.VarValueContains("=").PrintItems()
}
```

## Result
```
=== PrintItems ============================================
LOGNAME=someuser
GOPATH=/home/someuser/go
... a lot of lines ...
USERNAME=someuser
HOME=/home/someuser
USER=someuser
... a lot of lines ...
CONF_FIRST=conf_first
CONF_PROTOCOL=http
CONF_HOST=example
CONF_DOMAIN=com
CONF_LAST=conf_last

=== get VarPrefix 'CONF' ==================================
CONF_FIRST=conf_first
CONF_PROTOCOL=http
CONF_HOST=example
CONF_DOMAIN=com
CONF_LAST=conf_last

=== get VarContains 'ICE' =================================
SESSION_MANAGER=local/somehost:@/tmp/.ICE-unix/2998,unix/somehost:/tmp/.ICE-unix/2998

=== get VarKeyContains 'SESS' =============================
GNOME_SHELL_SESSION_MODE=ubuntu
XDG_SESSION_DESKTOP=ubuntu
DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus
SESSION_MANAGER=local/somehost:@/tmp/.ICE-unix/2998,unix/somehost:/tmp/.ICE-unix/2998
XDG_SESSION_CLASS=user
GDMSESSION=ubuntu
GNOME_DESKTOP_SESSION_ID=this-is-deprecated
DESKTOP_SESSION=ubuntu
XDG_SESSION_TYPE=wayland

=== get VarValueContains '=' ==============================
XMODIFIERS=@im=ibus
DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus
MEMORY_PRESSURE_WRITE=c29tZSAySomeThingDAwMAA=
```