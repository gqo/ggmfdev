package langsupport

import (
	"net/http"

	"golang.org/x/text/language"
)

var serverLangs = []language.Tag{
	language.AmericanEnglish,
	language.Japanese,
}

var matcher = language.NewMatcher(serverLangs)

func DetermineLanguage(r *http.Request) language.Tag {
	lang, _ := r.Cookie("lang")
	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(matcher, lang.Value, accept)

	return tag
}
