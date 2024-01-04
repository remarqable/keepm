**Stop/Remove the container if running**  
docker stop keepm  
docker rm keepm  

**Remove the image**  
docker rmi -f `docker images --format="{{.Repository}} {{.ID}}" | grep "^keepm " | cut -d' ' -f2`  
 
(docker system prune --force --all)

**Build new image and verify**  
docker build --progress=plain -t keepm  -f keepm.dockerfile .  
docker images  

**Run container and verify**  
docker run --name keepm -h keepm-docker -d  -p 5432:5432 -p 8080:8080  -t keepm  
docker ps  
docker logs keepm  

**Connect with psql**  
export PGPASSWORD=keepm  
psql -h localhost -U keepm -d keepm  

**Optionally attach with the running container**  
docker exec -it keepm  /bin/bash  

Open http://localhost:8080/  


docker tag <image id> asim95/keepm:latest
  357  docker push asim95/keepm


