package main

import "google.golang.org/grpc"

func main() {
	grpc.Dial("0.0.0.0:56001", grpc.WithInsecure())
}
