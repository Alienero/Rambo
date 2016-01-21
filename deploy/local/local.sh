service_name='mysql_backend'
num=0
force=false
USAGE='Usage: $0 {init {num}|start|stop|remove [-f]|ports {num}}'

init_service() {
	docker-compose up -d
	docker-compose scale $service_name=$num
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
			exit 1
		fi
		ports
		;;
	*  ) 
		echo $USAGE
        	exit 1
        ;;
esac