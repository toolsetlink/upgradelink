#!/bin/bash
set -euo pipefail

##############################################################################
# 1. 基础路径配置
##############################################################################
PROJECT_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
echo "=== 项目根目录: $PROJECT_ROOT ==="

# 日志目录
LOG_DIR="$PROJECT_ROOT/logs"
mkdir -p "$LOG_DIR" || { echo "错误: 创建日志目录失败！"; exit 1; }

##############################################################################
# 2. 服务配置（配置多API代理）
##############################################################################

SERVICES=(
  "upgradelink-admin-core-rpc:go:$PROJECT_ROOT/upgradelink-admin-core/server/rpc:./core-rpc -f=./etc/core.yaml:9101"
  "upgradelink-admin-file:go:$PROJECT_ROOT/upgradelink-admin-file/server:./file -f=./etc/fms.yaml:9102"
  "upgradelink-admin-message:go:$PROJECT_ROOT/upgradelink-admin-message/server:./message -f=./etc/mcms.yaml:9106"
  "upgradelink-admin-upgrade:go:$PROJECT_ROOT/upgradelink-admin-upgrade/server:./upgrade -f=./etc/upgrade.yaml:8181"
  "upgradelink-admin-core-api:go:$PROJECT_ROOT/upgradelink-admin-core/server/api:./core-api -f=./etc/core.yaml:9100"

#  # 前端服务，使用检测到的前端服务器命令
#  "admin-ui:frontend:$PROJECT_ROOT/upgradelink-admin-ui/apps/simple-admin-core:$FRONTEND_SERVER_CMD --directory dist --port 8080 \
#    --proxy /sys-api/=$CORE_API_URL \
#    --proxy /fms-api/=$FILE_SERVICE_URL \
#    --proxy /upgrade/=$UPGRADE_SERVICE_URL"
)


##############################################################################
# 3. 核心工具函数
##############################################################################

parse_service() {
  local service=$1
  IFS=':' read -r name type work_dir run_cmd port <<< "$service"
}

# 检查端口是否被占用
is_port_used() {
  local port=$1
  if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1; then
    return 0
  else
    return 1
  fi
}

# 通过端口停止服务
stop_service_by_port() {
  local port=$1
  local pid=$(lsof -Pi :$port -sTCP:LISTEN -t)

  if [ -n "$pid" ]; then
    echo "停止使用端口 $port 的进程 (PID: $pid)..."
    if kill "$pid"; then
      # 等待进程完全终止
      local wait_count=0
      while is_port_used "$port" && [ $wait_count -lt 10 ]; do
        sleep 1
        ((wait_count++))
      done

      if is_port_used "$port"; then
        echo "警告: 端口 $port 上的进程未能正常停止，尝试强制终止..."
        kill -9 "$pid" >/dev/null 2>&1
      else
        echo "端口 $port 上的进程已成功停止"
      fi
    else
      echo "警告: 停止端口 $port 上的进程失败"
    fi
  fi
}

# 检查端口是否已处理过（兼容旧版本Bash）
is_port_processed() {
  local port=$1
  local ports=$2
  # 检查端口是否在已处理列表中
  if [[ " $ports " =~ " $port " ]]; then
    return 0
  else
    return 1
  fi
}

# 停止所有服务（使用兼容旧版本Bash的方式）
stop_all_services() {
  echo -e "\n=== 开始停止所有已运行的服务 ==="

  # 使用字符串存储已处理的端口，用空格分隔
  local processed_ports=""

  for service in "${SERVICES[@]}"; do
    parse_service "$service"
    # 检查端口是否已经处理过
    if ! is_port_processed "$port" "$processed_ports"; then
      # 添加到已处理列表
      processed_ports="$processed_ports $port"
      stop_service_by_port "$port"
    fi
  done

  echo "=== 服务停止操作完成 ==="
}

