# Go Web Crawler

This project implements a web crawler in Go, designed to crawl web pages starting from a given URL and extract links.

## Features

- **Crawl**: Initiates a crawl starting from a specified URL.
- **Sitemap**: Retrieves a sitemap of all crawled URLs.
- **CLI Tool**: Provides a command-line interface (CLI) for easy interaction.

## Project Structure

The project is structured as follows:

- **`main.go`**: Entry point for the web crawler server.
- **`internal/`**: Contains main project modules.
  - **`crawler/`**: Implements the web crawler logic.
  - **`extractor/`**: Handles link extraction from HTML pages.
  - **`reader/`**: Reads web pages and extracts links.
  - **`connector/`**: Connects to websites and fetches HTML content.
  - **`util/`**: Utility functions (e.g., URL normalization).
  - **`logger/`**: logger Utility functions.
- **`pkg/api/`**: HTTP API handlers.
- **`go.mod`, `go.sum`**: Go module files.

## Getting Started

### Prerequisites

- Go programming language (version 1.17 or higher)
- Docker (optional, for containerization)

### Installation

1. Extract the project files:

   ```bash
    tar -xvf webcrawler.tar.gz
    cd webcrawler/
   ```

2. Install dependencies and Build the project:

   ```bash
   make build
   ```

### Usage

#### Running Locally

To start the web crawler server locally, use:

```bash
go run main.go
```
Or use the binary build by make file
```
./webcrawler
```

#### CLI Commands

The CLI tool supports the following commands:

- **`crawl run `**:  Start a crawl and retrieve the sitemap once
- **`crawl sitemap`**: Retrieve the sitemap of crawled URLs.

Example usage:

```bash
./webcrawler crawl run --endpoint https://example.com
./webcrawler crawl sitemap --id 213213
./webcrawler crawl start --endpoint https://example.com
```

### Docker

To run the web crawler in a Docker container:

1. Build the Docker image:

   ```bash
   docker build -t webcrawler .
   
   or 

   docker pull nageshnode/webcrawler:server
   ```
  
2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 webcrawler
   ```

Replace `8080:8080` with the desired host port and container port if necessary.

### Testing

Run tests using:

```bash
go test ./...
```
