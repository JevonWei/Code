package main

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Users struct {
	ID   int    `json:id`
	Name string `json:name`
}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   Users
}

//type StringMap []map[string]string
type StringMap map[int]Users

func (m StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return nil
	}

	for k, v := range m {
		//e.Encode(xmlMapEntry{XMLName: xml.Name{Local: string(k)}, Value: v})
		//e.Encode(xmlMapEntry{XMLName: k, Value: v})
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: string(k)}, Value: v})
		//e.Encode(v)
	}

	return e.EncodeToken(start.End())
}

func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		//var e xmlMapEntry
		var e Users

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		//(*m)[e.XMLName.Local] = e.Value
		//(*m)[] = e
		for i := 0; i <= len(*m); i++ {
			*m = e
		}

	}
	return nil

}

func main() {
	/*
		json.Marshal   序列化
		json.Unmarshal 反序列化
	*/

	//users := []map[string]string{{"name": "Dan", "Addr": "shanghai"}, {"name": "Ran", "Addr": "Henan"}}
	users := map[int]Users{1: {ID: 1, Name: "DAN"}, 2: {ID: 2, Name: "RAN"}}
	buf, _ := xml.MarshalIndent(StringMap(users), " ", "    ")
	fmt.Println(string(buf))

	stringMap := make(map[int]Users)
	err := xml.Unmarshal(buf, (*StringMap)(&stringMap))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringMap)

}
