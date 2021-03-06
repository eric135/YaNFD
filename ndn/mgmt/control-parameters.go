/* YaNFD - Yet another NDN Forwarding Daemon
 *
 * Copyright (C) 2020-2021 Eric Newberry.
 *
 * This file is licensed under the terms of the MIT License, as found in LICENSE.md.
 */

package mgmt

import (
	"errors"
	"strconv"

	"github.com/named-data/YaNFD/ndn"
	"github.com/named-data/YaNFD/ndn/tlv"
)

// ControlParameters represents the parameters of a management command.
type ControlParameters struct {
	Name                          *ndn.Name
	FaceID                        *uint64
	URI                           *ndn.URI
	LocalURI                      *ndn.URI
	Origin                        *uint64
	Cost                          *uint64
	Capacity                      *uint64
	Count                         *uint64
	BaseCongestionMarkingInterval *uint64
	DefaultCongestionThreshold    *uint64
	MTU                           *uint64
	Flags                         *uint64
	Mask                          *uint64
	Strategy                      *ndn.Name
	ExpirationPeriod              *uint64
	FacePersistency               *uint64
}

// MakeControlParameters creates an empty ControlParameters.
func MakeControlParameters() *ControlParameters {
	c := new(ControlParameters)
	return c
}

// DecodeControlParameters decodes a ControlParameters from the wire.
func DecodeControlParameters(wire *tlv.Block) (*ControlParameters, error) {
	if wire == nil {
		return nil, errors.New("wire is unset")
	}

	if wire.Type() != tlv.ControlParameters {
		return nil, tlv.ErrUnexpected
	}

	c := new(ControlParameters)

	wire.Parse()
	var err error
	for _, elem := range wire.Subelements() {
		switch elem.Type() {
		case tlv.Name:
			if c.Name != nil {
				return nil, errors.New("duplicate Name")
			}
			c.Name, err = ndn.DecodeName(elem)
			if err != nil {
				return nil, errors.New("unable to decode Name: " + err.Error())
			}
		case tlv.FaceID:
			if c.FaceID != nil {
				return nil, errors.New("duplicate FaceId")
			}
			c.FaceID = new(uint64)
			*c.FaceID, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode FaceId: " + err.Error())
			}
		case tlv.URI:
			if c.URI != nil {
				return nil, errors.New("duplicate Uri")
			}
			c.URI = ndn.DecodeURIString(string(elem.Value()))
			if err != nil {
				return nil, errors.New("unable to decode Uri: " + err.Error())
			}
		case tlv.LocalURI:
			if c.LocalURI != nil {
				return nil, errors.New("duplicate LocalUri")
			}
			c.LocalURI = ndn.DecodeURIString(string(elem.Value()))
			if err != nil {
				return nil, errors.New("unable to decode LocalUri: " + err.Error())
			}
		case tlv.Origin:
			if c.Origin != nil {
				return nil, errors.New("duplicate Origin")
			}
			c.Origin = new(uint64)
			*c.Origin, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode Origin: " + err.Error())
			}
		case tlv.Cost:
			if c.Cost != nil {
				return nil, errors.New("duplicate Cost")
			}
			c.Cost = new(uint64)
			*c.Cost, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode Cost: " + err.Error())
			}
		case tlv.Capacity:
			if c.Capacity != nil {
				return nil, errors.New("duplicate FaceId")
			}
			c.Capacity = new(uint64)
			*c.Capacity, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode Capacity: " + err.Error())
			}
		case tlv.Count:
			if c.Count != nil {
				return nil, errors.New("duplicate Count")
			}
			c.Count = new(uint64)
			*c.Count, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode Count: " + err.Error())
			}
		case tlv.BaseCongestionMarkingInterval:
			if c.BaseCongestionMarkingInterval != nil {
				return nil, errors.New("duplicate BaseCongestionMarkingInterval")
			}
			c.BaseCongestionMarkingInterval = new(uint64)
			*c.BaseCongestionMarkingInterval, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode BaseCongestionMarkingInterval: " + err.Error())
			}
		case tlv.DefaultCongestionThreshold:
			if c.FaceID != nil {
				return nil, errors.New("duplicate DefaultCongestionThreshold")
			}
			c.DefaultCongestionThreshold = new(uint64)
			*c.DefaultCongestionThreshold, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode DefaultCongestionThreshold: " + err.Error())
			}
		case tlv.MTU:
			if c.MTU != nil {
				return nil, errors.New("duplicate Mtu")
			}
			c.MTU = new(uint64)
			*c.MTU, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("Unable to decode Mtu: " + err.Error())
			}
		case tlv.Flags:
			if c.Flags != nil {
				return nil, errors.New("duplicate Flags")
			}
			c.Flags = new(uint64)
			*c.Flags, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode Flags: " + err.Error())
			}
		case tlv.Mask:
			if c.Mask != nil {
				return nil, errors.New("duplicate Mask")
			}
			c.Mask = new(uint64)
			*c.Mask, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode Mask: " + err.Error())
			}
		case tlv.Strategy:
			if c.Strategy != nil {
				return nil, errors.New("duplicate Strategy")
			}
			elem.Parse()
			if len(elem.Subelements()) != 1 {
				return nil, errors.New("unable to decode Strategy: Expecting one subelement, found " + strconv.Itoa(len(elem.Subelements())))
			}
			c.Strategy, err = ndn.DecodeName(elem.Subelements()[0])
			if err != nil {
				return nil, errors.New("unable to decode Strategy: " + err.Error())
			}
		case tlv.ExpirationPeriod:
			if c.ExpirationPeriod != nil {
				return nil, errors.New("duplicate ExpirationPeriod")
			}
			c.ExpirationPeriod = new(uint64)
			*c.ExpirationPeriod, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode ExpirationPeriod: " + err.Error())
			}
		case tlv.FacePersistency:
			if c.FacePersistency != nil {
				return nil, errors.New("duplicate FacePersistency")
			}
			c.FacePersistency = new(uint64)
			*c.FacePersistency, err = tlv.DecodeNNIBlock(elem)
			if err != nil {
				return nil, errors.New("unable to decode FacePersistency: " + err.Error())
			}
		default:
			if tlv.IsCritical(elem.Type()) {
				return nil, tlv.ErrUnrecognizedCritical
			}
		}
	}

	return c, nil
}

