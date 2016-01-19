# gophpfpm [![Travis](https://travis-ci.org/yookoala/gophpfpm.svg)][travis] [![GoDoc](https://godoc.org/github.com/yookoala/gophpfpm?status.svg)][godoc]

**gophpfpm** is a minimalistic php-fpm process manager written
in [go][golang].

It generates config file for a simple php-fpm process with 1 pool
and listen to 1 address only.

This is a fringe case, I know. Just hope it might be useful for
someone else.

[godoc]: https://godoc.org/github.com/yookoala/gophpfpm
[travis]: https://travis-ci.org/yookoala/gophpfpm
[golang]: https://golang.org

Usage
-----

```go
package main

import "github.com/yookoala/gophpfpm"

func main() {

  phpfpm := gophpfpm.New("/usr/sbin/php5-fpm")

  // config to save pidfile, log to "/home/foobar/var"
  // also have the socket file "/home/foobar/var/php-fpm.sock"
  phpfpm.SetPrefix("/home/foobar/var")

  // save the config file to basepath + "/etc/php-fpm.conf"
  phpfpm.SaveConfig(basepath + "/etc/php-fpm.conf")
  phpfpm.Start()

  go func() {

    // do something that needs phpfpm
    // ...
    phpfpm.Stop()

  }()

  // will wait for phpfpm to exit
  phpfpm.Wait()

}

```

License
-------

This software is license under [MIT License][mit-license]. You
may find [a copy of the license][license] in this repository.

[mit-license]: https://opensource.org/licenses/MIT
[license]: /LICENSE
