package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/common"
	"ether"
	"contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"strings"

)
var(
	tpl *template.Template
	alloc map[common.Address]core.GenesisAccount
	contract  *contracts.SimpleStorage
	contractAddr common.Address
	sim *backends.SimulatedBackend
	authMain *bind.TransactOpts
	TransOpt []*bind.TransactOpts

	accs map[string]*bind.TransactOpts
	unicAccs []string
	userNumber int
	account *bind.TransactOpts
	dataToTemplate User
)

type User struct {
	Name string
	Visa string
	Nationality string
	Age string
	Speaks string
	MedCondition []string
	MedMedications []string

}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	alloc = make(core.GenesisAlloc)
	TransOpt,sim = ether.CreateAccs(alloc)
	contractAddr, contract = ether.CreateContract(TransOpt[0],sim)
	accs = make(map[string]*bind.TransactOpts)

}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.POST("/userData", userData)
	mux.GET("/userChange", userChange)
	http.ListenAndServe(":8080", mux)

}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req.ParseForm()
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func userData(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	req.ParseForm()
	if len(req.Form)!=1 && len(req.Form)!=0{
		switch{
		case req.Form["visa"][0]!="":
			contract.SetVisa(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["visa"][0])
			fmt.Println("Mining Visa\n")
			sim.Commit()
			fallthrough
		case req.Form["nationality"][0]!="":
			contract.SetNationality(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["nationality"][0])
			fmt.Println("Mining Nationality\n")
			sim.Commit()
			fallthrough
		case req.Form["age"][0]!="":
			contract.SetAge(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["age"][0])
			fmt.Println("Mining Age\n")
			sim.Commit()
			fallthrough
		case req.Form["speaks"][0]!="":
			contract.SetSpeaks(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["speaks"][0])
			fmt.Println("Mining Speaks\n")
			sim.Commit()
			fallthrough
		case req.Form["medCondition"][0]!="":
			contract.SetMedCondition(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["medCondition"][0])
			fmt.Println("Mining MedicalConditions\n")
			sim.Commit()
			fallthrough
		case req.Form["medMedication"][0]!="":
			contract.SetMedMedication(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["medMedication"][0])
			fmt.Println("Mining MedicalMedications\n")
			sim.Commit()
		}

	}else{
		var data string
		if len(req.Form)==0{
			data=dataToTemplate.Name
		} else{
			data = req.Form["name"][0]
		}

		if data=="John Doe"{
			account=TransOpt[0]
		}else{
			inputHex := ether.StringToHex(data)
			account=getAddress(inputHex)

		}


	}



	//Get values from contract
	name, _ := contract.GetName(&bind.CallOpts{
		From:account.From,
	})
	if name ==""{
			contract.SetName(&bind.TransactOpts{
				From:     account.From,
				Signer:   account.Signer,
				GasLimit: uint64(2381623),
				Value:    big.NewInt(0),
			}, req.Form["name"][0])
			fmt.Println("Mining Name\n")
			name = req.Form["name"][0]
			sim.Commit()
	}

	visa, _ := contract.GetVisa(&bind.CallOpts{
		From:account.From,
	})
	nationality, _ := contract.GetNationality(&bind.CallOpts{
		From:account.From,
	})
	age, _ := contract.GetAge(&bind.CallOpts{
		From:account.From,
	})
	speaks, _ := contract.GetSpeaks(&bind.CallOpts{
		From:account.From,
	})
	medConditionStr, _ := contract.GetMedCondition(&bind.CallOpts{
		From:account.From,
	})
	medCondition:=strings.Split(medConditionStr," ")
	medMedicationsStr, _:=contract.GetMedMedication(&bind.CallOpts{
		From:account.From,
	})
	medMedications:=strings.Split(medMedicationsStr," ")







	dataToTemplate=User{name, visa,
						nationality,age,
						speaks,medCondition,
						medMedications,}
	sim.Commit()

	err := tpl.ExecuteTemplate(w, "userData.gohtml", dataToTemplate)
	HandleError(w, err)

}

func userChange(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	err := tpl.ExecuteTemplate(w, "userChange.gohtml", dataToTemplate)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

//getting address for hex string
func getAddress(str string) (*bind.TransactOpts){
	if _,ok:=accs[str];ok{
		return accs[str]
	}
	userNumber+=1
	accs[str]= TransOpt[userNumber]
	return accs[str]
}