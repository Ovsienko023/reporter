package config

import "flag"

// add flags when needed
var ConfigPathFlag = flag.String("config", DefaultConfigPath, "path to the config file")
