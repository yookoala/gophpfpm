package gophpfpm_test

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/yookoala/gophpfpm"
)

var basepath, pathToPhpFpm string

func init() {
	var err error
	basepath, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	basepath = path.Join(basepath, "_test")

	pathToPhpFpm = "/usr/sbin/php5-fpm"
}

func TestNew(t *testing.T) {
	path := pathToPhpFpm
	process := gophpfpm.NewProcess(path)
	if want, have := path, process.Exec; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}

func TestProcess_SetPrefix(t *testing.T) {
	path := pathToPhpFpm
	process := gophpfpm.NewProcess(path)
	process.SetDatadir(basepath + "/var")
	if want, have := basepath+"/var/phpfpm.pid", process.PidFile; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := basepath+"/var/phpfpm.error_log", process.ErrorLog; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := basepath+"/var/phpfpm.sock", process.Listen; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}

func TestProcess_Address(t *testing.T) {
	var network, address string
	process := &gophpfpm.Process{}

	process.Listen = "192.168.123.456:12345"
	network, address = process.Address()
	if want, have := "tcp", network; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}
	if want, have := "192.168.123.456:12345", address; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}

	process.Listen = "12345"
	network, address = process.Address()
	if want, have := "tcp", network; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}
	if want, have := ":12345", address; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}

	process.Listen = "hello.sock"
	network, address = process.Address()
	if want, have := "unix", network; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}
	if want, have := "hello.sock", address; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}

	process.Listen = "/path/to/hello.sock"
	network, address = process.Address()
	if want, have := "unix", network; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}
	if want, have := "/path/to/hello.sock", address; want != have {
		t.Errorf("expected %#v; got %#v", want, have)
	}

}

func TestProcess_StartStop(t *testing.T) {
	path := pathToPhpFpm
	process := gophpfpm.NewProcess(path)
	process.SetDatadir(basepath + "/var")
	process.SaveConfig(basepath + "/etc/phpfpm_test_startstop.conf")

	var err error
	var stdout, stderr io.ReadCloser

	if stdout, stderr, err = process.Start(); err != nil {
		t.Errorf("unexpected error: %#v", err.Error())
		if stdout != nil {
			stdout.Close()
		}
		if stderr != nil {
			stderr.Close()
		}
		return
	}
	defer stdout.Close()
	defer stderr.Close()

	go func() {
		// do something that needs phpfpm
		// ...
		if err := process.Stop(); err != nil {
			panic(err)
		}
	}()

	if _, err := process.Wait(); err != nil {
		t.Errorf("unexpected error: %#v", err.Error())
	}
}

func ExampleProcess() {

	process := gophpfpm.NewProcess(pathToPhpFpm)

	// SetDatadir equals to running these 3 settings:
	// process.PidFile  = basepath + "/phpfpm.pid"
	// process.ErrorLog = basepath + "/phpfpm.error_log"
	// process.Listen   = basepath + "/phpfpm.sock"
	process.SetDatadir(basepath + "/var")

	// save the config file to basepath + "/etc/php-fpm.conf"
	process.SaveConfig(basepath + "/etc/php-fpm.conf")
	process.Start()

	go func() {
		// do something that needs phpfpm
		// ...
		process.Stop()
	}()

	process.Wait()

	// Output:
}
