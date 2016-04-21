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

func (ps *WeightStrategy) Init(){

}

func (ps *WeightStrategy) Redirect(w http.ResponseWriter,r *http.Request){
	return
}

