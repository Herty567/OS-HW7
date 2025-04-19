package main

import (
    "fmt"
    "os"
    "osproject/raids"
    "osproject/shared"
)

// Initializes and returns a list of 5 disk files for RAID simulation.
func initDisks() []*shared.Disk {
    var disks []*shared.Disk
    for i := 0; i < 5; i++ {
        d, err := shared.OpenDisk(fmt.Sprintf("disk%d.dat", i), i)
        if err != nil {
            panic(err)
        }
        disks = append(disks, d)
    }
    return disks
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run . [raid0|raid1|raid4|raid5]")
        return
    }

    disks := initDisks()
    var r raids.RAID

    // Selects the RAID implementation based on the command-line argument
    switch os.Args[1] {
    case "raid0":
        r = raids.NewRAID0(disks)
    case "raid1":
        r = raids.NewRAID1(disks)
    case "raid4":
        r = raids.NewRAID4(disks)
    case "raid5":
        r = raids.NewRAID5(disks)
    default:
        fmt.Println("Invalid RAID level. Choose: raid0, raid1, raid4, raid5")
        return
    }

    // Run the benchmark with the selected RAID configuration
    BenchmarkRAID(r, 100)
}
