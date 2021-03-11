
1. 安装docker

2. 下载consul镜像
docker pull consul

3. 运行镜像
docker run -d --name=consul -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0

    -server: 作为服务端
    --bootstrap: 取消节点选举，自己在内部默认为主节点
    -ui: 启动ui界面
    -client: 客户端监听地址

4. 编写自动注册服务的服务

5. 将服务添加到服务入口接口