package mock

import (
	"time"

	"github.com/mauleyzaola/issue-tracker/server/domain"
	"github.com/mauleyzaola/issue-tracker/server/operations/database"
	"github.com/mauleyzaola/tecutils"
)

func FileItemNoData() *domain.FileItem {
	t := &domain.FileItem{}
	t.DateCreated = time.Now()
	t.MimeType = "text/html"
	t.Name = tecutils.UUID() + ".txt"
	t.Extension = ".txt"
	return t
}

func FileBytes() []byte {
	content := `this is the content for testing a text file`
	return []byte(content)
}

func FileItemCreate(tx interface{}, op *database.DbOperations, item *domain.FileItem) error {
	return op.FileItemDb.Create(tx, item)
}
