package raids

// RAID interface defines the basic operations for all RAID levels
type RAID interface {
    Write(blockNum int, data []byte) error
    Read(blockNum int) ([]byte, error)
}
