package gophpfpm

import "github.com/go-ini/ini"

// Process describes a minimalistic php-fpm config
// that runs only 1 pool
type Process struct {

	// path to php-fpm executable
	Exec string

	// path to the config file
	ConfigFile string

	// The address on which to accept FastCGI requests.
	// Valid syntaxes are: 'ip.add.re.ss:port', 'port',
	// '/path/to/unix/socket'. This option is mandatory for each pool.
	Listen string

	// path of the PID file
	PidFile string

	// path of the error log
	ErrorLog string
}

// New creates a new process descriptor
func New(phpFpm string) *Process {
	return &Process{
		Exec: phpFpm,
	}
}

// SaveConfig generates config file according to the
// process attributes
func (proc *Process) SaveConfig(path string) {
	proc.ConfigFile = path
	proc.Config().SaveTo(proc.ConfigFile)
}

// Config generates an minimalistic config ini file
// in *ini.File format. You may then use SaveTo(path)
// to save it
func (proc *Process) Config() (f *ini.File) {
	f = ini.Empty()
	f.NewSection("global")
	f.Section("global").NewKey("pid", proc.PidFile)
	f.Section("global").NewKey("error_log", proc.ErrorLog)
	f.NewSection("www")
	f.Section("www").NewKey("listen", proc.Listen)
	f.Section("www").NewKey("pm", "dynamic")
	f.Section("www").NewKey("pm.max_children", "5")
	f.Section("www").NewKey("pm.start_servers", "2")
	f.Section("www").NewKey("pm.min_spare_servers", "1")
	f.Section("www").NewKey("pm.max_spare_servers", "3")
	return
}

// SetPrefix sets default config values according
// with reference to the folder prefix
func (proc *Process) SetPrefix(prefix string) {
}

// Start starts the php-fpm process
// in foreground mode instead of daemonize
func (proc *Process) Start() {
}

// Stop stops the php-fpm process with SIGINT
// instead of killing
func (proc *Process) Stop() {
}

// Wait wait for the process to finish
func (proc *Process) Wait() {
}
