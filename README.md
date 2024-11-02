# Your Project Name

## Overview
This project integrates the power of Wolfram Alpha, Wit.ai, and Slack to provide a Slack bot that can answer complex queries using natural language processing. It demonstrates using Go for real-time communication and data processing.

## Features
- Integration with Slack using `slacker`.
- Natural language processing with `wit.ai`.
- Querying and obtaining results from Wolfram Alpha.

## Prerequisites
Before running this project, you will need:
- A Slack account with bot permissions.
- API keys for Wolfram Alpha and Wit.ai.

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://yourrepo.com/yourproject.git
   cd yourproject

2.Set up Environment Variables Create a .env file in the root directory and populate it with the following:


SLACK_BOTUSER_TOKEN=your_slack_botuser_token
SLACK_SOCKET_TOKEN=your_slack_socket_token
WITAI_SERVER_ACCESS_TOKEN=your_witai_server_access_token
WOLFRAM_APPID=your_wolfram_appid

3. go get .


go run main.go

This will activate the bot, listening to Slack commands. You can interact with the bot by sending commands directly through Slack.

Example Commands
"What is the capital of India?"
"What is the population of China?"
