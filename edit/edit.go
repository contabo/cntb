package edit

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

type Editor struct {
	editor string
}

func NewEditor(editorEnvs []string) Editor {

	var editorCmd string

	// check envs
	var editorCmdFromEnv string
	for _, editorEnv := range editorEnvs {
		editorCmdFromEnv = os.Getenv(editorEnv)
		// use first found
		if editorCmdFromEnv != "" {
			editorCmd = editorCmdFromEnv
			log.Debugf("Found editor %v from env[%v]", editorCmd, editorEnv)
			break
		}
	}

	// fallback to defaults if nothing is to be found in env
	if editorCmd == "" {
		switch runtime.GOOS {
		case "linux", "darwin":
			editorCmd = "vi"
		case "windows":
			editorCmd = "notepad"
		default:
			log.Panic(fmt.Sprintf("Unsupported OS detected: %v", runtime.GOOS))
		}
		log.Debugf("Fall back to default editor '%v' for OS '%v'", editorCmd, runtime.GOOS)
	}
	return Editor{
		editor: editorCmd,
	}
}

func (ed Editor) Start(fileName string) {
	if ed.editor == "" {
		log.Errorf("Editor not specified, unable to start editor for file %v", fileName)
	}
	absFileName, err := filepath.Abs(fileName)
	if err != nil {
		log.Errorf("Somethig went wrong when normalizing file %v; %v", fileName, err)
	}
	// args := e.args(abs)
	cmd := exec.Command(ed.editor, absFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		log.Errorf("Could not launch cmd %v %v", absFileName, absFileName)
	}
}

func (ed Editor) Edit(content []byte) ([]byte, error) {
	tmpFile, err := ioutil.TempFile("", "cntb.edit")
	if err != nil {
		log.Errorf("Could not create working file %v; %v", tmpFile, err)
	}
	log.Debugf("Created working file %v", tmpFile.Name())
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tmpFile.Write(content)
	ed.Start(tmpFile.Name())

	return ioutil.ReadFile(tmpFile.Name())
}
