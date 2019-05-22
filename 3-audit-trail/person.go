package main

import (
	"errors"
	"time"
)

type AuditTrail struct {
	Action    string
	Actor     string
	Timestamp time.Time
	Diffs     []Diff
}

type Diff struct {
	Field string
	Old   interface{}
	New   interface{}
}


type Job struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Salary      uint    `json:"salary,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (p Job) Fields() []string {
	return []string{"ID", "Name", "Salary", "CreatedAt"}
}

func (p Job) Type() string {
	return "Job"
}

func (p Job) Values() []interface{} {
	return []interface{}{
		        p.ID,
		      p.Name,
		 p.Salary,
		 p.CreatedAt,
	}
}

type Person struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (p Person) Fields() []string {
	return []string{"ID", "Name", "CreatedAt"}
}

func (p Person) Type() string {
	return "Person"
}

func (p Person) Values() []interface{} {
	return []interface{}{
		        p.ID,
		      p.Name,
		 p.CreatedAt,
	}
}

type Auditor interface {
	Fields() []string
	Type() string
	Values() []interface{}
}

func Compare(a, b Auditor) (at *AuditTrail, err error) {
	at = &AuditTrail{}
	if a.Type() != b.Type() {
		err = errors.New("Different type cannot compare")
		return
	}
	v1, v2 := a.Values(), b.Values()
	for i, key := range a.Fields() {
		if v1[i] != v2[i] {
			at.Diffs = append(at.Diffs, Diff{key, v1[i], v2[i]})
		}
	}
	return at, err
}
