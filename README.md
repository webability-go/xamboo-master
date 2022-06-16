Xamboo-Master for GO v1
=============================

The full manual is on xamboo github:

https://github.com/webability-go/xamboo

Xamboo-Master is a system plugin for the Xamboo system to administrate the application with a web interface
for the xamboo configuration and maintenance, and xmodules installation.


TO DO
=======================

- Components admin and components config admin
- Engines admin
- Datasources container admin
- Datasources admin
- Listeners admin
- Hosts admin
- plugins admin
- Xmodules admin

Version Changes Control
=======================

V0.1.2 - 2022-06-16
-----------------------
- Creation of appEntries to access app functions
- Needed code rewritten to use the appEntries

V0.1.1 - 2022-06-14
-----------------------
- removed bridge support (not anymore needed)
- All the code rewriten without bridge, and with new XDommask hooks

V0.1.0 - 2022-06-09
-----------------------
- Made compatible with Xamboo 1.6
. Nade compatible with xmodules/base 2022-06-09

V0.0.7 - 2021-04-22
-----------------------
- Changes to meet the new xmodules structure: remove the bridge, use directly the functions into the application.
- Enhance the administration pages with a better display and representation of objects into the system.

V0.0.6 - 2021-03-09
-----------------------
- HostWriter requeststat parameter changed to lowers (the official parameter from the Component is in lowercase)


V0.0.5 - 2021-02-24
-----------------------
- Bug corrected on new option to reload configuration in app.so

V0.0.4 - 2021-02-21
-----------------------
- New option added to the menu to reload the configuration in the server

V0.0.3 - 2021-02-12
-----------------------
- Adaptation to Xamboo v1.5.0 not compatible with previous versions

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
