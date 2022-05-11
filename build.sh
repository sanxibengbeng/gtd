 #build rpc/transform
 docker build -t trans:v2 -f rpc/transform/Dockerfile .

 #build shorturl
 docker build -t short:v1 -f Dockerfile .
 