@UTF-8

Xamboo-Master for GO v1
=============================

Xamboo-Master is a system plugin for the Xamboo system to administrate the application with a web interface
It contains 2 sites:
- The master site, for the xamboo configuration and maintenance, and xmodules installation.
- The admin site, to administrate the xmodules and others.

Xamboo is the result of over 15 years of manufacturing engineering frameworks, originally written for PHP 7+ and now ported to GO 1.8+

It is a very high quality framework for CMS, made in GO 1.8 or higher, fully object-oriented and strong to distribute code into Web portals with heavy load and REST APIs optimization.

Xamboo is freeware, and uses several other freeware components (XConfig, XCore)

INSTALATION AND COMPILATION
=============================

After installing the Xamboo system, creates a master directory and download the master:

For instance, if you have installed the xamboo in /home/sites/xamboo

$ cd /home/sites/xamboo
$ mkdir master
$ git init

Pull the last verion of Xamboo-Master

$ git pull https://github.com/webability-go/xamboo-master.git

Edit the master/config.json file to adapt the listeners and hosts to your IPs and domain.
Link the master/config.json to your main system config.json adding it in the "include":[] section

Adds the compilation of the master app to your compilation process.

...
go build --buildmode=plugin -o ./master/app/app.so master/app/*.go
...


Compile and restart your xamboo.

Enter in the master web site to configure your access credentials.

The admin credentials are created as default with system / manager when the 'user' xmodule is installed. Change them inmediatly.


TO DO
=======================


Version Changes Control
=======================

V0.0.2 - 2020-01-04
-----------------------
- Menu ordered by different types of objects.
- Main config editor done
- module page nearly done

V0.0.1 - 2020-12-22
-----------------------
- Main first realease
- Basic operation available: installation, main data editable


Manual:
=======================

- If you want to help converting the manual from text into .md file, you are most welcome.
- Translations are also welcome

Refer to the Xamboo main manual configuration file to learn about the options of configuration.
