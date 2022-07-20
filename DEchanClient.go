package main


import (
    "encoding/hex"
	"fmt"
    "bufio"
	"log"
    "os"
    "sort"
	"strings"
    //"strconv"
	//"path/filepath"
	//"time"
    "github.com/deroproject/derohe/rpc"
    "github.com/ybbus/jsonrpc"
)

type BoardConfig struct{
    walletRPCClient jsonrpc.RPCClient
    daemonRPCClient jsonrpc.RPCClient
    SCID            string 
}

func main() {
    def := BoardConfig{}
    var daemonEndpoint = "127.0.0.1:40402" //TODO: Add CLI parameters to change these
    def.daemonRPCClient = jsonrpc.NewClient("http://" + daemonEndpoint + "/json_rpc")
    var walletEndpoint = "127.0.0.1:40403"
    def.walletRPCClient = jsonrpc.NewClient("http://" + walletEndpoint + "/json_rpc")
    def.SCID = "a69f718ffce3de8d06aeb78cce450afcf9a639cc5a8168cfc6c0f8f1346723bb"

    scMap := getSCinfo(def)
    
    var running = 1
    for {
        var threads = listThreads(scMap)
        fmt.Print("Type T to display threads, type a thread number to display its messages: ")
        input := bufio.NewScanner(os.Stdin)
        input.Scan()
        if input.Text() == "T"{
            fmt.Println(formatThreads(scMap, threads))
            
        }
        fmt.Println(listAndFormatReplies(scMap, input.Text()))

        if running == 0 {
            break
        }
    }

}

func getSCinfo(m BoardConfig) map[string]interface{} { //
    var scstr *rpc.GetSC_Result
    getSC := rpc.GetSC_Params{SCID: m.SCID, Variables: true}
	err := m.daemonRPCClient.CallFor(&scstr, "getsc", getSC)
	if err != nil {
		log.Printf("[getSCinfo] getting SC tx err %s\n", err)
	}
    return scstr.VariableStringKeys
}

func listThreads (m map[string]interface{}) map[int]string {
    threadList := make(map[int]string)
    var c = 0
    for key := range m {
        if key != "C" && key != "owner" && key != "IsBoard" { //These should be be the only non-post variables :)
            var split = strings.Split(key, ":r")
            if split[1] == "0" {
                threadList[c] = split[0]
            c++
            }
        }
        
    }
    //Sort the threads, oldest to newest
    sortArr := []string{}
    for _, elm := range threadList {
        sortArr = append(sortArr, elm)
    }
    sort.Strings(sortArr)
    var d = 0
    for _, elm := range sortArr{
        threadList[d] = elm
        d++
    }
    
    return threadList

}   

func listAndFormatReplies (m map[string]interface{}, thread string) string {
    replyList := make(map[int]string)
    var c = 0
    for key, message := range m {
        str := fmt.Sprint(message)
        bs, err := hex.DecodeString(str)
        if err != nil {
        	panic(err)
        }
        if key != "C" && key != "owner" && key != "IsBoard" { 
            var split = strings.Split(key, ":r")
            if string(split[1]) == thread {
                replyList[c] = split[0] + " " + string(bs)
            c++
            }
        }
        
    }
    

    //Sort the replies, oldest to newest
    sortArr := []string{}
    for _, elm := range replyList {
        sortArr = append(sortArr, elm)
    }
    sort.Strings(sortArr)
    var d = 0
    for _, elm := range sortArr{
        replyList[d] = elm
        d++
    }
    var finalString = ""
    for _, msg := range replyList {
        finalString += string(msg) + "\n"
    }
    return finalString
}

func formatThreads(m map[string]interface{}, threads map[int]string) string {
    finalString := ""
    for _, key := range threads {
        for key2, msg := range m {
            str := fmt.Sprint(msg)
            bs, err := hex.DecodeString(str)
            if err != nil {
            	panic(err)
            }
            var split = strings.Split(key2, ":r")
            if string(split[0]) == key {
                finalString += string(key) + "- " + string(bs) + "\n"
            }
        }
    }
    return finalString
}