# 启动Go服务
start_go_service() {
  local name=$1
  local work_dir=$2
  local run_cmd=$3
  local port=$4
  local log_file="$LOG_DIR/$name.log"

  if is_port_used "$port"; then
    echo "警告: $name 端口 $port 已被占用，服务可能已在运行"
    return 0
  fi

  echo "启动 $name (端口: $port) ..."
  if ! cd "$work_dir"; then
    echo "错误: 进入工作目录 $work_dir 失败"
    return 1
  fi

  # 检查二进制文件是否存在
  local binary_name=$(echo "$run_cmd" | awk '{print $1}')
  if [ ! -f "$binary_name" ]; then
    echo "错误: $name 启动失败：未找到二进制文件 $binary_name"
    return 1
  fi

  # 清空日志文件并启动服务
  rm -f "$log_file"
  nohup $run_cmd > "$log_file" 2>&1 &

  # 等待服务启动
  echo "等待 $name 启动..."
  local max_wait=10
  local wait_count=0

  while [ $wait_count -lt $max_wait ]; do
    if is_port_used "$port"; then
      echo "$name 启动成功 (端口: $port)"
      echo "日志文件: $log_file"
      return 0
    fi
    sleep 1
    ((wait_count++))
  done

  echo "错误: $name 启动超时，请检查日志: $log_file"
  return 1
}

# 启动前端服务
start_frontend_service() {
  local name=$1
  local work_dir=$2
  local run_cmd=$3
  local port=$4
  local log_file="$LOG_DIR/$name.log"

  if is_port_used "$port"; then
    echo "警告: $name 端口 $port 已被占用，服务可能已在运行"
    return 0
  fi

  echo "启动 $name (端口: $port) ..."
  if ! cd "$work_dir"; then
    echo "错误: 进入工作目录 $work_dir 失败"
    return 1
  fi

  # 检查dist目录是否存在
  if [ ! -d "dist" ]; then
    echo "错误: $name 启动失败：未找到dist目录"
    return 1
  fi

  # 清空日志文件并启动服务
  rm -f "$log_file"
  nohup $run_cmd > "$log_file" 2>&1 &

  # 等待服务启动
  echo "等待 $name 启动..."
  local max_wait=15
  local wait_count=0

  while [ $wait_count -lt $max_wait ]; do
    if is_port_used "$port"; then
      echo "$name 启动成功 (端口: $port)"
      echo "日志文件: $log_file"
      echo "访问地址: http://localhost:$port"
      return 0
    fi
    sleep 1
    ((wait_count++))
  done

  echo "错误: $name 启动超时，请检查日志: $log_file"
  return 1
}

##############################################################################
# 4. 显示服务端口信息
##############################################################################
show_service_ports() {
  echo -e "\n=== 服务端口信息 ==="
  for service in "${SERVICES[@]}"; do
    parse_service "$service"
    local status="未运行"
    if is_port_used "$port"; then
      status="运行中"
    fi
    echo "  $name: 端口 $port - $status"
  done
}

##############################################################################
# 5. 主流程：先停止所有服务，再启动服务
##############################################################################

# 先停止所有已运行的服务
stop_all_services

echo -e "\n=== 开始按顺序启动所有预编译服务 ==="

# 显示当前端口状态
show_service_ports

# 启动服务
previous_service_success=0

for service in "${SERVICES[@]}"; do
  # 检查前一个服务是否启动成功，如果失败则退出
  if [ $previous_service_success -ne 0 ]; then
    echo -e "\n由于上一个服务启动失败，终止后续服务启动"
    exit 1
  fi

  parse_service "$service"
  echo -e "\n------------------------------"
  echo "准备启动: $name"

  case $type in
    go)
      if start_go_service "$name" "$work_dir" "$run_cmd" "$port"; then
        previous_service_success=0
      else
        previous_service_success=1
      fi
      ;;
    frontend)
      if start_frontend_service "$name" "$work_dir" "$run_cmd" "$port"; then
        previous_service_success=0
      else
        previous_service_success=1
      fi
      ;;
    *)
      echo "错误: 未知服务类型: $type"
      previous_service_success=1
      ;;
  esac
done

##############################################################################
# 6. 启动完成提示
##############################################################################
echo -e "\n------------------------------"
echo "=== 服务启动流程完成 ==="
echo -e "\n最终服务状态:"
show_service_ports

echo -e "\n管理命令:"
echo "查看日志: tail -f $LOG_DIR/<服务名>.log"
echo "停止服务: 使用 pkill -f \"服务进程名\" 或 kill 对应进程"
echo "查看进程: ps aux | grep -E \"(core-rpc|file|message|core-api|upgrade|http.server)\" | grep -v grep"
