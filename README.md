## What
Golang executable programs that reads an yaml file and do the mouse and keyboard actions.

## Why
The input for a computer is all about mouse and keyboard, if have a program does it, we can do a lot of fancy things.
For instance, testing software or website. Yes, for example, Selenium is good for web testing, but what happens for non-web stuff? Some other programs can do it but mostly requires a game engine that requires build etc.

This is a small app with a minimum learning curve. You only need to create a yaml file, without a build or development environment, you can get the job done.

## How
There are two executable files in the project, track is used for recording mouse position, and master is the program to do the job.

Here is an example:
```bash
./master sleep.yaml
```
worked on mac screen 1450x900
```yaml
- mouse: 29 12
- click:
- mouse: 29 196
- click:
```
windows / linux switch program and input username/password and login
```yaml
- keytab: tab alt
- keyup: shift
- sleep: 100
- typestr: <username>
- keytab: tab
- typestr: <password>
- keytab: tab
- keytab: enter
```
You get the idea, it let you control your mouse and keyboard, with an yaml file, without a programming environment.

## Use
You can download the latest build from https://github.com/privapps/mkinputs/tree/latest-binaries
Thanks golang's cross platform build, you should be able to find executables for most popular os/archetecture. You might need cygwin for windows.

If you find this program is very slow in windows, blame on your antivirus software.

## Under the neat
This is project is based on https://github.com/go-vgo/robotgo and it only use subset of it, see for details.
And here for all keys https://github.com/go-vgo/robotgo/blob/master/docs/keys.md