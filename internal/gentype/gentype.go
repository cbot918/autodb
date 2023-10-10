package gentype

import (
	"fmt"
	"os"
	"strings"

	"github.com/cbot918/autodb/internal"
)

type Gentype struct {
	DBM     *internal.DBMetadata
	MTs     []ModelType
	pkgName string
}

type ModelType struct {
	FileName string
	Content  string
}

func NewGentype(dbm *internal.DBMetadata) *Gentype {
	return &Gentype{
		DBM:     dbm,
		pkgName: "types",
	}
}

func (g *Gentype) InitContent() error {

	for _, table := range g.DBM.TableMetadata {

		c := "package " + g.pkgName + "\n\n"

		c += "type " + strings.ToUpper(table.Name[:1]) + table.Name[1:] + " struct {\n"
		c += g.GetColumnData(table.ColumnMetadata)
		c += "\n}"

		g.MTs = append(g.MTs, ModelType{
			FileName: table.Name + ".go",
			Content:  c,
		})

	}
	return nil
}

func (g *Gentype) GetColumnData(columns []internal.ColumnMetadata) string {
	rows := ""
	tm := NewTypeMap()
	for _, c := range columns {
		var t string

		t = tm[c.DBType]
		if c.Nullable == "YES" {
			t = "sql.NullString"
		}

		dbtype := fmt.Sprintf(`'db:"%s"'`, c.Name)
		dbtype = strings.Replace(dbtype, "'", "`", 2)
		rows += fmt.Sprintf(`%s %s %s
	`,
			strings.ToUpper(c.Name[:1])+c.Name[1:],
			t,
			dbtype,
		)

	}
	return rows
}

func (g *Gentype) Create() error {

	// create types folder
	err := os.Mkdir(g.pkgName, 0755)
	if err != nil {
		return err
	}

	// generate content
	for _, mt := range g.MTs {
		fd, err := os.Create(mt.FileName)
		if err != nil {
			return err
		}
		defer fd.Close()

		_, err = fd.Write([]byte(mt.Content))
		if err != nil {
			return err
		}

		err = move(mt.FileName, g.pkgName+"/"+mt.FileName)
		if err != nil {
			return err
		}
	}
	return nil
}

func move(source string, dest string) error {
	return os.Rename(source, dest)
}
