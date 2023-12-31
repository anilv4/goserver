# GoServer

GoServer is a simple yet powerful HTTP server written in Go. It serves HTML files from a specified directory and logs all HTTP interactions, including successful responses and errors. This server is ideal for situations where you need to quickly deploy a lightweight, configurable web server to serve static HTML content.

Purpose of this application is for testing multiple Openshift and K8s deployment scenarios.

## Features

- **Serve HTML Files:** Dynamically serves HTML files from a specified directory.
- **HTTP Logging:** Logs details of all HTTP requests including client IP, request type, path, response status, and processing time.
- **Customizable via Command-Line Flags:** Easily configure the server's listening port and home directory for HTML files.
- **Error Handling:** Properly handles file-not-found scenarios, permission issues, and other file reading errors, responding with appropriate HTTP status codes.

## Installation

To install and run GoServer, follow these steps:

1. **Compile the Source Code:**
   - Ensure you have Go installed on your system.
   - Clone or download the repository.
   - Navigate to the source directory and compile the code using:
     ```
     go build -o goserver
     ```

2. **Run the Server:**
   - Run the server using:
     ```
     ./goserver
     ```
   - Optionally, use command-line flags to customize the server's behavior.

## Usage

### Command-Line Flags

- `--port=<port>`: Define the listening port (default is 8080).
- `--home=<directory>`: Specify the home directory from where HTML files are served (default is `/tmp/home`).
- `--logging=<true|false>`: Enable or disable HTTP access logging (default is true).

### Accessing Content

- Access the content by navigating to `http://<server-address>:<port>/<filename>`.
- The server looks for files with a `.html` extension in the specified home directory.
- Special endpoints `/os` and `/hostname` serve the system's OS release info and hostname, respectively.

### Example

Start the server with a custom port and home directory:

```
./goserver --port=9000 --home=/path/to/html/files
```

Access an HTML file (`example.html`) located in `/path/to/html/files`:

http://localhost:9000/example

## Logging

- Logs are printed to the console.
- Each log entry includes the server hostname, client IP, HTTP request method, request path, response status code, and the time taken to process the request in milliseconds.

## Error Responses

- `200`: Returned if the requested HTML file is available.
- `404 Not Found`: Returned if the requested HTML file is not found.
- `501 Not Implemented`: Returned for permission issues while accessing files.
- `500 Internal Server Error`: Returned for all other file reading errors.

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](#) for open issues or open a new issue to discuss changes or features you would like to add.

## License

Distributed under the MIT License. See `LICENSE` for more information.

