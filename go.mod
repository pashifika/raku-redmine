module raku-redmine

go 1.17

require (
	fyne.io/fyne/v2 v2.1.2
	github.com/goccy/go-json v0.8.1
	github.com/pashifika/rollingFiles v2.2.1+incompatible
	github.com/pashifika/util v0.2.4
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
	golang.org/x/image v0.0.0-20200430140353-33d19683fad8
	gopkg.in/ini.v1 v1.66.2
	gorm.io/driver/sqlite v1.2.6
	gorm.io/gorm v1.22.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fredbi/uri v0.0.0-20181227131451-3dcfdacbaaf3 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-gl/gl v0.0.0-20210813123233-e4099ee2221f // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211213063430-748e38ca8aec // indirect
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/goki/freetype v0.0.0-20181231101311-fa8a33aabaff // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/srwiley/oksvg v0.0.0-20200311192757-870daf9aa564 // indirect
	github.com/srwiley/rasterx v0.0.0-20200120212402-85cb7272f5e9 // indirect
	github.com/yuin/goldmark v1.3.8 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace fyne.io/fyne/v2 v2.1.2 => ../fyne
