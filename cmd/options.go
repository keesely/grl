package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/urfave/cli/v2"
)

type COpts struct {
	Method    string `flag:"method" alias:"X" default:"GET" usage:"<METHOD>HTTP method"`
	Data      string `flag:"data" alias:"d" usage:"<DATA>HTTP POST data"`
	Form      string `flag:"form" alias:"F" usage:"<KEY=VALUE>HTTP POST data"`
	Header    string `flag:"header" alias:"H" usage:"<HEADER>HTTP header"`
	UserAgent string `flag:"user-agent" alias:"A" usage:"<STRING>HTTP user agent header"`
	HTTP10    bool   `flag:"http10" usage:"Use HTTP/1.0"`
	HTTP11    bool   `flag:"http11" usage:"Use HTTP/1.1"`
	HTTP2     bool   `flag:"http2" usage:"Use HTTP/2"`
	Insecure  bool   `flag:"insecure" usage:"Ignore certificate validation"`
	Timeout   int    `flag:"timeout" usage:"<SECONDS>Set request timeout"`
	MaxTime   int    `flag:"max-time" usage:"<SECONDS>Set request max time"`
	Proxy     string `flag:"proxy" usage:"<PROXY_URL>Set proxy address"`
	ProxyUser string `flag:"proxy-user" usage:"<USER:PASSWORD>Set proxy authentication information"`
	ProxyHead string `flag:"proxy-header" usage:"<HEADER>Set proxy request header"`
	User      string `flag:"user" alias:"u" usage:"<USER:PASSWORD>Set authentication information"`
	Head      bool   `flag:"head" alias:"I" usage:"Get response header only"`
	Output    string `flag:"output" alias:"o" usage:"<FILE>Output response content to file"`
	Verbose   bool   `flag:"verbose" alias:"V" usage:"Output detailed information"`
	Silent    bool   `flag:"silent" alias:"s" usage:"Silent mode, no output"`
	Parallel  int    `flag:"parallel" alias:"p" usage:"<NUM>Set parallel number"`
}

func getTypeKeys(val interface{}) []string {
	var keys []string
	t := reflect.TypeOf(val)
	for i := 0; i < t.NumField(); i++ {
		keys = append(keys, t.Field(i).Name)
	}
	return keys
}

// 解析标签为键值对
func parseTags(tag reflect.StructTag) map[string]string {
	stag := fmt.Sprintf("%v", tag)
	tags := make(map[string]string)
	var keys []string
	// 去除首尾空格
	stag = strings.TrimSpace(stag)
	// 去除首尾引号
	stag = strings.Trim(stag, "`")
	parts := strings.Split(stag, " ")
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if len(kv) == 2 {
			kt := strings.TrimSpace(kv[0])
			keys = append(keys, kt)
		}
	}
	for _, key := range keys {
		tags[key] = tag.Get(key)
	}
	return tags
}

type TypeDesc struct {
	Key   string
	Type  string
	Value interface{}
	Tags  map[string]string
}

func getTypeDesc(val *COpts) []TypeDesc {
	var desc []TypeDesc
	t := reflect.TypeOf(*val)
	v := reflect.ValueOf(val).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		x := TypeDesc{
			Key:  f.Name,
			Type: f.Type.Name(),
			// 对象的地址引用
			Value: v.Field(i).Addr().Interface(),
		}

		x.Tags = parseTags(f.Tag)
		desc = append(desc, x)
	}
	return desc
}

func getDescTag(desc TypeDesc, tag string, def ...interface{}) interface{} {
	if v, ok := desc.Tags[tag]; ok {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

func typeDescToCmdFlags(opts COpts, desc []TypeDesc) []cli.Flag {
	var cmdFlags []cli.Flag
	for _, x := range desc {
		dType, key := x.Type, getDescTag(x, "flag").(string)

		alias, aliases, asOk := getDescTag(x, "alias", "").(string), []string{}, false
		if alias != "" {
			aliases = strings.Split(alias, ",")
			asOk = true
		}

		switch dType {
		case "string":
			flag := &cli.StringFlag{
				Name:        key,
				Value:       getDescTag(x, "default", "").(string),
				Usage:       (getDescTag(x, "usage", "").(string)),
				Destination: x.Value.(*string),
			}
			if asOk == true {
				flag.Aliases = aliases
			}
			cmdFlags = append(cmdFlags, flag)
		case "bool":
			flag := &cli.BoolFlag{
				Name:        key,
				Value:       getDescTag(x, "default", false).(bool),
				Usage:       (getDescTag(x, "usage", "").(string)),
				Destination: x.Value.(*bool),
			}
			if asOk == true {
				flag.Aliases = aliases
			}
			cmdFlags = append(cmdFlags, flag)
		case "int", "int64", "int32", "int16", "int8":
			flag := &cli.IntFlag{
				Name:        key,
				Value:       getDescTag(x, "default", 0).(int),
				Usage:       (getDescTag(x, "usage", "").(string)),
				Destination: x.Value.(*int),
			}
			if asOk == true {
				flag.Aliases = aliases
			}
			cmdFlags = append(cmdFlags, flag)
		case "float64", "float32":
			flag := &cli.Float64Flag{
				Name:        key,
				Value:       getDescTag(x, "default", 0).(float64),
				Usage:       (getDescTag(x, "usage", "").(string)),
				Destination: x.Value.(*float64),
			}
			if asOk == true {
				flag.Aliases = aliases
			}
			cmdFlags = append(cmdFlags, flag)
		default:
			continue
		}
	}
	return cmdFlags
}
