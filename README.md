Go Port Scanner
============

Port scanner I made with Go programming language.

Usage
-----

1.  Clone the project to a local directory or download the ZIP file and extract it:

    git clone https://github.com/BaverTorun/PortScanner.git

3.  Install the required dependencies (not needed if you have Golang installed):

    go get github.com/BaverTorun/PortScanner/port

5.  Edit the \`main.go\` file to update the target hostname:

    results := port.InitialScan("target-hostname")

7.  Run the code:

    go run main.go
