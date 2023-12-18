package prompt

import (
	"errors"
	"fmt"
	"regexp"
)



func GeneratePromptText(prompt string) string {
	templateText := fmt.Sprintf("#Please make a seeder file based on the migration in the information\n#Purpose\nSave me the trouble of making a seeder file\n#Information\n%s#Rules\nPlease answer with just the code.Output in a sql file\n#Output#.", prompt) 

	return templateText
}

func ValidatePromptText(prompt string) error {
    // マイグレーションのパターンを定義
    pattern := `(?i)CREATE\s+TABLE\s+\)`
		
    matched, err := regexp.MatchString(pattern, prompt)
		if err != nil {
			return err
		}
    if !matched {
				err = errors.New("マイグレーションの形式が違います")
        return err
    }
		return nil
}
