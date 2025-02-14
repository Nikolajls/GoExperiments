package iterations_range_fun

import (
	"errors"
	"fmt"
)

func SumOverSliceInt(numbers []int) int {
	sum := 0
	for index, value := range numbers {
		s := fmt.Sprintf("Index:%v\tValue:%v", index, value)
		fmt.Println(s)
		sum += value
	}

	return sum
}

func AddNumbersFromStartToEnd(startNo int, endNo int) int {
	sum := 0
	for i := startNo; i <= endNo; i++ {
		sum += i
	}
	return sum
}

func GetMapDictionaryContainKeyValueIteration(dict map[string]string, key string) (string, error) {
	for kvKey, kvValue := range dict {
		if kvKey == key {
			return kvValue, nil
		}
	}
	return "", errors.New("key not found")
}

func GetMapDictionaryContainKeyValueLookup(dict map[string]string, key string) (string, error) {
	if valueFromMap, keyPresentInMap := dict[key]; keyPresentInMap {
		return valueFromMap, nil
	}
	return "", errors.New("key not found")
}
