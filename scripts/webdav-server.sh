#!/bin/bash

go_executable="$HOME/webdav-server/webdavserver"

ERR_LOG_PATH="$HOME/webdav-server/log/error_log.txt"
SYS_LOG_PATH="$HOME/webdav-server/log/system_log.txt"

local_ip=$(ifconfig | awk '/inet / && !/127.0.0.1/ {print $2}')

WEBDAV_URL="http://$local_ip:8080/webdav/"

PID_FILE="/tmp/go_server.pid"
current_time=""

start_server() {
    
    current_time=$(date +"%Y-%m-%d %H:%M:%S")

    if [ -f "$PID_FILE" ]; then
        echo "Server is already running."
    else

        echo "$go_executable"

        nohup $go_executable start > /dev/null 2>&1 &

        echo $! > $PID_FILE
        echo "Server started."

        echo "Connect to the server" 
        echo "Go to Finder hold command and press K" 
        echo "In the top box paste $WEBDAV_URL" 
        echo "Connect as guest, note: data is not encrypted" 

    fi
    
}

stop_server() {
    current_time=$(date +"%Y-%m-%d %H:%M:%S")
    if [ -f "$PID_FILE" ]; then
        kill $(cat $PID_FILE)
        rm -f $PID_FILE
        echo "$current_time - Server stopped"
    else
        echo "Server is not running."
    fi
}

status_server() {
    if [ -f "$PID_FILE" ]; then
        echo "Server is running."
        echo "Executable path: $go_executable"
        echo "WebDAV URL: $WEBDAV_URL"
    else
        echo "Server is not running."
    fi
}

errlog_server() {
    if [ -f "$PID_FILE" ]; then
        cat "$ERR_LOG_PATH"
    else
        echo "Server is not running."
    fi
}

serverlog_server() {
    if [ -f "$PID_FILE" ]; then
        cat "$SYS_LOG_PATH"
    else
        echo "Server is not running."
    fi
}

commands_server() {
    echo "start      stop           status"
    echo "errlog     serverlog"
}

# Command line argument processing
case $1 in
    start)
        start_server
        ;;
    stop)
        stop_server
        ;;
    status)
        status_server
        ;;
    commands)
        commands_server
        ;;
    errlog)
        errlog_server
        ;;
    serverlog) 
        serverlog_server
        ;; 
    *)
        echo "Usage: webdav-server.sh <options>"
        commands_server
        exit 1
        ;;
esac