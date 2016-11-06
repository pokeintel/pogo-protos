// Code generated by protoc-gen-go.
// source: inventory_item.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ItemId int32

const (
	ItemId_ITEM_UNKNOWN                   ItemId = 0
	ItemId_ITEM_POKE_BALL                 ItemId = 1
	ItemId_ITEM_GREAT_BALL                ItemId = 2
	ItemId_ITEM_ULTRA_BALL                ItemId = 3
	ItemId_ITEM_MASTER_BALL               ItemId = 4
	ItemId_ITEM_POTION                    ItemId = 101
	ItemId_ITEM_SUPER_POTION              ItemId = 102
	ItemId_ITEM_HYPER_POTION              ItemId = 103
	ItemId_ITEM_MAX_POTION                ItemId = 104
	ItemId_ITEM_REVIVE                    ItemId = 201
	ItemId_ITEM_MAX_REVIVE                ItemId = 202
	ItemId_ITEM_LUCKY_EGG                 ItemId = 301
	ItemId_ITEM_INCENSE_ORDINARY          ItemId = 401
	ItemId_ITEM_INCENSE_SPICY             ItemId = 402
	ItemId_ITEM_INCENSE_COOL              ItemId = 403
	ItemId_ITEM_INCENSE_FLORAL            ItemId = 404
	ItemId_ITEM_TROY_DISK                 ItemId = 501
	ItemId_ITEM_X_ATTACK                  ItemId = 602
	ItemId_ITEM_X_DEFENSE                 ItemId = 603
	ItemId_ITEM_X_MIRACLE                 ItemId = 604
	ItemId_ITEM_RAZZ_BERRY                ItemId = 701
	ItemId_ITEM_BLUK_BERRY                ItemId = 702
	ItemId_ITEM_NANAB_BERRY               ItemId = 703
	ItemId_ITEM_WEPAR_BERRY               ItemId = 704
	ItemId_ITEM_PINAP_BERRY               ItemId = 705
	ItemId_ITEM_SPECIAL_CAMERA            ItemId = 801
	ItemId_ITEM_INCUBATOR_BASIC_UNLIMITED ItemId = 901
	ItemId_ITEM_INCUBATOR_BASIC           ItemId = 902
	ItemId_ITEM_POKEMON_STORAGE_UPGRADE   ItemId = 1001
	ItemId_ITEM_ITEM_STORAGE_UPGRADE      ItemId = 1002
)

