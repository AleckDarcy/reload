package observation

import (
	"fmt"
	"github.com/AleckDarcy/reload/core/context_bus/core/encoder"
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
	"github.com/AleckDarcy/reload/core/context_bus/public"
	"os"
	"time"
)

// the Do function
// finalize observation

// Do function returns number of log, trace span and metric entries
func (c *Configure) Do(ed *cb.EventData) (cntL, cntT, cntM int) {
	cntL = (*LoggingConfigure)(c.Logging).Do(ed)
	cntT = (*TracingConfigure)(c.Tracing).Do(ed)
	cntM = len(c.Metrics)
	for _, metric := range c.Metrics {
		(*MetricsConfigure)(metric).Do(ed)
	}

	return
}

func (c *TimestampConfigure) Do() {
	if c != nil {
		// todo default
		return
	}
}

func (c *StackTraceConfigure) Do() {
	if c != nil {
		// todo default
		return
	}
}

func (c *LoggingConfigure) Do(ed *cb.EventData) int {
	if c == nil {
		return 0
	}

	er := ed.Event

	e := newEvent()
	e.buf = encoder.JSONEncoder.BeginObject(e.buf)

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "level")
	e.buf = encoder.JSONEncoder.AppendString(e.buf, "info")

	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "time")

	e.buf = encoder.JSONEncoder.BeginString(e.buf)
	if ts := c.Timestamp; ts == nil {
		e.buf = time.Unix(0, er.When.Time).AppendFormat(e.buf, public.TIME_FORMAT_DEFAULT)
	} else {
		e.buf = time.Unix(0, er.When.Time).AppendFormat(e.buf, ts.Format)
	}
	e.buf = encoder.JSONEncoder.EndString(e.buf)

	// do message
	e.buf = encoder.JSONEncoder.AppendKey(e.buf, "message")
	msg := er.What.Application.GetMessage()
	paths := er.What.Application.GetPaths()
	values := make([]interface{}, len(paths))
	var err error
	for i, path := range paths {
		values[i], err = er.What.GetValue(path)
		if err != nil {
			values[i] = fmt.Sprintf("!error(%s)", err.Error())
		}
	}
	e.buf = encoder.JSONEncoder.AppendString(e.buf, fmt.Sprintf(msg, values...))

	// do tag
	if len(c.Attrs) != 0 {
		e.buf = encoder.JSONEncoder.AppendKey(e.buf, "tags")
		e.buf = DoTagFaster(e.buf, c.Attrs, er)

		//tags := DoTag(c.Attrs, er)
		//e.buf = encoder.JSONEncoder.AppendTags(e.buf, tags)
	}

	e.buf = encoder.JSONEncoder.EndObject(e.buf)
	str := string(e.buf)
	e.finalize()

	(*TimestampConfigure)(c.Timestamp).Do()
	(*StackTraceConfigure)(c.Stacktrace).Do()

	switch c.Out {
	case cb.LogOutType_LogOutType_:
		// omit print
	case cb.LogOutType_Stdout:
		fmt.Fprintln(os.Stdout, str)
	case cb.LogOutType_Stderr:
		fmt.Fprintln(os.Stderr, str)
	case cb.LogOutType_File:

	default:
		fmt.Fprintln(os.Stdout, str)
	}

	return 1
}

func (c *TracingConfigure) Do(ed *cb.EventData) int {
	if c == nil {
		return 0
	}

	if prev := ed.GetPreviousEventData(c.PrevName); prev != nil {
		tags := DoTag(c.Attrs, ed.Event)
		if len(tags) != 0 {
			fmt.Printf("todo tracing span(\"%s\", %s)=%d (from %s to %s)\n",
				c.Name, tags, ed.Event.When.Time-prev.Event.When.Time, prev.Event.Recorder.Name, ed.Event.Recorder.Name)
		} else {
			fmt.Printf("todo tracing span(\"%s\")=%d (from %s to %s)\n",
				c.Name, ed.Event.When.Time-prev.Event.When.Time, prev.Event.Recorder.Name, ed.Event.Recorder.Name)
		}

		// todo push span

		return 1
	}

	fmt.Println("previous event not found", c.PrevName)

	return 0
}

func (c *MetricsConfigure) Do(ed *cb.EventData) int {
	if c == nil {
		return 0
	}

	labels := DoTag(c.Attrs, ed.Event)

	switch c.Type {
	case cb.MetricType_Counter:
		vec := MetricVecStore.getCounter(c.OptsId)
		if vec == nil { // todo report error
			fmt.Println("counter vec not found for Opts", c.OptsId)

			return 0
		}

		fmt.Println(c)
		vec.With(labels).Inc()

		if len(labels) != 0 {
			fmt.Printf("todo metrics Counter(\"%s\", %s)\n", c.Name, labels)
		} else {
			fmt.Printf("todo metrics Counter(\"%s\")\n", c.Name)
		}
	case cb.MetricType_Gauge:
		fmt.Println("todo MetricsConfigure do", c.Name)
	case cb.MetricType_Histogram:
		if prev := ed.GetPreviousEventData(c.PrevName); prev != nil {
			vec := MetricVecStore.getHistogram(c.OptsId)
			if vec == nil { // todo report error
				fmt.Println("histogram vec not found for Opts", c.OptsId)

				return 0
			}

			vec.With(labels).Observe(float64(ed.Event.When.Time-prev.Event.When.Time) / float64(time.Millisecond))

			if len(labels) != 0 {
				fmt.Printf("todo metrics Histogram(\"%s\", %s)=%d (from %s to %s)\n",
					c.Name, labels, ed.Event.When.Time-prev.Event.When.Time, prev.Event.Recorder.Name, ed.Event.Recorder.Name)
			} else {
				fmt.Printf("todo metrics Histogram(\"%s\", {})=%d (from %s to %s)\n",
					c.Name, ed.Event.When.Time-prev.Event.When.Time, prev.Event.Recorder.Name, ed.Event.Recorder.Name)
			}
		} else {
			fmt.Println("previous event not found", c.PrevName)
		}
	case cb.MetricType_Summary:
		fmt.Println("todo MetricsConfigure do", c.Name)
	}

	return 1
}

func DoTag(cfg []*cb.AttributeConfigure, er *cb.EventRepresentation) map[string]string {
	tags := map[string]string{}

	for _, path := range cfg {
		str, err := er.What.GetValue(path.Path)

		if err == nil {
			tags[path.Name] = str
		}
	}

	return tags
}

func DoTagFaster(dst []byte, cfg []*cb.AttributeConfigure, er *cb.EventRepresentation) []byte {
	dst = encoder.JSONEncoder.BeginObject(dst)

	for _, path := range cfg {
		str, err := er.What.GetValue(path.Path)

		if err == nil {
			dst = encoder.JSONEncoder.AppendKey(dst, path.Name)
			dst = encoder.JSONEncoder.AppendString(dst, str)
		}
	}

	return dst
}
