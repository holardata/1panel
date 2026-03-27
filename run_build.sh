# 1. 修改 Makefile 中的 GOOS=linux GOARCH=arm64
# 2. 修改 cmd/server/cmd/version.go 版本号
make build_on_arm64
rm -rf ./deploy/1panel/1panel && cp build/1panel ./deploy/1panel/
NOW=$(date +%Y%m%d_%H%M%S)

TAR_OPTIONS=(
    --exclude=".DS_Store"
    --exclude="._*"
    --no-xattrs
    --no-mac-metadata
)

cd ./deploy/ && tar "${TAR_OPTIONS[@]}" -zcvf 1panel_${NOW}.tar.gz 1panel && md5sum 1panel_${NOW}.tar.gz && cp 1panel_${NOW}.tar.gz ~/Desktop/

md5sum 1panel/1panel

# 将1panel_${NOW}.tar.gz 上传到 10.10.10.10 服务器的 /opt/1panel/ 目录下
# tar -zxvf 1panel_${NOW}.tar.gz -C /opt/
# 重启服务
# systemctl restart 1panel
