package component

import "os"

type FileComponent struct {
	name      string
	directory *os.File
}

func (comp *FileComponent) createNode(uri string) {

}