var ItemId_name = map[int32]string{
	0:    "ITEM_UNKNOWN",
	1:    "ITEM_POKE_BALL",
	2:    "ITEM_GREAT_BALL",
	3:    "ITEM_ULTRA_BALL",
	4:    "ITEM_MASTER_BALL",
	101:  "ITEM_POTION",
	102:  "ITEM_SUPER_POTION",
	103:  "ITEM_HYPER_POTION",
	104:  "ITEM_MAX_POTION",
	201:  "ITEM_REVIVE",
	202:  "ITEM_MAX_REVIVE",
	301:  "ITEM_LUCKY_EGG",
	401:  "ITEM_INCENSE_ORDINARY",
	402:  "ITEM_INCENSE_SPICY",
	403:  "ITEM_INCENSE_COOL",
	404:  "ITEM_INCENSE_FLORAL",
	501:  "ITEM_TROY_DISK",
	602:  "ITEM_X_ATTACK",
	603:  "ITEM_X_DEFENSE",
	604:  "ITEM_X_MIRACLE",
	701:  "ITEM_RAZZ_BERRY",
	702:  "ITEM_BLUK_BERRY",
	703:  "ITEM_NANAB_BERRY",
	704:  "ITEM_WEPAR_BERRY",
	705:  "ITEM_PINAP_BERRY",
	801:  "ITEM_SPECIAL_CAMERA",
	901:  "ITEM_INCUBATOR_BASIC_UNLIMITED",
	902:  "ITEM_INCUBATOR_BASIC",
	1001: "ITEM_POKEMON_STORAGE_UPGRADE",
	1002: "ITEM_ITEM_STORAGE_UPGRADE",
}
var ItemId_value = map[string]int32{
	"ITEM_UNKNOWN":                   0,
	"ITEM_POKE_BALL":                 1,
	"ITEM_GREAT_BALL":                2,
	"ITEM_ULTRA_BALL":                3,
	"ITEM_MASTER_BALL":               4,
	"ITEM_POTION":                    101,
	"ITEM_SUPER_POTION":              102,
	"ITEM_HYPER_POTION":              103,
	"ITEM_MAX_POTION":                104,
	"ITEM_REVIVE":                    201,
	"ITEM_MAX_REVIVE":                202,
	"ITEM_LUCKY_EGG":                 301,
	"ITEM_INCENSE_ORDINARY":          401,
	"ITEM_INCENSE_SPICY":             402,
	"ITEM_INCENSE_COOL":              403,
	"ITEM_INCENSE_FLORAL":            404,
	"ITEM_TROY_DISK":                 501,
	"ITEM_X_ATTACK":                  602,
	"ITEM_X_DEFENSE":                 603,
	"ITEM_X_MIRACLE":                 604,
	"ITEM_RAZZ_BERRY":                701,
	"ITEM_BLUK_BERRY":                702,
	"ITEM_NANAB_BERRY":               703,
	"ITEM_WEPAR_BERRY":               704,
	"ITEM_PINAP_BERRY":               705,
	"ITEM_SPECIAL_CAMERA":            801,
	"ITEM_INCUBATOR_BASIC_UNLIMITED": 901,
	"ITEM_INCUBATOR_BASIC":           902,
	"ITEM_POKEMON_STORAGE_UPGRADE":   1001,
	"ITEM_ITEM_STORAGE_UPGRADE":      1002,
}

func (x ItemId) String() string {
	return proto.EnumName(ItemId_name, int32(x))
}
func (ItemId) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type ItemType int32

const (
	ItemType_ITEM_TYPE_NONE              ItemType = 0
	ItemType_ITEM_TYPE_POKEBALL          ItemType = 1
	ItemType_ITEM_TYPE_POTION            ItemType = 2
	ItemType_ITEM_TYPE_REVIVE            ItemType = 3
	ItemType_ITEM_TYPE_MAP               ItemType = 4
	ItemType_ITEM_TYPE_BATTLE            ItemType = 5
	ItemType_ITEM_TYPE_FOOD              ItemType = 6
	ItemType_ITEM_TYPE_CAMERA            ItemType = 7
	ItemType_ITEM_TYPE_DISK              ItemType = 8
	ItemType_ITEM_TYPE_INCUBATOR         ItemType = 9
	ItemType_ITEM_TYPE_INCENSE           ItemType = 10
	ItemType_ITEM_TYPE_XP_BOOST          ItemType = 11
	ItemType_ITEM_TYPE_INVENTORY_UPGRADE ItemType = 12
)

var ItemType_name = map[int32]string{
	0:  "ITEM_TYPE_NONE",
	1:  "ITEM_TYPE_POKEBALL",
	2:  "ITEM_TYPE_POTION",
	3:  "ITEM_TYPE_REVIVE",
	4:  "ITEM_TYPE_MAP",
	5:  "ITEM_TYPE_BATTLE",
	6:  "ITEM_TYPE_FOOD",
	7:  "ITEM_TYPE_CAMERA",
	8:  "ITEM_TYPE_DISK",
	9:  "ITEM_TYPE_INCUBATOR",
	10: "ITEM_TYPE_INCENSE",
	11: "ITEM_TYPE_XP_BOOST",
	12: "ITEM_TYPE_INVENTORY_UPGRADE",
}
var ItemType_value = map[string]int32{
	"ITEM_TYPE_NONE":              0,
	"ITEM_TYPE_POKEBALL":          1,
	"ITEM_TYPE_POTION":            2,
	"ITEM_TYPE_REVIVE":            3,
	"ITEM_TYPE_MAP":               4,
	"ITEM_TYPE_BATTLE":            5,
	"ITEM_TYPE_FOOD":              6,
	"ITEM_TYPE_CAMERA":            7,
	"ITEM_TYPE_DISK":              8,
	"ITEM_TYPE_INCUBATOR":         9,
	"ITEM_TYPE_INCENSE":           10,
	"ITEM_TYPE_XP_BOOST":          11,
	"ITEM_TYPE_INVENTORY_UPGRADE": 12,
}

