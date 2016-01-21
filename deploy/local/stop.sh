echo "stop mysql ..."
docker-compose stop db0 db1 db2 db3 
echo "stopped."
echo "remove mysql container"
docker-compose rm -f db0 db1 db2 db3 
echo "removed."
