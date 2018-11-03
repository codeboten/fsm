package main

import (
	"context"
	"fmt"
	"os"

	"github.com/honeycombio/beeline-go"

	"github.com/looplab/fsm"
)

func getStateMachine(ctx context.Context) *fsm.FSM {
	return fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"scan": func(e *fsm.Event) {
				_, span := beeline.StartSpan(ctx, e.Event)
				defer span.Send()
				fmt.Println("after_scan: " + e.FSM.Current())
			},
			"working": func(e *fsm.Event) {
				_, span := beeline.StartSpan(ctx, e.Event)
				defer span.Send()
				fmt.Println("working: " + e.FSM.Current())
			},
			"situation": func(e *fsm.Event) {
				_, span := beeline.StartSpan(ctx, e.Event)
				defer span.Send()
				fmt.Println("situation: " + e.FSM.Current())
			},
			"finish": func(e *fsm.Event) {
				_, span := beeline.StartSpan(ctx, e.Event)
				defer span.Send()
				fmt.Println("finish: " + e.FSM.Current())
			},
		},
	)
}

func run() {
	ctx, span := beeline.StartSpan(context.Background(), "init")
	defer span.Send()

	fsm := getStateMachine(ctx)

	fmt.Println(fsm.Current())

	err := fsm.Event("scan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + fsm.Current())

	err = fsm.Event("working")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2:" + fsm.Current())

	err = fsm.Event("situation")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3:" + fsm.Current())

	err = fsm.Event("finish")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("4:" + fsm.Current())
	defer beeline.Flush(ctx)
}

func main() {
	beeline.Init(beeline.Config{
		WriteKey:    os.Getenv("HONEYCOMB_KEY"),
		Dataset:     os.Getenv("HONEYCOMB_DATASET"),
		ServiceName: "admiral",
	})
	run()
}
