package main

import "fmt"

func ExampleRedisRunOk() {
	//*DryRunFlag = false
	DryRunPush("", "1234566789")
	fmt.Println(redisDockerRun())
	// Output:
	// docker pull library/redis:5
	// docker run -d -p 6379:6379 --rm --name iosdk-redis --hostname redis library/redis:5 --requirepass password
}

func ExampleRedisRunKo() {
	//*DryRunFlag = false
	DryRunPush("cannot pull library/redis:5")
	fmt.Println(1, redisDockerRun())
	DryRunPush("", "!cannot start")
	fmt.Println(2, redisDockerRun())
	// Output:
	// docker pull library/redis:5
	// 1 cannot pull library/redis:5
	// docker pull library/redis:5
	// docker run -d -p 6379:6379 --rm --name iosdk-redis --hostname redis library/redis:5 --requirepass password
	// 2 cannot start redis: cannot start
}

func ExampleRedisDockerDestroy() {
	// *DryRunFlag = false
	DryRunPush("iosdk-redis")
	RedisDestroy()
	// Output:
	// docker stop iosdk-redis
	// Destroying Redis: iosdk-redis
}
