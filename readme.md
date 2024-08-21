# TCP Port Scanner
## Overview
This is a simple TCP port scanner written in Go. It scans a range of ports (1-1024) on a specified host to identify open ports. The tool is designed to be lightweight and efficient, using concurrency to speed up the scanning process.

## Features
* Scans ports 1 through 1024 on a specified host.
* Concurrent scanning using Goroutines.
* Displays a list of open ports.


## Usage
To use this tool, you need to have Go installed on your system.
### 1) Clone the repository (if not already cloned):
```bash
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
```
### 2) Build the application:
```bash
go build -o tcpScanner
```
### 3) Run the application
You need to specify the host you want to scan. For example, to scan ports on example.com:
```bash
./tcpScanner example.com
```
The application will then scan ports 1 through 1024 and print out any that are open.

## Disclaimer 
This code is provided for educational purposes only. You are responsible for using it ethically and legally. You must use this code only on targets for which you have explicit permission. Unauthorized use of this code on systems you do not own or control is illegal and may result in severe consequences. The author assumes no responsibility for any misuse or damage caused by this code.

## Contribution
Feel free to open an issue or submit a pull request if you have any suggestions or improvements.