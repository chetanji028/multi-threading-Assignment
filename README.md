# Distributed File Storage Server

A Golang-based distributed file storage server with API support for uploading, retrieving, and downloading files. The application uses PostgreSQL for storage and supports multithreading for efficient file operations.

## Features

- **Upload File**: Split files into multiple parts and store them in the database in parallel.
- **Get File Data**: Retrieve metadata or parts information of uploaded files.
- **Download File**: Fetch file parts in parallel and merge them to reconstruct the original file.
- **API Documentation**: Available via Swagger/OpenAPI.
- **Dockerized**: Easily deploy using Docker Compose.

## Requirements

- Docker
- Docker Compose

## Setup and Run

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/distributed-file-storage.git
cd distributed-file-storage


