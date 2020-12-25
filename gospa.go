package gospa

import (
	"net/http"
	"os"
	"path/filepath"
)

type spaHandler struct {
	rootPath  string
	indexFile string
}

func (spa *spaHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := filepath.Join(spa.rootPath, r.URL.Path) // filepath.Join cleans the result
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		// if path does not exists || is a directory, serve index
		http.ServeFile(rw, r, filepath.Join(spa.rootPath, spa.indexFile))
		return
	}
	// now its safe to serve the requestes file
	http.ServeFile(rw, r, path)
}

// NewSpaHandler serves the SPA located in 'rootPath' at "/".
// If the spa should not be servede at "/", use "NewSpaHandlerAt" instead.
// It serves any requested file inside 'rootPath' if it exists
// - 'indexFile' otherwise or if a directory is requested
//
// - 'indexFile' is relative to 'rootPath'
func NewSpaHandler(rootPath string, indexFile string) http.Handler {
	return &spaHandler{rootPath, indexFile}
}

// NewSpaHandlerAt just passes the spaHandler through 'http.StripPrefix()'
func NewSpaHandlerAt(pathPrefix string, rootPath string, indexFile string) http.Handler {
	return http.StripPrefix(pathPrefix, &spaHandler{rootPath, indexFile})
}
