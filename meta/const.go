package meta

const (
	Perfix = "/rambo"

	// users' config.
	UserInfo      = Perfix + "/user_info"
	Scheme        = "/scheme"
	Password      = "/password"
	DB            = "/db"
	Tables        = "/tables" // DB's sub dir.
	ChildBackends = "/child_backends"

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
