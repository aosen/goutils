// Package config provides for parse config file.
package utils

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	globalContent   map[string]string
	sectionContents map[string]map[string]string
	sections        []string
}

func NewConfig() *Config {
	return &Config{
		globalContent:   make(map[string]string),
		sectionContents: make(map[string]map[string]string),
	}
}

// Load reads config file and returns an initialized Config.
func (self *Config) Load(configFile string) *Config {
	stream, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic("config read file error : " + configFile + "\n")
	}
	self.LoadString(string(stream))
	return self
}

// Save writes config content to a config file.
func (self *Config) Save(configFile string) error {
	return ioutil.WriteFile(configFile, []byte(self.String()), 0777)
}

func (self *Config) Clear() {
	self.globalContent = make(map[string]string)
	self.sectionContents = make(map[string]map[string]string)
	self.sections = nil
}

func (self *Config) LoadString(s string) error {
	lines := strings.Split(s, "\n")
	section := ""
	for _, line := range lines {
		line = strings.Trim(line, emptyRunes)
		if line == "" || line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			if lineLen := len(line); line[lineLen-1] == ']' {
				section = line[1 : lineLen-1]
				sectionAdded := false
				for _, oldSection := range self.sections {
					if section == oldSection {
						sectionAdded = true
						break
					}
				}
				if !sectionAdded {
					self.sections = append(self.sections, section)
				}
				continue
			}
		}
		pair := strings.SplitN(line, "=", 2)
		if len(pair) != 2 {
			return errors.New("bad config file syntax")
		}
		key := strings.Trim(pair[0], emptyRunes)
		value := strings.Trim(pair[1], emptyRunes)
		if section == "" {
			self.globalContent[key] = value
		} else {
			if _, ok := self.sectionContents[section]; !ok {
				self.sectionContents[section] = make(map[string]string)
			}
			self.sectionContents[section][key] = value
		}
	}
	return nil
}

func (self *Config) String() string {
	s := ""
	for key, value := range self.globalContent {
		s += key + "=" + value + "\n"
	}
	for section, content := range self.sectionContents {
		s += "[" + section + "]\n"
		for key, value := range content {
			s += key + "=" + value + "\n"
		}
	}
	return s
}

func (self *Config) StringWithMeta() string {
	s := "__sections__=" + strings.Join(self.sections, ",") + "\n"
	return s + self.String()
}

func (self *Config) GlobalHas(key string) bool {
	if _, ok := self.globalContent[key]; ok {
		return true
	}
	return false
}

func (self *Config) GlobalGet(key string) string {
	return self.globalContent[key]
}

func (self *Config) GlobalSet(key string, value string) {
	self.globalContent[key] = value
}

func (self *Config) GlobalGetInt(key string) int {
	value := self.GlobalGet(key)
	if value == "" {
		return 0
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}

func (self *Config) GlobalGetInt64(key string) int64 {
	value := self.GlobalGet(key)
	if value == "" {
		return 0
	}
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

func (self *Config) GlobalGetDuration(key string) time.Duration {
	return time.Duration(self.GlobalGetInt(key)) * time.Second
}

func (self *Config) GlobalGetDeadline(key string) time.Time {
	return time.Now().Add(time.Duration(self.GlobalGetInt(key)) * time.Second)
}

func (self *Config) GlobalGetSlice(key string, separator string) []string {
	result := []string{}
	value := self.GlobalGet(key)
	if value != "" {
		for _, part := range strings.Split(value, separator) {
			result = append(result, strings.Trim(part, emptyRunes))
		}
	}
	return result
}

func (self *Config) GlobalGetSliceInt(key string, separator string) []int {
	result := []int{}
	value := self.GlobalGetSlice(key, separator)
	for _, part := range value {
		int, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		result = append(result, int)
	}
	return result
}

func (self *Config) GlobalContent() map[string]string {
	return self.globalContent
}

func (self *Config) Sections() []string {
	return self.sections
}

func (self *Config) HasSection(section string) bool {
	if _, ok := self.sectionContents[section]; ok {
		return true
	}
	return false
}

func (self *Config) SectionHas(section string, key string) bool {
	if !self.HasSection(section) {
		return false
	}
	if _, ok := self.sectionContents[section][key]; ok {
		return true
	}
	return false
}

func (self *Config) SectionGet(section string, key string) string {
	if content, ok := self.sectionContents[section]; ok {
		return content[key]
	}
	return ""
}

func (self *Config) SectionSet(section string, key string, value string) {
	if content, ok := self.sectionContents[section]; ok {
		content[key] = value
	} else {
		content = make(map[string]string)
		content[key] = value
		self.sectionContents[section] = content
	}
}

func (self *Config) SectionGetInt(section string, key string) int {
	value := self.SectionGet(section, key)
	if value == "" {
		return 0
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}

func (self *Config) SectionGetDuration(section string, key string) time.Duration {
	return time.Duration(self.SectionGetInt(section, key)) * time.Second
}

func (self *Config) SectionGetSlice(section string, key string, separator string) []string {
	result := []string{}
	value := self.SectionGet(section, key)
	if value != "" {
		for _, part := range strings.Split(value, separator) {
			result = append(result, strings.Trim(part, emptyRunes))
		}
	}
	return result
}

func (self *Config) SectionContent(section string) map[string]string {
	return self.sectionContents[section]
}

func (self *Config) SectionContents() map[string]map[string]string {
	return self.sectionContents
}

const emptyRunes = " \r\t\v"
