package db

import (
    "database/sql"

    "github.com/rs/zerolog/log"
)

type enhancedStatement struct {
    *sql.Stmt
}

func (stmt *enhancedStatement) Finalize() {
    err := stmt.Close()
    if err != nil {
        log.Error().Err(err).Msg("Unable to close statement")
    }
}

type enhancedResultSet struct {
    *sql.Rows
}

func (rs *enhancedResultSet) fetchRow() (map[string]any, error) {
    columns, err := rs.Columns()
    if err != nil {
        return nil, err
    }

    scanRow := make([]any, len(columns))
    rowPointers := make([]any, len(columns))
    for i, _ := range scanRow {
        rowPointers[i] = &scanRow[i]
    }

    if err := rs.Scan(rowPointers...); err != nil {
        return nil, err
    }

    row := make(map[string]any)
    for i, column := range columns {
        row[column] = scanRow[i]
    }

    return row, nil
}

func (rs *enhancedResultSet) Finalize() {
    err := rs.Close()
    if err != nil {
        log.Error().Err(err).Msg("Unable to close result set")
    }
}
