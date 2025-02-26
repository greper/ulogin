package utility

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
	"strings"
)

func ConvertList(sourceList interface{}, targetList interface{}, mapping ...map[string]string) {
	if err := gconv.Structs(sourceList, targetList, mapping...); err != nil {
		panic(err)
	}
}

func Convert(source interface{}, target interface{}) {
	if err := gconv.Struct(source, target); err != nil {
		if strings.HasPrefix(err.Error(), "convert params from") && strings.HasSuffix(err.Error(), "to \"map[string]interface{}\" failed") {
			return
		}
		glog.Panic(context.Background(), "convert error", err)
	}
}
func ArrayStringToInt(strArr []string) []int64 {
	res := make([]int64, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.ParseInt(val, 10, 64)
	}

	return res
}
func ArrayInt64ToString(intArr []int64) []string {
	res := make([]string, len(intArr))

	for index, val := range intArr {
		res[index] = strconv.FormatInt(val, 10)
	}

	return res
}

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}
