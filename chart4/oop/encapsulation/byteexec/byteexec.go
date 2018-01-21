package byteexec


import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/elitecodegroovy/go-core/util/logger"
	"github.com/elitecodegroovy/go-core/util/filepersist"
)

var (
	logTime = true
	log = logger.NewStdLogger(logTime, true , false, false, true)

	fileMode = os.FileMode(0744)

	initMutex sync.Mutex
)

// Exec is a handle to an executable that can be used to create an exec.Cmd
// using the Command method. Exec is safe for concurrent use.
type Exec struct {
	Filename string
}

// New creates a new Exec using the program stored in the provided data, at the
// provided filename (relative or absolute path allowed). If the path given is
// a relative path, the executable will be placed in one of the following
// locations:
//
// On Windows - %APPDATA%/byteexec
// On OSX - ~/Library/Application Support/byteexec
// All Others - ~/.byteexec
//
// Creating a new Exec can be somewhat expensive, so it's best to create only
// one Exec per executable and reuse that.
//
// WARNING - if a file already exists at this location and its contents differ
// from data, Exec will attempt to overwrite it.
func New(data []byte, filename string) (*Exec, error) {
	log.Tracef("Creating new at %v", filename)
	// Use initMutex to synchronize file operations by this process
	initMutex.Lock()
	defer initMutex.Unlock()

	var err error
	if !filepath.IsAbs(filename) {
		filename, err = inStandardDir(filename)
		if err != nil {
			return nil, err
		}
	}
	filename = renameExecutable(filename)
	log.Tracef("Placing executable in %s", filename)

	err = filepersist.Save(filename, data, fileMode)
	if err != nil {
		return nil, err
	}
	log.Debugf("File saved, returning new Exec")
	return newExec(filename)
}

// Command creates an exec.Cmd using the supplied args.
func (be *Exec) Command(args ...string) *exec.Cmd {
	return exec.Command(be.Filename, args...)
}

func newExec(filename string) (*Exec, error) {
	absolutePath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	return &Exec{Filename: absolutePath}, nil
}

func inStandardDir(filename string) (string, error) {
	folder, err := pathForRelativeFiles()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(folder, fileMode)
	if err != nil {
		return "", fmt.Errorf("Unable to make folder %s: %s", folder, err)
	}
	return filepath.Join(folder, filename), nil
}

func inHomeDir(filename string) (string, error) {
	log.Tracef("Determining user's home directory")
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Unable to determine user's home directory: %s", err)
	}
	return filepath.Join(usr.HomeDir, filename), nil
}

