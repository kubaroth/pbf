// Copyright 2017 the original author or authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pbfparser

//go:generate stringer -type=MemberType

import (
	"fmt"
	"math"
	"time"
)

// DecimalDegrees is the decimal degree representation of a longitude or latitude.
type DecimalDegrees float64

// NanoDecimalDegrees is one billionth of a degree.
const NanoDecimalDegrees DecimalDegrees = 1e-9

// MemberType is an enumeration of relation types.
type MemberType int

const (
	// NODE denotes that the member is a node
	NODE MemberType = iota

	// WAY denotes that the member is a way
	WAY

	// RELATION denotes that the member is a relation
	RELATION
)

// BoundingBox is simply a bounding box.
type BoundingBox struct {
	Left   DecimalDegrees
	Right  DecimalDegrees
	Top    DecimalDegrees
	Bottom DecimalDegrees
}

// Header is the contents of the OpenStreetMap PBF data file.
type Header struct {
	BoundingBox                      *BoundingBox
	RequiredFeatures                 []string
	OptionalFeatures                 []string
	WritingProgram                   string
	Source                           string
	OsmosisReplicationTimestamp      time.Time
	OsmosisReplicationSequenceNumber int64
	OsmosisReplicationBaseURL        string
}

// Info represents information common to Node, Way, and Relation elements.
type Info struct {
	Version   int32
	UID       int32
	Timestamp time.Time
	Changeset int64
	UserSID   string
	Visible   bool
}

// Node represents a specific point on the earth's surface defined by its
// latitude and longitude. Each node comprises at least an id number and a
// pair of coordinates.
type Node struct {
	ID   int64
	Tags map[string]string
	Info *Info
	Lat  DecimalDegrees
	Lon  DecimalDegrees
}

// Way is an ordered list of between 2 and 2,000 nodes that define a polyline.
type Way struct {
	ID      int64
	Tags    map[string]string
	Info    *Info
	NodeIds []int64
}

// Member represents an element that
type Member struct {
	ID   int64
	Type MemberType
	Role string
}

// Relation is a multi-purpose data structure that documents a relationship
// between two or more data elements (nodes, ways, and/or other relations).
type Relation struct {
	ID      int64
	Tags    map[string]string
	Info    *Info
	Members []Member
}

func (d DecimalDegrees) String() string {
	val := math.Abs(float64(d))
	degrees := int(math.Floor(val))
	minutes := int(math.Floor(60 * (val - float64(degrees))))
	seconds := 3600 * (val - float64(degrees) - (float64(minutes) / 60))

	return fmt.Sprintf("%d\u00B0 %d' %f\"", degrees, minutes, seconds)
}