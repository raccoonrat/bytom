// generated by "go generate"; DO NOT EDIT

package langreg

import (
	"errors"
	"fmt"
)

// LangCodeInfo returns the English and native language in its script for a
// given string, and an error if any.  If there are more than one official
// names for the language (either English or native), they are separated by a
// semi-colon (;)
// Language codes should always be lowercase, and this is enforced.
func LangCodeInfo(s string) (english, native string, err error) {

	// codes have to be two characters long
	if len(s) != 2 {
		return "", "",
			errors.New("ISO 639-1 language codes must be 2 characters long")
	}

	switch s[0] {

	case 'a':
		switch s[1] {
		case 'a':
			return "Afar", "Afaraf", nil
		case 'b':
			return "Abkhazian", "аҧсуа бызшәа; аҧсшәа", nil
		case 'e':
			return "Avestan", "avesta", nil
		case 'f':
			return "Afrikaans", "Afrikaans", nil
		case 'k':
			return "Akan", "Akan", nil
		case 'm':
			return "Amharic", "አማርኛ", nil
		case 'n':
			return "Aragonese", "aragonés", nil
		case 'r':
			return "Arabic", "العربية", nil
		case 's':
			return "Assamese", "অসমীয়া", nil
		case 'v':
			return "Avaric", "авар мацӀ; магӀарул мацӀ", nil
		case 'y':
			return "Aymara", "aymar aru", nil
		case 'z':
			return "Azerbaijani", "azərbaycan dili", nil
		}

	case 'b':
		switch s[1] {
		case 'a':
			return "Bashkir", "башҡорт теле", nil
		case 'e':
			return "Belarusian", "беларуская мова", nil
		case 'g':
			return "Bulgarian", "български език", nil
		case 'h':
			return "Bihari languages", "भोजपुरी", nil
		case 'i':
			return "Bislama", "Bislama", nil
		case 'm':
			return "Bambara", "bamanankan", nil
		case 'n':
			return "Bengali", "বাংলা", nil
		case 'o':
			return "Tibetan", "བོད་ཡིག", nil
		case 'r':
			return "Breton", "brezhoneg", nil
		case 's':
			return "Bosnian", "bosanski jezik", nil
		}

	case 'c':
		switch s[1] {
		case 'a':
			return "Catalan; Valencian", "català; valencià", nil
		case 'e':
			return "Chechen", "нохчийн мотт", nil
		case 'h':
			return "Chamorro", "Chamoru", nil
		case 'o':
			return "Corsican", "corsu; lingua corsa", nil
		case 'r':
			return "Cree", "ᓀᐦᐃᔭᐍᐏᐣ", nil
		case 's':
			return "Czech", "čeština; český jazyk", nil
		case 'u':
			return "Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic", "ѩзыкъ словѣньскъ", nil
		case 'v':
			return "Chuvash", "чӑваш чӗлхи", nil
		case 'y':
			return "Welsh", "Cymraeg", nil
		}

	case 'd':
		switch s[1] {
		case 'a':
			return "Danish", "dansk", nil
		case 'e':
			return "German", "Deutsch", nil
		case 'v':
			return "Divehi; Dhivehi; Maldivian", "ދިވެހި", nil
		case 'z':
			return "Dzongkha", "རྫོང་ཁ", nil
		}

	case 'e':
		switch s[1] {
		case 'e':
			return "Ewe", "Eʋegbe", nil
		case 'l':
			return "Greek, Modern (1453-)", "ελληνικά", nil
		case 'n':
			return "English", "English", nil
		case 'o':
			return "Esperanto", "Esperanto", nil
		case 's':
			return "Spanish; Castilian", "español; castellano", nil
		case 't':
			return "Estonian", "eesti; eesti keel", nil
		case 'u':
			return "Basque", "euskara; euskera", nil
		}

	case 'f':
		switch s[1] {
		case 'a':
			return "Persian", "فارسی", nil
		case 'f':
			return "Fulah", "Fulfulde; Pulaar; Pular", nil
		case 'i':
			return "Finnish", "suomi; suomen kieli", nil
		case 'j':
			return "Fijian", "vosa Vakaviti", nil
		case 'o':
			return "Faroese", "føroyskt", nil
		case 'r':
			return "French", "français; langue française", nil
		case 'y':
			return "Western Frisian", "Frysk", nil
		}

	case 'g':
		switch s[1] {
		case 'a':
			return "Irish", "Gaeilge", nil
		case 'd':
			return "Gaelic; Scottish Gaelic", "Gàidhlig", nil
		case 'l':
			return "Galician", "galego", nil
		case 'n':
			return "Guarani", "Avañeẽ", nil
		case 'u':
			return "Gujarati", "ગુજરાતી", nil
		case 'v':
			return "Manx", "Gaelg; Gailck", nil
		}

	case 'h':
		switch s[1] {
		case 'a':
			return "Hausa", "(Hausa) هَوُسَ", nil
		case 'e':
			return "Hebrew", "עברית", nil
		case 'i':
			return "Hindi", "हिन्दी; हिंदी", nil
		case 'o':
			return "Hiri Motu", "Hiri Motu", nil
		case 'r':
			return "Croatian", "hrvatski jezik", nil
		case 't':
			return "Haitian; Haitian Creole", "Kreyòl ayisyen", nil
		case 'u':
			return "Hungarian", "magyar", nil
		case 'y':
			return "Armenian", "Հայերեն", nil
		case 'z':
			return "Herero", "Otjiherero", nil
		}

	case 'i':
		switch s[1] {
		case 'a':
			return "Interlingua (International Auxiliary language Association)", "Interlingua", nil
		case 'd':
			return "Indonesian", "Bahasa Indonesia", nil
		case 'e':
			return "Interlingue; Occidental", "Interlingue; Occidental", nil
		case 'g':
			return "Igbo", "Asụsụ Igbo", nil
		case 'i':
			return "Sichuan Yi; Nuosu", "ꆈꌠ꒿ Nuosuhxop", nil
		case 'k':
			return "Inupiaq", "Iñupiaq; Iñupiatun", nil
		case 'o':
			return "Ido", "Ido", nil
		case 's':
			return "Icelandic", "Íslenska", nil
		case 't':
			return "Italian", "italiano", nil
		case 'u':
			return "Inuktitut", "ᐃᓄᒃᑎᑐᑦ", nil
		}

	case 'j':
		switch s[1] {
		case 'a':
			return "Japanese", "日本語(にほんご)", nil
		case 'v':
			return "Javanese", "basa Jawa", nil
		}

	case 'k':
		switch s[1] {
		case 'a':
			return "Georgian", "ქართული", nil
		case 'g':
			return "Kongo", "Kikongo", nil
		case 'i':
			return "Kikuyu; Gikuyu", "Gĩkũyũ", nil
		case 'j':
			return "Kuanyama; Kwanyama", "Kuanyama", nil
		case 'k':
			return "Kazakh", "қазақ тілі", nil
		case 'l':
			return "Kalaallisut; Greenlandic", "kalaallisut; kalaallit oqaasii", nil
		case 'm':
			return "Central Khmer", "ខ្មែរ; ខេមរភាសា; ភាសាខ្មែរ", nil
		case 'n':
			return "Kannada", "ಕನ್ನಡ", nil
		case 'o':
			return "Korean", "한국어; 조선어", nil
		case 'r':
			return "Kanuri", "Kanuri", nil
		case 's':
			return "Kashmiri", "कश्मीरी; كشميري\u200e", nil
		case 'u':
			return "Kurdish", "Kurdî; كوردی\u200e", nil
		case 'v':
			return "Komi", "коми кыв", nil
		case 'w':
			return "Cornish", "Kernewek", nil
		case 'y':
			return "Kirghiz; Kyrgyz", "Кыргызча; Кыргыз тили", nil
		}

	case 'l':
		switch s[1] {
		case 'a':
			return "Latin", "latine; lingua latina", nil
		case 'b':
			return "Luxembourgish; Letzeburgesch", "Lëtzebuergesch", nil
		case 'g':
			return "Ganda", "Luganda", nil
		case 'i':
			return "Limburgan; Limburger; Limburgish", "Limburgs", nil
		case 'n':
			return "Lingala", "Lingála", nil
		case 'o':
			return "Lao", "ພາສາລາວ", nil
		case 't':
			return "Lithuanian", "lietuvių kalba", nil
		case 'u':
			return "Luba-Katanga", "Tshiluba", nil
		case 'v':
			return "Latvian", "latviešu valoda", nil
		}

	case 'm':
		switch s[1] {
		case 'g':
			return "Malagasy", "fiteny malagasy", nil
		case 'h':
			return "Marshallese", "Kajin M̧ajeļ", nil
		case 'i':
			return "Maori", "te reo Māori", nil
		case 'k':
			return "Macedonian", "македонски јазик", nil
		case 'l':
			return "Malayalam", "മലയാളം", nil
		case 'n':
			return "Mongolian", "монгол", nil
		case 'r':
			return "Marathi", "मराठी", nil
		case 's':
			return "Malay", "bahasa Melayu; بهاس ملايو\u200e", nil
		case 't':
			return "Maltese", "Malti", nil
		case 'y':
			return "Burmese", "ဗမာစာ", nil
		}

	case 'n':
		switch s[1] {
		case 'a':
			return "Nauru", "Ekakairũ Naoero", nil
		case 'b':
			return "Bokmål, Norwegian; Norwegian Bokmål", "Norsk bokmål", nil
		case 'd':
			return "Ndebele, North; North Ndebele", "isiNdebele", nil
		case 'e':
			return "Nepali", "नेपाली", nil
		case 'g':
			return "Ndonga", "Owambo", nil
		case 'l':
			return "Dutch; Flemish", "Nederlands; Vlaams", nil
		case 'n':
			return "Norwegian Nynorsk; Nynorsk, Norwegian", "Norsk nynorsk", nil
		case 'o':
			return "Norwegian", "Norsk", nil
		case 'r':
			return "Ndebele, South; South Ndebele", "isiNdebele", nil
		case 'v':
			return "Navajo; Navaho", "Diné bizaad; Dinékʼehǰí", nil
		case 'y':
			return "Chichewa; Chewa; Nyanja", "chiCheŵa; chinyanja", nil
		}

	case 'o':
		switch s[1] {
		case 'c':
			return "Occitan (post 1500); Provençal", "occitan; lenga dòc", nil
		case 'j':
			return "Ojibwa", "ᐊᓂᔑᓈᐯᒧᐎᓐ", nil
		case 'm':
			return "Oromo", "Afaan Oromoo", nil
		case 'r':
			return "Oriya", "ଓଡ଼ିଆ", nil
		case 's':
			return "Ossetian; Ossetic", "ирон æвзаг", nil
		}

	case 'p':
		switch s[1] {
		case 'a':
			return "Panjabi; Punjabi", "ਪੰਜਾਬੀ; پنجابی\u200e", nil
		case 'i':
			return "Pali", "पाऴि", nil
		case 'l':
			return "Polish", "język polski; polszczyzna", nil
		case 's':
			return "Pushto; Pashto", "پښتو", nil
		case 't':
			return "Portuguese", "português", nil
		}

	case 'q':
		switch s[1] {
		case 'u':
			return "Quechua", "Runa Simi; Kichwa", nil
		}

	case 'r':
		switch s[1] {
		case 'm':
			return "Romansh", "rumantsch grischun", nil
		case 'n':
			return "Rundi", "Ikirundi", nil
		case 'o':
			return "Romanian; Moldavian; Moldovan", "limba română", nil
		case 'u':
			return "Russian", "русский язык", nil
		case 'w':
			return "Kinyarwanda", "Ikinyarwanda", nil
		}

	case 's':
		switch s[1] {
		case 'a':
			return "Sanskrit", "संस्कृतम्", nil
		case 'c':
			return "Sardinian", "sardu", nil
		case 'd':
			return "Sindhi", "सिन्धी; سنڌي، سندھی\u200e", nil
		case 'e':
			return "Northern Sami", "Davvisámegiella", nil
		case 'g':
			return "Sango", "yângâ tî sängö", nil
		case 'i':
			return "Sinhala; Sinhalese", "සිංහල", nil
		case 'k':
			return "Slovak", "slovenčina; slovenský jazyk", nil
		case 'l':
			return "Slovenian", "slovenski jezik; slovenščina", nil
		case 'm':
			return "Samoan", "gagana faa Samoa", nil
		case 'n':
			return "Shona", "chiShona", nil
		case 'o':
			return "Somali", "Soomaaliga; af Soomaali", nil
		case 'q':
			return "Albanian", "Shqip", nil
		case 'r':
			return "Serbian", "српски језик", nil
		case 's':
			return "Swati", "SiSwati", nil
		case 't':
			return "Sotho, Southern", "Sesotho", nil
		case 'u':
			return "Sundanese", "Basa Sunda", nil
		case 'v':
			return "Swedish", "Svenska", nil
		case 'w':
			return "Swahili", "Kiswahili", nil
		}

	case 't':
		switch s[1] {
		case 'a':
			return "Tamil", "தமிழ்", nil
		case 'e':
			return "Telugu", "తెలుగు", nil
		case 'g':
			return "Tajik", "тоҷикӣ;toğikī;تاجیکی\u200e", nil
		case 'h':
			return "Thai", "ไทย", nil
		case 'i':
			return "Tigrinya", "ትግርኛ", nil
		case 'k':
			return "Turkmen", "Türkmen; Түркмен", nil
		case 'l':
			return "Tagalog", "Wikang Tagalog; ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔", nil
		case 'n':
			return "Tswana", "Setswana", nil
		case 'o':
			return "Tonga (Tonga Islands)", "faka Tonga", nil
		case 'r':
			return "Turkish", "Türkçe", nil
		case 's':
			return "Tsonga", "Xitsonga", nil
		case 't':
			return "Tatar", "татар теле; tatar tele", nil
		case 'w':
			return "Twi", "Twi", nil
		case 'y':
			return "Tahitian", "Reo Tahiti", nil
		}

	case 'u':
		switch s[1] {
		case 'g':
			return "Uighur; Uyghur", "Uyƣurqə; ئۇيغۇرچە\u200e", nil
		case 'k':
			return "Ukrainian", "українська мова", nil
		case 'r':
			return "Urdu", "اردو", nil
		case 'z':
			return "Uzbek", "O‘zbek; Ўзбек; أۇزبېك\u200e", nil
		}

	case 'v':
		switch s[1] {
		case 'e':
			return "Venda", "Tshivenḓa", nil
		case 'i':
			return "Vietnamese", "Tiếng Việt", nil
		case 'o':
			return "Volapük", "Volapük", nil
		}

	case 'w':
		switch s[1] {
		case 'a':
			return "Walloon", "walon", nil
		case 'o':
			return "Wolof", "Wollof", nil
		}

	case 'x':
		switch s[1] {
		case 'h':
			return "Xhosa", "isiXhosa", nil
		}

	case 'y':
		switch s[1] {
		case 'i':
			return "Yiddish", "ייִדיש", nil
		case 'o':
			return "Yoruba", "Yorùbá", nil
		}

	case 'z':
		switch s[1] {
		case 'a':
			return "Zhuang; Chuang", "Saɯ cueŋƅ; Saw cuengh", nil
		case 'h':
			return "Chinese", "中文(Zhōngwén); 汉语; 漢語", nil
		case 'u':
			return "Zulu", "isiZulu", nil
		}
	}
	return "", "",
		fmt.Errorf("%q is not a valid ISO-639-1 language code", s)
}
