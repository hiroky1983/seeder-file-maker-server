package usecase

import (
	"seeder-app/domain/repository"
	"seeder-app/moq/fakerepository"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestPrompt_Prompt(t *testing.T) {
	type fields struct {
		promptRepo repository.Prompt
	}
	type args struct {
		c    *openai.Client
		prom string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				promptRepo: &fakerepository.PromptMock{
					PromptFunc: func(client *openai.Client, s string) (string, error) {
						return "test", nil
					},
				},
		},
			args: args{
				prom: "CREATE TABLE test (id int);",
			},
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prompt{
				promptRepo: tt.fields.promptRepo,
			}
			got, err := p.Prompt(tt.args.c, tt.args.prom)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prompt.Prompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Prompt.Prompt() = %v, want %v", got, tt.want)
			}
		})
	}
}
