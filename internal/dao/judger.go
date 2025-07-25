// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"spark-oj-server/internal/dao/internal"
)

// judgerDao is the data access object for the table judger.
// You can define custom methods on it to extend its functionality as needed.
type judgerDao struct {
	*internal.JudgerDao
}

var (
	// Judger is a globally accessible object for table judger operations.
	Judger = judgerDao{internal.NewJudgerDao()}
)

// Add your custom methods and functionality below.
