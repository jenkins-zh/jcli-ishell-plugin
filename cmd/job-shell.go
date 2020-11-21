package cmd

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/jenkins-zh/jenkins-cli/app"
	"github.com/jenkins-zh/jenkins-cli/app/cmd/common"
	jCLI "github.com/jenkins-zh/jenkins-cli/client"
)

const (
	Client  = "client"
	JobName = "jobName"
)

// NewJobCmd create a command to deal with the job
func NewJobCmd(args []string) (shell *ishell.Shell) {
	shell = ishell.New()

	jclient := &jCLI.JobClient{
		JenkinsCore: jCLI.JenkinsCore{
			RoundTripper: nil,
		},
	}
	if _, err := getCurrentJenkinsAndClient(&(jclient.JenkinsCore)); err != nil {
		shell.Println(fmt.Errorf("cannot get the Jenkins Client: %v", err))
		return
	}
	shell.Set(Client, jclient)

	shell.Println("interactive Jenkins job shell")
	shell.AddCmd(&ishell.Cmd{
		Name: "job",
		Help: "set or print current job name",
		Func: func(c *ishell.Context) {
			currentJobName := shell.Get(JobName)
			if len(c.Args) > 0 {
				shell.Set(JobName, c.Args[0])
			} else if currentJobName == nil {
				c.Println("job name cannot be empty")
			} else {
				c.Printf("current job name: %v\n", shell.Get(JobName))
			}
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "search",
		Help: "search all jobs",
		Func: func(c *ishell.Context) {
			client := shell.Get(Client).(*jCLI.JobClient)
			var items []jCLI.JenkinsItem
			var err error

			if items, err = client.Search("", "", 0, 10); err == nil {
				output := common.OutputOption{
					Writer: &ShellWriter{Shell:c,},
					Columns: "Name,DisplayName,Type,URL",
				}
				err = output.OutputV2(items)
			}

			if err != nil {
				c.Println(err)
			}
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "build",
		Help: "trigger current job",
		Func: func(c *ishell.Context) {
			client := shell.Get(Client).(*jCLI.JobClient)
			if err := client.Build(fmt.Sprintf("%v", shell.Get(JobName))); err != nil {
				c.Println(err)
			}
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "context",
		Help: "switch context between different Jenkins",
		Func: func(c *ishell.Context) {
			jenkinsList := []string{}
			for _, cfg := range config.JenkinsServers {
				jenkinsList = append(jenkinsList, cfg.Name)
			}

			choice := c.MultiChoice(jenkinsList, "Which Jenkins do you want to choose?")
			jclient := &jCLI.JobClient{
				JenkinsCore: jCLI.JenkinsCore{
					RoundTripper: nil,
				},
			}
			getClient(&config.JenkinsServers[choice], &(jclient.JenkinsCore))
			shell.Set(Client, jclient)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "current",
		Help: "show the current Jenkins",
		Func: func(c *ishell.Context) {
			c.Println("current Jenkins is", config.Current)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "history",
		Help: "show the history of job builds",
		Func: func(c *ishell.Context) {
			client := shell.Get(Client).(*jCLI.JobClient)
			var builds []*jCLI.JobBuild
			var err error

			if builds, err = client.GetHistory(fmt.Sprintf("%v", shell.Get(JobName))); err == nil {
				output := common.OutputOption{
					Writer: &ShellWriter{Shell:c,},
					Columns: "DisplayName,Building,Result",
				}
				err = output.OutputV2(builds)
			}

			if err != nil {
				c.Print(err)
			}
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "version",
		Help: "show the version of this plugin",
		Func: func(c *ishell.Context) {
			c.Printf("Version: %s\n", app.GetVersion())
			c.Printf("Last Commit: %s\n", app.GetCommit())
			c.Printf("Build Date: %s\n", app.GetDate())
		},
	})
	return
}

// ShellWriter combines ishell and io.Writer
type ShellWriter struct {
	Shell *ishell.Context
}

// Write outputs the data
func (s * ShellWriter) Write(p []byte) (n int, err error) {
	s.Shell.Print(string(p))
	return
}