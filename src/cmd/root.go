package cmd

//TODO: 
//extend usage with API
import (
	"github.com/spf13/cobra"
	"github.com/pterm/pterm"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
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
			
			resp,err := http.Get("https://pokeapi.co/api/v2/pokemon/" + args[0])
			if err != nil{
				log.Fatal(err)
			}

			body,readerr := ioutil.ReadAll(resp.Body)
			if readerr != nil {
				log.Fatal(err)
			}
		
			var pokemon Pokemon
			jsonerr := json.Unmarshal(body,&pokemon)
			if jsonerr != nil{
				log.Fatal(jsonerr)
			}
			var pokeAbilities string = ""
			for i:= 0; i<len(pokemon.Abilities);i++{
				if i<len(pokemon.Abilities)-1{
					pokeAbilities+=pokemon.Abilities[i].Ability.Name+","
				}else{
					pokeAbilities+=pokemon.Abilities[i].Ability.Name
				}
			}
			var pokeTypes string = ""
			for i:=0;i<len(pokemon.Types);i++{
				if i<len(pokemon.Types)-1{
					pokeTypes += pokemon.Types[i].Type.Name + ","
				}else{
					pokeTypes+= pokemon.Types[i].Type.Name
				}
			}

			
			tableData := pterm.TableData{
				{"Name","Types","Abilities"},
				{pokemon.Name,pokeTypes,pokeAbilities},

			}
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()

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
			Data := [11][]string{} 
			Data[0] = columns
			sumpoints := [6]int{}
			datacount := 1
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
					Data[datacount] = append(Data[datacount][:],pokemon.Name,pokemon.Types[0].Type.Name)
					
					for i:=0;i<len(pokemon.Stats);i++{
						Data[datacount] = append(Data[datacount][:],strconv.Itoa(pokemon.Stats[i].Amount))
						sumpoints[i] += pokemon.Stats[i].Amount
					}

					if p == len(allPokemons[j])-1 {
						Data[datacount] = append(Data[datacount][:],"None")
					} else {
						Data[datacount] = append(Data[datacount][:],allPokemons[j][p+1])
					}
					datacount++
				}
			}
			Data[datacount] = append(Data[datacount][:],"Aggregated","None")
			for i:=0;i<len(sumpoints);i++{
				Data[datacount] =append(Data[datacount][:],strconv.Itoa(sumpoints[i]))
			}
			Data[datacount] = append(Data[datacount][:],"None")
			tableData := pterm.TableData{
				Data[0],
				Data[1],
				Data[2],
				Data[3],
				Data[4],
				Data[5],
				Data[6],
				Data[7],
				Data[8],
				Data[9],
				Data[10],
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
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
			
				data := [4][2]string{}
				data[0] = [2]string{allPokemons[i][0] + " Pokemons","Attack"}
				
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
					data[p] = [2]string{pokemon.Name,strconv.Itoa(pokeAttack)}
				}
				tableData := pterm.TableData{
					data[0][:],data[1][:],data[2][:],data[3][:],
				}
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
			}

		},
	}

	cmd.AddCommand(pokemon)
	cmd.AddCommand(attack)
	cmd.AddCommand(stats)
	return cmd
}

