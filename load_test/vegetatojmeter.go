package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TestResults struct {
	XMLName xml.Name `xml:"testResults"`
	Version string   `xml:"version,attr"`
	Samples []Sample `xml:"httpSample"`
}

type Sample struct {
	Label        string `xml:"lb,attr"`
	TimeStamp    int64  `xml:"ts,attr"`
	Success      bool   `xml:"s,attr"`
	Elapsed      int64  `xml:"t,attr"`
	ResponseCode int    `xml:"rc,attr"`
}

type VegetaResult struct {
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
}

func WriteJMeter(filename string, vegetaResults []VegetaResult) {
	_, label := filepath.Split(filename)
	index := strings.LastIndex(label, ".")
	label = label[:index]

	result := &TestResults{
		Version: "1.2",
		Samples: make([]Sample, len(vegetaResults)),
	}

	for i := 0; i < len(vegetaResults); i++ {
		result.Samples[i].Label = label
		result.Samples[i].TimeStamp = vegetaResults[i].Timestamp.UTC().UnixNano() / int64(time.Millisecond)
		result.Samples[i].Elapsed = vegetaResults[i].Latency / int64(time.Millisecond)
		result.Samples[i].ResponseCode = vegetaResults[i].Code

		if vegetaResults[i].Code > 199 && vegetaResults[i].Code < 300 {
			result.Samples[i].Success = true
		}
	}

	buffer := &bytes.Buffer{}
	buffer.WriteString(xml.Header)

	encoder := xml.NewEncoder(buffer)
	encoder.Indent("", "    ")
	if err := encoder.Encode(result); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(filename, buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func ReadVegeta(filename string) []VegetaResult {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	results := []VegetaResult{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := VegetaResult{}
		json.Unmarshal([]byte(scanner.Text()), &result)
		results = append(results, result)
	}

	return results
}

func main() {
	jsonFilename := ""
	jmeterFilename := ""

	flag.StringVar(&jsonFilename, "vegeta", "", "Vegeta JSON filename")
	flag.StringVar(&jmeterFilename, "jmeter", "", "jMeter output filename")
	flag.Parse()

	vegetaResults := ReadVegeta(jsonFilename)
	WriteJMeter(jmeterFilename, vegetaResults)
}
