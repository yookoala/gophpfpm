package gophpfpm_test

import (
	"os"
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
}

func TestNew(t *testing.T) {
	path := "/usr/sbin/php5-fpm"
	phpfpm := gophpfpm.New(path)
	if want, have := path, phpfpm.Exec; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}

func ExampleProcess() {

	phpfpm := gophpfpm.New("/usr/sbin/php5-fpm")

	// config to save pidfile, log to basepath + "/var"
	// also have the socket file basepath + "/var/php-fpm.sock"
	phpfpm.SetPrefix(basepath + "/var")

	// save the config file to basepath + "/etc/php-fpm.conf"
	phpfpm.SaveConfig(basepath + "/etc/php-fpm.conf")
	phpfpm.Start()

	go func() {
		// do something that needs phpfpm
		// ...
		phpfpm.Stop()
	}()

	phpfpm.Wait()

	// Output:
}
