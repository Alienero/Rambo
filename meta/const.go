package meta

const (
	Perfix = "/rambo"

	// users' config.

	UserInfo = Perfix + "/user_info"
	Password = "/password"
	DB       = "/db"
	Backends = "/Backends" // DB's sub dir.
	Scheme   = "/scheme"   // DB's sub dir.
	Config   = "/config"   // DB's sub dir.
	Tables   = "/tables"   // DB's sub dir.

	// MysqlInfo is Mysql backends' config.
	// MysqlInfo is a dir.
	// each backend node's config record in a name which named
	// by the backend's name.
	MysqlInfo = Perfix + "/mysql_info"

	// partition type.
	Hash  = "hash"
	Range = "range"
	CHash = "chash"
)

// Error code
const (
	NotFoud      = 100
	NotEqual     = 101
	AlreadyExist = 105
)
