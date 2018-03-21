package handle


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/1851616111/util/message"
	"superBrain/model"
	"github.com/ethereum/go-ethereum/log"
)

func BegainPlay(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	PlayID := r.FormValue("playId")
	if PlayID == " " {
		log.Error("PlayID is nil")
		return
	}
	PlayInfo, err := model.GetBegainPlay(PlayID)
	if err != nil {
		log.Error("Get Begain Play Info Error")
		return
	}
	if err := json.NewEncoder(w).Encode(PlayInfo); err != nil {
		message.InnerError(w)
		return
	}
	return
}

func EndPlay(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	PlayID := r.FormValue("playId")
	if PlayID == " " {
		log.Error("PlayID is nil")
		return
	}

	Code, err := model.GetEndPlay(PlayID)
	if err != nil {
		log.Error("Get End Play Info Error")
		return
	}

	if err := json.NewEncoder(w).Encode(Code); err != nil {
		message.InnerError(w)
		return
	}
	return
}
