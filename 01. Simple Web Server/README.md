# Simple Web Server in Go

This is a basic HTTP web server implementation in Go that demonstrates fundamental web development concepts including static file serving, form handling, and custom route handlers.

## Project Structure

```
01. Simple Web Server/
├── main.go          # Main server implementation
├── static/          # Static files directory
│   ├── index.html   # Home page
│   └── form.html    # Form page
└── README.md        # This file
```

## Features

- **Static File Serving**: Serves HTML files from the `/static` directory
- **Form Handling**: Processes POST requests with form data
- **Custom Route Handlers**: Implements custom endpoints with specific functionality
- **Error Handling**: Proper HTTP error responses for invalid requests

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to the project directory:
   ```cmd
   cd "01. Simple Web Server"
   ```
3. Run the server:
   ```cmd
   go run main.go
   ```
4. Open your browser and visit:
   - `http://localhost:8080` - Home page (serves static files)
   - `http://localhost:8080/hello` - Hello endpoint
   - `http://localhost:8080/form.html` - Form page
   - `http://localhost:8080/submit` - Form submission endpoint (POST only)

## Code Tutorial: Understanding `main.go`

### Package Declaration and Imports

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)
```

- `package main`: Declares this as an executable program (not a library)
- Import statements bring in necessary packages:
  - `fmt`: For formatted I/O operations
  - `log`: For logging errors
  - `net/http`: For HTTP server functionality

### Function: `helloHandler`

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello World!")
}
```

**Purpose**: Handles requests to the `/hello` endpoint.

**Parameters**:
- `w http.ResponseWriter`: Used to write the HTTP response back to the client
- `r *http.Request`: Contains information about the incoming HTTP request

**Logic**:
1. **Path Validation**: Checks if the requested path is exactly `/hello`
2. **Method Validation**: Ensures only GET requests are allowed
3. **Response**: Sends "Hello World!" to the client

**Key Concepts**:
- HTTP status codes (`http.StatusNotFound`)
- Request validation
- Writing responses with `fmt.Fprint()`

### Function: `formHandler`

```go
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful!\n")

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}
```

**Purpose**: Processes form submissions from HTML forms.

**Logic**:
1. **Parse Form Data**: `r.ParseForm()` extracts form data from the request
2. **Error Handling**: Checks for parsing errors and responds appropriately
3. **Extract Values**: Gets specific form fields using `r.FormValue()`
4. **Response**: Displays the submitted form data back to the user

**Key Concepts**:
- Form data parsing
- Error handling with Go's error interface
- Formatted string output with `fmt.Fprintf()`
- Extracting form values by name

### Function: `main`

```go
func main() {
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/submit", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

**Purpose**: Entry point of the program; sets up and starts the HTTP server.

**Setup Process**:
1. **Static File Server**: Creates a file server for the `./static/` directory
2. **Route Registration**:
   - `http.Handle("/", fileServer)`: Serves static files from root path
   - `http.HandleFunc("/submit", formHandler)`: Routes `/submit` to form handler
   - `http.HandleFunc("/hello", helloHandler)`: Routes `/hello` to hello handler
3. **Server Start**: Starts the server on port 8080

**Key Concepts**:
- `http.FileServer()`: Built-in static file serving
- `http.Handle()` vs `http.HandleFunc()`: Different ways to register handlers
- `http.ListenAndServe()`: Starts the HTTP server
- Error handling with `log.Fatal()`

## Learning Points

### 1. **Handler Functions**
Go HTTP handlers follow the signature: `func(http.ResponseWriter, *http.Request)`
- `ResponseWriter`: Write response data
- `Request`: Read request data

### 2. **Routing**
- `http.HandleFunc()`: Register function handlers for specific paths
- `http.Handle()`: Register handler objects (like file servers)

### 3. **Static Files**
- `http.FileServer()`: Serves files from a directory
- `http.Dir()`: Specifies the directory to serve from

### 4. **Request Processing**
- Check request methods (`r.Method`)
- Validate request paths (`r.URL.Path`)
- Parse form data (`r.ParseForm()`)
- Extract form values (`r.FormValue()`)

### 5. **Response Writing**
- `fmt.Fprint()`: Write plain text
- `fmt.Fprintf()`: Write formatted text
- `http.Error()`: Send HTTP error responses

### 6. **Error Handling**
- Go's explicit error handling pattern
- Using `log.Fatal()` for unrecoverable errors
- Proper HTTP status codes

## Common Issues

- **Port already in use**: Change the port number in `http.ListenAndServe()`
- **Static files not loading**: Ensure the `static/` directory exists and contains your HTML files
- **Form not submitting**: Check that form method is POST and action points to `/submit`
