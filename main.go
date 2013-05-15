package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"twilio-go/nouns"
	"twilio-go/verbs"
)

func main() {
	primary()

	secondary()
	fmt.Println("")
}

func primary() {
	fmt.Println("\nPrimary verbs")

	say := verbs.Say{
		Voice:    "man",
		Language: "en-gb",
		Loop:     1,
		Text:     "Hello world!"}

	xmlout, err := xml.MarshalIndent(say, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	/* Add Play */
	play := verbs.Play{}

	gather := verbs.Gather{}

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	gather.AddSay(say)

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	gather.AddPause(10)

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	gather.RemoveSay()

	gather.AddPlay(play)

	xmlout, err = xml.MarshalIndent(gather, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(xmlout)
	fmt.Println("")

	sip := nouns.Sip{}
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
		WaitUrl:       "http://www.wait.com",
		WaitUrlMethod: "POST",
		QueueName:     "newqueue"}
	v.Action = "http://www.something.com"
	v.Method = "POST"

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
