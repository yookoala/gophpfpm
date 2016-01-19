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
	phpfpm := gophpfpm.New(path)
	if want, have := path, phpfpm.Exec; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}

func ExampleProcess() {

	process := gophpfpm.New("/usr/sbin/php5-fpm")

	// config to save pidfile, log to basepath + "/var"
	// also have the socket file basepath + "/var/php-fpm.sock"
	process.SetPrefix(basepath + "/var")

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
