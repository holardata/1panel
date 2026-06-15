#!/bin/bash

# 默认分支
BRANCH="master"

# 检查 /holarbox_sysinfo 文件是否存在
if [ -f /holarbox_sysinfo ]; then
    # 尝试解析 JSON 文件获取 version 字段
    # 使用 grep 和 sed 进行简单解析（避免依赖 jq）
    VERSION=$(grep -o '"version"[[:space:]]*:[[:space:]]*"[^"]*"' /holarbox_sysinfo | sed 's/.*"version"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/')
    
    # 如果 version 是 t4000 或 t5000，则切换到对应分支
    if [ "$VERSION" = "t4000" ]; then
        BRANCH="th-4k"
        echo "检测到硬件版本: $VERSION，使用分支: $BRANCH"
    elif [ "$VERSION" = "t5000" ]; then
        BRANCH="th-master"
        echo "检测到硬件版本: $VERSION，使用分支: $BRANCH"
    else
        echo "检测到硬件版本: $VERSION，使用默认分支: $BRANCH"
    fi
else
    echo "未找到 /holarbox_sysinfo 文件，使用默认分支: $BRANCH"
fi

rm -rf /opt/1panel/resource/apps/appstore-localApps

git clone --depth=1 -b $BRANCH https://jezzhu:be1d663aeb9f1738016099a92827d4e6@gitee.com/jezzhu/1PanelAppStore /opt/1panel/resource/apps/appstore-localApps \
  && rm -rf /opt/1panel/resource/apps/local/* \
  && cp -rf /opt/1panel/resource/apps/appstore-localApps/apps/* /opt/1panel/resource/apps/local/ \
  && rm -rf /opt/1panel/resource/apps/appstore-localApps \
  || echo "git 操作或后续步骤失败，已停止，未清理本地"
