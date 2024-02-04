package cmd

//TODO: Beautify output with pterm 
//extend usage with API
import (
	"github.com/spf13/cobra"
	//"github.com/pterm/pterm"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
)

type Ability struct{
	Name string `json:"name"`
}
type Abilities struct{
	Ability Ability `json:"ability"`
}

type Pokemon struct{
	Name string `json:"name"`
	Abilities []Abilities `json:"abilities"`
	Types []PokemonTypes `json:"types"`
	Stats []StatType `json:"stats"`
}
type StatType struct{
	Stat Ability `json:"name"`
	Amount int `json:"base_stat"`
}


type PokemonTypes struct{
	Type PokemonType `json:type`
}
type PokemonType struct{
	Name string `json:"name"`
}

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flypoke",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			},
	}

	pokemon := &cobra.Command{
		Use:   "pokemon",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			
			cmd.Println("Flypoke pokemon get " + args[0])

			resp,err := http.Get("https://pokeapi.co/api/v2/pokemon/" + args[0])
			if err != nil{
				log.Fatal(err)
			}
			cmd.Println(resp)
			body,readerr := ioutil.ReadAll(resp.Body)
			if readerr != nil {
				log.Fatal(err)
			}
		
			var pokemon Pokemon
			jsonerr := json.Unmarshal(body,&pokemon)
			if jsonerr != nil{
				cmd.Println(jsonerr)
				log.Fatal(jsonerr)
			}

			cmd.Println("Name from pokemon is: " + pokemon.Name)
			cmd.Println("Abilities")
			for i:= 0; i<len(pokemon.Abilities);i++{
				cmd.Println(pokemon.Abilities[i].Ability.Name)
			}
			for i:=0;i<len(pokemon.Types);i++{
				cmd.Println(pokemon.Types[i].Type.Name)
			}

			},
	}

	stats := &cobra.Command{
		Use:   "stats",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			columns := []string{"Pokemon","Type","Health Points","Attack","Special Attack","Defense","Special Defense","Speed","Evolution"}
			firePokemons := []string{"charmander","charmeleon","charizard"}
			waterPokemons := []string{"squirtle","wartortle","blastoise"}
			grassPokemons := []string{"bulbasaur","ivysaur","venusaur"}

			allPokemons := [][]string{firePokemons,waterPokemons,grassPokemons}

			cmd.Println(strings.Join(columns,","))
			sumpoints := [6]int{}
			for j:=0;j<len(allPokemons);j++{
				for p:=0;p<len(allPokemons[j]);p++{
					resp,err := http.Get("https://pokeapi.co/api/v2/pokemon/" + allPokemons[j][p])
					if err!= nil{
						log.Fatal(err)
					}

					body,ioerr := ioutil.ReadAll(resp.Body)
					if ioerr != nil{
						log.Fatal(ioerr)
					}
					var pokemon Pokemon
					jsonerr := json.Unmarshal(body,&pokemon)
					if jsonerr != nil{
						log.Fatal(jsonerr)
					}
					points := [6]string{}
					for i:=0;i<len(pokemon.Stats);i++{
						points[i] = strconv.Itoa(pokemon.Stats[i].Amount)
						sumpoints[i] += pokemon.Stats[i].Amount
					}
					pokeNameType := []string{pokemon.Name,pokemon.Types[0].Type.Name}
					
					if p == len(allPokemons[j])-1 {
						cmd.Println(strings.Join(pokeNameType,","),",",strings.Join(points[:],","),",","None")
					} else {
						cmd.Println(strings.Join(pokeNameType,","),",",strings.Join(points[:],","),",",allPokemons[j][p+1])
					}

				}
				
				
			}
			sumpointsStr := [6]string{}
			for i:=0;i<len(sumpoints);i++{
				sumpointsStr[i] = strconv.Itoa(sumpoints[i])
			}
			cmd.Println("Aggregated,None,",strings.Join(sumpointsStr[:],","),",None")
		},
	}

	attack := &cobra.Command{
		Use: "attack",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			firePokemons := []string{"🔥","charmander","charmeleon","charizard"}
			waterPokemons := []string{"💧","squirtle","wartortle","blastoise"}
			grassPokemons := []string{"🌵","bulbasaur","ivysaur","venusaur"}

			allPokemons := [][]string{firePokemons,waterPokemons,grassPokemons}

			for i:=0;i<len(allPokemons);i++{
				cmd.Println(allPokemons[i][0] + " Pokemons")
				cmd.Println("----------")
				for p:=1;p<len(allPokemons[i]);p++{
					resp,err := http.Get("https://pokeapi.co/api/v2/pokemon/"+allPokemons[i][p])
					if err!= nil{
						log.Fatal(err)
					}

					body,readerr := ioutil.ReadAll(resp.Body)
					if readerr != nil{
						log.Fatal(readerr)
					}
					var pokemon Pokemon
					jsonerr := json.Unmarshal(body,&pokemon)
					if jsonerr != nil{
						log.Fatal(jsonerr)
					}
					pokeAttack := pokemon.Stats[1].Amount
					cmd.Println(pokemon.Name,pokeAttack)

				}
			}





		},
	}
	// Register your commands here
	//how to add rootcommand
	//cmd.AddCommand(cmd)
	cmd.AddCommand(pokemon)
	cmd.AddCommand(attack)
	cmd.AddCommand(stats)
	return cmd
}