// Encode encodes a ControlParameters.
func (c *ControlParameters) Encode() (*tlv.Block, error) {
	wire := tlv.NewEmptyBlock(tlv.ControlParameters)

	if c.Name != nil {
		wire.Append(c.Name.Encode())
	}
	if c.FaceID != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.FaceID, *c.FaceID))
	}
	if c.URI != nil {
		wire.Append(tlv.NewBlock(tlv.URI, []byte(c.URI.String())))
	}
	if c.LocalURI != nil {
		wire.Append(tlv.NewBlock(tlv.LocalURI, []byte(c.LocalURI.String())))
	}
	if c.Origin != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.Origin, *c.Origin))
	}
	if c.Cost != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.Cost, *c.Cost))
	}
	if c.Capacity != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.Capacity, *c.Capacity))
	}
	if c.Count != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.Count, *c.Count))
	}
	if c.BaseCongestionMarkingInterval != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.BaseCongestionMarkingInterval, *c.BaseCongestionMarkingInterval))
	}
	if c.DefaultCongestionThreshold != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.DefaultCongestionThreshold, *c.DefaultCongestionThreshold))
	}
	if c.MTU != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.MTU, *c.MTU))
	}
	if c.Flags != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.Flags, *c.Flags))
	}
	if c.Mask != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.Mask, *c.Mask))
	}
	if c.Strategy != nil {
		strategyBlock := tlv.NewEmptyBlock(tlv.Strategy)
		strategyBlock.Append(c.Strategy.Encode())
		wire.Append(strategyBlock)
	}
	if c.ExpirationPeriod != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.ExpirationPeriod, *c.ExpirationPeriod))
	}
	if c.FacePersistency != nil {
		wire.Append(tlv.EncodeNNIBlock(tlv.FacePersistency, *c.FacePersistency))
	}

	wire.Encode()
	return wire, nil
}
