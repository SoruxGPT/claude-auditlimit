package api

var (
	MsgPlus429 = `
	{
		"error": {
		  "message": "You have sent too many messages to the model. Please try again later."
		}
	  }
	`

	MsgMod400 = `
	{
		"error": {
		  "message": "This content may violate [OpenAI Usage Policies](https://openai.com/policies/usage-policies)."
		}
	}
	`
)
