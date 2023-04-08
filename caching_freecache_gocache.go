/*
show coocood/freecache example using eko/gocache

In this example, we first create a FreeCache instance using the freecache.NewCache function and passing in the cache capacity in bytes.

We then create a cache manager instance using the cache.NewManager function, and add the FreeCache to it using the Add method.

We define a key, an expiration time, and some data to be cached. We then try to get the data from cache using the Get method of the cache manager. 
If the data exists in cache, we print a message and use the cached data. Otherwise, we fetch the data from the database (not shown in this example), 
set the data to cache using the Set method of the cache manager, and use the fetched data.

Note that the Eko GoCache library provides additional features such as cache tags, cache group, and cache chain. You may need to modify the code to 
fit your specific requirements.
*/

package main

import (
    "fmt"
    "time"
    "github.com/eko/gocache/v2/cache"
    "github.com/eko/gocache/v2/cache/freecache"
)

func main() {
    // Create a FreeCache cache instance with 50 MB capacity
    freeCache := freecache.NewCache(50 * 1024 * 1024)

    // Create a cache manager instance and add FreeCache cache to it
    cacheManager := cache.NewManager()
    cacheManager.Add(freeCache)

    // Define a key, expiration time, and data to be cached
    key := "mykey"
    expiration := 30 * time.Minute
    data := "Hello, world!"

    // Try to get data from cache
    val, err := cacheManager.Get(key)
    if err == nil {
        // Data exists in cache
        fmt.Println("Data found in cache:", val)
    } else {
        // Data does not exist in cache, fetch from database
        fmt.Println("Data not found in cache. Fetching from database...")

        // TODO: Fetch data from database

        // Set data to cache
        err := cacheManager.Set(&cache.Item{
            Key:        key,
            Value:      []byte(data),
            Expiration: expiration,
        })
        if err != nil {
            panic(err)
        }

        fmt.Println("Data set to cache:", data)
    }
}
