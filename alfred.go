package alfred

import (
	"encoding/xml"
	"log"
)

// XMLHeader is definition for XML format.
const XMLHeader = `<?xml version="1.0"?>`

// Item contains value list item.
type Item struct {
	UID          string     `xml:"uid,attr,omitempty"`
	Arg          string     `xml:"arg,attr,omitempty"`
	Valid        bool       `xml:"valid,attr,omitempty"`
	Autocomplete string     `xml:"autocomplete,attr,omitempty"`
	File         string     `xml:"file,attr,omitempty"`
	Type         string     `xml:"type,attr,omitempty"`
	Title        string     `xml:"title"`
	Subtitles    []Subtitle `xml:"subtitle,omitempty"`
	Icons        []Icon     `xml:"icon,omitempty"`
	Texts        []Text     `xml:"text,omitempty"`
}

// Subtitle contains values for Subtitle.
type Subtitle struct {
	Value string `xml:",chardata"`
	Mod   string `xml:"attr,omitempty"`
}

// Text contains values for Text.
type Text struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr,omitempty"`
}

// Icon contains values for Icon.
type Icon struct {
	Value string `xml:",chardata"`
	Type  string `xml:"type,attr,omitempty"`
}

// Marshal output XML for Alfred as string.
func Marshal(items []Item) string {
	str, err := xml.MarshalIndent(
		struct {
			XMLName struct{} `xml:"items"`
			Items   []Item   `xml:"item"`
		}{Items: items}, "", "  ")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return XMLHeader + "\n" + string(str) + "\n"
}

// AddSubtitle add to Subtitles.
func (item *Item) AddSubtitle(value string, mod string) {
	item.Subtitles = append(item.Subtitles, Subtitle{
		Value: value,
		Mod:   mod,
	})
}

// AddText add to Texts.
func (item *Item) AddText(value string, typename string) {
	item.Texts = append(item.Texts, Text{
		Value: value,
		Type:  typename,
	})
}

// AddIcon add to Icons.
func (item *Item) AddIcon(value string, typename string) {
	item.Icons = append(item.Icons, Icon{
		Value: value,
		Type:  typename,
	})
}
