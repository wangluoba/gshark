#!/bin/bash
echo 'guest01:XIONGHUI123' | sudo chpasswd
sudo usermod -aG wheel guest01
