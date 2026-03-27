#!/bin/bash
# */5 * * * *  bash /opt/1panel/update_localapp.sh

rm -rf /opt/1panel/resource/apps/appstore-localApps

git clone --depth=1 -b master https://gitee.com/jezzhu/1PanelAppStore /opt/1panel/resource/apps/appstore-localApps \
  && rm -rf /opt/1panel/resource/apps/local/* \
  && cp -rf /opt/1panel/resource/apps/appstore-localApps/apps/* /opt/1panel/resource/apps/local/ \
  && rm -rf /opt/1panel/resource/apps/appstore-localApps \
  || echo "git 操作或后续步骤失败，已停止，未清理本地"