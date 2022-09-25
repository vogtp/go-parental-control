package build

import (
	"embed"
	"path"

	"github.com/vogtp/go-hcl"
)

var (
	//go:embed *
	assetData embed.FS
)

func ListFiles() {
	out, err := getAllFilenames(&assetData, "")
	if err != nil {
		hcl.Warnf("Cannot list files: %v", err)
	}
	for _, l := range out {
		hcl.Infof("Binary: %v", l)
	}
}

func getAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}
