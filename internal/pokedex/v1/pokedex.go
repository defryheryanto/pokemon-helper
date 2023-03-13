package pokedex

import (
	"context"
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
	return colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36"),
		colly.AllowURLRevisit(),
	)
}

func NewService(options ...Option) *Service {
	scraper := &Service{}
	scraper.c = defaultCollector()

	for _, option := range options {
		option(scraper)
	}

	return scraper
}

func (s *Service) GetAllPokedex(ctx context.Context, search string) []*pokemon.PokemonData {
	pokemons := []*pokemon.PokemonData{}

	s.c.OnHTML("table[id=pokedex]", func(h *colly.HTMLElement) {
		h.DOM.Find("tr").Each(func(i int, s *goquery.Selection) {
			pokemonName := s.Find(".cell-name").ChildrenFiltered(".ent-name").Text()
			search := h.Request.Ctx.Get("search")
			if pokemonName != "" && strings.Contains(strings.ToUpper(pokemonName), strings.ToUpper(search)) {
				pokemonData := &pokemon.PokemonData{}
				pokemonData.Name = pokemonName

				types := []pokemontype.IType{}
				s.Find(".cell-icon").ChildrenFiltered(".type-icon").Each(func(i int, s2 *goquery.Selection) {
					types = append(types, pokemontype.Type(s2.Text()))
				})
				pokemonData.Types = types

				status := &pokemon.Status{}
				s.Find(".cell-num").Each(func(i int, s2 *goquery.Selection) {
					if i == 1 {
						status.HP, _ = strconv.Atoi(s2.Text())
					} else if i == 2 {
						status.Attack, _ = strconv.Atoi(s2.Text())
					} else if i == 3 {
						status.Defense, _ = strconv.Atoi(s2.Text())
					} else if i == 4 {
						status.SpecialAttack, _ = strconv.Atoi(s2.Text())
					} else if i == 5 {
						status.SpecialDefense, _ = strconv.Atoi(s2.Text())
					} else if i == 6 {
						status.Speed, _ = strconv.Atoi(s2.Text())
					}
				})
				status.CalculateTotal()
				pokemonData.BaseStatus = status
				pokemons = append(pokemons, pokemonData)
			}
		})
	})

	s.c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("search", search)
	})

	s.c.Visit(POKEDEX_SOURCE + "all")
	return pokemons
}

func (s *Service) GetPokedex(ctx context.Context, pokemonName string) *pokemon.PokemonData {
	var data *pokemon.PokemonData

	s.c.OnHTML("#main", func(h *colly.HTMLElement) {
		data = &pokemon.PokemonData{}

		data.Name = h.DOM.ChildrenFiltered("h1").First().Text()

		h.DOM.Find(".grid-col.span-md-6.span-lg-4").Each(func(i int, s *goquery.Selection) {
			if s.ChildrenFiltered("h2").Text() == "Pokédex data" {
				s.Find("tr").Each(func(i int, s2 *goquery.Selection) {
					if s2.ChildrenFiltered("th").Text() == "Type" && data.Types == nil {
						types := []pokemontype.IType{}
						for _, pokemonType := range strings.Split(s2.ChildrenFiltered("td").First().Text(), " ") {
							newType := pokemontype.Type(strings.ReplaceAll(pokemonType, "\n", ""))
							if newType != "" {
								types = append(types, newType)
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
