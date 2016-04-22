//weight strategy
package strategy

import(
	"net/http"
)

type WeightStrategy struct {}


func NewWeightStrategy() (sty *WeightStrategy) {
	sty = &WeightStrategy{}
	return
}

func (sty *WeightStrategy) Init(){

}

func (sty *WeightStrategy) Redirect(w http.ResponseWriter,r *http.Request){

}

