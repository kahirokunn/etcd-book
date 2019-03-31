package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//#@@range_begin(kv)
	chapter2KV := namespace.NewKV(client.KV, "/chapter2")
	//#@@range_end(kv)

	//#@@range_begin(get)
	client.Put(context.TODO(), "/chapter2/ns/1", "test")
	resp, _ := chapter2KV.Get(context.TODO(), "/ns/1")
	fmt.Printf("%s: %s\n", resp.Kvs[0].Key, resp.Kvs[0].Value)
	//#@@range_end(get)

	//#@@range_begin(put)
	chapter2KV.Put(context.TODO(), "/ns/2", "hoge")
	resp, _ = client.Get(context.TODO(), "/chapter2/ns/2")
	fmt.Printf("%s: %s\n", resp.Kvs[0].Key, resp.Kvs[0].Value)
	//#@@range_end(put)

	//#@@range_begin(namespace)
	client.KV = namespace.NewKV(client.KV, "/chapter2")
	client.Watcher = namespace.NewWatcher(client.Watcher, "/chapter2")
	client.Lease = namespace.NewLease(client.Lease, "/chapter2")
	//#@@range_end(namespace)
}
