type Tuple struct {
    value string
    timestamp int
}

type TimeMap struct {
    data map[string][]Tuple
}

func Constructor() TimeMap {
    tm := TimeMap{}
    tm.data = make(map[string][]Tuple)
    return tm
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    this.data[key] = append(this.data[key], Tuple{value, timestamp})
    fmt.Println(this.data[key])
}


func binSearch(data []Tuple, i int, j int, timestamp int) string {

    if i > j {
        return ""
    }

    mid := int((i+j)/2)

    if timestamp == data[mid].timestamp {
        return data[mid].value
    } else if timestamp < data[mid].timestamp {
        return binSearch(data, i, mid-1, timestamp)
    } else {
        if mid == len(data)-1 || timestamp < data[mid+1].timestamp {
            return data[mid].value
        }
        return binSearch(data, mid+1, j, timestamp)
    }

}


func (this *TimeMap) Get(key string, timestamp int) string {
    
    dataSlice := this.data[key]

    ans := binSearch(dataSlice, 0, len(dataSlice)-1, timestamp)
    
    return ans
}
