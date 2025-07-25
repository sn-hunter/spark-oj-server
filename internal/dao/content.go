// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"spark-oj-server/internal/dao/internal"
)

// contentDao is the data access object for the table content.
// You can define custom methods on it to extend its functionality as needed.
type contentDao struct {
	*internal.ContentDao
}

var (
	// Content is a globally accessible object for table content operations.
	Content = contentDao{internal.NewContentDao()}
)

// Add your custom methods and functionality below.
