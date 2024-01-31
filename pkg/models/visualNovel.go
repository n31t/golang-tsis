package models

import (
	"strings"
)

type VisualNovel struct {
	Title       string `json:"title"`
	Developer   string `json:"developer"`
	ReleaseDate string `json:"release_date"`
	Platform    string `json:"platform"`
	Image       string `json:"image"`
}

var visualNovels = []VisualNovel{
	VisualNovel{Title: "Fate/Stay Night", Developer: "Type-Moon", ReleaseDate: "January 30, 2004", Platform: "Windows"},
	VisualNovel{Title: "Fate/Hollow Ataraxia", Developer: "Type-Moon", ReleaseDate: "October 28, 2005", Platform: "Windows"},
	VisualNovel{Title: "Fate/Zero", Developer: "Nitroplus", ReleaseDate: "December 29, 2006", Platform: "Windows"},
	VisualNovel{Title: "Fate/Extra", Developer: "Type-Moon", ReleaseDate: "July 22, 2010", Platform: "PlayStation Portable"},
	VisualNovel{Title: "Fate/Extra CCC", Developer: "Type-Moon", ReleaseDate: "March 28, 2013", Platform: "PlayStation Portable"},
	VisualNovel{Title: "Fate/Extra CCC Fox Tail", Developer: "Type-Moon", ReleaseDate: "March 26, 2013", Platform: "Manga"},
	VisualNovel{Title: "Fate/Extra CCC: SE.RA.PH", Developer: "Type-Moon", ReleaseDate: "July 12, 2017", Platform: "PlayStation Vita"},
	VisualNovel{Title: "Fate/Grand Order", Developer: "Type-Moon", ReleaseDate: "July 30, 2015", Platform: "Android, iOS"},
	VisualNovel{Title: "Fate/Grand Order Arcade", Developer: "Type-Moon", ReleaseDate: "July 26, 2018", Platform: "Arcade"},
	VisualNovel{Title: "Fate/Extella", Developer: "Type-Moon", ReleaseDate: "November 10, 2016", Platform: "PlayStation 4, PlayStation Vita, Microsoft Windows, Nintendo Switch"},
	VisualNovel{Title: "Fate/Extella Link", Developer: "Type-Moon", ReleaseDate: "June 7, 2018", Platform: "PlayStation 4, PlayStation Vita, Microsoft Windows, Nintendo Switch"},
	VisualNovel{Title: "Fate/Extella: The Umbral Star", Developer: "Type-Moon", ReleaseDate: "July 25, 2017", Platform: "PlayStation 4, PlayStation Vita, Microsoft Windows, Nintendo Switch"},
	VisualNovel{Title: "Fate/Unlimited Codes", Developer: "Capcom", ReleaseDate: "June 18, 2009", Platform: "PlayStation Portable, Arcade"},
	VisualNovel{Title: "Fate/tiger colosseum", Developer: "Capcom", ReleaseDate: "September 13, 2007", Platform: "PlayStation Portable"},
	VisualNovel{Title: "Fate/tiger colosseum Upper", Developer: "Capcom", ReleaseDate: "August 28, 2008", Platform: "PlayStation Portable"},

	VisualNovel{Title: "Tsukihime", Developer: "Type-Moon", ReleaseDate: "December 18, 2000", Platform: "Windows"},
	VisualNovel{Title: "Kagetsu Tohya", Developer: "Type-Moon", ReleaseDate: "August 10, 2001", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood", Developer: "Type-Moon", ReleaseDate: "December 2002", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood Re-ACT", Developer: "Type-Moon", ReleaseDate: "March 2004", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood Act Cadenza", Developer: "Type-Moon", ReleaseDate: "August 2005", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood Act Cadenza Ver.B", Developer: "Type-Moon", ReleaseDate: "March 2006", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood Actress Again", Developer: "Type-Moon", ReleaseDate: "August 2008", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood Actress Again Current Code", Developer: "Type-Moon", ReleaseDate: "July 2009", Platform: "Windows"},
	VisualNovel{Title: "Melty Blood Actress Again Current Code Ver. 1.07", Developer: "Type-Moon", ReleaseDate: "December 2010", Platform: "Windows"},

	VisualNovel{Title: "Kara no Kyokai", Developer: "Type-Moon", ReleaseDate: "December 1998", Platform: "Windows"},
	VisualNovel{Title: "Kara no Kyokai: The Garden of Sinners", Developer: "Type-Moon", ReleaseDate: "December 2001", Platform: "Windows"},

	VisualNovel{Title: "Mahotsukai no Yoru", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
	VisualNovel{Title: "Mahotsukai no Yoru: Witch on the Holy Night", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-san-kan - A Sunny Day", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-yon-kan - The Blue Grimoire", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-go-kan - The Eve", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Light novel"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-roku-kan - The First Evening", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-nana-kan - The Rainy Day", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Light novel"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-hachi-kan - The Last Episode", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
	VisualNovel{Title: "Mahotsukai no Yoru: Dai-kyū-kan - The Blue Grimoire", Developer: "Type-Moon", ReleaseDate: "April 12, 2012", Platform: "Windows"},
}

func GetAllVisualNovels() []VisualNovel {
	return visualNovels
}

func GetVisualNovelByTitle(title string) (VisualNovel, bool) {
	for _, novel := range visualNovels {
		novelTitleWithoutSpaces := strings.ReplaceAll(novel.Title, " ", "")
		novelTitleWithoutSpaces = strings.ReplaceAll(novelTitleWithoutSpaces, "/", "")
		if novelTitleWithoutSpaces == title {
			return novel, true
		}
	}
	return VisualNovel{}, false
}
