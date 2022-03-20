package pokedex

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/defry256/pokemon-helper/internal/pokemon"
	"github.com/defry256/pokemon-helper/internal/pokemontype"
	"github.com/gocolly/colly/v2"
)

const (
	POKEDEX_SOURCE = "https://pokemondb.net/pokedex/"
)

type Service struct {
	c *colly.Collector
}
type Option func(*Service)

func CollyCollectorOption(c *colly.Collector) Option {
	return func(ps *Service) {
		ps.c = c
	}
}

func defaultCollector() *colly.Collector {
	return colly.NewCollector()
}

func NewService(options ...Option) *Service {
	scraper := &Service{}
	scraper.c = defaultCollector()

	for _, option := range options {
		option(scraper)
	}

	return scraper
}

func (s *Service) GetPokedex(pokemonName string) *pokemon.PokemonData {
	var data *pokemon.PokemonData

	s.c.OnHTML("#main", func(h *colly.HTMLElement) {
		data = &pokemon.PokemonData{}

		data.Name = h.DOM.ChildrenFiltered("h1").First().Text()

		h.DOM.Find(".grid-col.span-md-6.span-lg-4").Each(func(i int, s *goquery.Selection) {
			if s.ChildrenFiltered("h2").Text() == "Pokédex data" {
				s.Find("tr").Each(func(i int, s2 *goquery.Selection) {
					if s2.ChildrenFiltered("th").Text() == "Type" && data.Types == nil {
						types := []pokemontype.IType{}
						fmt.Println(s2.ChildrenFiltered("td").Html())
						for _, pokemonType := range strings.Split(s2.ChildrenFiltered("td").First().Text(), " ") {
							if pokemonType != "" {
								types = append(types, pokemontype.Type(strings.ReplaceAll(pokemonType, "\n", "")))
							}
						}
						data.Types = types
					}
				})
			}
		})
		h.DOM.Find(".grid-col.span-md-12.span-lg-8").Each(func(i int, s *goquery.Selection) {
			if s.ChildrenFiltered("h2").Text() == "Base stats" && data.BaseStatus == nil {
				status := &pokemon.Status{}
				s.Find("tr").Each(func(i int, s2 *goquery.Selection) {
					statusName := s2.ChildrenFiltered("th").Text()
					statusValue := s2.ChildrenFiltered("td").First().Text()
					if statusName == "Attack" {
						status.Attack, _ = strconv.Atoi(statusValue)
					} else if statusName == "Defense" {
						status.Defense, _ = strconv.Atoi(statusValue)
					} else if statusName == "HP" {
						status.HP, _ = strconv.Atoi(statusValue)
					} else if statusName == "Sp. Atk" {
						status.SpecialAttack, _ = strconv.Atoi(statusValue)
					} else if statusName == "Sp. Def" {
						status.SpecialDefense, _ = strconv.Atoi(statusValue)
					} else if statusName == "Speed" {
						status.Speed, _ = strconv.Atoi(statusValue)
					}
				})
				status.CalculateTotal()
				data.BaseStatus = status
			}
		})
	})

	s.c.Visit(POKEDEX_SOURCE + pokemonName)

	return data
}
