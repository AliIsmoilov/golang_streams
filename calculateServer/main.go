package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "app/calculatepb"
)

type server struct {
	pb.UnimplementedCalculateServiceServer
}

func (s server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {

	fmt.Println("Unary method request ->>>>>", req)

	return &pb.SumResponse{
		Result: req.GetFirstNumber() + req.GetSecondNumber(),
	}, nil
}

func (s server) SquareRoot(ctx context.Context, req *pb.Squareroot) (*pb.Squareroot, error) {

	fmt.Println("Unary method request ->>>>>", req)

	return &pb.Squareroot{
		Number: int32(math.Sqrt(float64(req.Number))),
	}, nil
}

func (s server) PrimeDecomposition(req *pb.PrimeDecompositionRequest, stream pb.CalculateService_PrimeDecompositionServer) error {

	fmt.Println("Server-side method request ->>>>>", req)

	primeNumer := req.GetNumber()
	divide := int32(2)

	for primeNumer > 1 {

		if primeNumer%divide == 0 {
			stream.Send(
				&pb.PrimeDecompositionResponse{
					Result: divide,
				},
			)

			primeNumer /= divide
		} else {
			divide++
		}
	}

	return nil
}

func (s server) ComputeAvarage(stream pb.CalculateService_ComputeAvarageServer) error {

	sum := 0
	size := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("Client-side method request ->>>>>", req)

		sum += int(req.GetNumber())
		size++

		time.Sleep(time.Second * 1)
	}

	err := stream.SendAndClose(&pb.ComputeAvarageResponse{
		Avarage: float64(sum) / float64(size),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s server) FindMaximum(stream pb.CalculateService_FindMaximumServer) error {

	var max int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		log.Println("FindMaximum => ", req)

		time.Sleep(time.Second * 1)

		if req.GetNumber() > max {
			max = req.GetNumber()

			stream.Send(&pb.FindMaximumResponse{
				Maximum: float64(max),
			})
		}
	}

	return nil
}

func (s server) PerfectNumber(req *pb.PerfectNumberRequest, stream pb.CalculateService_PerfectNumberServer) error {

	fmt.Println("Server-side method request ->>>>>", req)

	Number := req.GetNumber()
	sum := 0

	for i := 1; i <= int(Number); i++ {
		sum = 0
		for j := 1; j <= i/2; j++ {
			if i%j == 0 {
				sum += j
			}
		}

		if sum == i {
			fmt.Println(sum)
			stream.Send(
				&pb.PerfectNumberResponse{
					Number: int32(sum),
				},
			)
		}
	}

	return nil
}

func (s server) FindMinimum(stream pb.CalculateService_FindMinimumServer) error {

	var min int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if count == 0{
			min = req.Number
			count++
		}

		log.Println("FindMinimum => ", req)

		time.Sleep(time.Second * 1)

		if req.GetNumber() < min {
			min = req.GetNumber()

			stream.Send(&pb.FindMinimumResponse{
				Minimum: float64(min),
			})
		}
	}

	return nil
}

func (s server) TotalNumber(stream pb.CalculateService_TotalNumberServer) error {

	sum := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("Client-side method request ->>>>>", req)

		sum += int(req.GetNumber())

		time.Sleep(time.Second * 1)
	}

	err := stream.SendAndClose(&pb.TotalNumberResponse{
		Total: float64(sum),
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculateServiceServer(s, &server{})

	fmt.Println("Calculate server running....")

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
