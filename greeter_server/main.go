    
package main

import (
	"io"
	"log"
	"net"


	pb "grp/service"
	"os"
	"crypto/sha1"
	"encoding/hex"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) SayHello(srv pb.Chat_SayHelloServer) error {
	
	log.Println("start new server")
	ctx := srv.Context()

	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// continue if number reveived from stream
		// less than max
		baseStr := req.Data;
		h := sha1.New()
		h.Write([]byte(baseStr))
		hashStr := hex.EncodeToString(h.Sum(nil))

		// update max and send it to stream
		host_name,err:= os.Hostname()
		resp := pb.Hash{Data: hashStr, Id:host_name}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new max=%s", hashStr+" "+host_name)
	}
}

func main() {
	
	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterChatServer(s, server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}