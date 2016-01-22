service_name='mysql_backend'
num=0
force=false
etcd=""
USAGE='Usage: {init {num} {etcd_host}|start|stop|remove [-f]|ports {num}}'

init_service() {
	docker-compose up -d
	docker-compose scale $service_name=$num
	# wait mysql set up.
	# sleep 15
	# write confif to etcd.
	for (( i = 1; i <= $num; i++ )); do
		host=$(docker-compose port --index=$i $service_name 3306)
		go run local.go -user=root -password=123456 -name=db_$i -host=$host -etcd_host=$etcd
	done
}

start() {
	docker-compose start $service_name
}

stop() {
	docker-compose stop $service_name
}

remove() {
	if $force; then
		stop
	fi
	docker-compose rm $service_name
}

ports() {
	echo "Print bakcends' host"
	for (( i = 1; i <= $num; i++ )); do
		docker-compose port --index=$i $service_name 3306
	done
}

case $3 in
	"")
		;;
	*)
		etcd=$3
		;;
esac

case $2 in
	-f)
		force=true	
		;;
	"")
		num=0
		;;
	*) 
		num=$2
        ;;
esac

case $1 in
	init)
		if [[ $num == 0 ]]; then
			echo "num is 0"
			echo $USAGE
			exit 1
		fi
		if [[ $etcd == "" ]]; then
			echo "etcd host is null!"
			echo $USAGE
			exit 1
		fi
		init_service	
		;;
	start)
		start
		;;
	stop)
		stop
		;;
	remove)
		remove
		;;
	ports)	
		if [[ $num == 0 ]]; then
			echo "num is 0"
			echo $USAGE
			exit 1
		fi
		ports
		;;
	*) 
		echo $USAGE
        	exit 1
        ;;
esac