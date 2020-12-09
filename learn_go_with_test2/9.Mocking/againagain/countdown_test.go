package countdown

import (
    "testing"
    "time"
)

func TestCountdown(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        sleepTime := 5 * time.Second

        spyTimeSleeper := &SpyTimeSleeper{}
        sleeper := &ConfigurableSleeper{sleepTime, spyTimeSleeper}
        sleeper.CSleep()

        if sleepTime != spyTimeSleeper.sleptTime {
            t.Errorf("got %v want %v", spyTimeSleeper.sleptTime, sleepTime)
        }
    })
}
