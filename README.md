# Webdav-server

Webdav-server implements WebDAV (Web Distributed Authoring and Versioning) allowing for file sharing, editing. and versioning.
Specifically webdav-server is set up to integrate with linux like operating systems for the ability to use as a "cloud storage" using the native file managment system. 
The goal of this repository is to abstract the complexity when creating a WebDAV server, for use in home enviroments. 
# Prerequsites

Make sure Go is installed and check any firewalls permissions

# Installation

```bash 
git clone https://github.com/braxtonhardman/webdav-server.git
```

Make sure ~/webdav-server/ is located in the /Users/"CurrentUser"/ directory.
This is necessary in order for the script to work. 

To execute commands make sure working directory is scripts directory. 

```bash 
cd ~/webdav-server/scripts/ 
```

Then to start a server use
 
```bash 
./webdav-server.sh start 
```

For more commands use 

```bash 
./webdav-server.sh commands
```

For ease of use to execute commands from anywhere in system. 
Make sure current directory is ~/webdav-server/scripts/ then move the script to /usr/local/bin 

```bash 
sudo mv webdav-server.sh /usr/local/bin
```

Now commands can be executed as following: 

```bash 
webdav-server.sh <options> 
```
