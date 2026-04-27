package reporter

import (
	"fmt"
	"sort"

	"go-wg/internal/entity"
)

func Print(sites []entity.Site) {
	sorted := make([]entity.Site, len(sites))
	copy(sorted, sites)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Latency < sorted[j].Latency
	})

	for _, s := range sorted {
		if s.OK() {
			fmt.Printf("Success calling service [%s] [%.2fs]\n", s.URL, s.Latency.Seconds())
			continue
		}
		if s.Err != nil {
			fmt.Printf("Failure calling service [%s] [%.2fs] error=%v\n", s.URL, s.Latency.Seconds(), s.Err)
			continue
		}
		fmt.Printf("Failure calling service [%s] [%.2fs] status=%d\n", s.URL, s.Latency.Seconds(), s.Status)
	}
}
