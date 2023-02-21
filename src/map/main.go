package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//counts()
	//countFiles()
	//readFile2Map("/Users/ssg/Documents/sources/ceshiwd/querySettlementDetail.sh")
	json2Struct1()
}

func readFile2Map(fileName string) {
	if "" == fileName {
		return
	}
	var countMap = make(map[string]int)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range strings.Split(string(data), "\n") {
			countMap[line]++
		}
	}

	for k, v := range countMap {
		fmt.Println("name=", k, ",count=", v)
	}
	//
	fmt.Println("------------转json-------------")
	bt, _ := json.Marshal(countMap)

	jsonStr := string(bt)
	fmt.Println(jsonStr)

	fmt.Println("------------json转map-------------")
	map2 := make(map[string]interface{})
	error := json.Unmarshal([]byte(jsonStr), &map2)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(map2)
	fmt.Println(map2["美宁"])
	fmt.Println(map2["宁夏"])

}

func json2Struct1() {
	var user = UserInfo{"zhangsan", "重庆", "13883482726"}
	fmt.Println(user)

	fmt.Println("--------------user to json--------------")
	bytes, _ := json.Marshal(user)
	jsonStr := string(bytes)
	fmt.Println(jsonStr)

	fmt.Println("--------------json to user--------------")
	userInfo := new(UserInfo)
	error := json2Struct(jsonStr, userInfo)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(userInfo)
		fmt.Println(userInfo.Name, userInfo.Mobile)
	}

}

func json2Struct(str string, v interface{}) error {
	if str == "" {
		return errors.New("str is null")
	}
	return json.Unmarshal([]byte(str), &v)
}

type UserInfo struct {
	Name   string
	Addr   string
	Mobile string
}

func countFiles() {
	var counts = make(map[string]int)
	files := os.Args[1:]

	fmt.Println(files)
	fmt.Println(len(files))
	if len(files) == 0 {
		countFile(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Println("arg:", arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println("read file error", err)
			} else {
				countFile(f, counts)
			}
		}
	}

	fmt.Println(counts)
}

func countFile(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func counts() {
	var countMap = make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		countMap[input.Text()]++ // 等价于 line := input.Text(); countMap[line] = countMap[line] + 1
	}

	for key, val := range countMap {
		if val > 1 {
			fmt.Println("line=", key, ",val=", val)
		}
	}

	fmt.Println(countMap)
}

func testInit() {
	var a map[int]string
	// 申请内存空间，map存放元素无序
	a = make(map[int]string, 10)
	a[1001] = "张三"
	a[1002] = "李四"
	a[1009] = "王五"

	fmt.Println(a)

	//初始化2
	b := make(map[int]string)
	b[1002] = "李四"
	b[1009] = "王五"
	fmt.Println(b)

	//初始化3
	c := map[int]string{
		20221: "yangsong",
		20222: "刘彻",
	}
	c[1001] = "美宁"
	fmt.Println(c)
	mapOperate(c)

	a2, flag := c[20221]
	fmt.Println(a2)
	fmt.Println(flag)

	mapTest()

}

func mapTest() {
	fmt.Println("map operate function ---- map[string][map]")
	gradeMap := make(map[string]map[int]string)
	gradeMap["class1"] = make(map[int]string)
	gradeMap["class1"][10001] = "xiao wang"
	gradeMap["class1"][10002] = "xiao li"
	gradeMap["class1"][10003] = "xiao zhang"

	gradeMap["class2"] = make(map[int]string)
	gradeMap["class2"][20001] = "xiao wang2"
	gradeMap["class2"][20002] = "xiao li2"
	gradeMap["class2"][20003] = "xiao zhang2"

	for k1, v1 := range gradeMap {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("key=%v,val=%v\t\n", k2, v2)
		}

		fmt.Println()
	}

}

func mapOperate(channelMap map[int]string) {
	fmt.Println("map operate function ---- delete key")
	delete(channelMap, 20222)
	fmt.Println(channelMap)

	for k, v := range channelMap {
		fmt.Printf("key=%v,val=%v\t", k, v)
	}
}
