package meta

const (
	Perfix     = "/rambo"
	ProxyNodes = Perfix + "/proxy_nodes"

	// users' config.

	UserInfo = Perfix + "/user_info"
	Password = "/password"
	DB       = "/db"
	Backends = "/backends" // DB's sub dir.   // Scheme   = "/scheme"   // DB's sub dir.
	Config   = "/config"   // DB's sub dir.
	Tables   = "/tables"   // DB's sub dir.

	// MysqlInfo is Mysql backends' config.
	// MysqlInfo is a dir.
	// each backend node's config record in a name which named
	// by the backend's name.
	MysqlInfo = Perfix + "/mysql_info"

	// partition type

	Hash  = "hash"
	Range = "range"
	CHash = "chash"

	// DDL etcd config

	DDLInfo   = Perfix + "/ddl_info"
	Masters   = "/masters"    // k: user id, v: node addr
	TaskQueue = "/task_queue" // k: user id, v: ddl excute plan
	Lock      = "/lock"       // k: user/db/table, v: "lock"
)

// etcd error code
const (
	NotFound     = 100
	NotEqual     = 101
	AlreadyExist = 105
)

// etcd action
const (
	Delete = "delete"
	Set    = "set"
	Create = "create"
	Update = "update"
	Expire = "expire"
	Get    = "get"
	CAS    = "compareAndSwap"
	CAD    = "compareAndDelete"
)
