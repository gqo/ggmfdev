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

	var tagIndex int

	if lang == nil {
		_, tagIndex = language.MatchStrings(matcher, accept)
	} else {
		_, tagIndex = language.MatchStrings(matcher, lang.Value, accept)
	}

	// Why? See: https://stackoverflow.com/questions/49997766/language-matchstrings-returns-garbage
	tag := serverLangs[tagIndex]

	return tag
}
