myServer := grpc.NewServer(
	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(AuthInterceptor))
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(interface{}, error){
	conn, err := grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewAuthClient(ctx, &pb.ValidateToken{Token: token}
		//shippyserver и дальше как там
}