func (x ItemType) String() string {
	return proto.EnumName(ItemType_name, int32(x))
}
func (ItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type ItemAward struct {
	ItemId    ItemId `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	ItemCount int32  `protobuf:"varint,2,opt,name=item_count,json=itemCount" json:"item_count,omitempty"`
}

func (m *ItemAward) Reset()                    { *m = ItemAward{} }
func (m *ItemAward) String() string            { return proto.CompactTextString(m) }
func (*ItemAward) ProtoMessage()               {}
func (*ItemAward) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type ItemData struct {
	ItemId ItemId `protobuf:"varint,1,opt,name=item_id,json=itemId,enum=POGOProtos.Inventory.Item.ItemId" json:"item_id,omitempty"`
	Count  int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Unseen bool   `protobuf:"varint,3,opt,name=unseen" json:"unseen,omitempty"`
}

func (m *ItemData) Reset()                    { *m = ItemData{} }
func (m *ItemData) String() string            { return proto.CompactTextString(m) }
func (*ItemData) ProtoMessage()               {}
func (*ItemData) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func init() {
	proto.RegisterType((*ItemAward)(nil), "POGOProtos.Inventory.Item.ItemAward")
	proto.RegisterType((*ItemData)(nil), "POGOProtos.Inventory.Item.ItemData")
	proto.RegisterEnum("POGOProtos.Inventory.Item.ItemId", ItemId_name, ItemId_value)
	proto.RegisterEnum("POGOProtos.Inventory.Item.ItemType", ItemType_name, ItemType_value)
}

func init() { proto.RegisterFile("inventory_item.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x56, 0xd3, 0x4e,
	0x14, 0xc7, 0x69, 0xd3, 0x7f, 0x0c, 0xfc, 0x60, 0xb8, 0x14, 0x28, 0x3f, 0x15, 0x01, 0x37, 0x1c,
	0x16, 0x5d, 0xe8, 0xce, 0xdd, 0x24, 0x1d, 0xea, 0x9c, 0x26, 0x99, 0x9c, 0xc9, 0x04, 0x1a, 0x36,
	0x73, 0x50, 0x8a, 0x76, 0x41, 0xcb, 0x81, 0xa2, 0x87, 0x07, 0xd0, 0xbd, 0x7f, 0x5e, 0xc0, 0x07,
	0x70, 0xe3, 0xde, 0x7f, 0x2b, 0x8f, 0x3e, 0x82, 0xfa, 0x02, 0xba, 0x77, 0xef, 0xc9, 0x64, 0x1a,
	0x1a, 0x8f, 0xae, 0xdc, 0x74, 0x9a, 0xcf, 0x7c, 0xef, 0xdc, 0xb9, 0xf7, 0x7e, 0x13, 0x54, 0xef,
	0x0f, 0x1e, 0xf6, 0x06, 0xa3, 0xe1, 0xe9, 0x85, 0xea, 0x8f, 0x7a, 0xc7, 0xcd, 0x93, 0xd3, 0xe1,
	0x68, 0x08, 0xab, 0x01, 0x6f, 0xf3, 0x20, 0xf9, 0x7b, 0xd6, 0x64, 0x63, 0x41, 0x93, 0x8d, 0x7a,
	0xc7, 0x9b, 0x47, 0x68, 0x3a, 0x59, 0xc9, 0xa3, 0x83, 0xd3, 0x43, 0xb8, 0x8d, 0xaa, 0x49, 0x94,
	0xea, 0x1f, 0x36, 0x0a, 0xeb, 0x85, 0xad, 0xb9, 0x9b, 0x1b, 0xcd, 0xbf, 0x46, 0xea, 0x1f, 0x76,
	0x28, 0x2a, 0x7d, 0xbd, 0xc2, 0x35, 0x84, 0x74, 0xec, 0xbd, 0xe1, 0xf9, 0x60, 0xd4, 0x28, 0xae,
	0x17, 0xb6, 0xca, 0x62, 0x3a, 0x21, 0x4e, 0x02, 0x36, 0x47, 0xa8, 0x96, 0x04, 0xb4, 0x0e, 0x46,
	0x07, 0xff, 0x94, 0xa6, 0x8e, 0xca, 0x93, 0x19, 0xd2, 0x07, 0x58, 0x46, 0x95, 0xf3, 0xc1, 0x59,
	0xaf, 0x37, 0x68, 0x58, 0xeb, 0x85, 0xad, 0x9a, 0x30, 0x4f, 0xdb, 0x1f, 0xcb, 0xa8, 0x92, 0x1e,
	0x00, 0x18, 0xcd, 0x32, 0x49, 0x3d, 0x15, 0xf9, 0x1d, 0x9f, 0xef, 0xf9, 0x78, 0x0a, 0x00, 0xcd,
	0x69, 0x12, 0xf0, 0x0e, 0x55, 0x36, 0x71, 0x5d, 0x5c, 0x80, 0x45, 0x34, 0xaf, 0x59, 0x5b, 0x50,
	0x22, 0x53, 0x58, 0xcc, 0x60, 0xe4, 0x4a, 0x41, 0x52, 0x68, 0x41, 0x1d, 0x61, 0x0d, 0x3d, 0x12,
	0x4a, 0x2a, 0x52, 0x5a, 0x82, 0x79, 0x34, 0x63, 0xce, 0x94, 0x8c, 0xfb, 0xb8, 0x07, 0x4b, 0x68,
	0x41, 0x83, 0x30, 0x0a, 0xa8, 0x18, 0xe3, 0xa3, 0x0c, 0xdf, 0x89, 0x27, 0xf0, 0xfd, 0x2c, 0x93,
	0x47, 0xba, 0x63, 0xf8, 0x00, 0xb0, 0x39, 0x53, 0xd0, 0x5d, 0xb6, 0x4b, 0xf1, 0xa7, 0x02, 0xd4,
	0x27, 0x64, 0x86, 0x7e, 0x4e, 0xee, 0x9e, 0xd6, 0xe3, 0x46, 0x4e, 0x27, 0x56, 0xb4, 0xdd, 0xc6,
	0xaf, 0x8a, 0xf0, 0x3f, 0x5a, 0xd2, 0x90, 0xf9, 0x0e, 0xf5, 0x43, 0xaa, 0xb8, 0x68, 0x31, 0x9f,
	0x88, 0x18, 0x3f, 0xb5, 0x60, 0x05, 0x41, 0x6e, 0x2f, 0x0c, 0x98, 0x13, 0xe3, 0x67, 0x16, 0x2c,
	0x9b, 0xdb, 0x8d, 0x37, 0x1c, 0xce, 0x5d, 0xfc, 0xdc, 0x82, 0x06, 0x5a, 0xcc, 0xf1, 0x1d, 0x97,
	0x0b, 0xe2, 0xe2, 0x17, 0x56, 0x96, 0x5b, 0x0a, 0x1e, 0xab, 0x16, 0x0b, 0x3b, 0xf8, 0xa7, 0x05,
	0x80, 0xfe, 0xd3, 0xb0, 0xab, 0x88, 0x94, 0xc4, 0xe9, 0xe0, 0x2f, 0xa5, 0x4c, 0xd8, 0x55, 0x2d,
	0xba, 0x93, 0x1c, 0x82, 0xbf, 0x4e, 0x42, 0x8f, 0x09, 0xe2, 0xb8, 0x14, 0x7f, 0x2b, 0x65, 0x45,
	0x0a, 0xb2, 0xbf, 0xaf, 0x6c, 0x2a, 0x44, 0x8c, 0xdf, 0x94, 0x33, 0x6a, 0xbb, 0x51, 0xc7, 0xd0,
	0xb7, 0x65, 0x58, 0x32, 0xc3, 0xf0, 0x89, 0x4f, 0x6c, 0x83, 0xdf, 0x5d, 0xe2, 0x3d, 0x1a, 0x10,
	0x61, 0xf0, 0xfb, 0x4b, 0x1c, 0x30, 0x9f, 0x04, 0x06, 0x7f, 0x28, 0x67, 0xd5, 0x85, 0x01, 0x75,
	0x18, 0x71, 0x95, 0x43, 0x3c, 0x2a, 0x08, 0x7e, 0x59, 0x81, 0x1b, 0x68, 0x6d, 0x5c, 0x77, 0x64,
	0x13, 0xc9, 0x93, 0x71, 0x87, 0xcc, 0x51, 0x91, 0xef, 0x32, 0x8f, 0x49, 0xda, 0xc2, 0x8f, 0xab,
	0xb0, 0x8a, 0xea, 0x7f, 0x12, 0xe1, 0x27, 0x55, 0xd8, 0x40, 0x57, 0x33, 0xa7, 0x79, 0xdc, 0x57,
	0xa1, 0xe4, 0x82, 0xb4, 0xa9, 0x8a, 0x82, 0xb6, 0x20, 0x2d, 0x8a, 0xbf, 0x57, 0x61, 0x0d, 0xad,
	0xa6, 0xd1, 0xfa, 0x06, 0xbf, 0xed, 0xff, 0xa8, 0x6e, 0xbf, 0x2e, 0xa6, 0x2f, 0x90, 0xbc, 0x38,
	0xe9, 0x65, 0xce, 0x95, 0x71, 0x40, 0x95, 0xcf, 0x7d, 0x8a, 0xa7, 0x60, 0xd9, 0x0c, 0x53, 0xb3,
	0x24, 0x91, 0x71, 0xf4, 0xd8, 0xa7, 0x86, 0x6b, 0x4f, 0x15, 0xf3, 0xd4, 0x58, 0xc8, 0x82, 0x05,
	0x33, 0x30, 0x4d, 0x3d, 0x12, 0xe0, 0x52, 0x5e, 0x68, 0x13, 0x29, 0x5d, 0x8a, 0xcb, 0xf9, 0x0b,
	0xec, 0x70, 0xde, 0xc2, 0x95, 0xbc, 0xd2, 0xf4, 0xae, 0x9a, 0x57, 0x6a, 0x63, 0xd4, 0x60, 0xc5,
	0x34, 0x5a, 0xb3, 0xac, 0x5d, 0x78, 0x3a, 0x7b, 0x2b, 0xc6, 0x1b, 0xda, 0x1f, 0x28, 0x5f, 0x5a,
	0x37, 0x50, 0x36, 0xe7, 0xa1, 0xc4, 0x33, 0x70, 0x1d, 0x5d, 0x99, 0x94, 0xef, 0x52, 0x5f, 0x72,
	0x11, 0x67, 0x5d, 0x9b, 0xb5, 0x6b, 0xfb, 0x15, 0xfd, 0x01, 0x3c, 0xbb, 0x9b, 0xae, 0xb7, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x95, 0x01, 0xbf, 0x07, 0x20, 0x05, 0x00, 0x00,
}
