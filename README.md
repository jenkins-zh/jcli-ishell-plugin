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
