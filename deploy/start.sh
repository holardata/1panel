#!/bin/bash
# 部署说明: 
# 1. 把 conf目录下的 app.yaml 放到 /opt/1panel/conf 目录下
# 2. 二进制 1panel, 1panel.service 和 start.sh 都放到 /opt/1panel/ 目录下
PROJ_NAME="1panel"

TARGET_DIR="/lib/systemd/system"
SERVICE_NAME="${PROJ_NAME}.service"
SOURCE_PATH="/opt/1panel/${SERVICE_NAME}"
SOURCE_BIN="/opt/1panel/${PROJ_NAME}"

# 检查是否传入了stop参数
if [ "$1" = "stop" ]; then
    echo "Stopping $SERVICE_NAME..."
    systemctl stop "$SERVICE_NAME"
    echo "$SERVICE_NAME stopped."
    ps auxf | grep -v "grep" |grep $PROJ_NAME
    # ps auxf |grep java |grep $PROJ_NAME |grep -v grep | awk '{print $2}' |xargs kill -9
    exit 0
fi

if [ -f "$SOURCE_BIN" ]; then
    rm -rf /usr/local/bin/$PROJ_NAME && \
    mv "$SOURCE_BIN" /usr/local/bin/ && \
    chmod +x /usr/local/bin/$PROJ_NAME
    if [ $? -eq 0 ]; then
        echo "Successfully moved and set executable permission for $PROJ_NAME"
    else
        echo "Failed to move or set permission for $PROJ_NAME"
        exit 1
    fi
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
