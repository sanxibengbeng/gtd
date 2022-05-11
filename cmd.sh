docker network create --subnet=192.167.0.0/16 shorturlnet

docker run --net shorturlnet --ip 192.167.0.167   -d -p 7379:6379 redis --requirepass "aiya"

docker run -it --net shorturlnet --ip 192.167.0.168 -p 3379:2379 -p 3380:2380 --env ALLOW_NONE_AUTHENTICATION=yes  -d bitnami/etcd

#docker run rpc/transform use link to connet fail
docker run -d -v /Users/baidu/projects/shorturl/trans_etc/:/app/etc --rm trans:v3   

#use docker subnet to connect
docker run -it  --net bridge -v /Users/baidu/projects/shorturl/trans_etc/:/app/etc --rm trans:v3   


 # 生成 DockerFile
 goctl docker --go shorturl.go --base alpine 