package main

import (
    "encoding/csv"
    "os"
    "fmt"
    "io"
    "strconv"
)

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        fmt.Println("First argument should be a path to a *.csv file.")
        os.Exit(1)
    }
    file_path := args[0]


    var in_file string = file_path
    f, err := os.Open(in_file)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    /*
     * Load up records, if stock
     *  [0] part number
     *  [1] manufacturer brand
     *  [2] quantity
     *  [3] warehouse location
     *  [4] lacuna
     */
    records := [][]string{}
    r := csv.NewReader(f) // default delimiter `,'
    for i := 0;; i++{
        record, err := r.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if record[2] <= strconv.Itoa(0) {
            for i := 0; i < len(record); i++ {
                fmt.Printf("[%d] - %s\n", i, record[i])
            }
            continue // no stock - ignore this record
        }

        records = append(records, record)
    }

    /*
     * TODO
     * Parse records
     */

    /*
     * TODO
     * Write parsed records
     */
}
