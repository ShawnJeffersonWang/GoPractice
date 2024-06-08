//package kitex
//
//import (
//	kitex "GoAlgorithm"
//	kargs "GoAlgorithm/args"
//	"GoAlgorithm/log"
//	"GoAlgorithm/thriftgo"
//	"bytes"
//	"io/ioutil"
//	"os"
//	"os/exec"
//	"strings"
//)
//
//var args kargs.Arguments
//
//func main() {
//
//	// run as kitex
//	args.ParseArgs(kitex.Version)
//
//	out := new(bytes.Buffer)
//	cmd := args.BuildCmd(out)
//	err := cmd.Run()
//	if err != nil {
//		if args.Use != "" {
//			out := strings.TrimSpace(out.String())
//			if strings.HasSuffix(out, thriftgo.TheUseOptionMessage) {
//				goto NormalExit
//			}
//		}
//		os.Exit(1)
//	}
//NormalExit:
//	if args.IDLType == "thrift" {
//		cmd := "go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0"
//		argv := strings.Split(cmd, " ")
//		err := exec.Command(argv[0], argv[1:]...).Run()
//
//		res := "Done"
//		if err != nil {
//			res = err.Error()
//		}
//		log.Warn("Adding apache/thrift@v0.13.0 to go.mod for generated code ..........", res)
//	}
//
//	// remove kitex.yaml generated from v0.4.4 which is renamed as kitex_info.yaml
//	if args.ServiceName != "" {
//		DeleteKitexYaml()
//	}
//}
//
//func DeleteKitexYaml() {
//	// try to read kitex.yaml
//	data, err := ioutil.ReadFile("kitex.yaml")
//	if err != nil {
//		if !os.IsNotExist(err) {
//			log.Warn("kitex.yaml, which is used to record tool info, is deprecated, it's renamed as kitex_info.yaml, you can delete it or ignore it.")
//		}
//		return
//	}
//	// if kitex.yaml exists, check content and delete it.
//	if strings.HasPrefix(string(data), "kitexinfo:") {
//		err = os.Remove("kitex.yaml")
//		if err != nil {
//			log.Warn("kitex.yaml, which is used to record tool info, is deprecated, it's renamed as kitex_info.yaml, you can delete it or ignore it.")
//		} else {
//			log.Warn("kitex.yaml, which is used to record tool info, is deprecated, it's renamed as kitex_info.yaml, so it's automatically deleted now.")
//		}
//	}
//}
