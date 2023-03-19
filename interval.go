package main

import (
	"sort";
	"fmt"
)

type Interval struct {
    start int
    end int
}

func mergeIntervals(intervals []Interval) []Interval {
    if len(intervals) == 0 {
        return []Interval{}
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })
    merged := []Interval{intervals[0]}

    for i := 1; i < len(intervals); i++ {
        if intervals[i].start <= merged[len(merged)-1].end {
            merged[len(merged)-1].end = max(merged[len(merged)-1].end, intervals[i].end)
        } else {
            merged = append(merged, intervals[i])
        }
    }

    return merged
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func main(){
	intervals := []Interval{
    {1, 10},
    {2, 9},
    {5, 11},
    {3, 4},
}
	merger :=(mergeIntervals(intervals))

	for _, interval := range merger {
    fmt.Printf("(%d, %d) ", interval.start, interval.end)
}
}
