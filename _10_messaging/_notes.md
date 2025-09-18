# start docker 
docker compose up -d

# create topic with 3 partitions:
docker exec -it <kafka-container-id> \
  kafka-topics --create \
  --topic orders \
  --bootstrap-server localhost:9092 \
  --partitions 3 \
  --replication-factor 1

# verify
docker exec -it <kafka-container-id> \
  kafka-topics --describe \
  --topic orders \
  --bootstrap-server localhost:9092

# close docker 
docker compose down