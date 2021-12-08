package index

import (
	"fmt"
)

type Index struct {
	Name string
	Id   string
}

func (ix *Index) Parse(fi *FileInfo) string {
	ix.Name = splitName(fi.Name)
	ix.Id = splitName(fi.Name) + "_" + fi.ModTime
	if fi.IsShow {
		return fmt.Sprintf("%s[%s][%s]\n", fi.Indent, ix.Name, sha_256(ix.Id, "animeic"))
		// return fmt.Sprintf("%s[%s][%s]\n", fi.Indent, ix.Name, ix.Id)
	} else {
		return fmt.Sprintf("%s%s\n", fi.Indent, ix.Name)
	}

	// WriteIndex(ixf, str)
	// log.Println(str)
}

func (ix *Index) SampleParse(fi *FileInfo) string {
	ix.Name = splitName(fi.Name)
	ix.Id = fi.Path + "/" + fi.Name
	if fi.IsShow {
		return fmt.Sprintf("%s[%s](%s)\n", fi.Indent, ix.Name, ix.Id)
	} else {
		return fmt.Sprintf("%s%s\n", fi.Indent, ix.Name)
	}

}
