# ssh 执行多行命令
user=root
ip=10.202.40.94
ssh $user@"$ip" << EOF

[ -f "$pid_file" ] && cat $pid_file| xargs kill -9 && rm -rf "$pid_file"
exit 0
EOF