package mysql

// MySQL type informations.
const (
	TypeDecimal byte = iota
	TypeTiny
	TypeShort
	TypeLong
	TypeFloat
	TypeDouble
	TypeNull
	TypeTimestamp
	TypeLonglong
	TypeInt24
	TypeDate
	TypeDuration /* Original name was TypeTime, renamed to Duration to resolve the conflict with Go type Time.*/
	TypeDatetime
	TypeYear
	TypeNewDate
	TypeVarchar
	TypeBit
)

// MySQL type informations.
const (
	TypeNewDecimal byte = iota + 0xf6
	TypeEnum
	TypeSet
	TypeTinyBlob
	TypeMediumBlob
	TypeLongBlob
	TypeBlob
	TypeVarString
	TypeString
	TypeGeometry
)
