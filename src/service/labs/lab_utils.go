package labs

import "fmt"

// LimitOption is a function type that modifies AutodriverLimits
type LimitOption func(*AutodriverLimits)

// AutodriverLimits limits to construct the autodriver cli args
// more info - https://github.com/UB-CSE-IT/Autolab-Public-Documentation/blob/d6c2fb902d8cc0534794f20d6ef0bab871fe2602/The%20autograding%20process.md#running-the-job
// todo verify assumptions
type AutodriverLimits struct {
	// limits the number of processes that can be started
	ULimit int
	// sets the maximum file size that can be created in bytes I assume
	MaxFileSize int
	// limits the size of the output in bytes I assume
	MaxOutputSize int
	// sets the job timeout in seconds I assume
	Timeout int
}

var (
	// DefaultAutoDriverLimits limits taken from
	// https://github.com/UB-CSE-IT/Autolab-Public-Documentation/blob/d6c2fb902d8cc0534794f20d6ef0bab871fe2602/The%20autograding%20process.md#running-the-job
	DefaultAutoDriverLimits = AutodriverLimits{
		ULimit:        100,
		MaxFileSize:   104857600,
		Timeout:       20,
		MaxOutputSize: 1024000,
	}
)

// CreateLeviathanEntryCommand command for normal graders made for leviathan
func CreateLeviathanEntryCommand(cmd string) string {
	return "cd autolab;" + cmd
}

// CreateTangoEntryCommand builds the autodriver command for tango compatibility
// defaults to DefaultAutoDriverLimits, if no limits is passed
func CreateTangoEntryCommand(opts ...LimitOption) string {
	limits := &DefaultAutoDriverLimits // Start with default limits
	for _, opt := range opts {         // Apply all options
		opt(limits)
	}

	return fmt.Sprintf(""+
		"chown autolab:autolab autolab; chmod 755 autolab;"+ // mode perm since CopyToContainer switches to root and messes up perms
		"su autolab -c \"autodriver -u %d -f %d -t %d -o %d autolab > output/feedback 2>&1\";"+
		"cat output/feedback;", // print to stdout so that leviathan can parse
		limits.ULimit, limits.MaxFileSize, limits.Timeout, limits.MaxOutputSize,
	)
}

// WithTimeout changes the DefaultAutoDriverLimits timeout
func WithTimeout(timeout int) LimitOption {
	return func(limits *AutodriverLimits) {
		limits.Timeout = timeout
	}
}
