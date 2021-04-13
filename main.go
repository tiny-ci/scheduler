package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tiny-ci/core/db"
	"github.com/tiny-ci/core/parser"
	"github.com/tiny-ci/core/pipe"
	"github.com/tiny-ci/core/schema"
	"log"
)

func newPipeline(message string) error {
	bmsg := []byte(message)

	err := schema.ValidateNotification(bmsg)
	if err != nil {
		log.Println("notification validation error")
		return err
	}

	ntf, err := parser.ParseNotification(bmsg)
	if err != nil {
		log.Println("notification parser error")
		return nil
	}

	configContent, err := pipe.Fetch(&pipe.GitRef{
		Name:  ntf.Info.RefName,
		URL:   ntf.Info.URL,
		Hash:  ntf.Info.CommitHash,
		IsTag: ntf.Info.IsTag,
	})

	if err != nil {
		log.Println("fetch error")
		return err
	}

	if err == nil && configContent == nil {
		log.Println("pipe config not found")
		return nil
	}

	config, err := parser.ParsePipeConfig(configContent.Bytes())
	if err != nil {
		log.Println("parser error")
		return err
	}

	var marshalledConfig []byte
	marshalledConfig, err = json.Marshal(config)
	if err != nil {
		log.Println("marshall error")
		return err
	}

	err = schema.ValidateConfig(marshalledConfig)
	if err != nil {
		log.Println("validation error")
		return err
	}

	matchedJobs, err := pipe.Filter(config, ntf.Info.RefName)
	if err != nil {
		log.Println("filter error")
		return err
	}

	if len(matchedJobs) == 0 {
		log.Println("no jobs to be processed")
		return nil
	}

	rdb, err := db.New("localhost:6379")
	if err != nil {
		log.Println("redis conn error")
		return err
	}

	err = rdb.Populate(ntf, &matchedJobs)
	if err != nil {
		log.Println("redis population error")
		return err
	}

	return nil
}

func handler(request events.SNSEvent) {
	if len(request.Records) == 0 {
		log.Println("nothing to process")
		return
	}

	for _, record := range request.Records {
		switch subject := record.SNS.Subject; subject {
		case "new_pipeline":
			err := newPipeline(record.SNS.Message)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	lambda.Start(handler)
}
