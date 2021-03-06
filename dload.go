/*
动态加载配置文件, 需要继承DLoad来实现自己的热文件加载类
2016-03-15
@aosen
*/

package goutils

import "sync"

const (
    BUF0 = 0
    BUF1 = 1
)

type DLoad struct {
    buf0   interface{}
    buf1   interface{}
    state  int32 //0            使用0  1 使用1
    locker *sync.Mutex
}

func NewDLoad() *DLoad {
        return &DLoad{
            //默认使用one
            state:  0,
            locker: new(sync.Mutex),
        }
    }

func (self *DLoad) getState() int32 {
    return self.state
}

    //如果state为0 则设置为1 如果为1则设置为0
func (self *DLoad) setState(s int32) {
    self.state = s
}

//动态加载配置文件æ¹法
func (self *DLoad) Load(filename string, fn func(string) (interface{}, error)) error {
    self.locker.Lock()
    defer self.locker.Unlock()
    tmp, err := fn(filename)
    if err != nil {
        return err
    }
    if self.getState() == BUF0 {
        self.buf1 = tmp
        self.setState(BUF1)
    } else if self.getState() == BUF1 {
        self.buf0 = tmp
        self.setState(BUF0)
    }
    return nil
}

//获取配置文件中的map
func (self *DLoad) Get() interface{} {
    if self.getState() == BUF0 {
        return self.buf0
    } else if self.getState() == BUF1 {
        return self.buf1
    } else {
        return self.buf0
    }
}
