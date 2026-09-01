# 1. 修改 Makefile 中的 GOOS=linux GOARCH=arm64
# 2. 修改 cmd/server/cmd/version.go 版本号
set -e
cd "$(dirname "$0")"
(cd ./cmd/server && go generate)
make build_on_amd64
rm -rf ./deploy/1panel/1panel && cp build/1panel ./deploy/1panel/
NOW=$(date +%Y%m%d_%H%M%S)

TAR_OPTIONS=(
    --exclude=".DS_Store"
    --exclude="._*"
    --no-xattrs
    --no-mac-metadata
)

cd ./deploy/ && tar "${TAR_OPTIONS[@]}" -zcvf 1panel_${NOW}.tar.gz 1panel

md5sum 1panel/1panel && cp 1panel/1panel ~/Desktop/
