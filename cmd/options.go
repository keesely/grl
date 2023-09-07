package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/urfave/cli/v2"
)

type COpts struct {
	Method    string          `flag:"method" alias:"X" default:"GET" usage:"<METHOD>HTTP method"`
	Data      string          `flag:"data" alias:"d" usage:"<DATA>HTTP POST data"`
	Form      string          `flag:"form" alias:"F" usage:"<KEY=VALUE>HTTP POST data"`
	Header    cli.StringSlice `flag:"header" alias:"H" usage:"<HEADER>HTTP header"`
	UserAgent string          `flag:"user-agent" alias:"A" usage:"<STRING>HTTP user agent header"`
	HTTP10    bool            `flag:"http10" usage:"Use HTTP/1.0"`
	HTTP11    bool            `flag:"http11" usage:"Use HTTP/1.1"`
	HTTP2     bool            `flag:"http2" usage:"Use HTTP/2"`
	Insecure  bool            `flag:"insecure" usage:"Ignore certificate validation"`
	Timeout   int             `flag:"timeout" usage:"<SECONDS>Set request timeout"`
	MaxTime   int             `flag:"max-time" usage:"<SECONDS>Set request max time"`
	Proxy     string          `flag:"proxy" usage:"<PROXY_URL>Set proxy address"`
	ProxyUser string          `flag:"proxy-user" usage:"<USER:PASSWORD>Set proxy authentication information"`
	ProxyHead string          `flag:"proxy-header" usage:"<HEADER>Set proxy request header"`
	User      string          `flag:"user" alias:"u" usage:"<USER:PASSWORD>Set authentication information"`
	Head      bool            `flag:"head" alias:"I" usage:"Get response header only"`
	Output    string          `flag:"output" alias:"o" usage:"<FILE>Output response content to file"`
	Verbose   bool            `flag:"verbose" alias:"V" usage:"Output detailed information"`
	Silent    bool            `flag:"silent" alias:"s" usage:"Silent mode, no output"`
	Parallel  int             `flag:"parallel" alias:"p" usage:"<NUM>Set parallel number"`
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
		dType := x.Type

		switch dType {
		case "string":
			flag := makeConvert(x, cli.StringFlag{})
			flag.Destination = x.Value.(*string)
			cmdFlags = append(cmdFlags, &flag)
		case "bool":
			flag := makeConvert(x, cli.BoolFlag{})
			flag.Destination = x.Value.(*bool)
			cmdFlags = append(cmdFlags, &flag)
		case "int", "int64", "int32", "int16", "int8":
			flag := makeConvert(x, cli.IntFlag{})
			flag.Destination = x.Value.(*int)
			cmdFlags = append(cmdFlags, &flag)
		case "float64", "float32":
			flag := makeConvert(x, cli.Float64Flag{})
			flag.Destination = x.Value.(*float64)
			cmdFlags = append(cmdFlags, &flag)
		case "cli.StringSlice":
			fmt.Println("cli.StringSlice Hit")
		default:
			switch x.Value.(type) {
			case *[]string, *cli.StringSlice:
				flag := makeConvert(x, cli.StringSliceFlag{})
				flag.Destination = x.Value.(*cli.StringSlice)
				cmdFlags = append(cmdFlags, &flag)
			default:
				break
			}
			continue
		}
	}
	return cmdFlags
}

type MixedFlag interface {
	cli.StringFlag | cli.BoolFlag | cli.DurationFlag | cli.PathFlag | cli.TimestampFlag |
	cli.IntFlag | cli.Float64Flag |
	cli.MultiStringFlag | cli.MultiInt64Flag | cli.MultiIntFlag | cli.MultiFloat64Flag |
	cli.StringSliceFlag | cli.IntSliceFlag
} 

type MixedFlagValue interface {
	string | bool | 
	int | int64 | int32 | int16 | int8 |
	float64 | float32 |
	cli.Timestamp | cli.StringSlice | cli.IntSlice
}

func makeCliFlag (options TypeDesc) cli.GenericFlag {
	flag := cli.GenericFlag{
		Name: getDescTag(options, "flag", options.Key).(string),
		Usage: getDescTag(options, "usage", "").(string),
	}

	alias, aliases, asOk := getDescTag(options, "alias", "").(string), []string{}, false
	if alias != "" {
		aliases = strings.Split(alias, ",")
		asOk = true
	}
	if asOk == true {
		flag.Aliases = aliases
	}
	return flag
}

func makeConvert[T MixedFlag] (in TypeDesc, out T) T {
	generic := makeCliFlag(in)
	reflected := reflect.ValueOf(&out).Elem()
	for _, key := range []string{
		"Name", "Category", "Useage", "Aliases", "EnvVars",
	} {
		value := reflect.ValueOf(generic).FieldByName(key)
		if value.IsValid() {
			reflected.FieldByName(key).Set(value)
		}
	}
	return out
}

