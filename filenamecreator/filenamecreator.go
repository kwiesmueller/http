package filenamecreator

import "os"

type FilenameCreator interface {
	Create(directory *os.File) string
}
