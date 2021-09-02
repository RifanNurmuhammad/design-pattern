package proxy

import (
	"fmt"
	"time"
)

// Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object.
// A proxy controls access to the original object,
// allowing you to perform something either before or after the request gets through to the original object.

type ThirdPartyLogistics interface {
	TrackingAwb() []string
}

type ThirdPartyLogisticsCall struct{}

func (t ThirdPartyLogisticsCall) TrackingAwb() []string {
	<-time.After(300 * time.Millisecond)
	return []string{
		"Order Pickup",
		"Order Deliver",
		"Order Received",
	}
}

type CacheLogistics struct {
	service ThirdPartyLogistics
	cache   []string
}

func (c *CacheLogistics) start() {
	for {
		select {
		case <-time.After(5 * time.Second):
			c.cache = []string{}
		}
	}
}

func (c *CacheLogistics) TrackingAwb() []string {
	if len(c.cache) == 0 {
		// set cache when cache empty
		c.cache = c.service.TrackingAwb()
		fmt.Println("get from thirdparty")
	}

	fmt.Println("get from cache")
	return c.cache
}

func Init() {
	t := ThirdPartyLogisticsCall{}
	cache := CacheLogistics{
		service: t,
	}

	go cache.start()

	// first call, return from thirdparty
	fmt.Println(cache.TrackingAwb())
	// second call, return from cache
	fmt.Println(cache.TrackingAwb())

	// after 6 second will set empty cache, because `cache.start()` clean cache 5 second
	<-time.After(7 * time.Second)

	fmt.Println(cache.TrackingAwb())
	fmt.Println(cache.TrackingAwb())
}
