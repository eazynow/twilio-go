package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"twilio-go/verbs"
)

func main() {
	primary()

	secondary()
	fmt.Println("")
}

func primary() {
	fmt.Println("\nPrimary verbs")

	say := Say{
		Voice:    "man",
		Language: "en-gb",
		Loop:     1,
		Text:     "Hello world!"}

	xmlout, err := xml.MarshalIndent(say, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)

	sip := verbs.Sip{}
	sip.Username = "user1"
	sip.Uri = "Test"
	sip.Password = "pass1"
	sip.Url = "http://www.test.com"
	sip.Method = "POST"

	xmlout, err = xml.MarshalIndent(sip, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")
}

func secondary() {
	fmt.Println("\nSecondary verbs")

	v := verbs.Enqueue{
		Action:        "http://www.something.com",
		Method:        "POST",
		WaitUrl:       "http://www.wait.com",
		WaitUrlMethod: "POST",
		QueueName:     "newqueue"}
	v.Action = "POST"

	xmlout, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Leave{}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Hangup{}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Redirect{Url: "../nextInstructions", Method: "POST"}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Reject{}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Reject{Reason: "not sure"}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	xmlout, err = xml.MarshalIndent(verbs.Pause{Length: 30}, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

}
