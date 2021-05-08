package main

import (
	"fmt"
	"github.com/tamalsaha/goembed-demo/hub/resourcedescriptors"
	"io/fs"
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

	fs.WalkDir(f, ".", func(filename string, e fs.DirEntry, err error) error {
		if !e.IsDir() {
			fmt.Println(filename)
		}
		return err
	})
	fmt.Println("--------------------------------------------------")

	data, err := f.ReadFile("app.k8s.io/v1beta1/applications.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
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
