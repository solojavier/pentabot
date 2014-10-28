# Pentabot

This a game written with Gobot to move sphero using different types of controls.

## Hardware used

* Arduino
* Sphero
* Spark
* LeapMotion
* iPhone
* Pebble
* PS3 Joystick
* LED's
* HC SR04 - Ranging module

## Software used

* Gobot (http://gobot.io)
* Commander (http://commander.io)
* Watchbot (http://watchbot.io)

## How it works

* To get started you need to charge up the sphero by shaking it. It will change color and leds will turn on as power accumulates.
* Once the sphero is charged turns blue and now you can control it using commander app. (Command set: https://gist.githubusercontent.com/solojavier/17ddb42f4ce6d0d8a242/raw/05e0fa63e9b9bad91405a73519ae5f1f4a9dbd16/command_set.json)
* When you get closer to the goal, you will now be able to control it using the Joystick.
* Next time you reach the goal, you will be able to control it with pebble (watchbot accelerometer)
* Finally if you reach the goal again, you will be able to control it with your hands (Leap Motion)

## Configuration

Better configuration options (and instructions) need to be defined. This was developed for Robotops Talk Demo.

If you are interested in running it and need further information, feel free to open an issue.
