[![](https://goreportcard.com/badge/jenkins-zh/jcli-ishell-plugin)](https://goreportcard.com/report/jenkins-zh/jcli-ishell-plugin)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/jenkins-zh/jcli-ishell-plugin)
[![GitHub release](https://img.shields.io/github/release/jenkins-zh/jcli-ishell-plugin.svg?label=release)](https://github.com/jenkins-zh/jcli-ishell-plugin/releases/latest)
![GitHub All Releases](https://img.shields.io/github/downloads/jenkins-zh/jcli-ishell-plugin/total)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/jenkins-zh/jcli-ishell-plugin)

# jcli-ishell-plugin
This is a plugin for [Jenkins CLI](https://github.com/jenkins-zh/jenkins-cli/) which allows you to manage your jobs in an interactive way.

# Get started
```
jcli config plugin fetch
jcli config plugin install ishell
➜  ~ jcli ishell
interactive Jenkins job shell
>>> help

Commands:
  build        trigger current job
  clear        clear the screen
  context      switch context between different Jenkins
  current      show the current Jenkins
  exit         exit the program
  help         display help
  history      show the history of job builds
  job          set or print current job name
  search       search all jobs
  version      show the version of this plugin


>>>
```

# More plugins
You can find more plugins from [here](https://github.com/jenkins-zh/jcli-plugins).
