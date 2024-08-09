#!/bin/bash

# 输出调试信息
echo "开始执行脚本" >&2

# 添加 guest02 用户
sudo adduser guest02 --disabled-password --gecos "" >&2

# 修改 guest02 用户密码
echo 'guest02:XIONGHUI123' | sudo chpasswd >&2

# 将 guest02 用户添加到 wheel 组
sudo usermod -aG wheel guest02 >&2

# 输出完成信息
echo "脚本执行完毕" >&2
