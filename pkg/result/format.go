package result

type Format int

const (
	_ Format = iota
	FormatProtobuf
	FormatJSON
)

func (f Format) String() string {
	switch f {
	case FormatJSON:
		return "json"
	case FormatProtobuf:
		return "proto"
	}
	return "unset"
}
