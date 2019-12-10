package tool

import (
	"log"
	"runtime"
)

func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%v(mb) HeapIdle:%v(mb) HeapReleased:%v(mb)", float64(ms.Alloc)/1000/1000, float64(ms.HeapIdle)/1000/1000, float64(ms.HeapReleased)/1000/1000)
}
