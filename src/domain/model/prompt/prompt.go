package prompt

import (
	"errors"
	"fmt"
	"regexp"
)



func GeneratePromptText(prompt string) string {
	templateText := fmt.Sprintf(`
	# Template
## Request
- Create an SQL seeder file

## Purpose
- To save the effort of creating a seeder file

## Information
- %s
## Rules
- Please respond with code only
- Output should be in the format of an SQL dump file`,
	 prompt) 

	return templateText
}

func ValidatePromptText(prompt string) error {
    // マイグレーションのパターンを定義
    pattern := `(?i)CREATE\s+TABLE\s+(IF\s+NOT\s+EXISTS\s+)?\w+\s*\([^\)]+\)`
		
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
