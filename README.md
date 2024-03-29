# plog
Golang log library, It is a plug-in log library, so the name plog is the abbreviation of plugin-in log.


## example
### simple example
```go
import (
	"github.com/tianxingpan/plog"
)

func Example() {
	l := plog.WithFields("uid", "10012")

	l.Trace("helloworld")
	l.Debug("helloworld")
	l.Info("helloworld")
	l.Warn("helloworld")
	l.Error("helloworld")
	l.Tracef("helloworld")
	l.Debugf("helloworld")
	l.Infof("helloworld")
	l.Warnf("helloworld")
	l.Errorf("helloworld")
	// Output:
}
```

### Examples used in services
**Config**
```yaml
# app.yaml
log:
  - writer: file       
    level: debug       
    writer_config:     
      log_path: ./logs/
      filename: app.log
      roll_type: size  
      max_age: 30      
      max_size: 100    
      max_backups: 20  
      compress:  false 
```
**Code**
```go
// test service, file: main.go

package main

import (
    "io/ioutil"

    "github.com/tianxingpan/plog"
    yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Log   plog.Config `yaml:"log" json:"log"`
}

func initConfig(f string) (*Config, error) {
	var conf Config
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func main() {
    conf, err := initConfig('./app.yaml')  // config initial
	if err != nil {
		panic(err.Error())
	}
	if err := plog.Init(conf.Log); err != nil {
		panic(err.Error())
	}
    plog.Info("plog Init ok!")
}
```