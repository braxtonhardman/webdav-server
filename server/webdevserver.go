package webdevserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"github.com/braxtonhardman/webdav-server/log"
	"golang.org/x/net/webdav"
)

func Start() { 

    
    
    // Starting logger to be used later 
    logger.Start() 
    logger.LogSystem("Server Started")

    user, _ := user.Current()

    // Root Directory of application 
    dataDir := filepath.Join(user.HomeDir, "webdav-server")
    logDir := filepath.Join(dataDir, "log")
                        
    
    // Check if the root directory /home/johndoer/webdavserver/log exists
    if _, err := os.Stat(dataDir); os.IsNotExist(err) {
        logger.LogError(err)
        os.Exit(1)
    }

    // Check if the root directory /home/johndoe/webdavserver/log exists
    if _, err := os.Stat(logDir); os.IsNotExist(err){ 
        log.Fatal(err)
    }
    
    // Specifies that all routes containg prefix /webdav/ will be handled here
    handler := &webdav.Handler{
        Prefix:     "/webdav/",
		// Creates a filesystem implementation on that directory 
        FileSystem: webdav.Dir(dataDir),
        LockSystem: webdav.NewMemLS(),
    }
    
    logger.LogSystem("Handler Created in " + dataDir) 

    http.HandleFunc("/webdav/", func(w http.ResponseWriter, r *http.Request) {
		// Specifies which methods the server supports, needed for the client to see because mac excpects to match these 
        w.Header().Set("Allow", "OPTIONS, GET, HEAD, POST, PUT, DELETE, PROPFIND, MKCOL, COPY, MOVE, LOCK, UNLOCK")
		// Specifies which version of webDav are suppored by the server 
        w.Header().Set("DAV", "1, 2")
		// Specifies which type of content is returned 
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        logger.LogSystem("Method " + r.Method +  " from address " + r.RemoteAddr)
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Checking to make sure the method being used matches are server and if the path is correct 
        if r.Method == http.MethodGet && r.URL.Path == "/webdav/" {
            w.Header().Set("Content-Type", "text/html")
            // Reads
            entries, err := os.ReadDir(dataDir)
            if err != nil {
                http.Error(w, "Unable to read directory", http.StatusInternalServerError)
                return
            }

            w.WriteHeader(http.StatusOK)
            w.Write([]byte("<html><body><ul>"))
            for _, entry := range entries {
                w.Write([]byte("<li>" + entry.Name() + "</li>"))
            }
            w.Write([]byte("</ul></body></html>"))
            return
        }

        handler.ServeHTTP(w, r)
    })

    addr := "0.0.0.0:8080"

    logger.LogSystem("WebDAV server listening on: " + addr) 

    err := http.ListenAndServe(addr, nil)

    if err != nil {
        logger.LogError(err)
        log.Fatal(err)
        fmt.Println("Error Occured starting server")
        os.Exit(1)
    }
    
}


