# About

A straightforward command-line interface (CLI) that counts the number of words in a text file provided as an argument, crafted in Go with the urfave cli library.

# How to run

To run the CLI, use the following command:

```
$ go run main.go -f tmp.txt
```

`Observation:` you need to specify the file path by using the flag `--file` or `-f`.

To display the help menu, utilize the flag `--help` or `-h`:

```
$ go run main.go --help
```

# Reference(s)