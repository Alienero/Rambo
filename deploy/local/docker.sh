# remove mysql backends
etcdctl --endpoint "http://192.168.99.100:4001" rm --recursive /rambo/ddl_info                                           
etcdctl --endpoint "http://192.168.99.100:4001" rm --recursive /rambo/mysql_info
etcdctl --endpoint "http://192.168.99.100:4001" rm --recursive /rambo/user_info/user/db