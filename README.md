# Trend-ring

## 介绍


## docker

运行容器
```bash
docker run -it -v /root/project/:/root/project/ --name home centos_go:7.9.2009 /bin/bash
 
--name 为容器指定一个名称
-v 将宿主机目录挂载到容器内  ；格式：宿主机目录:容器目录
-e 设置环境变量，该环境变量看覆盖Dockerfile中的ENV环境变量
-p 需要手动指定一个或多个映射端口号，格式为：主机(宿主)端口:容器端口 -p 80:8080
-P Docker会随机映射一个 49000~49900 的端口到内部容器开放的网络端口
-it 其中，-i以交互模式运行容器 ；-t为容器重新分配一个伪输入终端
-d 后台运行容器，并返回容器ID （没有此参数，容器在前台窗口运行，窗口关，随之关）
 
-it 开启新的终端，以交互方式进入容器内部（尾部要指定/bin/bash方式）；和-d一起使用后，将不会进入容器内部。
 
注：该写法相当于把docker run命令和docker exec -it 命令合并的效果（区别是：前者run的同时，进入了容器内部，如果exit退出容器，容器状态(通过docker ps -a查看)立马会变成Exited状态；后者就不一样了，以此方式进入容器，然后exit退出容器，不会主动影响容器的原有状态）。
 
 
```

update.sh
```zsh
docker pull origin 分支名
go build -o 打包二进制文件名
kill -9 "$(pgrep -f 打包二进制文件名)" #杀死之前的进程
chmod +x /二进制文件全路径 # 给目标文件加权
nohup /二进制文件全路径 -c /项目配置文件全路径 > start.log  2>&1 & # 使用nohup后台运行
```
