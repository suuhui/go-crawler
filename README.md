# go-crawler

## elasticsearch安装运行
```shell script
docker pull elasticsearch:7.6.2
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2
```