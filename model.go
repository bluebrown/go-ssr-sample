package main

type Link struct {
	Label string
	Href  string
}

type Sidebar struct {
	Title string
	Links []Link
}

type Article struct {
	Title string
	Body  string
}

type Page struct {
	DocTitle  string
	PageTitle string
	Sidebar   Sidebar
	Aside     []Article
}

func newHomePageModel() Page {
	return Page{
		DocTitle:  "SSR",
		PageTitle: "Home",
		Sidebar: Sidebar{
			Title: "SSR",
			Links: []Link{
				{Label: "Home", Href: "/"},
				{Label: "About", Href: "/about"},
			},
		},
		Aside: []Article{
			{
				Title: "Side Stuff",
				Body:  "This is some post on the side, maybe we can display resources here.",
			},
			{
				Title: "More Things",
				Body:  "Render All the posts",
			},
		},
	}
}
