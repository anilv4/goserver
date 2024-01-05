package main

import (
    "context"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

func main() {
    var port, homeDir string
    var logging bool

    flag.StringVar(&port, "port", "8080", "Define the listening port")
    flag.StringVar(&homeDir, "home", "/tmp/home", "Define the home directory")
    flag.BoolVar(&logging, "logging", true, "Enable or disable HTTP access logging")
    flag.Parse()

    printStartupMessage(port, homeDir, logging)

    handler := func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()

        rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        r = r.WithContext(context.WithValue(r.Context(), responseWriterKey, rw))

        defer func() {
            if logging {
                logRequest(r, rw.statusCode, startTime)
            }
        }()

        path := r.URL.Path[1:]
        if path == "" {
            path = "hostname"
        }

        var filename string
        switch path {
        case "os":
            filename = "/etc/os-release"
        case "hostname":
            filename = "/etc/hostname"
        default:
            filename = filepath.Join(homeDir, path+".html")
        }

        content, err := ioutil.ReadFile(filename)
        if err != nil {
            if os.IsNotExist(err) {
                rw.statusCode = http.StatusNotFound
            } else if os.IsPermission(err) {
                rw.statusCode = http.StatusNotImplemented
            } else {
                rw.statusCode = http.StatusInternalServerError
            }
            http.Error(rw, http.StatusText(rw.statusCode), rw.statusCode)
            return
        }

        rw.Write(content)
    }

    http.Handle("/", http.HandlerFunc(handler))

    fmt.Printf("Starting server at port %s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %s\n", err.Error())
    }
}

func printStartupMessage(port, homeDir string, logging bool) {
    fmt.Println("Starting GoServer...")
    fmt.Println("A simple HTTP server to serve HTML files with Go.")
    fmt.Printf("Listening Port: %s (default 8080)\n", port)
    fmt.Printf("Home Directory: %s (default /tmp/home)\n", homeDir)
    fmt.Printf("Logging Enabled: %t (default true)\n", logging)
    fmt.Println("Use --port=<port>, --home=<directory>, and --logging=<true|false> to configure.")
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

const responseWriterKey = "responseWriter"

func logRequest(r *http.Request, statusCode int, startTime time.Time) {
    duration := time.Since(startTime)
    clientIP := getClientIP(r)
    hostname, err := os.Hostname()
    if err != nil {
        hostname = "unknown"
    }

    log.Printf("%s - %s - %s %s - %d - %dms\n",
        hostname, clientIP, r.Method, r.URL.Path, statusCode, duration.Milliseconds())
}

func getClientIP(r *http.Request) string {
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return "unknown"
    }
    return ip
}
