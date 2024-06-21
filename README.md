# Webdav-server

Webdav-server implements WebDAV (Web Distributed Authoring and Versioning) allowing for file sharing, editing. and versioning.
Specifically webdav-server is set up to integrate with linux like operating systems for the ability to use as a "cloud storage" using the native file managment system. 
The goal of this repository is to abstract the complexity when creating a WebDAV server, for use in home enviroments. 

# Installation

'''bash 
git clone https://github.com/braxtonhardman/webdav-server.git
''' 

Then change directory to the webdav-server directory

'''bash 
cd webdav-server
'''

In order for script to work to handle CLI webdav-server must be in /Users/current_users/ directory. 
Move folder to specified directory 

'''bash 
mv ./ /Users/current_user/ 
'''

To execute commands 

'''bash 
cd webdav-server/scripts/ 
'''

Then use
 
'''bash 
./webdav-server.sh start 
'''

to start a server. 
For more commands use 

'''bash 
./webdav-server.sh commands
'''

For ease of use to execute commands from anywhere in system. 
Make sure current directory is ~/webdav-server/scripts/ then move the script to /usr/local/bin 

'''bash 
sudo mv webdav-server.sh /usr/local/bin
'''

Now commands can be executed as following: 

'''bash 
webdav-server <options> 
'''
