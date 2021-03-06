package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := CreateLRUCache(5)
	cache.AddElement(1, "1")
	cache.AddElement(2, "1")
	cache.AddElement(3, "1")
	cache.AddElement(4, "1")
	cache.AddElement(5, "1")
	if _, ok := cache.GetElement(1); ok {
		fmt.Println("命中 1")
	} else {
		fmt.Println("没命中 1")
	}
	cache.AddElement(6, "6")
	if _, ok := cache.GetElement(1); ok {
		fmt.Println("命中 1")
	} else {
		fmt.Println("没命中 1")
	}
	if _, ok := cache.GetElement(2); ok {
		fmt.Println("命中 2")
	} else {
		fmt.Println("没命中 2")
	}
	cache.AddElement(2, "2")
	cache.AddElement(7, "7")
	if _, ok := cache.GetElement(3); ok {
		fmt.Println("命中 3")
	} else {
		fmt.Println("没命中 3")
	}
	if _, ok := cache.GetElement(4); ok {
		fmt.Println("命中 4")
	} else {
		fmt.Println("没命中 4")
	}
}
