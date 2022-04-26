# Core Utils - Generic utils, like lodash

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/core-utils)](https://pkg.go.dev/github.com/go-zoox/core-utils)
[![Build Status](https://github.com/go-zoox/core-utils/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/core-utils/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/core-utils)](https://goreportcard.com/report/github.com/go-zoox/core-utils)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/core-utils/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/core-utils?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/core-utils.svg)](https://github.com/go-zoox/core-utils/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/core-utils.svg?label=Release)](https://github.com/go-zoox/core-utils/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/core-utils
```

## Getting Started

```go
// see test cases
```

## Spec

### Array

```go
func Filter[K any](values []K, fn func(value K, index int) bool) []K
func Find[T comparable](elements []T, compare func(element T, index int) bool) (T, bool)
func FindIndex[T comparable](collection []T, value T) int
func ForEach[K any](values []K, fn func(value K, index int))
func GroupBy[T any, K comparable](values []T, iteratee func(value T, index int) K) map[K][]T
func Includes[K comparable](values []K, value K) bool
func IndexOf[T comparable](collection []T, value T) int
func Last[T any](collection []T) T
func LastIndexOf[T comparable](collection []T, value T) int
func Map[K any, R any](values []K, fn func(value K, index int) R) []R
func Max[T Number](collection []T) T
func Min[T Number](collection []T) T
func Reduce[K any, R any](values []K, fn func(all R, value K, index int) R, initialValue R) R
func Repeat[T any](count int, value T) []T
func Reverse[T any](collection []T) []T
func Shuffle[T any](collection []T) []T
func UniqueBy[T any, U comparable](collection []T, iteratee func(value T, index int) U) []T
type Number interface{ ... }
```

### Object

```go
func Get[K comparable, V any](object map[K]V, key string) V
func Indexes[K comparable, V any](object map[K]V) map[string]V
func Keys[K comparable, V any](object map[K]V) []K
func Map[K comparable, V any](object map[K]V, iteratee func(V, K) V) map[K]V
func Omit[K comparable, V any](object map[K]V, keys ...K) map[K]V
func OmitBy[K comparable, V any](object map[K]V, fn func(V, K) bool) map[K]V
func Pick[K comparable, V any](object map[K]V, keys ...K) map[K]V
func PickBy[K comparable, V any](object map[K]V, fn func(V, K) bool) map[K]V
func ToMap(in any, tag ...string) (map[string]any, error)
func Values[K comparable, V any](object map[K]V) []V
func Walk[K comparable, V any](object map[K]V, fn func(V, K, string))
func WalkByDeepFirst[K comparable, V any](object map[K]V, fn func(V, K, string), parent string)
```

## License
GoZoox is released under the [MIT License](./LICENSE).
