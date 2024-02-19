package util

import "log"

func LogError(err error, messages ...string) error {
    if err != nil {
        if len(messages) > 0 {
            log.Fatalf("%s: %v", messages[0], err)
        } else {
            log.Fatal(err)
        }
    }
    return err
}

func LogPrint(v ...interface{}) {
    log.Println(v...)
}