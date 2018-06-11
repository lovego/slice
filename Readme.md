# slice
a union find sets data structure based on map and list in golang.

[![Build Status](https://travis-ci.org/lovego/slice.svg?branch=master)](https://travis-ci.org/lovego/slice)
[![Coverage Status](https://coveralls.io/repos/github/lovego/slice/badge.svg?branch=master)](https://coveralls.io/github/lovego/slice?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/slice)](https://goreportcard.com/report/github.com/lovego/slice)
[![GoDoc](https://godoc.org/github.com/lovego/slice?status.svg)](https://godoc.org/github.com/lovego/slice)

## Install
`$ go get github.com/lovego/slice`

## Usage
```go

fmt.Println(ContainsString([]string{"1", "2"}, "1")) // Prints true
  fmt.Println(ContainsInt([]int{1, 2}, 1))             // Prints true
  fmt.Println(ContainsInt64([]int64{1, 2}, int64(1)))  // Prints true

  fmt.Println(IndexString([]string{"2", "1", "3", "1"}, "1"))     // Prints 1
  fmt.Println(LastIndexString([]string{"2", "1", "3", "1"}, "1")) // Prints 3
  fmt.Println(IndexInt([]int{2, 1, 3, 1}, 1))                     // Prints 1
  fmt.Println(LastIndexInt([]int{2, 1, 3, 1}, 1))                 // Prints 3
  fmt.Println(IndexInt64([]int64{2, 1, 3, 1}, 1))                 // Prints 1
  fmt.Println(LastIndexInt64([]int64{2, 1, 3, 1}, 1))             // Prints 3

  fmt.Println(IntersectInt64([]int64{2, 1, 3, 1}, []int64{1, 3, 4})) // Prints [1,3]

  fmt.Println(Remove([]interface{}{"2", 1, "3", true}, []interface{}{"1", "3", true}))    // Prints ["2",1]
  fmt.Println(RemoveStrings([]string{"2", "1", "3", "true"}, []string{"1", "3", "true"})) // Prints ["2"]
  fmt.Println(RemoveInts([]int{2, 1, 3}, []int{1, 3}))                                    // Prints [2]

```

## Documentation
[https://godoc.org/github.com/lovego/slice](https://godoc.org/github.com/lovego/slice)
