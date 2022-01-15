package snippets

import (
	"go.uber.org/zap"
)

type stackMsg string

func printMultiline() {
	input := "line1\n\tline2\n\tline3\n\tline4\n\t..."
	println(input)

	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()

		r := recover()
		if r != nil {
			s := zap.Stack("stack").String
			logger.Sugar().Info("in recover", zap.Any("recover", r), zap.String("stack", s))
			return
		}
		logger.Info("in defer")
	}()

	panicer()
}

func (s stackMsg) String() string {
	return "test\n\ttest2"
}

func panicer() {

	panic("I'm panic!")
}
