steps to run mysql server from codebase:


delete data dir:
cd ~/CS/mysql-server/build

rm -rf data/*
bin/mysqld --initialize-insecure \
  --basedir=$(pwd) \
  --datadir=$(pwd)/data

rebuild:
cd ~/mysql-server/build
make -j$(sysctl -n hw.ncpu) mysqld

run server:
cd ~/mysql-dev/build
bin/mysqld \
  --basedir=$(pwd) \
  --datadir=$(pwd)/data \
  --socket=$(pwd)/mysql-debug.sock \
  --port=3307 \
  --pid-file=$(pwd)/mysql-debug.pid

client:
~/mysql-dev/bin/mysql -uroot -h127.0.0.1 -P3307