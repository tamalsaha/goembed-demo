package main

import (
	"fmt"
	"github.com/tamalsaha/goembed-demo/hub/resourcedescriptors"
	"io/fs"
	"path/filepath"
)

func main() {
	f := resourcedescriptors.FS()
	//entries, err := f.ReadDir("apps")
	//if err != nil {
	//	panic(err)
	//}
	//for _, e := range entries {
	//	fmt.Println(e.Name(), e.Type(), e.IsDir())
	//}
	fs.WalkDir(f, ".", func(path string, e fs.DirEntry, err error) error {
		if !e.IsDir() {
			fmt.Println(filepath.Join(path, e.Name()))
		}
		return err
	})
}
//
//func main() {
//	useOS := len(os.Args) > 1 && os.Args[1] == "live"
//	http.Handle("/", http.FileServer(getFileSystem(useOS)))
//	http.ListenAndServe(":8888", nil)
//}
//
//func getFileSystem(useOS bool) http.FileSystem {
//	if useOS {
//		log.Print("using live mode")
//		return http.FS(os.DirFS("static"))
//	}
//
//	log.Print("using embed mode")
//	fsys, err := fs.Sub(catalog.FS, "raw")
//	if err != nil {
//		panic(err)
//	}
//
//	return http.FS(fsys)
//}
