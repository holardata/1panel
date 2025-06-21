#!/bin/bash
# */5 * * * *  bash /opt/1panel/update_localapp.sh

git clone -b master https://gitee.com/jezzhu/1PanelAppStore /opt/1panel/resource/apps/local/appstore-localApps

cp -rf /opt/1panel/resource/apps/local/appstore-localApps/apps/* /opt/1panel/resource/apps/local/

rm -rf /opt/1panel/resource/apps/local/appstore-localApps
