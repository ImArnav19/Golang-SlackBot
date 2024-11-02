package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Edw590/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"

	witai "github.com/wit-ai/wit-go/v2"
)

func printEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	godotenv.Load(".env")

	//Bot code
	bot := slacker.NewClient(os.Getenv("SLACK_BOTUSER_TOKEN"), os.Getenv("SLACK_SOCKET_TOKEN"))

	witClient := witai.NewClient(os.Getenv("WITAI_SERVER_ACCESS_TOKEN"))

	wolfram_client := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APPID")}

	go printEvents(bot.CommandEvents())

	bot.Command("<message>", &slacker.CommandDefinition{
		Description: "Get a response from Wit.ai",
		Examples:    []string{"What is the capital of India", "What is the population of China"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			query := request.Param("message")

			msg, _ := witClient.Parse(&witai.MessageRequest{
				Query: query,
			})

			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := value.String()

			res, err := wolfram_client.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(value)
			response.Reply(res)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
