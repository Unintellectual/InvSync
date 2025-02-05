# Offline Inventory Management Server

A Go-based server that connects devices in an offline environment. The server supports managing and querying inventory data from a MySQL database, designed to handle up to 10,000 concurrent processes. It uses the Gin web framework and MySQL for data storage.

## Features

- **Offline-capable**: Designed to work in environments without an internet connection.
- **Concurrency**: Efficiently handles up to 10,000 concurrent processes, making it suitable for large-scale operations.
- **MySQL Integration**: Stores inventory data in a MySQL database, including food items, quantities, prices, categories, and expiry dates.

## Prerequisites

Before running the server, you need:

1. **Go 1.18+**: To build and run the Go server.
2. **MySQL 8.0+**: For the database and data storage.

## Installation

### 1. Clone the Repository

```sh
git clone https://github.com/yourusername/inventory-server.git
cd inventory-server
