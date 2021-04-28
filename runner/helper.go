package runner

import (
	"fmt"

	constant "github.com/lutece-awesome/osiris-next/constant"
	fm "github.com/lutece-awesome/osiris-next/file"
)

var FileManager, err = fm.NewFileManager(*constant.Work)

func init(){
	if ( err != nil ){
		panic( fmt.Sprintf("can not initialize the file manager with error %v" , err ))
	}
}