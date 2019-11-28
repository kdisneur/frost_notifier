package main

import (
	"context"
	"flag"
	"fmt"
	"frostnotifier/internal"
	"os"
)

func main() {
	var config internal.Config
	parseFlagsAndArguments(&config)

	ctx := context.Background()

	if err := internal.Main(ctx, &config); err != nil {
		fail(err.Error())
	}
}

func parseFlagsAndArguments(c *internal.Config) {
	flag.Usage = func() {
		usage := `Usage: %s [options] countrycode zipcode recipient

Description

  The frost-notifier command line fetches the weather forecast for a night and
  sends a text message if there are risks of frost.

  If the current hour is between 7pm-8am we look for the current "night" (the
  current 7pm-8am time range), if the current hour is after 8am, we look for
  the next "night" (the next 7pm-8am time range).

  The command line keeps a local cache folder storing the last text-message
  sent in order to not resend the text message several times for the same night.

Arguments

  init
    override the default command. it's used to initialize the credentials.json
	file based on user input
  countrycode
	iso country code (e.g. "fr", "be", ...)
  zipcode
	locale zip code (e.g. "59000")
  recipient
	contact information (e.g. +33601010101)

Options

`
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&c.Debug, "v", false, "enable debug mode")
	flag.StringVar(&c.ConfigFile, "f", "credentials.json", "path to the JSON credentials file")
	flag.StringVar(&c.Language, "l", "en", "language used to send the notifications")

	flag.Parse()

	if flag.Arg(0) == "init" {
		if err := internal.InitializeCredentialsCmd(c.ConfigFile); err != nil {
			fail(err.Error())
		}

		os.Exit(0)
	}

	if err := c.Load(); err != nil {
		fail(fmt.Sprintf("invalid configuration: %s", err.Error()))
	}

	if flag.NArg() != 3 {
		failWithUsage("recipient number is a required argument")
	}

	c.CountryCode = flag.Arg(0)
	c.ZipCode = flag.Arg(1)
	c.Recipient = flag.Arg(2)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func failWithUsage(message string) {
	fmt.Fprintln(os.Stderr, message)
	flag.Usage()
	os.Exit(1)
}
