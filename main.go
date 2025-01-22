package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rishusahu23/fam-youtube/config"
	youtubePb "github.com/rishusahu23/fam-youtube/gen/api/youtube"
	"github.com/rishusahu23/fam-youtube/youtube/wire"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
)

// Combined gRPC and HTTP server using cmux
func startCombinedServer(ctx context.Context, conf *config.Config) {
	// Create a listener for the shared port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.Server.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", conf.Server.GrpcPort, err)
	}

	// Create a cmux instance
	m := cmux.New(lis)

	// Match connections for gRPC
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	// Match connections for HTTP (gRPC-Gateway and custom HTTP)
	httpL := m.Match(cmux.Any())

	// Start gRPC server
	go func() {
		startGrpcServer(ctx, grpcL, conf)
	}()

	// Start HTTP server
	go func() {
		startGrpcHttpServer(ctx, httpL, conf)
	}()

	// Start cmux serving
	log.Printf("Starting combined gRPC and HTTP server on :%v", conf.Server.GrpcPort)
	if err := m.Serve(); err != nil {
		log.Fatalf("cmux server error: %v", err)
	}
}

// gRPC server setup
func startGrpcServer(ctx context.Context, lis net.Listener, conf *config.Config) {
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 100), // Max receive message size (100 MB)
		grpc.MaxSendMsgSize(1024 * 1024 * 100), // Max send message size (100 MB)
	}

	s := grpc.NewServer(opts...)
	//mongoClient := mongo.GetMongoClient(ctx, conf)
	//redisClient := redis2.GetRedisClient(conf)
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conf.PostgresConfig.User,
		conf.PostgresConfig.Password,
		conf.PostgresConfig.Host,
		conf.PostgresConfig.Port,
		conf.PostgresConfig.DBName,
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	youtubePb.RegisterYoutubeServiceServer(s, wire.InitialiseYoutubeService(conf, db))

	log.Printf("gRPC server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

// gRPC-Gateway server setup
func startGrpcHttpServer(ctx context.Context, lis net.Listener, conf *config.Config) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := youtubePb.RegisterYoutubeServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%v", conf.Server.GrpcPort), opts)
	if err != nil {
		log.Fatalf("failed to register service handler: %v", err)
	}

	httpServer := &http.Server{
		Handler: mux, // Use the mux directly
	}

	log.Printf("HTTP server running")
	if err := httpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve HTTP server: %v", err)
	}
}

func main() {
	ctx := context.Background()
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Start combined gRPC and HTTP server
	startCombinedServer(ctx, conf)

	// Block main goroutine indefinitely (this will keep the servers running)
	select {}
}
