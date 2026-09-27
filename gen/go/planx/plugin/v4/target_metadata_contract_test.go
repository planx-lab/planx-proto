package pluginv4_test

import (
	"encoding/json"
	"strings"
	"testing"

	pb "github.com/planx-lab/planx-proto/gen/go/planx/plugin/v4"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestDiscoverSchemaKeepsTargetMetadataOptional(t *testing.T) {
	response := (&pb.DiscoverSchemaResponse{}).ProtoReflect().Descriptor()
	field := response.Fields().ByName(protoreflect.Name("target_metadata"))
	if field == nil || field.Number() != 3 || field.Message() == nil {
		t.Fatal("DiscoverSchemaResponse needs additive target metadata at field 3")
	}
	columns := field.Message().Fields().ByName(protoreflect.Name("columns"))
	keys := field.Message().Fields().ByName(protoreflect.Name("usable_upsert_keys"))
	complete := field.Message().Fields().ByName(protoreflect.Name("usable_upsert_keys_complete"))
	if columns == nil || keys == nil || complete == nil || columns.Cardinality() != protoreflect.Repeated || keys.Cardinality() != protoreflect.Repeated {
		t.Fatal("target metadata must separate column policy, usable upsert keys, and completeness")
	}
}

func TestTargetMetadataUsesCurrentHTTPJSONFieldNames(t *testing.T) {
	response := &pb.DiscoverSchemaResponse{TargetMetadata: &pb.TargetTableMetadata{
		Columns:                  []*pb.TargetColumnPolicy{{Name: "id"}},
		UsableUpsertKeysComplete: true,
	}}
	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{`"target_metadata"`, `"usable_upsert_keys_complete"`} {
		if !strings.Contains(string(encoded), name) {
			t.Fatalf("HTTP JSON lost %s: %s", name, encoded)
		}
	}
	if strings.Contains(string(encoded), `"has_default":false`) {
		t.Fatalf("wire omission of false policy flags changed: %s", encoded)
	}
}
