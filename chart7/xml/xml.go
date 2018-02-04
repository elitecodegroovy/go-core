package xml

/**
XML fields that have tags containing "-" will not be printed.

If a tag contains "name,attr", it uses name as the attribute name and the field value as the value, like version in
the above example.

If a tag contains ",attr", it uses the field's name as the attribute name and the field value as its value.

If a tag contains ",chardata", it prints character data instead of element.

If a tag contains ",innerxml", it prints the raw value.

If a tag contains ",comment", it prints it as a comment without escaping, so you cannot have "--" in its value.

If a tag contains "omitempty", it omits this field if its value is zero-value, including false, 0, nil pointer or
nil interface, zero length of array, slice, map and string.
 */
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Staff struct {
	XMLName 	xml.Name 	`xml:"staff"`
	ID 			int 		`xml:"id"`
	FirstName 	string 		`xml:"firstname"`
	LastName 	string 		`xml:"lastname"`
	UserName 	string 		`xml:"username"`
}

type Company struct {
	XMLName 	xml.Name 	`xml:"company"`
	Staffs 		[]Staff  	`xml:"staff"`
}

func Write2XMLFile(v Company, filename string)(bool, error) {
	data, err := xml.MarshalIndent(&v, "  ", "    ")
	if err != nil {
		fmt.Printf("xml.MarshalIndent error : %s \n", err.Error())
		return false, err
	}
	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		fmt.Errorf("WriteFile error: %s", err.Error())
		return false, err
	}
	return true , nil
}


func ReadFromXMLFile(filename string)(Company, error) {
	var c Company
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}
	err = xml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

