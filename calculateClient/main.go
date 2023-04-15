package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "app/calculatepb"
)

func main() {

	conn, err := grpc.Dial(
		"localhost:9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		panic(err)
	}

	stub := pb.NewCalculateServiceClient(conn)

	// doUnary(stub)
	// doServerStream(stub)
	// doClientStream(stub)
	// doBiDirectionalStream(stub)
	// SquareRoot(stub)
	// PerfectNumber(stub)
	// FindMinimum(stub)
	TotalNumber(stub)
}

func doUnary(stub pb.CalculateServiceClient) {

	resp, err := stub.Sum(
		context.Background(),
		&pb.SumRequest{
			FirstNumber:  10,
			SecondNumber: 15,
		},
	)

	if err != nil {
		log.Println("unary error:", err)
		return
	}

	fmt.Println("Unary result ->>>>", resp)
}


func doServerStream(stub pb.CalculateServiceClient) {
	
	stream, err := stub.PrimeDecomposition(
		context.Background(),
		&pb.PrimeDecompositionRequest{
			Number: 40,
		},
	)

	if err != nil {
		log.Println("doServerStream => PrimeDecomposition error:", err)
		return
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("doServerStream => PrimeDecomposition => Recv error:", err)
			return
		}

		fmt.Println("doServerStream result ->>>>", resp.Result)
		time.Sleep(time.Second * 1)
	}
}

func doClientStream(stub pb.CalculateServiceClient) {
	
	stream, err := stub.ComputeAvarage(context.Background())

	if err != nil {
		log.Println("doClientStream => ComputeAvarage error:", err)
		return
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {

		err = stream.Send(&pb.ComputeAvarageRequest{
			Number: int32(number),
		})

		if err != nil {
			log.Println("doClientStream => ComputeAvarage => Send error:", err)
			return
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("doClientStream => ComputeAvarage => CloseAndRecv error:", err)
		return
	}

	fmt.Println("doClientStream => ComputeAvarage => ", resp)
}

func doBiDirectionalStream(stub pb.CalculateServiceClient) {

	stream, err := stub.FindMaximum(context.Background())

	if err != nil {
		log.Println("doBiDirectionalStream => FindMaximum error:", err)
		return
	}

	numbers := []int{1, 20, 3, 5, 60, 7, 8, 40, 9, 10}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {

			stream.Send(&pb.FindMaximumRequest{
				Number: int32(number),
			})
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println("doBiDirectionalStream => FindMaximum => Recv error:", err)
				return
			}

			fmt.Println(resp)
		}

		close(waitc)
	}()

	<-waitc
}


// ------------------------------------------------------------------------------------------

func PerfectNumber(stub pb.CalculateServiceClient) {

	stream, err := stub.PerfectNumber(
		context.Background(),
		&pb.PerfectNumberRequest{
			Number: 1000,
		},
	)

	if err != nil {
		log.Println("PerfectNumber =>  error:", err)
		return
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("PerfectNumber => Recv error:", err)
			return
		}

		fmt.Println("PerfectNumber result ->>>>", resp.Number)
		time.Sleep(time.Second * 1)
	}
}

func SquareRoot(stub pb.CalculateServiceClient) {

	resp, err := stub.SquareRoot(
		context.Background(),
		&pb.Squareroot{
			Number: 9,
		},
	)

	if err != nil {
		log.Println("SquareRoot error:", err)
		return
	}

	fmt.Println("SquareRoot result ->>>>", resp)
}

func FindMinimum(stub pb.CalculateServiceClient) {

	stream, err := stub.FindMinimum(context.Background())

	if err != nil {
		log.Println("FindMinimum error:", err)
		return
	}

	numbers := []int{10, 20, 9, 5, 60, 7, 8, 40, 3, 1}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {

			stream.Send(&pb.FindMinimumRequest{
				Number: int32(number),
			})
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println("FindMinimum => Recv error:", err)
				return
			}

			fmt.Println(resp)
		}

		close(waitc)
	}()

	<-waitc
}

func TotalNumber(stub pb.CalculateServiceClient) {

	stream, err := stub.TotalNumber(context.Background())

	if err != nil {
		log.Println("TotalNumber error:", err)
		return
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {

		err = stream.Send(&pb.TotalNumberRequest{
			Number: int32(number),
		})

		if err != nil {
			log.Println("TotalNumber => Send error:", err)
			return
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("TotalNumber => CloseAndRecv error:", err)
		return
	}

	fmt.Println("TotalNumber => ", resp)
}