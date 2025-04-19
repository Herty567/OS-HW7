package main

import (
    "fmt"
    "os"
    "osproject/raids"
    "osproject/shared"
)

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
        fmt.Println("Usage: go run main.go [raid0|raid1|raid4|raid5]")
        return
    }

    disks := initDisks()
    var r raids.RAID

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

    BenchmarkRAID(r, 100)
}
