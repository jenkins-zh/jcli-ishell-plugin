/**
 * These code lines should be moved into github.com/jenkins-zh/jenkins-cli at sometime
 */
package cmd

import (
	"fmt"
	"github.com/jenkins-zh/jenkins-cli/app/cmd/keyring"
	"github.com/jenkins-zh/jenkins-cli/client"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	appCfg "github.com/jenkins-zh/jenkins-cli/app/config"
)

var config *appCfg.Config
func getCurrentJenkins() (cfg *appCfg.JenkinsServer, err error) {
	if err = loadDefaultConfig(); err == nil {
		cfg = findJenkinsByName(config.Current)
	}
	return
}

func getClient(jenkins *appCfg.JenkinsServer, jClient *client.JenkinsCore) {
	jClient.URL = jenkins.URL
	jClient.UserName = jenkins.UserName
	jClient.Token = jenkins.Token
	jClient.Proxy = jenkins.Proxy
	jClient.ProxyAuth = jenkins.ProxyAuth
	jClient.InsecureSkipVerify = jenkins.InsecureSkipVerify
}

func getCurrentJenkinsAndClient(jClient *client.JenkinsCore) (jenkins *appCfg.JenkinsServer, err error) {
	if jenkins, err = getCurrentJenkins(); err == nil && jenkins != nil {
		getClient(jenkins, jClient)
	}
	return
}

func findJenkinsByName(name string) (jenkinsServer *appCfg.JenkinsServer) {
	if config == nil {
		return
	}

	for _, cfg := range config.JenkinsServers {
		if cfg.Name == name {
			jenkinsServer = &cfg
			break
		}
	}
	return
}

func getDefaultConfigPath() (configPath string, err error) {
	var userHome string
	userHome, err = homedir.Dir()
	if err == nil {
		configPath = fmt.Sprintf("%s/.jenkins-cli.yaml", userHome)
	}
	return
}

func loadDefaultConfig() (err error) {
	var configPath string
	if configPath, err = getDefaultConfigPath(); err == nil {
		if _, err = os.Stat(configPath); err == nil {
			err = loadConfig(configPath)
		}
	}
	return
}

func loadConfig(path string) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal([]byte(content), &config)

		keyring.LoadTokenFromKeyring(config)
	}
	return
}
