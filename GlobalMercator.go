package GlobalMercator

import (
	"fmt"
	"math"
	"sync"
)

var RADIUS = 6378137
var TILESIZE = 256
var PI2 = math.Pi * 2

type GlobalMercator struct {
	tileSize int
	initialResolution float64
	originShift float64
}

var once sync.Once
var instance *GlobalMercator

func getInstance() *GlobalMercator {
	once.Do(func() {
		instance = &GlobalMercator{TILESIZE, PI2 * float64(RADIUS) / float64(TILESIZE), PI2 * float64(RADIUS)/2}
	})

	return instance
}

func (gm GlobalMercator) test (){
	fmt.Println(gm.tileSize)
	fmt.Println(gm.initialResolution)
	fmt.Println(gm.originShift)
}
