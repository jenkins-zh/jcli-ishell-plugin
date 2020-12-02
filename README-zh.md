# jcli-ishell-plugin
这是 [Jenkins CLI](https://github.com/jenkins-zh/jenkins-cli/) 的一个插件，它可以让你以交互式命令的方式管理你的 Jenkins 任务。

# 快速开始
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

# 更多插件
你可以从[这里](https://github.com/jenkins-zh/jcli-plugins)找到更多的插件。
