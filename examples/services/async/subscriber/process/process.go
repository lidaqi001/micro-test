package process

import (
	"context"
	"fmt"
	"log"
	"sxx-go-micro/examples/proto/event"
	"sxx-go-micro/plugins/wrapper/trace/jaeger"
	"time"
)

func CallSing(ctx context.Context, event *event.DemoEvent) error {
	time.Sleep(time.Second * 1)
	out := fmt.Sprintf("Got sub：%v", event)
	log.Println(out)
	// 因为异步消息没有response对象
	// 所以手动记录tarce
	sp := jaeger.NewSpan(ctx)
	sp.SetResponse(out, nil, false)
	return nil
}

func SingEvent(ctx context.Context, event *event.DemoEvent) error {
	time.Sleep(time.Second * 1)
	out := fmt.Sprintf("Got sub：%v", event)
	log.Println(out)
	// 因为异步消息没有response对象
	// 所以手动记录tarce
	sp := jaeger.NewSpan(ctx)
	sp.SetResponse(out, nil, false)
	return nil
}