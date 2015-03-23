package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	dec := json.NewDecoder(os.Stdin)
	for {
		var kv map[string]interface{}
		if err := dec.Decode(&kv); err != nil {
			log.Println(err)
			return
		}
		for k, v := range kv {
			if k == "_TRANSPORT" {
				switch v {
				case "journal": encodeJournalRecord(kv)
				case "kernel" : encodeKernelRecord(kv)
				case "stdout": encodeStdoutRecord(kv)
				case "syslog": encodeSyslogRecord(kv)
				default: log.Printf("unexpected _TRANSPORT: %v\n", v)
				}
			}
		}
	}
}

func encodeJournalRecord(v map[string]interface{}) {
	m := make(map[string]interface{})
	m["_TRANSPORT"] = "journal"
	m["_HOSTNAME"] = v["_HOSTNAME"]
	m["SYSLOG_IDENTIFIER"] = v["SYSLOG_IDENTIFIER"]
	m["MESSAGE"] = v["MESSAGE"]
	m["_SOURCE_REALTIME_TIMESTAMP"] = v["_SOURCE_REALTIME_TIMESTAMP"]
        enc := json.NewEncoder(os.Stdout)
        if err := enc.Encode(&m); err != nil {
            log.Println(err)
        }
}

func encodeKernelRecord(v map[string]interface{}) {
        m := make(map[string]interface{})
	m["_TRANSPORT"] = "kernel"
        m["_HOSTNAME"] = v["_HOSTNAME"]
        m["SYSLOG_IDENTIFIER"] = v["SYSLOG_IDENTIFIER"]
        m["MESSAGE"] = v["MESSAGE"]
        m["_SOURCE_REALTIME_TIMESTAMP"] = v["_SOURCE_REALTIME_TIMESTAMP"]
        enc := json.NewEncoder(os.Stdout)
        if err := enc.Encode(&m); err != nil {
            log.Println(err)
        }
}

func encodeStdoutRecord(v map[string]interface{}) {
        m := make(map[string]interface{})
	m["_TRANSPORT"] = "stdout"
        m["_HOSTNAME"] = v["_HOSTNAME"]
        m["SYSLOG_IDENTIFIER"] = v["SYSLOG_IDENTIFIER"]
        m["MESSAGE"] = v["MESSAGE"]
        m["_SOURCE_REALTIME_TIMESTAMP"] = v["_SOURCE_REALTIME_TIMESTAMP"]
        enc := json.NewEncoder(os.Stdout)
        if err := enc.Encode(&m); err != nil {
            log.Println(err)
        }
}

func encodeSyslogRecord(v map[string]interface{}) {
        m := make(map[string]interface{})
	m["_TRANSPORT"] = "syslog"
        m["_HOSTNAME"] = v["_HOSTNAME"]
        m["SYSLOG_IDENTIFIER"] = v["SYSLOG_IDENTIFIER"]
        m["MESSAGE"] = v["MESSAGE"]
        m["_SOURCE_REALTIME_TIMESTAMP"] = v["_SOURCE_REALTIME_TIMESTAMP"]
        enc := json.NewEncoder(os.Stdout)
        if err := enc.Encode(&m); err != nil {
            log.Println(err)
        }
}
