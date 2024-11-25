package stat

import (
	"fmt"
	"go-demo-6/configs"
	"go-demo-6/pkg/middleware"
	"go-demo-6/pkg/response"
	"net/http"
	"time"
)

const GroupByDay = "day"
const GroupByMonth = "month"

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}

		fmt.Println(from, to, by)

		stats, err := handler.StatRepository.GetStats(by, from, to)
		if err != nil {
			http.Error(w, "Invalid params", http.StatusBadRequest)
			return
		}

		response.Json(w, stats, 200)
	}
}
