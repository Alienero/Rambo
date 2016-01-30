package meta

const (
	Perfix = "/rambo"

	// all database.
	Databases = Perfix + "/databases" // style: [user name]_[database name]

	// users' config.

	UserInfo = Perfix + "/user_info"
	Backends = "/Backends" // UserInfo's sud dir.
	Scheme   = "/scheme"
	Password = "/password"
	DB       = "/db"
	Config   = "/config" // DB's sub dir.
	Tables   = "/tables" // DB's sub dir.

	// Mysql backends' config.
	// MysqlInfo is a dir.
	// each backend node's config record in a name which named
	// by the backend's name.
	MysqlInfo = Perfix + "/mysql_info"

	// partition type.
	Hash  = "hash"
	Range = "range"
)

// Error code
const (
	NotFoud      = 100
	NotEqual     = 101
	AlreadyExist = 105
)
