package configuration

import "flag"

// add flags when needed
var ConfigPathFlag = flag.String("configuration", DefaultConfigPath, "path to the configuration file")
