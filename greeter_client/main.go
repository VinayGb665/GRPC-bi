package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"fmt"
	pb "grp/service"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	// dail server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewChatClient(conn)
	stream, err := client.SayHello(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	var hash string;
	ctx := stream.Context()
	done := make(chan bool)
	start := time.Now()
	
	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go func() {
		for {
			// generate random nummber and send it to stream
			fmt.Scanf("%s\n",&hash);
			request := pb.Message{Data:hash}
			if err:= stream.Send(&request); err != nil{
				log.Fatalf("Error sending %v",err)
			}
			start = time.Now()
			log.Printf("%s sent", request.Data)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			 
			hash = resp.Data
			fromName := resp.Id
			host_name,err := os.Hostname()
			log.Printf("Data recieved hash:%s\nFrom:%s At:%s In:%s \n", hash,fromName,host_name,time.Since(start))
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("finished with hash=%s", hash)
}