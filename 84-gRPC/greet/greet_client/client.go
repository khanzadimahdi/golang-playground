package main

import (
	"context"
	"dummy_grpc/greet/greetpb"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	log.Println("starting client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	greetClient := greetpb.NewGreetServiceClient(conn)

	// doUnary(greetClient)
	// doServerStreaming(greetClient)
	// doClientStreaming(greetClient)
	// doBiDiStreaming(greetClient)
	doUnaryWithDeadline(greetClient, 5*time.Second) // should complete
	doUnaryWithDeadline(greetClient, 1*time.Second) // should timeout
}

func doUnary(greetClient greetpb.GreetServiceClient) {
	greetReq := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mahdi",
			LastName:  "khanzadi",
		},
	}

	greetRes, err := greetClient.Greet(context.Background(), greetReq)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("response from Greet: %v", greetRes.Result)

	// calculatorClient := calculatorpb.NewCaculatorClient(conn)
	// sumbReq := &calculatorpb.SumRequest{
	// 	FirstNumber:  5,
	// 	SecondNumber: 10,
	// }
	// sumRes, err := calculatorClient.Sum(context.Background(), sumbReq)
	// if err != nil {
	// 	log.Fatalf("error while calling Sum RPC: %v", err)
	// }
	// log.Printf("response from Greet: %v", sumRes.SumResult)
}

func doServerStreaming(greetClient greetpb.GreetServiceClient) {

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mahdi",
			LastName:  "Khanzadi",
		},
	}

	resStream, err := greetClient.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreatManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()

		if err == io.EOF {
			// we've reached the end of stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(greetClient greetpb.GreetServiceClient) {
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mahdi",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ali",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Akbar",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Asghar",
			},
		},
	}

	stream, err := greetClient.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over our slice and send each message individually
	for _, req := range requests {
		log.Printf("sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1000 * time.Microsecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet:%v", err)
	}

	log.Printf("LongGreet response: %v", res.GetResult())
}

func doBiDiStreaming(greetClient greetpb.GreetServiceClient) {
	// we create a stream by invoking the client
	stream, err := greetClient.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	wait := make(chan struct{})

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mahdi",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ali",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Akbar",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Asghar",
			},
		},
	}

	// we send a bunch of messages to the client go (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}

		stream.CloseSend()
	}()

	// we recieve a bunch of messages form the client (go routine)
	go func() {
		// function to recieve a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}

			fmt.Printf("recieved: %v", res.GetResult())
		}
		close(wait)
	}()

	// block until everything is done
	<-wait
}

func doUnaryWithDeadline(greetClient greetpb.GreetServiceClient, timeout time.Duration) {
	greetReq := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mahdi",
			LastName:  "khanzadi",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	greetRes, err := greetClient.GreetWithDeadline(ctx, greetReq)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v", statusErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err) 
		}
		return
	}
	log.Printf("response from Greet: %v", greetRes.Result)
}
