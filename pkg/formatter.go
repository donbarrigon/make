package formatter

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var IrregularPlurals = map[string]string{
	"person":     "people",
	"child":      "children",
	"foot":       "feet",
	"tooth":      "teeth",
	"mouse":      "mice",
	"man":        "men",
	"woman":      "women",
	"ox":         "oxen",
	"cactus":     "cacti",
	"focus":      "foci",
	"analysis":   "analyses",
	"thesis":     "theses",
	"crisis":     "crises",
	"diagnosis":  "diagnoses",
	"appendix":   "appendices",
	"vertex":     "vertices",
	"index":      "indices",
	"matrix":     "matrices",
	"axis":       "axes",
	"basis":      "bases",
	"fungus":     "fungi",
	"radius":     "radii",
	"alumnus":    "alumni",
	"curriculum": "curricula",
	"datum":      "data",
	"medium":     "media",
	"forum":      "fora",
	"bacterium":  "bacteria",
	"syllabus":   "syllabi",
	"criterion":  "criteria",
	"aquarium":   "aquaria",
	"stadium":    "stadia",
	"stimulus":   "stimuli",
	"die":        "dice",
	"formula":    "formulae",
	"genus":      "genera",
	"bison":      "bison",    // no cambia
	"deer":       "deer",     // no cambia
	"sheep":      "sheep",    // no cambia
	"salmon":     "salmon",   // no cambia
	"aircraft":   "aircraft", // no cambia
	"series":     "series",   // no cambia
	"species":    "species",  // no cambia
	"fish":       "fish",     // no cambia
	"trousers":   "trousers", // no cambia
	"scissors":   "scissors", // no cambia
	"clothes":    "clothes",  // no cambia
	"news":       "news",     // no cambia
}

// splitWords separa una cadena en palabras sin importar el formato original.
func splitWords(s string) []string {
	if s == "" {
		return nil
	}

	// Normaliza separadores comunes
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			prev := rune(s[i-1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) {
				result.WriteRune(' ')
			}
		}
		result.WriteRune(r)
	}

	// Pasa todo a minúsculas y divide
	words := strings.Fields(strings.ToLower(result.String()))
	return words
}

// Capitaliza la primera letra de una palabra (Unicode-safe)
var titleCaser = cases.Title(language.Und)

func capitalize(s string) string {
	return titleCaser.String(s)
}

// 🐍 snake_case
func ToSnakeCase(s string) string {
	words := splitWords(s)
	return strings.Join(words, "_")
}

// 🐫 camelCase
func ToCamelCase(s string) string {
	words := splitWords(s)
	for i := range words {
		if i == 0 {
			continue
		}
		words[i] = capitalize(words[i])
	}
	return strings.Join(words, "")
}

// 🐘 PascalCase
func ToPascalCase(s string) string {
	words := splitWords(s)
	for i := range words {
		words[i] = capitalize(words[i])
	}
	return strings.Join(words, "")
}

// 🍢 kebab-case
func ToKebabCase(s string) string {
	words := splitWords(s)
	return strings.Join(words, "-")
}

// Pluralize pluraliza una cadena según las reglas estándar de pluralización en inglés.
func Pluralize(word string) string {
	if word == "" {
		return ""
	}

	if plural, exists := IrregularPlurals[word]; exists {
		return plural
	}

	if strings.HasSuffix(word, "y") {
		if len(word) > 1 && isVowel(rune(word[len(word)-2])) {
			return word + "s"
		}
		return word[:len(word)-1] + "ies"
	}

	if strings.HasSuffix(word, "s") ||
		strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "z") ||
		strings.HasSuffix(word, "ch") ||
		strings.HasSuffix(word, "sh") {
		return word + "es"
	}

	if strings.HasSuffix(word, "f") {
		return word[:len(word)-1] + "ves"
	}
	if strings.HasSuffix(word, "fe") {
		return word[:len(word)-2] + "ves"
	}

	return word + "s"
}

// PluralizeIdentifier pluraliza la última palabra de un identificador
// manteniendo el formato original (snake_case, camelCase, PascalCase, kebab-case).
func PluralizeIdentifier(s string) string {
	if s == "" {
		return ""
	}

	// Detectar el formato original
	format := detectFormat(s)

	// Separar en palabras
	words := splitWords(s)
	if len(words) == 0 {
		return s
	}

	// Pluralizar solo la última palabra
	words[len(words)-1] = Pluralize(words[len(words)-1])

	// Reconstruir en el formato original
	switch format {
	case "snake":
		return ToSnakeCase(strings.Join(words, " "))
	case "camel":
		return ToCamelCase(strings.Join(words, " "))
	case "pascal":
		return ToPascalCase(strings.Join(words, " "))
	case "kebab":
		return ToKebabCase(strings.Join(words, " "))
	default:
		return strings.Join(words, " ")
	}
}

// detectFormat detecta el formato del identificador
func detectFormat(s string) string {
	if strings.Contains(s, "_") {
		return "snake"
	}
	if strings.Contains(s, "-") {
		return "kebab"
	}

	// Detectar camelCase vs PascalCase
	hasUpper := false
	for i, r := range s {
		if unicode.IsUpper(r) {
			hasUpper = true
			if i == 0 {
				return "pascal"
			}
			return "camel"
		}
	}

	// Si no tiene mayúsculas ni separadores, asumir snake_case
	if !hasUpper {
		return "snake"
	}

	return "unknown"
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aiueo", r)
}
