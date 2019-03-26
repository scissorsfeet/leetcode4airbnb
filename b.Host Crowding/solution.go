package b_Host_Crowding

import (
	"strings"
)

func hostCrowding(pageNum int, results []string) []string {
	resultsNum := len(results)
	resultsSlice := make([][]string, 0, resultsNum)
	hostid2Row := make(map[string][][]string)
	for _, result := range results {
		row := strings.Split(result, ",")
		hostid := row[0]
		if hostid2Row[hostid] == nil {
			hostid2Row[hostid] = make([][]string, 0, 1)
		}
		hostid2Row[hostid] = append(hostid2Row[hostid], row)
		resultsSlice = append(resultsSlice, row)
	}

	res := make([]string, 0, resultsNum)
	remainResultsSlice := make([][]string, 0)
	hostidMap := make(map[string]bool, len(hostid2Row))
	initHostidMap(hostid2Row, hostidMap)
	count := 0
	m := make(map[string]bool)
	for len(resultsSlice) > 0 || len(remainResultsSlice) > 0 {
		if ((len(hostidMap) == 0 && count < pageNum) || (len(resultsSlice) == 0)) && len(remainResultsSlice) > 0 {
			for len(remainResultsSlice) > 0 {
				result := remainResultsSlice[0]
				remainResultsSlice = remainResultsSlice[1:]
				res = append(res, strings.Join(result, ","))
				count++
				if count == pageNum {
					initHostidMap(hostid2Row, hostidMap)
					count = 0
					res = append(res, "")
					m = make(map[string]bool)
					resultsSlice = append(remainResultsSlice, resultsSlice...)
					remainResultsSlice = [][]string{}
					break
				}
			}
		}

		if len(resultsSlice) > 0 {
			row := resultsSlice[0]
			resultsSlice = resultsSlice[1:]
			hostid := row[0]
			if _, ok := m[hostid]; !ok {
				count++
				m[hostid] = true
				delete(hostidMap, hostid)
				res = append(res, strings.Join(row, ","))
				if count == pageNum {
					initHostidMap(hostid2Row, hostidMap)
					count = 0
					res = append(res, "")
					m = make(map[string]bool)
					resultsSlice = append(remainResultsSlice, resultsSlice...)
					remainResultsSlice = [][]string{}
				}
			} else {
				remainResultsSlice = append(remainResultsSlice, row)
			}
		}
	}

	return res
}

func initHostidMap(hostid2Row map[string][][]string, hostidMap map[string]bool) {
	for hostid := range hostid2Row {
		hostidMap[hostid] = true
	}
}
