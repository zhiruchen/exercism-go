package letter

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(letterList []string) FreqMap {
	var chs = make(chan FreqMap, len(letterList))
	var mm = make(FreqMap)

	go func(ll []string, ch chan FreqMap) {
		for _, s := range ll {
			ch <- Frequency(s)
		}
		close(ch)
	}(letterList, chs)

	for c := range chs {
		for k, v := range c {
			vv, ok := mm[k]
			if ok {
				mm[k] = vv + v
			} else {
				mm[k] = v
			}
		}
	}

	return mm
}
