package index

import (
	"fmt"
)

type IndexLink struct {
	Id       string
	PathFile string
}

func (ixl *IndexLink) Parse(fi *FileInfo) string {
	ixl.PathFile = fi.Path + "/" + fi.Name
	ixl.Id = splitName(fi.Name) + "_" + fi.ModTime
	if fi.IsShow {
		return fmt.Sprintf("[%s]: %s\n", sha_256(ixl.Id, "animeic"), ixl.PathFile)
	} else {
		return ""
	}

}
