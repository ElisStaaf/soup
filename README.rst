Soup - Grepping tool
====================
.. image:: https://img.shields.io/badge/Build%20(Fedora)-passing-2a7fd5?logo=fedora&logoColor=2a7fd5&style=for-the-badge
   :alt: Build = Passing
   :target: https://github.com/ElisStaaf/soup
.. image:: https://img.shields.io/badge/Version-1.0.0-2dd245?style=for-the-badge
   :alt: Version = 1.0.0
   :target: https://github.com/ElisStaaf/soup
.. image:: https://img.shields.io/badge/Language-Go-20c9df?logo=Go&style=for-the-badge
   :alt: Language = Go
   :target: https://github.com/ElisStaaf/soup

I got tired of using stable command line tools like `grep`, so i made an unstable one!
It works... I think? But hey! It's written in go!

Requirements
------------
* `go`_
* `make`_
* `git`_ or `gh`_

Install
-------
Firstly, you wanna clone the repo:

.. code:: bash

   # git
   git clone https://github.com/ElisStaaf/soup
   # gh
   gh repo clone ElisStaaf/soup

Then, you want to build the executable:

.. code:: bash

   sudo make install

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

.. code:: bash

   bash> soup -re -n "Hello World!" main.go
   main.go:8     fmt.Println("Hello World!")

Conclusion
----------
So, yeah, my version is infinitely better because it's written in go, and go is awesome. If you say that anything other than go is best for these kinds of
things, i will find you. Anyway, feel free to contribute and enjoy the project!

.. _`go`: https://go.dev/doc/install
.. _`git`: https://git-scm.com/downloads 
.. _`gh`: https://github.com/cli/cli#installation
.. _`make`: https://www.gnu.org/software/make
