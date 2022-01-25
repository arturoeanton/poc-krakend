package notify

import (
	"log"

	"github.com/arturoeanton/gocommons/utils"
	"github.com/fsnotify/fsnotify"
)

type ObserverNotify struct {
	Filename string
	Content  string
	dev      bool
	Watcher  *fsnotify.Watcher
	fxWrite  func(observer *ObserverNotify)
	fxCreate func(observer *ObserverNotify)
	fxRemove func(observer *ObserverNotify)
	fxRename func(observer *ObserverNotify)
	fxChmod  func(observer *ObserverNotify)
}

func (o *ObserverNotify) Dev(f bool) {
	o.dev = f
}

func (o *ObserverNotify) FxCreate(fxCreate func(observer *ObserverNotify)) *ObserverNotify {
	o.fxCreate = fxCreate
	return o
}
func (o *ObserverNotify) FxWrite(fxWrite func(observer *ObserverNotify)) *ObserverNotify {
	o.fxWrite = fxWrite
	return o
}
func (o *ObserverNotify) FxRemove(fxRemove func(observer *ObserverNotify)) *ObserverNotify {
	o.fxRemove = fxRemove
	return o
}
func (o *ObserverNotify) FxRename(fxRename func(observer *ObserverNotify)) *ObserverNotify {
	o.fxRename = fxRename
	return o
}
func (o *ObserverNotify) FxChmod(fxChmod func(observer *ObserverNotify)) *ObserverNotify {
	o.fxChmod = fxChmod
	return o
}

func NewObserverNotify(filename string) *ObserverNotify {
	content, _ := utils.FileToString(filename)
	observer := &ObserverNotify{
		Filename: filename,
		Content:  content,
		dev:      false,
		fxWrite:  func(observer *ObserverNotify) {},
		fxCreate: func(observer *ObserverNotify) {},
		fxRemove: func(observer *ObserverNotify) {},
		fxRename: func(observer *ObserverNotify) {},
		fxChmod:  func(observer *ObserverNotify) {},
	}
	return observer
}

func (o *ObserverNotify) Run() {
	go func() {
		var err error
		o.Watcher, err = fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer o.Watcher.Close()

		done := make(chan bool)
		go func() {
			for {
				select {
				case event, ok := <-o.Watcher.Events:
					if !ok {
						return
					}
					switch {
					case event.Op&fsnotify.Write == fsnotify.Write:
						if o.dev {
							log.Println("WRITE", event.Name)
						}
						o.fxWrite(o)
					case event.Op&fsnotify.Create == fsnotify.Create:
						if o.dev {
							log.Println("CREATE", event.Name)
						}
						o.fxCreate(o)
					case event.Op&fsnotify.Remove == fsnotify.Remove:
						if o.dev {
							log.Println("REMOVE", event.Name)
						}
						o.fxRemove(o)
					case event.Op&fsnotify.Rename == fsnotify.Rename:
						if o.dev {
							log.Println("RENAME", event.Name)
						}
						o.fxRename(o)
					case event.Op&fsnotify.Chmod == fsnotify.Chmod:
						if o.dev {
							log.Println("CHMOD", event.Name)
						}
						o.fxChmod(o)
					}
					if o.dev {
						log.Println("event:", event)
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						if o.dev {
							log.Println("modified file:", event.Name)
						}
					}
				case err, ok := <-o.Watcher.Errors:
					if !ok {
						return
					}
					log.Println("error:", err)
				}
			}
		}()

		err = o.Watcher.Add(o.Filename)
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}()
}
