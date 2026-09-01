GOBUILD="go build"
GOCLEAN="go clean"
# GOARCH=${shell go env GOARCH}
# GOOS=${shell go env GOOS}
GOOS="darwin"
GOARCH="arm64"

BASE_PATH=$(pwd)
#echo "$BASE_PATH"
BUILD_PATH="${BASE_PATH}/build"
WEB_PATH="${BASE_PATH}/frontend"
SERVER_PATH="${BASE_PATH}/backend"
MAIN="${BASE_PATH}/cmd/server/main.go"
APP_NAME="1panel"
ASSERT_PATH="${BASE_PATH}/cmd/server/web/assets"


cd ${SERVER_PATH} \
&& GOOS=darwin GOARCH=amd64 ${GOBUILD} -trimpath -ldflags '-s -w'  -o ${BUILD_PATH}/${APP_NAME} ${MAIN}