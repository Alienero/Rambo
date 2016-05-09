package server

import "hash/crc32"

func FindForKey(key string, shardNum int) int {
	return int(crc32.ChecksumIEEE([]byte(key)) % uint32(shardNum))
}
