package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			name: "new logger json",
			args: args{
				cfg: Config{
					File: File{
						IsActive: true,
						LogFile:  "log/app.log",
						Format:   "json",
					},
					Console: Console{
						Format: "json",
					},
				},
			},
			want: &logger{
				log: &slog.Logger{
					ChannelName:    "",
					FlushInterval:  0,
					LowerLevelName: false,
					ReportCaller:   false,
					CallerSkip:     0,
					CallerFlag:     0,
					BackupArgs:     false,
					TimeClock:      nil,
					ExitFunc:       nil,
					PanicFunc:      nil,
				},
			},
		},
		{
			name: "new logger text",
			args: args{
				cfg: Config{
					File: File{
						IsActive: true,
						LogFile:  "log/app.log",
						Format:   "text",
					},
					Console: Console{
						Format: "text",
					},
				},
			},
			want: &logger{
				log: &slog.Logger{
					ChannelName:    "",
					FlushInterval:  0,
					LowerLevelName: false,
					ReportCaller:   false,
					CallerSkip:     0,
					CallerFlag:     0,
					BackupArgs:     false,
					TimeClock:      nil,
					ExitFunc:       nil,
					PanicFunc:      nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.cfg)
			assert.IsType(t, tt.want, got)
		})
	}
}

func Test_logger(t *testing.T) {
	fileHandler := handler.MustRotateFile("/tmp/app.log", rotatefile.EveryDay, handler.WithLogLevels(slog.AllLevels))
	fileHandler.SetFormatter(slog.NewJSONFormatter())

	consoleHandler := handler.NewConsoleHandler(slog.AllLevels)
	consoleHandler.SetFormatter(slog.NewJSONFormatter())

	log := slog.New()
	log.AddHandlers(fileHandler, consoleHandler)
	defer log.Close()

	type fields struct {
		log *slog.Logger
	}
	type args struct {
		msg    string
		fields map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "all levels",
			fields: fields{
				log: log,
			},
			args: args{
				msg: "hello world",
				fields: map[string]interface{}{
					"key": map[string]string{
						"child1": "value1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				log: tt.fields.log,
			}
			l.Info(tt.args.msg, tt.args.fields)
			l.Warn(tt.args.msg, tt.args.fields)
			l.Error(tt.args.msg, tt.args.fields)
			l.Debug(tt.args.msg, tt.args.fields)
		})
	}
}
