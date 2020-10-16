package client

import (
	// "ch15/series" //src是GOPATH的一部分，所以目录结构是从src以后写起
	"learn/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonaciiSerie(5))
	// t.Log(series.square(5)) //cannot refer to unexported name series.square  undefined: series.square
}
