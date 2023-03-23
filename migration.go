package main

import "line-bot-app/apiClient"

func migration() {
	apiClient.CreateRecipeMaster()
}
