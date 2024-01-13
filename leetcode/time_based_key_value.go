package leetcode

type TimeMap struct {
	m map[string]map[int]string
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string]map[int]string)}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	if _, ok := this.m[key]; !ok {
		this.m[key] = make(map[int]string)
	}
	this.m[key][timestamp] = value

}

func (this *TimeMap) Get(key string, timestamp int) string {

	if _, ok := this.m[key][timestamp]; !ok {
		return ""
	} else {
		return this.m[key][timestamp]
	}
}
