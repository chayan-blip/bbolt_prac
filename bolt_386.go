package bolt

// maxMapSize represents the largest map size supported by Bolt.
const maxMapSize = 0x7FFFFFFF // 2GB

// maxAllocSize is the sie used when creating pointers.
const maxAllocSize = 0xFFFFFFF

// Are unaligned oad/stores broken on this arch?
var brokenUnaligned = false
