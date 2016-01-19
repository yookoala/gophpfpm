package gophpfpm_test

import (
	"os"
	"path"
	"testing"

	"github.com/yookoala/gophpfpm"
)

var basepath string

func init() {
	var err error
	basepath, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	basepath = path.Join(basepath, "_test")
}

func TestNew(t *testing.T) {
	path := "/usr/sbin/php5-fpm"
	phpfpm := gophpfpm.NewProcess(path)
	if want, have := path, phpfpm.Exec; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}

func ExampleProcess() {

	process := gophpfpm.NewProcess("/usr/sbin/php5-fpm")

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
