# SGEmu

SGEmu is a Shattered Galaxy server emulator by ThreeSix.

#### Requirements

Go Lang - for code compliation
MongoDB - for the server database

#### Getting started (Linux)

Download the `build.sh` script and execute it in order to download the source code and build the binaries.

This will create a Go path in your home directory - `$HOME/go` and checkout the code.

Binaries will be created at `$HOME/sgemu`

Packaging of additional files required by each binary hasn't been implemented yet.

**TODO: Improve this and add documentation for initial server setup**




Previous thread:
http://forum.ragezone.com/f111/devel...server-523885/


tinkle asked me to look on this one so I wanted to give a little push to this server development and I have coded a little server in Go (sorry that's not C# ><).

Anyway, encryption/decryption is done and I can get to the main screen.
Source and picture is in the attachments,here's the localhost client.

Feel free to rewrite this in C# and continue the emu.
That's just a side project so I won't work on it a lot, I just hope someone will grab it and continue or post here packets and info and I will code it in.

Information about Go:
The Go Programming Language
Can be executed without any libs or frameworks and also cross-platform, I'am compiling the project using Eclipse
goclipse - Eclipse-based IDE for Google Go - Google Project Hosting


Rep:
sgemu - A Shattered Galaxy server emulator - Google Project Hosting

Tutorial:
Code:
Download mango Downloads - MongoDB
put in C:\Mongo.
Download last rev from repo or Downloads - sgemu - A Shattered Galaxy server emulator - Google Project Hosting.
Run mangod from C:\Mongo\bin\mangod.exe using the args --dbpath C:\Mongo
so its C:\Mongo\bin\mangod.exe "--dbpath C:\Mongo".
Change host file (Windows\System32\drivers\etc) to 
127.0.0.1 sg0.kru.com 
127.0.0.1 bbs.sgalaxy.com.
Run main.exe from the server folder.
Run SG.exe.

Enjoy.
Notes:
I'm not the creator of SGLocalHost.exe.
This project is for educational purpose only and not for hosting/selling/etc therefore it should run on localhost only.
Development is slow.

Server commands(just type them in the console while the server is running):
exit - exits the console
cleardb - deletes all the users

edit host file to enter your server
127.0.0.1 sg0.kru.com
127.0.0.1 bbs.sgalaxy.com

