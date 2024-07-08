# Coreutils 

## Description 
project implementing basic unix commands with goLang


## Installation
To use these utilities, you need to have Go installed on your machine. After installing Go, you can clone this repository and build the binaries with the following commands:
```bash
git clone https://github.com/codescalersinternships/Coreutils-SalmaElsoly
cd cmd/<toolName>
go build <tool>.go
./<tool> <flags>
```
## Usage
Here are the usage instructions for each implemented command:

### 1. Cat
 * ` ./cat -n <filepath> `
-n : for numbered lines output

### 2. Echo
`./echo -n <string>`
* Print arguments to standard output.
* Add a -n flag to omit the trailing newline.

### 3. Tree
` ./tree -L <num> <filepath> `
* To display a recursive directory tree
* If no file specifed it display tree of current directory
* Print files and directories up to 'num' levels of depth
* Default level is 1.

### 4. Env
* ` ./env <variable> `
To get the value for specfic variable
* ` ./env <variable>=<value> `
To set the value for the specfied variable

### 5. Head
`./head -n <num> <filepath>`
* Display the first 'num' lines of a file.
* If 'num' is not specified, the default is to display the first 10 lines.

### 6. Tail
`./tail -n <num> <filepath>`
* Display the last 'num' lines of a file.
* If 'num' is not specified, the default is to display the last 10 lines.

### 7. Wc
`./wc <flags> <filepath>`
* Count lines, words, and characters in the input.
* Add -l, -w, and -c flags to display only lines, words, or characters respectively.

### 8. Yes
`./yes <string>`
* Repeatedly outputs the specified string, or "y" if no string is specified.

### 9. False
`./false`
* Exits with a non-zero status code to represent an unsuccessful operation or failure.

### 10. True
`./true`
* Exits with a zero status code to represent success.