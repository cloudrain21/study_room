package countdown

import (
    "time"
)

type Sleeper interface {
    Sleep(time.Duration)
}

type SpyTimeSleeper struct {
    sleptTime time.Duration
}

func (s *SpyTimeSleeper)Sleep(duration time.Duration) {
    s.sleptTime = duration
}

type ConfigurableSleeper struct {
    duration time.Duration
    Sleeper
}

func (c *ConfigurableSleeper)CSleep() {
    c.Sleep(c.duration)
}
