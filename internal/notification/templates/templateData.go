package templates

import (
	"fmt"
	"html"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/i18n"
)

const (
	DefaultFont            = "http://fonts.googleapis.com/css?family=Lato:200,300,400,600"
	DefaultFontFamily      = "-apple-system, BlinkMacSystemFont, Segoe UI, Lato, Arial, Helvetica, sans-serif"
	DefaultLogo            = "https://static.zitadel.ch/zitadel-logo-dark@3x.png"
	DefaultFontColor       = "#22292f"
	DefaultBackgroundColor = "#fafafa"
	DefaultPrimaryColor    = "#5282C1"
)

type TemplateData struct {
	Title           string
	PreHeader       string
	Subject         string
	Greeting        string
	Text            string
	Href            string
	ButtonText      string
	PrimaryColor    string
	BackgroundColor string
	FontColor       string
	LogoURL         string
	FontURL         string
	FontFamily      string

	IncludeFooter bool
	FooterText    string
}

func (data *TemplateData) Translate(translator *i18n.Translator, msgType string, args map[string]interface{}, langs ...string) {
	data.Title = translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessageTitle), args, langs...)
	data.PreHeader = translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessagePreHeader), args, langs...)
	data.Subject = translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessageSubject), args, langs...)
	data.Greeting = translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessageGreeting), args, langs...)
	data.Text = html.UnescapeString(translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessageText), args, langs...))
	data.ButtonText = translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessageButtonText), args, langs...)
	data.FooterText = translator.Localize(fmt.Sprintf("%s.%s", msgType, domain.MessageFooterText), args, langs...)
}
