# OS-HW7

# How to run:
Make sure to have all the files downloaded from here and in a folder named "OS-Project". Make sure Go is installed and initialized. Run "go run . raid 0" to simulate RAID0 and do that for raid 1, 4, and 5.

# Program Design:
It has the main.go file which handles which raid level and runs the benchmark. The benchmark.go file which reads and writes tests and prints out the timing stats. The disk.go which is an abstraction over the files backend. The raid.go file interfaces for all the RAID types. And lastly the raid.go files to simulate each RAID type. Each RAID simulates block based reads and writes and uses 5 virtual disks and then measures the write and read speeds across 100MB of data.

# Library:
This project uses the Go standard library such as os, fmt, time, and math/rand
