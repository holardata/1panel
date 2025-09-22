# 1. 修改 Makefile 中的 GOOS=linux GOARCH=arm64
# 2. 修改 cmd/server/cmd/version.go 版本号
make build_on_arm64
rm -rf ./deploy/1panel/1panel && cp build/1panel ./deploy/1panel/
NOW=$(date +%Y%m%d_%H%M%S)
cd ./deploy/ && tar -zcvf 1panel_${NOW}.tar.gz 1panel && md5sum 1panel_${NOW}.tar.gz && cp 1panel_${NOW}.tar.gz ~/Desktop/
# 将1panel_${NOW}.tar.gz 上传到 10.10.10.10 服务器的 /opt/1panel/ 目录下
# tar -zxvf 1panel_${NOW}.tar.gz -C /opt/
# 重启服务
# systemctl restart 1panel
