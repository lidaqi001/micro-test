package service

import (
	"context"
	"github.com/asim/go-micro/v3"
)

type Option func(opts *Options)

type Options struct {
	Context context.Context

	Name      string
	Advertise string
	Init      []micro.Option
	CallFunc  func(service micro.Service)

	// use rabbitmq as the broker driver
	Rabbitmq bool
}

type initKey struct{}
type callFuncKey struct{}
type serviceNameKey struct{}
type advertiseKey struct{}
type rabbitmqKey struct{}

func Name(val string) Option {
	return SetOption(serviceNameKey{}, val)
}

func RabbitmqBroker() Option {
	return SetOption(rabbitmqKey{}, true)
}

// registry node address
func Advertise(val string) Option {
	return SetOption(advertiseKey{}, val)
}

func Init(val []micro.Option) Option {
	return SetOption(initKey{}, val)
}

func CallFunc(val func(service micro.Service)) Option {
	return SetOption(callFuncKey{}, val)
}
