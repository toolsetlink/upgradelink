#!/bin/bash

# 定义项目根目录
PROJECT_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
echo "项目根目录: $PROJECT_ROOT"

# 定义日志目录和构建输出目录
LOG_DIR="$PROJECT_ROOT/logs"
BUILD_DIR="./"
mkdir -p "$LOG_DIR" "$BUILD_DIR" || { echo "创建目录失败"; exit 1; }

# 定义子项目信息
# 格式: "服务名:类型:源路径:构建命令:输出路径:运行命令"
SERVICES=(
  # Go 后端服务
  "upgradelink-admin-core-api:go:$PROJECT_ROOT/upgradelink-admin-core/server/api:CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/core-api core.go:$BUILD_DIR/core-api --config=./etc/core.yaml"
  "upgradelink-admin-core-rpc:go:$PROJECT_ROOT/upgradelink-admin-core/server/rpc:CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/core-rpc core.go:$BUILD_DIR/core-rpc --config=./etc/core.yaml"
  "upgradelink-admin-file:go:$PROJECT_ROOT/upgradelink-admin-file/server:CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/file fms.go:$BUILD_DIR/file --config=./etc/fms.yaml"
  "upgradelink-admin-message:go:$PROJECT_ROOT/upgradelink-admin-message/server:CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/message mcms.go:$BUILD_DIR/message --config=./etc/mcms.yaml"
  "upgradelink-admin-upgrade:go:$PROJECT_ROOT/upgradelink-admin-upgrade/server:CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/upgrade upgrade.go:$BUILD_DIR/upgrade --config=./etc/upgrade.yaml"

  # 前端项目
  "admin-ui:frontend:$PROJECT_ROOT/upgradelink-admin-ui:pnpm build:simple-admin-core"
)

# 记录所有启动的进程ID
PIDS=()

# 解析服务信息的函数
parse_service() {
  local service=$1
  IFS=':' read -r name type source_dir build_cmd output_path run_cmd <<< "$service"
}

# 检查文件是否需要更新（比较源文件和目标文件的修改时间）
needs_update() {
  local source_dir=$1
  local target=$2

  # 如果目标不存在，需要更新
  if [ ! -e "$target" ]; then
    return 0
  fi

  # 找到源目录中最新的文件
  local latest_source=$(find "$source_dir" -type f -printf "%T@ %p\n" 2>/dev/null | sort -n | tail -1 | cut -d' ' -f2-)

  # 比较最新源文件和目标文件的修改时间
  if [ -n "$latest_source" ] && [ "$latest_source" -nt "$target" ]; then
    return 0  # 需要更新
  else
    return 1  # 不需要更新
  fi
}

# 构建Go项目
build_go() {
  local name=$1
  local source_dir=$2
  local build_cmd=$3
  local output_file=$4

  echo "检查 $name 是否需要构建..."

  if needs_update "$source_dir" "$output_file"; then
    echo "构建 $name ..."
    if [ ! -d "$source_dir" ]; then
      echo "错误：目录 $source_dir 不存在"
      return 1
    fi
    cd "$source_dir" || { echo "进入目录 $source_dir 失败"; return 1; }

    # 安装依赖
    if [ -f "go.mod" ]; then
      echo "安装Go依赖: $source_dir"
      go mod download || { echo "Go依赖安装失败"; return 1; }
    fi

    # 执行构建命令，使用eval解析环境变量
    eval $build_cmd || { echo "$name 构建失败"; return 1; }
    echo "$name 构建成功"
  else
    echo "$name 无需构建，使用现有版本"
  fi

  return 0
}

# 构建前端项目
build_frontend() {
  local name=$1
  local source_dir=$2
  local build_cmd=$3
  local output_dir=$4

  echo "检查 $name 是否需要构建..."

  # 检查目录是否存在并进入
  if [ ! -d "$source_dir" ]; then
    echo "错误：目录 $source_dir 不存在"
    return 1
  fi
  cd "$source_dir" || { echo "进入目录 $source_dir 失败"; return 1; }

  # 安装依赖
  if [ -f "package.json" ]; then
    echo "安装/更新前端依赖: $source_dir"
    pnpm install || { echo "前端依赖安装失败"; return 1; }
  fi

  # 检查是否需要构建
  if [ ! -d "$output_dir" ] || needs_update "$source_dir/src" "$output_dir"; then
    echo "构建 $name ..."
    # 执行构建命令
    $build_cmd || { echo "$name 构建失败"; return 1; }
    echo "$name 构建成功"
  else
    echo "$name 无需构建，使用现有版本"
  fi

  return 0
}

# 构建所有服务
for service in "${SERVICES[@]}"; do
  parse_service "$service"

  case $type in
    go)
      build_go "$name" "$source_dir" "$build_cmd" "$output_path"
      ;;
    frontend)
      build_frontend "$name" "$source_dir" "$build_cmd" "$output_path"
      ;;
    *)
      echo "未知类型: $type 对于服务 $name"
      ;;
  esac
done
