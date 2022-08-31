package grpc

import (
	"context"
	"github.com/lschulzes/hex_arch/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateInputs(req *pb.OperationParameters) error {
	if req.GetA() == 0 || req.GetB() == 0 {
		return status.Error(codes.InvalidArgument, "missing required")
	}
	return nil
}

func (grpcA Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}
	if err := ValidateInputs(req); err != nil {
		return ans, err
	}
	result, err := grpcA.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Errorf(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: result,
	}
	return ans, nil
}

func (grpcA Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}
	if err := ValidateInputs(req); err != nil {
		return ans, err
	}
	result, err := grpcA.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return ans, status.Errorf(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: result,
	}
	return ans, nil
}

func (grpcA Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}
	if err := ValidateInputs(req); err != nil {
		return ans, err
	}
	result, err := grpcA.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Errorf(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: result,
	}
	return ans, nil
}

func (grpcA Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}
	if err := ValidateInputs(req); err != nil {
		return ans, err
	}
	result, err := grpcA.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Errorf(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: result,
	}
	return ans, nil
}
