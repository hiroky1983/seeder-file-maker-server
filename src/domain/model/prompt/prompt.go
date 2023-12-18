package prompt

import "fmt"

func GeneratePromptText(prompt string) string {
	templateText := fmt.Sprintf("#Please make a seeder file based on the migration in the information\n#Purpose\nSave me the trouble of making a seeder file\n#Information\n%s#Rules\nPlease answer with just the code.Output in a sql file\n#Output#.", prompt) 

	return templateText
}