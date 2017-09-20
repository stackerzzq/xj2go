package xj2go

import (
	"fmt"
	"path"
	"testing"
)

func Test_xmlToMap(t *testing.T) {
	fs := []string{
		"./testxml/xl/workbook.xml",
		"./testxml/xl/styles.xml",
		"./testxml/xl/_rels/workbook.xml.rels",
		"./testxml/xl/theme/theme1.xml",
		"./testxml/docProps/app.xml",
		"./testxml/docProps/core.xml",
	}

	for k, v := range fs {
		t.Run("xml to map"+string(k), func(t *testing.T) {
			xj := New(v)
			got, err := xj.xmlToMap("", nil)
			if err != nil {
				t.Errorf("xmlToMap() error = %v", err)
				return
			}
			// fmt.Printf("%#v\r\n", got)
			for key, val := range got {
				fmt.Printf("%s, %#v\r\n", key, val)
				fmt.Println()
				for k, v := range val.(map[string]interface{}) {
					fmt.Printf("%s, %#v\r\n", k, v)
				}
			}
		})
	}
}

func Test_XMLToStruct(t *testing.T) {
	fs := []string{
		"./testxml/xl/workbook.xml",
		"./testxml/xl/styles.xml",
		"./testxml/xl/_rels/workbook.xml.rels",
		"./testxml/xl/theme/theme1.xml",
		"./testxml/docProps/app.xml",
		"./testxml/docProps/core.xml",
	}

	pkg := "excel"
	for k, v := range fs {
		t.Run("leaf paths "+string(k), func(t *testing.T) {
			xj := New(v)
			filename := path.Base(v)
			xj.XMLToStruct(pkg+"/"+filename+".go", pkg)
			fmt.Println()
		})
	}
}
