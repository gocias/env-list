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
	fmt.Println("=== Print =================================================")
	ose.Print()
	fmt.Println("=== get VarWithPrefix 'CONF' ==============================")
	ose.VarWithPrefix("CONF").Print()
	fmt.Println("=== get VarContains 'ICE' =================================")
	ose.VarContains("ICE").Print()
	fmt.Println("=== get VarKeyContains 'SESS' =============================")
	ose.VarKeyContains("SESS").Print()
	fmt.Println("=== get VarValueContains '=' ==============================")
	ose.VarValueContains("=").Print()
	fmt.Println("=== sort ASC for VarWithPrefix 'CONF' =====================")
	ose.VarWithPrefix("CONF").SortAsc().Print()
	fmt.Println("=== sort DESC for VarWithPrefix 'CONF' ====================")
	ose.VarWithPrefix("CONF").SortDesc().Print()
}
```

## Result
```
=== Print =================================================
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

=== get VarWithPrefix 'CONF' ==============================
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

=== sort ASC for VarWithPrefix 'CONF' =====================
CONF_DOMAIN=com
CONF_FIRST=conf_first
CONF_HOST=example
CONF_LAST=conf_last
CONF_PROTOCOL=http

=== sort DESC for VarWithPrefix 'CONF' ====================
CONF_PROTOCOL=http
CONF_LAST=conf_last
CONF_HOST=example
CONF_FIRST=conf_first
CONF_DOMAIN=com
```