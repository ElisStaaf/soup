# Soup - Grepping tool
[![Version](https://img.shields.io/badge/Version-1.0.0-a53fc0)](https://github.com/ElisStaaf/soup)
[![Build](https://img.shields.io/badge/Build_(openSUSE)-passing-19e646?logo=opensuse&logoColor=19e646)](https://github.com/ElisStaaf/soup)
[![Language](https://img.shields.io/badge/Language-Go-20c9df?logo=Go)](https://github.com/ElisStaaf/soup)    
I got tired of using stable command line tools like `grep`, so i made an unstable one!
It works... I think? But hey! It's written in go!

Install
-------
Firstly, you wanna clone the repo:
```bash
# Git
git clone https://github.com/ElisStaaf/soup
# Gh
gh repo clone ElisStaaf/soup
```
Then, you want to build the executable:
```bash
./install.sh
# OR
go build soup.go
```
Then you're done!

Usage
-----
The parameters for soup are:
* query
* path
But you can *also* add flags to customize the parser:
* -n - display line number for non-binary files
* -re - treat query as a regex (regular expression)
Here's an example of what a valid soup invocation looks like:
```bash
$ soup -re -n Hello World! main.go
main.go:8     fmt.Println("Hello World!")
```

Conclusion
----------
So, yeah, my version is infinitely better because it's written in go, and go is awesome. If you say that anything other than go is best for these kinds of
things, i will find you. Anyway, feel free to contribute and enjoy the project!
