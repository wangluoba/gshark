#!/bin/bash

# 添加guest02用户
sudo adduser guest02 --disabled-password --gecos ""

# 修改guest02用户密码
echo 'guest02:XIONGHUI123' | sudo chpasswd

# 将guest02用户添加到wheel组
sudo usermod -aG wheel guest02
