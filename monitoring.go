package main

import (
	"strconv"
	"os"
	"strings"
	"time"
	"log"
)

type Monitoring struct {
	conf Config
	got Source
	saved int
	sent int
	dropped int
	lg log.Logger
	ch chan string
}
type Source struct {
	net int
	dir int
	retry int
}

func (m *Monitoring) generateOwnMonitoring(){
	hostname,_ := os.Hostname()
	hostnameForGraphite := strings.Replace(hostname, ".", "_", -1)
	path := m.conf.GrafsyPrefix + "."+ hostnameForGraphite + "." + m.conf.GrafsySuffix + ".grafsy"
	now := strconv.FormatInt(time.Now().Unix(),10)

	monitor_slice := []string{
		path + ".got.net " + strconv.Itoa(m.got.net) + " " + now,
		path + ".got.dir " + strconv.Itoa(m.got.dir) + " " + now,
		path + ".got.retry " + strconv.Itoa(m.got.retry) + " " + now,
		path + ".saved " + strconv.Itoa(m.saved) + " " + now,
		path + ".sent " + strconv.Itoa(m.sent) + " " + now,
		path + ".dropped " + strconv.Itoa(m.dropped) + " " + now,
	}

	for _, metric := range monitor_slice {
		select {
			case m.ch <- metric:
			default:
				m.lg.Printf("Too many metrics in the MON queue! This is very bad")
				m.dropped++
		}
	}

}

func (m *Monitoring) clean() *Monitoring{
	m.saved = 0
	m.sent = 0
	m.dropped = 0
	m.got = Source{0,0,0}
	return m
}

func (m *Monitoring) runMonitoring() {
	for ;; time.Sleep(60*time.Second) {
		m.generateOwnMonitoring()
		*m = *m.clean()
	}
}