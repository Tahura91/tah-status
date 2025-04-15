
# `tah-status` - Website Status Checker

A lightweight and easy-to-use tool built with **Go** to check the status and uptime of websites. With this tool, you can monitor whether a website is **up** or **down**, and analyze HTTP status codes for any given URL.

## 📦 **Features**

- 🚀 **Check Website Status**: Check whether a given website is online or offline.
- 💡 **Detailed HTTP Status Codes**: See exact HTTP status codes returned by the website (200, 404, etc.).
- 🔄 **Simple CLI**: Clean, easy-to-use command-line interface.
- ⏱ **Uptime Monitoring**: Continuous monitoring of multiple websites for uptime tracking.
- 🛠 **Built with Go**: Written in Go for speed and efficiency.

## 🖥️ **Installation**

To get started with `tah-status`, follow these easy steps:

### 1. **Clone the Repository**

   First, clone the repository to your local machine:

   ```bash
   git clone https://github.com/Tahura91/tah-status.git
   ```

### 2. **Install Dependencies**

   Make sure you have **Go** installed on your system. You can download it from the official website: [https://golang.org/dl/](https://golang.org/dl/).

   Install any dependencies by running:

   ```bash
   go mod tidy
   ```

### 3. **Build the Project**

   To build the project, run:

   ```bash
   go build
   ```

   This will generate an executable file.

## 🚀 **Usage**

### 1. **Check a Single URL**

To check the status of a single URL, use the following command:

```bash
go run main.go --url https://example.com
```

### 2. **Check Multiple URLs from a File**

You can also check multiple URLs by putting them in a file (`urls.txt`) and running the following command:

```bash
go run main.go --file urls.txt
```

Where `urls.txt` is a text file containing a list of URLs, one per line:

```
https://www.google.com
https://www.github.com
https://www.twitter.com
```

### 3. **Display Help**

To display a help message with all available options and usage instructions, run:

```bash
go run main.go --help
```

## 📈 **Example Output**

After running the program, the output will display the status of each website:

```bash
go run main.go --file urls.txt
>>
✅ https://www.google.com - 200 OK (0.57s)
✅ https://www.facebook.com - 200 OK (0.34s)
❌ https://www.twitter.com - 400 Bad Request (0.98s)
✅ https://www.github.com - 200 OK (0.20s)
✅ https://www.wikipedia.org - 200 OK (0.35s)
✅ https://www.youtube.com - 200 OK (0.23s)
✅ https://www.amazon.com - 200 OK (0.76s)
✅ https://www.reddit.com - 200 OK (0.83s)
✅ https://www.stackoverflow.com - 200 OK (1.11s)
✅ https://www.apple.com - 200 OK (0.15s)
✅ https://www.microsoft.com - 200 OK (0.32s)
```

## 🔧 **Configuration**

- **URLs List**: Make sure to list all the websites you want to check in the `urls.txt` file. Each URL should be on a new line.
- **HTTP Status Codes**: The program will output HTTP status codes such as:
  - `200 OK`: Website is up and running
  - `404 Not Found`: Website is down or page is missing
  - Other status codes: (e.g., `503 Service Unavailable`)

## 🌍 **Contributing**

We welcome contributions! If you'd like to improve `tah-status`, please feel free to fork the repository and submit a pull request.

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new feature'`)
5. Push to the branch (`git push origin feature/your-feature`)
6. Create a pull request

## 📞 **Author**

Created by **Tahura** – feel free to reach out on [LinkedIn](https://www.linkedin.com/in/tahura-hayath-483397243/) for any questions or feedback.
```

---

