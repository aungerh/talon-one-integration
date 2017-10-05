# Usage

1. download this repo
2. `cd ~/talon-one-integration`
3. `go run consumer.go`

alternatively: include the package in your source

1. `~> go get github.com/zyxan/talon-one-integration/TalonOneClient`
2. in your source: `import "github.com/zyxan/talon-one-integration/TalonOneClient"`
3. ????
4. profits

# Description

goal: **To get you familiar with Talon.One by interacting with it, and in the process to write some Go code**

Feel free to contact us with any questions at all times, if you think you'd like a more interactive means of communication than email we can create a channel on slack during the week to talk there.

Also you're welcome to check in partial updates on your progress so we can help/guide if needed or if you're struggling with anything in particular, or just to know where you are.

Finally, I'd like to keep the time frame a bit open for this as I understand you may have other things going on and need to get organized, so progress at your own pace. I encourage you to set your own deadline even if you have to move it later (it is totally expected for that to happen since you don't know the system yet).

# Tasks

- [X] set up slack channel
- [X] decide upon a deadline
- [X] major cleanup
- [X] test integration with Talon.one dashboard

# Resources

- https://developers.talon.one/
- https://help.talon.one/
- https://developers.talon.one/platforms/#shop-plugins (assignment 1 resource)

---

# Part 0: Credentials

Accessing the Talon.one dashboard. This frontend is where marketers are expected to create their applications, campaigns and promotion rules.

# Part 1: Writting an integration client in Go

The first program will be an integration client for Talon.One, name this package `TalonOneClient`. focus on the three event calls for:

- `update customer profile`
- `update customer session`
- `send events`

# Part 2: Set up your campaigns to try out your integration

## Pre-conditions

- `Campaign Manager` frontend -> create an Application with 3 different campaigns on it.

## TODO's *(a.k.a: steps needed to test the integration)*

- In the first campaign create a rule that will check that a session was created (use the condition in Events->Session Created) and pick any effect for it

- in the second campaign create a rule that will check for a characteristic of the Customer Profile, again you can pick any effect here.

- Finally before going into the third campaign go to the user menu in the top right, enter the Developer section and in Attributes go to create a new attribute. There create any attribute you prefer for the Event entity (gotcha: in the Event Type field press 'tab' after writing or it will not let you save, this is a bug yes :P ). Now go into the third campaign and make a rule with a condition that checks for that attribute. The attribute corresponds to a custom event created with the Event Type field, this is what you will use to actually send that event when integrating, and the attribute will be in the payload. Pick any effect for this rule.


# Part 3: Test your integration client from a CLI

## Requirements

The following will be to create another Go program that will use the TalonOneClient package to integrate with Talon.One.

This program should be a very simple command line program that will send three hard-coded events to trigger the effects of the conditions of the campaigns previously created, so it will have to update a customer session that fulfills the conditions of the first campaign (by just PUTting a session with a new integration id it will create it and fulfill the rule) and print out the effects, and the same with customer profiles and a custom event.

The program should be very short, requiring no interaction, just sending the events using the client and printing out the effects from the response.

---

# RAW

```
Hi Alex,
  first of all it was very nice to meet you on Friday, thanks for coming in, I had to leave before you but I'm happy to know you had a good time too!

If it's ok with you I'd like us to move forward with the next step, this would be to write some code that will interact with Talon.One.

What we'd want to achieve is to get you familiar with Talon.One by interacting with it, and in the process to write some Go code also.

For this we'd like you to create a Git repo in your GitHub account or somewhere we can see it :)

To clarify just in case you're wondering, none of the code written by you during this step will actually be used for anything without your permission.

Feel free to contact us with any questions at all times, if you think you'd like a more interactive means of communication than email we can create a channel on slack during the week to talk there.

Also you're welcome to check in partial updates on your progress so we can help/guide if needed or if you're struggling with anything in particular, or just to know where you are.

Finally, I'd like to keep the time frame a bit open for this as I understand you may have other things going on and need to get organized, so progress at your own pace. I encourage you to set your own deadline even if you have to move it later (it is totally expected for that to happen since you don't know the system yet).

If all of this sounds good to you we'll jump into the guide.

Now let's get started with Part 0: Check your credentials and access to the docs

  Use the following login to access https://demo.talon.one :
  email: alex@talon.one
  password: kiwi0120

  You are now looking at the initial page of our frontend. This frontend is where marketers are expected to create their applications, campaigns and promotion rules.

  Please have at hand the following documentation as you will refer to it many times (especially the developers one):
  https://developers.talon.one/
  https://help.talon.one/

  Now that you have confirmed you have access to the account and the docs we'll describe the actual task:

  You'll have to write 2 programs using Go language.

Part 1: Write an integration client in Go

  The first program will be an integration client for Talon.One, name this package TalonOneClient.
  I suggest you look at the currently existing plugins in
  https://developers.talon.one/platforms/#shop-plugins

  to get an idea of how it should look, especially the JS client should be a good one to look at first since it's quite simple. Don't worry about implementing the calls for referrals and customer profiles search, only focus on the three event calls for update customer profile, update customer session and send events.

Part 2: Set up your campaigns to try out your integration

  Before moving on please go to the Campaign Manager frontend, create an Application and then 3 different campaigns in it.

  In the first campaign create a rule that will check that a session was created (use the condition in Events->Session Created) and pick any effect for it.

  Then in the second campaign create a rule that will check for a characteristic of the Customer Profile, again you can pick any effect here.

  Finally before going into the third campaign go to the user menu in the top right, enter the Developer section and in Attributes go to create a new attribute. There create any attribute you prefer for the Event entity (gotcha: in the Event Type field press 'tab' after writing or it will not let you save, this is a bug yes :P ). Now go into the third campaign and make a rule with a condition that checks for that attribute. The attribute corresponds to a custom event created with the Event Type field, this is what you will use to actually send that event when integrating, and the attribute will be in the payload. Pick any effect for this rule.

  Now you have all the setup you need to try your integration!

Part 3: Test your integration client from a command line program

  The following will be to create another Go program that will use the TalonOneClient package to integrate with Talon.One.

  This program should be a very simple command line program that will send three hard-coded events to trigger the effects of the conditions of the campaigns previously created, so it will have to update a customer session that fulfills the conditions of the first campaign (by just PUTting a session with a new integration id it will create it and fulfill the rule) and print out the effects, and the same with customer profiles and a custom event.

  The program should be very short, requiring no interaction, just sending the events using the client and printing out the effects from the response.


That's it! Remember it is super fine to check-in partial updates, ask anything you need as much as you need, and keep the communication open :)

I hope you have fun with it.

Best,
Julian.
```
