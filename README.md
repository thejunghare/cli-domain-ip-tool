# Domain IP Lookup Tool

This tool is designed to find the IP addresses from a list of domains in a text file.

## Installation

To use this tool, you'll need to have Go installed on your system. You can download and install Go from the official Go website: https://golang.org

Once you have Go installed, you can build the tool by running the following command:

## Usage

To run the tool, use the following command:

```
domain-ip-tool --file <path-to-input-file>
```

Replace `<path-to-input-file>` with the actual path to your input file. The input file should be a text file containing a list of domains, with each domain on a separate line.

The tool will read the contents of the input file, resolve the IP addresses for each domain, and display the results in the terminal.

## Example

Let's say you have a text file named `domains.txt` with the following contents:

example.com

google.com

facebook.com

github.com

To find the IP addresses for these domains, you can run the following command:

```
domain-ip-tool --file domains.txt
```

The tool will process the file and display the domain and IP address pairs in the terminal.

Domain: example.com, IP: 93.184.216.34

Domain: google.com, IP: 142.250.185.46

Domain: facebook.com, IP: 31.13.66.35

Domain: github.com, IP: 140.82.112.4
