// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package db

import (
	"bytes"
	"io"

	"github.com/calmh/xdr"
)

/*

fileVersion Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                       Vector Structure                        \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Length of device                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   device (variable length)                    \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct fileVersion {
	Vector version;
	opaque device<>;
}

*/

func (o fileVersion) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.EncodeXDRInto(xw)
}

func (o fileVersion) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o fileVersion) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o fileVersion) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.EncodeXDRInto(xw)
	return []byte(aw), err
}

func (o fileVersion) EncodeXDRInto(xw *xdr.Writer) (int, error) {
	_, err := o.version.EncodeXDRInto(xw)
	if err != nil {
		return xw.Tot(), err
	}
	xw.WriteBytes(o.device)
	return xw.Tot(), xw.Error()
}

func (o *fileVersion) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.DecodeXDRFrom(xr)
}

func (o *fileVersion) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.DecodeXDRFrom(xr)
}

func (o *fileVersion) DecodeXDRFrom(xr *xdr.Reader) error {
	(&o.version).DecodeXDRFrom(xr)
	o.device = xr.ReadBytes()
	return xr.Error()
}

/*

versionList Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      Number of versions                       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\              Zero or more fileVersion Structures              \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct versionList {
	fileVersion versions<>;
}

*/

func (o versionList) EncodeXDR(w io.Writer) (int, error) {
	var xw = xdr.NewWriter(w)
	return o.EncodeXDRInto(xw)
}

func (o versionList) MarshalXDR() ([]byte, error) {
	return o.AppendXDR(make([]byte, 0, 128))
}

func (o versionList) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o versionList) AppendXDR(bs []byte) ([]byte, error) {
	var aw = xdr.AppendWriter(bs)
	var xw = xdr.NewWriter(&aw)
	_, err := o.EncodeXDRInto(xw)
	return []byte(aw), err
}

func (o versionList) EncodeXDRInto(xw *xdr.Writer) (int, error) {
	xw.WriteUint32(uint32(len(o.versions)))
	for i := range o.versions {
		_, err := o.versions[i].EncodeXDRInto(xw)
		if err != nil {
			return xw.Tot(), err
		}
	}
	return xw.Tot(), xw.Error()
}

func (o *versionList) DecodeXDR(r io.Reader) error {
	xr := xdr.NewReader(r)
	return o.DecodeXDRFrom(xr)
}

func (o *versionList) UnmarshalXDR(bs []byte) error {
	var br = bytes.NewReader(bs)
	var xr = xdr.NewReader(br)
	return o.DecodeXDRFrom(xr)
}

func (o *versionList) DecodeXDRFrom(xr *xdr.Reader) error {
	_versionsSize := int(xr.ReadUint32())
	if _versionsSize < 0 {
		return xdr.ElementSizeExceeded("versions", _versionsSize, 0)
	}
	o.versions = make([]fileVersion, _versionsSize)
	for i := range o.versions {
		(&o.versions[i]).DecodeXDRFrom(xr)
	}
	return xr.Error()
}
