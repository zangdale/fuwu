package conf

type FuWu struct {
	Version string   `json:"version" yaml:"version"`
	Server  []Server `json:"server" yaml:"server"`
}
type Server struct {
	Name        string `json:"name" yaml:"name"`
	RunCmd      string `json:"run_cmd" yaml:"runCmd"`
	LogMsgCount int32  `json:"log_msg_count" yaml:"logMsgCount"`
	LogDir      string `json:"log_dir" yaml:"logDir"`
	AutoRestart bool   `json:"auto_restart" yaml:"autoRestart"`
	StartRun    bool   `json:"start_run" yaml:"startRun"`
}

func GetFuwu() (*FuWu, error) {
	return &FuWu{}, nil

}
