#!/bin/bash

PROJ_NAME="1panel"

TARGET_DIR="/lib/systemd/system"
SERVICE_NAME="${PROJ_NAME}.service"
SOURCE_PATH="/opt/1panel/${SERVICE_NAME}"

# 检查是否传入了stop参数
if [ "$1" = "stop" ]; then
    echo "Stopping $SERVICE_NAME..."
    systemctl stop "$SERVICE_NAME"
    echo "$SERVICE_NAME stopped."
    ps auxf | grep -v "grep" |grep $PROJ_NAME
    # ps auxf |grep java |grep $PROJ_NAME |grep -v grep | awk '{print $2}' |xargs kill -9
    exit 0
fi

if [ ! -f "$TARGET_DIR/$SERVICE_NAME" ]; then
    cp "$SOURCE_PATH" "$TARGET_DIR/"
    systemctl enable $SERVICE_NAME
    systemctl daemon-reload
fi

# 检查服务的状态
SERVICE_STATUS=$(systemctl is-active "$SERVICE_NAME")

if [ "$SERVICE_STATUS" = "active" ]; then
    echo "$SERVICE_NAME is running. Restarting..."
    systemctl restart "$SERVICE_NAME"
else
    echo "$SERVICE_NAME is not running. Starting..."
    systemctl start "$SERVICE_NAME"
fi

ps auxf | grep -v "grep" |grep $PROJ_NAME
