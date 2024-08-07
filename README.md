# http-status-checker

`http-status-checker` is a command-line interface (CLI) tool to check the status of a website. It can be used to check the status of a website and get the response time of the website.

## Features

- Check the status of a website.
- Get the IP address of the website and the local machine
- Display the alias or canonical name of the specified URL 

## Technologies Used
- Golang

## Installation

To install the `http-status-checker` tool, clone this repository and build the project:

```sh
git clone https://github.com/Hrishikesh-Panigrahi/http-status-checker
cd http-status-checker
go build .
```

## Usage

### Basic Usage

Check the status of the specified URL:

```sh
./http-status-checker check [url]
```

### Custom Ping Count

Check the status and response time of the specified URL with specified pings:

```sh
./http-status-checker check [url] [pings]
```

### Check IP

Retrieve and display the IP address of the Local Machine:

```sh
./http-status-checker ip
```

### Check IP of the website

Retrieve and display the IP address of the specified URL:

```sh
./http-status-checker ip [url]
```

### Display the alias of the website

Resolve and display the alias or canonical name of the specified URL:

```sh
./http-status-checker alias [url]
```

### Help Command

To display the help message:

```sh
./http-status-checker --help
```

## Docker Usage

You can also run the tool using Docker. To do this, pull the Docker image and run it:

### Pulling the Docker Image

```sh
docker pull hrishikeshpanigrahi025/http-status-checker
```

### Running the Docker Container

To check the status of a website with a specific number of pings:

```sh
docker run hrishikeshpanigrahi025/http-status-checker check www.google.com 2
```

## Examples

### Checking Status with Default Pings

If you don't specify the number of pings, it will default to 4 pings:

```sh
./http-status-checker check www.google.com
```

### Checking Status with Custom Ping Count

Specify the number of pings:

```sh
./http-status-checker check www.google.com 10
```

### Displaying Help

```sh
./http-status-checker --help
```

## Flags

--help: Show help for the hsc command.
