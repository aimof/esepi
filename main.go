package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

var (
	// score : 一人の使用のみ想定しているため、直接書き換えている。
	// 複数人同時アクセスの場合には適宜対応する。
	score = scoreType{
		results: make([]bool, len(questions)),
	}

	// Result : 一人の使用のみ想定しているため、直接書き換えている。
	Result = ResultType{}

	endChan = make(chan bool)
)

type scoreType struct {
	startTime time.Time
	endTime   time.Time
	results   []bool
}

// ResultType : 結果画面に表示する要素の型
type ResultType struct {
	ID        string
	Accuracy  int
	TimeSpent string
}

func main() {
	srv := &http.Server{Addr: ":8880"}

	mux := http.NewServeMux()

	mux.HandleFunc("/bootstrap.min.css", bootstrapHandler)
	mux.HandleFunc("/", startHandler)
	mux.HandleFunc("/question", questionHandler)
	mux.HandleFunc("/end", endHanlder)

	srv.Handler = mux

	go end(srv)

	log.Fatalln(srv.ListenAndServe())
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("start").Parse(startTemplate)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	tmpl.Execute(w, nil)
}

func questionHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query()["number"]
	var (
		Q   Question
		err error
	)
	if len(number) < 1 {
		w.WriteHeader(500)
		return
	}

	Q.Number, err = strconv.Atoi(number[0])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	Q.Number++

	// 問題番号は1はじまりインデックス
	if Q.Number < 1 || Q.Number > len(questions)+1 {
		w.WriteHeader(500)
		return
	}

	// 計測開始
	if Q.Number == 1 && score.startTime.Before(time.Date(2019, 1, 1, 1, 0, 0, 0, time.UTC)) {
		score.startTime = time.Now()
		id, ok := r.URL.Query()["id"]
		if ok && len(id) == 1 {
			Result.ID = id[0]
		}
	}

	// 答え合わせ
	a, ok := r.URL.Query()["answer"]
	if ok && len(a) == 1 {
		if a[0] == questions[Q.Number-2].Answer {
			score.results[Q.Number-2] = true
		}
	}

	// 最終問題の後の場合resultへ
	if Q.Number == len(questions)+1 {
		resultHandler(w, r)
		return
	}

	Q.Question, Q.Resources, Q.Answer = questions[Q.Number-1].Question, questions[Q.Number-1].Resources, questions[Q.Number-1].Answer

	tmpl, err := template.New("question").Parse(questionTemplate)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	tmpl.Execute(w, Q)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	if score.endTime.Before(time.Date(2019, 1, 1, 1, 0, 0, 0, time.UTC)) {
		score.endTime = time.Now()
	}

	var sum int
	for _, rslt := range score.results {
		if rslt {
			sum++
		}
	}

	Result.Accuracy = sum * 100 / len(score.results)
	Result.TimeSpent = strconv.Itoa(int(math.Ceil(score.endTime.Sub(score.startTime).Seconds())))

	tmpl, err := template.New("result").Parse(resultTemplate)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	tmpl.Execute(w, Result)
}

func endHanlder(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("end").Parse(endTemplate)
	if err != nil {
		w.WriteHeader(500)
		close(endChan)
		return
	}
	tmpl.Execute(w, nil)
	close(endChan)
}

func bootstrapHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, bootstrap)
}

func end(srv *http.Server) {
	<-endChan
	srv.Close()
}
