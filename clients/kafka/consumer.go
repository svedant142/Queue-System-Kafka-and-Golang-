package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"message-queue-system/business/product/repository"
	"message-queue-system/db"
	"message-queue-system/domain/entity"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kafka "github.com/IBM/sarama"
)

const topic = "topic.product.id"

func InitConsumer(brokersUrl []string) (kafka.Consumer, error) {
	config := kafka.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := kafka.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Consume() {

	worker, err := InitConsumer([]string{"localhost:9092"})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, kafka.OffsetOldest)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Consumer started")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				go ProcessImages(context.Background() , msg.Value)
			case <-sigchan:
				fmt.Println("Interrupted")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	if err := worker.Close(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ProcessImages(ctx context.Context, value []byte)  error {
	var productID int64
	err := json.Unmarshal(value, &productID)
	if err != nil {
		fmt.Println(err.Error()) 
		return err
	}
	//get product image from DB by productID
	productImage, err := repository.NewProductRepo(db.DbClient).GetProductImage(ctx, productID)
	if err != nil {
		fmt.Println(err.Error()) 
		return err
	}
	
	// Download and compress image and store in  local path
	var compressedPath entity.URLS
	for idx, imgPath := range productImage {
		path	:= fmt.Sprintf("E:\\Vedant\\Images\\productID-%d-%d.png", productID, idx)
		err :=	downloadAndStoreImage(imgPath, path, productID)
		if err != nil {
			fmt.Println(err.Error()) 
			continue 
		}
		compressedPath = append(compressedPath, path)
	}

	//store path in compressed_product_images column
	err = repository.NewProductRepo(db.DbClient).UpdateImagePath(ctx, compressedPath, productID)
	if err != nil {
		fmt.Println(err.Error()) 
		return err
	}
	return nil
}

func downloadAndStoreImage(imagePath, outputPath string, productID int64) (error) {
	response, err := http.Get(imagePath)
	if err != nil {
		return err
	}
	defer response.Body.Close()


	//compressing
	img, _, err := image.Decode(response.Body)
	if err != nil {
		return err
	}
	
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}
