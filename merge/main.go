package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/pubgo/xerror"
	"google.golang.org/grpc/keepalive"
	"reflect"
	"time"
)

func Mergo(dst, src interface{}, opts ...func(*mergo.Config)) error {
	opts = append(opts, mergo.WithOverride, mergo.WithTypeCheck)
	return xerror.WrapF(mergo.Merge(dst, src, opts...), "\ndst: %#v\n\nsrc: %#v", dst, src)
}

func Map(dst, src interface{}, opts ...func(*mergo.Config)) error {
	opts = append(opts, mergo.WithOverrideEmptySlice, mergo.WithOverride, mergo.WithTypeCheck)
	return xerror.WrapF(mergo.Map(dst, src, opts...), "\ndst: %#v\n\nsrc: %#v", dst, src)
}

type KeepaliveParams struct {
	MaxConnectionAge      time.Duration `json:"max_connection_age"`
	MaxConnectionAgeGrace time.Duration `json:"max_connection_age_grace"`
	MaxConnectionIdle     time.Duration `json:"max_connection_idle"`
	Time                  time.Duration `json:"time"`
	Timeout               time.Duration `json:"timeout"`
}

func main() {
	var dt = make(map[string]interface{})
	xerror.Exit(Map(&dt, KeepaliveParams{
		MaxConnectionIdle:     2 * time.Second,
		MaxConnectionAge:      2 * time.Second,
		MaxConnectionAgeGrace: 2 * time.Second,
		//Time:                  2 * time.Second,
		//Timeout:               2 * time.Second,
	}))

	fmt.Printf("%#v\n", dt)

	ks := keepalive.ServerParameters{
		MaxConnectionIdle:     3 * time.Second,
		MaxConnectionAge:      3 * time.Second,
		MaxConnectionAgeGrace: 3 * time.Second,
		Time:                  3 * time.Second,
		Timeout:               3 * time.Second,
	}
	xerror.Exit(Map(&ks, dt, func(config *mergo.Config) {
		config.Transformers = timeTransfomer{}
	}))

	fmt.Printf("%#v\n", ks)
}

type timeTransfomer struct {
	overwrite bool
}

func (t timeTransfomer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	return func(dst, src reflect.Value) error {
		if src.IsZero() {
			return nil
		}

		if dst.CanSet() {
			dst.Set(src)
		}
		return nil
	}
}
