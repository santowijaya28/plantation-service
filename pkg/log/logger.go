package log

func Info(args ...interface{}) {
	sugarLog.Info(args...)
}

func Error(args ...interface{}) {
	sugarLog.Error(args...)
}

func Fatal(args ...interface{}) {
	sugarLog.Fatal(args...)
}

func Panic(args ...interface{}) {
	sugarLog.Panic(args...)
}

func Infof(template string, args ...interface{}) {
	sugarLog.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	sugarLog.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarLog.Fatalf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	sugarLog.Panicf(template, args...)
}

func Sync() {
	sugarLog.Sync()
}
