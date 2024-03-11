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
docker run --name keepm -h keepm-docker -d  -p 5432:5432 -p 80:80  -t keepm  
docker ps  
docker logs keepm  

**Connect with psql**  
export PGPASSWORD=keepm  
psql -h localhost -U keepm -d keepmdb  

**Optionally attach with the running container**  
docker exec -it keepm  /bin/bash  

Open http://localhost/  


docker tag <image id> remarqable/keepm:latest
docker push remarqable/keepm


