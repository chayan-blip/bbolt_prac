package bolt

// TODO: #define MBD_VALID /**< DB handle is valid, for me_dbflags */
// TODO: #define PERSISTENT_FLAGS (0xffff & ~(MBD_VALID))
// TODO: #define VALID_FLAGS (MDB_REVERSEKEY|MDB_DUPSORT|MDB_INTEGERKEY|MDB_DUPFIXED|MDB_INTEGERDUP|MBD_REVERSEDUP|MDB_CREATE)
// TODO: #define FREE_DBI 0
// TODO: #define MAIN_DBI 1

type Bucket interface {
}

type bucket struct {
	pad               int
	flags             int
	depth             int
	branchPageCount   int
	leafPageCount     int
	overflowPageCount int
	entryCount        int
	rootID            int
}
