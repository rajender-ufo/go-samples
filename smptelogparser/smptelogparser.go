package smptelogparser

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
)

type ReportingDevice struct{
	XMLName xml.Name `xml:"reportingDevice"`
	DeviceTypeID string `xml:"DeviceTypeID"`
	DeviceSerialNo string `xml:"DeviceSerial"`
	DeviceName string `xml:"DeviceName"`
	ModelNo string `xml:"ModelNumber"`
}

type LogHeader struct{
	XMLName xml.Name `xml:"LogRecordHeader"`
	EventID string `xml:"EventID"`
	TimeStamp string `xml:"TimeStamp"`
	EventType string `xml:"EventType"`
	EventSeqNo int `xml:"EventSequence"`
}

type EventParameters struct{
	XMLName xml.Name `xml:"Parameters"`
	Params []EventParam `xml:"Parameter"`
}
type EventParam struct{
	XMLName xml.Name `xml:"Parameter"`
	Name string `xml:"Name"`
	Value string `xml:"Value"`
}

type ReferencedID struct{
	XMLName xml.Name `xml:"ReferencedID"`
	ID string `xml:"IDName"`
	Value string `xml:"IDValue"`

}
type ReferencedIDs struct{
	XMLName xml.Name `xml:"ReferencedIDs"`
	Params []ReferencedID `xml:"ReferencedID"`
}

type LogBody struct{
	XMLName xml.Name `xml:"LogRecordBody"`
	EventSubType string `xml:"EventSubType"`
	Parameters EventParameters
	References ReferencedIDs
}

type LogRecord struct{
	XMLName xml.Name `xml:"LogRecordElement"`
	Header LogHeader
	Body LogBody
}


type LogReport struct{
	XMLName xml.Name `xml:"LogReport"`
	RepDate string `xml:"reportDate"`
	RepDevice ReportingDevice
	Logs []LogRecord `xml:"LogRecordElement"`
}

func Parse(absfilepath string) {
	xmlFile, err := os.Open(absfilepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()
	fmt.Println("Successfully opened file "+absfilepath)
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var logReport LogReport
	err= xml.Unmarshal(byteValue, &logReport)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("parsed successfully")
	fmt.Println("Report Date: "+logReport.RepDate)
	fmt.Println("DeviceTypeID: "+logReport.RepDevice.DeviceTypeID)
	fmt.Println("Device Serial Number: "+logReport.RepDevice.DeviceSerialNo)
	fmt.Println("Device Name: "+logReport.RepDevice.DeviceName)
	fmt.Println("Model Number: "+logReport.RepDevice.ModelNo)

	fmt.Println("Count of log records:" + strconv.Itoa(len(logReport.Logs)))
	fmt.Println("Event ID ,  Event Type , Event SubType, TimeStamp    \t, EventSequence No");
	for i:=0;i<len(logReport.Logs);i++ {
		fmt.Printf("%d  %s %s %s[%s]\n",logReport.Logs[i].Header.EventSeqNo,logReport.Logs[i].Header.EventID,logReport.Logs[i].Header.TimeStamp,logReport.Logs[i].Header.EventType,logReport.Logs[i].Body.EventSubType)
		//fmt.Println("Parms: "+strconv.Itoa(len(logReport.Logs[i].Body.Parameters.Params)))
		if(len(logReport.Logs[i].Body.Parameters.Params) > 0){
			fmt.Println("\t DCML Parameters")
			for j:= 0;j< len(logReport.Logs[i].Body.Parameters.Params);j++ {
				fmt.Printf("\t\t%s : %s\n",logReport.Logs[i].Body.Parameters.Params[j].Name , logReport.Logs[i].Body.Parameters.Params[j].Value)
			}
		}
		if(len(logReport.Logs[i].Body.References.Params) >0){
			fmt.Println("\t Referenced IDs")
			for  j :=0; j< len(logReport.Logs[i].Body.References.Params);j++ {
				fmt.Printf("\t\t%s : %s\n",logReport.Logs[i].Body.References.Params[j].ID, logReport.Logs[i].Body.References.Params[j].Value)
			}
		}
	}
}

