package model

import (
	"errors"
	"time"
)

// MetaverseType defines the allowed types for the Metaverse field
type MetaverseType string

const (
	MetaverseDCL     MetaverseType = "dcl"
	MetaverseHyperfy MetaverseType = "hyperfy"
	MetaverseIRL     MetaverseType = "irl"
)

// Attendance represents the structure of our attendance record
type Attendance struct {
	ID         string        `json:"id"`
	Address    string        `json:"address"`
	Created_At time.Time     `json:"created_at"`
	Metaverse  MetaverseType `json:"metaverse"`
	Location   string        `json:"location"`
}

// Validate checks if the Metaverse field has a valid value
func (a *Attendance) Validate() error {
	switch a.Metaverse {
	case MetaverseDCL, MetaverseHyperfy, MetaverseIRL:
		return nil
	default:
		return errors.New("Invalid metaverse type. Use dcl, hyperfy, or irl ")
	}
}
