package utils

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"strings"
)

func SanitizeDBError(err error) error {
	if err == nil {
		return nil
	}

	errorMsg := err.Error()

	// SQLSTATE Class Check
	switch {
	// Class 23 — Integrity Constraint Violation
	case strings.Contains(errorMsg, "SQLSTATE 23505"): // Unique violation
		return errors.New("a record with this value already exists. Please use a unique value")
	case strings.Contains(errorMsg, "SQLSTATE 23503"): // Foreign key violation
		return errors.New("operation failed due to a foreign key constraint. Please check related records")
	case strings.Contains(errorMsg, "SQLSTATE 23502"): // Not-null violation
		return errors.New("one or more required fields are missing")
	case strings.Contains(errorMsg, "SQLSTATE 23514"): // Check constraint violation
		return errors.New("a field value does not meet the required constraints")

	// Class 22 — Data Exception
	case strings.Contains(errorMsg, "SQLSTATE 22001"): // String data, right truncation
		return errors.New("input data is too long for one or more fields")
	case strings.Contains(errorMsg, "SQLSTATE 22007"): // Invalid datetime format
		return errors.New("invalid date or time format")
	case strings.Contains(errorMsg, "SQLSTATE 22003"): // Numeric value out of range
		return errors.New("numeric value is out of range. Please check your inputs")
	case strings.Contains(errorMsg, "SQLSTATE 22012"): // Division by zero
		return errors.New("attempted division by zero. Please check input values")

	// Class 42 — Syntax Error or Access Rule Violation
	case strings.Contains(errorMsg, "SQLSTATE 42601"): // Syntax error
		return errors.New("there is a syntax error in the SQL statement")
	case strings.Contains(errorMsg, "SQLSTATE 42883"): // Undefined function
		return errors.New("an unsupported operation was attempted in the database")
	case strings.Contains(errorMsg, "SQLSTATE 42P01"): // Undefined table
		return errors.New("database operation failed due to a missing table")
	case strings.Contains(errorMsg, "SQLSTATE 42703"): // Undefined column
		return errors.New("database operation failed due to a missing column")

	// Class 40 — Transaction Rollback
	case strings.Contains(errorMsg, "SQLSTATE 40001"): // Serialization failure
		return errors.New("transaction failed due to a concurrency issue. Please try again")
	case strings.Contains(errorMsg, "SQLSTATE 40003"): // Statement completion unknown
		return errors.New("the database transaction's status is unclear. Please retry the operation")

	// Class 08 — Connection Exception
	case strings.Contains(errorMsg, "SQLSTATE 08001"): // SQL client unable to establish connection
		return errors.New("unable to connect to the database. Please check connection settings")
	case strings.Contains(errorMsg, "SQLSTATE 08003"): // Connection does not exist
		return errors.New("database connection was lost. Please retry the operation")
	case strings.Contains(errorMsg, "SQLSTATE 08006"): // Connection failure
		return errors.New("a connection error occurred. Please verify network or server status")

	// Class 53 — Insufficient Resources
	case strings.Contains(errorMsg, "SQLSTATE 53100"): // Disk full
		return errors.New("database operation failed due to insufficient disk space")
	case strings.Contains(errorMsg, "SQLSTATE 53200"): // Out of memory
		return errors.New("database operation failed due to memory limits being reached")
	case strings.Contains(errorMsg, "SQLSTATE 53300"): // Too many connections
		return errors.New("the database has too many active connections. Please try again later")

	// Class 28 — Invalid Authorization Specification
	case strings.Contains(errorMsg, "SQLSTATE 28000"): // Invalid authorization specification
		return errors.New("database authorization failed. Please verify access credentials")

	// Class XX — Internal Error
	case strings.Contains(errorMsg, "SQLSTATE XX000"): // Internal error
		return errors.New("an internal database error occurred. Please contact support")

	// Additional Known GORM Errors
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.New("no matching record found")
	case errors.Is(err, gorm.ErrInvalidData):
		return errors.New("invalid data provided")
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return errors.New("invalid transaction operation")
	case errors.Is(err, gorm.ErrRegistered):
		return errors.New("database resource already registered")
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return errors.New("unsupported database driver")
	case errors.Is(err, gorm.ErrEmptySlice):
		return errors.New("no data provided for operation")
	}

	// Log the detailed error for internal tracking
	log.Printf("Detailed DB Error: %v\n", err)

	// General error message for unknown errors
	return errors.New("a database error occurred. Please try again later")
}
