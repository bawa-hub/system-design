# start docker 
docker compose up -d

# start/stop only single service
docker compose stop redis
docker compose start redis

# redis cli 
docker exec -it redis redis-cli

# close docker 
docker compose down