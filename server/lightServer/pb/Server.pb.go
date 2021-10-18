// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Server.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OrderState int32

const (
	OrderState_Pending  OrderState = 0
	OrderState_Cooking  OrderState = 1
	OrderState_Deliver  OrderState = 2
	OrderState_Complete OrderState = 3
)

var OrderState_name = map[int32]string{
	0: "Pending",
	1: "Cooking",
	2: "Deliver",
	3: "Complete",
}

var OrderState_value = map[string]int32{
	"Pending":  0,
	"Cooking":  1,
	"Deliver":  2,
	"Complete": 3,
}

func (x OrderState) String() string {
	return proto.EnumName(OrderState_name, int32(x))
}

func (OrderState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{0}
}

type RestaurantStatistics struct {
	Names                []*RestaurantName `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RestaurantStatistics) Reset()         { *m = RestaurantStatistics{} }
func (m *RestaurantStatistics) String() string { return proto.CompactTextString(m) }
func (*RestaurantStatistics) ProtoMessage()    {}
func (*RestaurantStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{0}
}

func (m *RestaurantStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestaurantStatistics.Unmarshal(m, b)
}
func (m *RestaurantStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestaurantStatistics.Marshal(b, m, deterministic)
}
func (m *RestaurantStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestaurantStatistics.Merge(m, src)
}
func (m *RestaurantStatistics) XXX_Size() int {
	return xxx_messageInfo_RestaurantStatistics.Size(m)
}
func (m *RestaurantStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_RestaurantStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_RestaurantStatistics proto.InternalMessageInfo

func (m *RestaurantStatistics) GetNames() []*RestaurantName {
	if m != nil {
		return m.Names
	}
	return nil
}

type RestaurantName struct {
	Path                 []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestaurantName) Reset()         { *m = RestaurantName{} }
func (m *RestaurantName) String() string { return proto.CompactTextString(m) }
func (*RestaurantName) ProtoMessage()    {}
func (*RestaurantName) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{1}
}

func (m *RestaurantName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestaurantName.Unmarshal(m, b)
}
func (m *RestaurantName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestaurantName.Marshal(b, m, deterministic)
}
func (m *RestaurantName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestaurantName.Merge(m, src)
}
func (m *RestaurantName) XXX_Size() int {
	return xxx_messageInfo_RestaurantName.Size(m)
}
func (m *RestaurantName) XXX_DiscardUnknown() {
	xxx_messageInfo_RestaurantName.DiscardUnknown(m)
}

var xxx_messageInfo_RestaurantName proto.InternalMessageInfo

func (m *RestaurantName) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *RestaurantName) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Restaurant struct {
	Name                 *RestaurantName `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url                  string          `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Image                string          `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Id                   int64           `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Restaurant) Reset()         { *m = Restaurant{} }
func (m *Restaurant) String() string { return proto.CompactTextString(m) }
func (*Restaurant) ProtoMessage()    {}
func (*Restaurant) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{2}
}

func (m *Restaurant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restaurant.Unmarshal(m, b)
}
func (m *Restaurant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restaurant.Marshal(b, m, deterministic)
}
func (m *Restaurant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restaurant.Merge(m, src)
}
func (m *Restaurant) XXX_Size() int {
	return xxx_messageInfo_Restaurant.Size(m)
}
func (m *Restaurant) XXX_DiscardUnknown() {
	xxx_messageInfo_Restaurant.DiscardUnknown(m)
}

var xxx_messageInfo_Restaurant proto.InternalMessageInfo

func (m *Restaurant) GetName() *RestaurantName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Restaurant) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Restaurant) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Restaurant) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Restaurant) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Menus struct {
	Owner                *RestaurantName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Menus                []*Menu         `protobuf:"bytes,2,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Menus) Reset()         { *m = Menus{} }
func (m *Menus) String() string { return proto.CompactTextString(m) }
func (*Menus) ProtoMessage()    {}
func (*Menus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{3}
}

func (m *Menus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menus.Unmarshal(m, b)
}
func (m *Menus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menus.Marshal(b, m, deterministic)
}
func (m *Menus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menus.Merge(m, src)
}
func (m *Menus) XXX_Size() int {
	return xxx_messageInfo_Menus.Size(m)
}
func (m *Menus) XXX_DiscardUnknown() {
	xxx_messageInfo_Menus.DiscardUnknown(m)
}

var xxx_messageInfo_Menus proto.InternalMessageInfo

func (m *Menus) GetOwner() *RestaurantName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Menus) GetMenus() []*Menu {
	if m != nil {
		return m.Menus
	}
	return nil
}

type Menu struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PriceAmount          int64    `protobuf:"varint,4,opt,name=priceAmount,proto3" json:"priceAmount,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Image                string   `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	Options              []*Menu  `protobuf:"bytes,10,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Menu) Reset()         { *m = Menu{} }
func (m *Menu) String() string { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()    {}
func (*Menu) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{4}
}

func (m *Menu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menu.Unmarshal(m, b)
}
func (m *Menu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menu.Marshal(b, m, deterministic)
}
func (m *Menu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menu.Merge(m, src)
}
func (m *Menu) XXX_Size() int {
	return xxx_messageInfo_Menu.Size(m)
}
func (m *Menu) XXX_DiscardUnknown() {
	xxx_messageInfo_Menu.DiscardUnknown(m)
}

var xxx_messageInfo_Menu proto.InternalMessageInfo

func (m *Menu) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Menu) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Menu) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Menu) GetPriceAmount() int64 {
	if m != nil {
		return m.PriceAmount
	}
	return 0
}

func (m *Menu) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Menu) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Menu) GetOptions() []*Menu {
	if m != nil {
		return m.Options
	}
	return nil
}

type Categories struct {
	Owner                *RestaurantName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Categories           []string        `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Categories) Reset()         { *m = Categories{} }
func (m *Categories) String() string { return proto.CompactTextString(m) }
func (*Categories) ProtoMessage()    {}
func (*Categories) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{5}
}

func (m *Categories) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Categories.Unmarshal(m, b)
}
func (m *Categories) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Categories.Marshal(b, m, deterministic)
}
func (m *Categories) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Categories.Merge(m, src)
}
func (m *Categories) XXX_Size() int {
	return xxx_messageInfo_Categories.Size(m)
}
func (m *Categories) XXX_DiscardUnknown() {
	xxx_messageInfo_Categories.DiscardUnknown(m)
}

var xxx_messageInfo_Categories proto.InternalMessageInfo

func (m *Categories) GetOwner() *RestaurantName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Categories) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

type Auth struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ServiceId            string   `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{6}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Auth) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

type User struct {
	Auth                 *Auth    `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{7}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mime                 string   `protobuf:"bytes,2,opt,name=mime,proto3" json:"mime,omitempty"`
	CreateAt             string   `protobuf:"bytes,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{8}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *File) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type OrderField struct {
	Menu                 int64    `protobuf:"varint,1,opt,name=menu,proto3" json:"menu,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderField) Reset()         { *m = OrderField{} }
func (m *OrderField) String() string { return proto.CompactTextString(m) }
func (*OrderField) ProtoMessage()    {}
func (*OrderField) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{9}
}

func (m *OrderField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderField.Unmarshal(m, b)
}
func (m *OrderField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderField.Marshal(b, m, deterministic)
}
func (m *OrderField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderField.Merge(m, src)
}
func (m *OrderField) XXX_Size() int {
	return xxx_messageInfo_OrderField.Size(m)
}
func (m *OrderField) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderField.DiscardUnknown(m)
}

var xxx_messageInfo_OrderField proto.InternalMessageInfo

func (m *OrderField) GetMenu() int64 {
	if m != nil {
		return m.Menu
	}
	return 0
}

func (m *OrderField) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OrderField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderField) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Order struct {
	Where                *RestaurantName `protobuf:"bytes,1,opt,name=where,proto3" json:"where,omitempty"`
	Who                  string          `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
	OrderAt              string          `protobuf:"bytes,3,opt,name=order_at,json=orderAt,proto3" json:"order_at,omitempty"`
	CookingAt            string          `protobuf:"bytes,4,opt,name=cooking_at,json=cookingAt,proto3" json:"cooking_at,omitempty"`
	DeliverAt            string          `protobuf:"bytes,5,opt,name=deliver_at,json=deliverAt,proto3" json:"deliver_at,omitempty"`
	CompleteAt           string          `protobuf:"bytes,6,opt,name=complete_at,json=completeAt,proto3" json:"complete_at,omitempty"`
	State                OrderState      `protobuf:"varint,7,opt,name=state,proto3,enum=pb.OrderState" json:"state,omitempty"`
	Id                   int64           `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	Purchased            []*OrderField   `protobuf:"bytes,10,rep,name=purchased,proto3" json:"purchased,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{10}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetWhere() *RestaurantName {
	if m != nil {
		return m.Where
	}
	return nil
}

func (m *Order) GetWho() string {
	if m != nil {
		return m.Who
	}
	return ""
}

func (m *Order) GetOrderAt() string {
	if m != nil {
		return m.OrderAt
	}
	return ""
}

func (m *Order) GetCookingAt() string {
	if m != nil {
		return m.CookingAt
	}
	return ""
}

func (m *Order) GetDeliverAt() string {
	if m != nil {
		return m.DeliverAt
	}
	return ""
}

func (m *Order) GetCompleteAt() string {
	if m != nil {
		return m.CompleteAt
	}
	return ""
}

func (m *Order) GetState() OrderState {
	if m != nil {
		return m.State
	}
	return OrderState_Pending
}

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetPurchased() []*OrderField {
	if m != nil {
		return m.Purchased
	}
	return nil
}

type OrderResult struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderResult) Reset()         { *m = OrderResult{} }
func (m *OrderResult) String() string { return proto.CompactTextString(m) }
func (*OrderResult) ProtoMessage()    {}
func (*OrderResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{11}
}

func (m *OrderResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResult.Unmarshal(m, b)
}
func (m *OrderResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResult.Marshal(b, m, deterministic)
}
func (m *OrderResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResult.Merge(m, src)
}
func (m *OrderResult) XXX_Size() int {
	return xxx_messageInfo_OrderResult.Size(m)
}
func (m *OrderResult) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResult.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResult proto.InternalMessageInfo

func (m *OrderResult) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type ManageCookingResult struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TimeAt               string   `protobuf:"bytes,2,opt,name=time_at,json=timeAt,proto3" json:"time_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManageCookingResult) Reset()         { *m = ManageCookingResult{} }
func (m *ManageCookingResult) String() string { return proto.CompactTextString(m) }
func (*ManageCookingResult) ProtoMessage()    {}
func (*ManageCookingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{12}
}

func (m *ManageCookingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManageCookingResult.Unmarshal(m, b)
}
func (m *ManageCookingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManageCookingResult.Marshal(b, m, deterministic)
}
func (m *ManageCookingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManageCookingResult.Merge(m, src)
}
func (m *ManageCookingResult) XXX_Size() int {
	return xxx_messageInfo_ManageCookingResult.Size(m)
}
func (m *ManageCookingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ManageCookingResult.DiscardUnknown(m)
}

var xxx_messageInfo_ManageCookingResult proto.InternalMessageInfo

func (m *ManageCookingResult) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ManageCookingResult) GetTimeAt() string {
	if m != nil {
		return m.TimeAt
	}
	return ""
}

type ManageDelieverResult struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TimeAt               string   `protobuf:"bytes,2,opt,name=time_at,json=timeAt,proto3" json:"time_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManageDelieverResult) Reset()         { *m = ManageDelieverResult{} }
func (m *ManageDelieverResult) String() string { return proto.CompactTextString(m) }
func (*ManageDelieverResult) ProtoMessage()    {}
func (*ManageDelieverResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{13}
}

func (m *ManageDelieverResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManageDelieverResult.Unmarshal(m, b)
}
func (m *ManageDelieverResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManageDelieverResult.Marshal(b, m, deterministic)
}
func (m *ManageDelieverResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManageDelieverResult.Merge(m, src)
}
func (m *ManageDelieverResult) XXX_Size() int {
	return xxx_messageInfo_ManageDelieverResult.Size(m)
}
func (m *ManageDelieverResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ManageDelieverResult.DiscardUnknown(m)
}

var xxx_messageInfo_ManageDelieverResult proto.InternalMessageInfo

func (m *ManageDelieverResult) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ManageDelieverResult) GetTimeAt() string {
	if m != nil {
		return m.TimeAt
	}
	return ""
}

type ManageCompleteResult struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	TimeAt               string   `protobuf:"bytes,2,opt,name=time_at,json=timeAt,proto3" json:"time_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManageCompleteResult) Reset()         { *m = ManageCompleteResult{} }
func (m *ManageCompleteResult) String() string { return proto.CompactTextString(m) }
func (*ManageCompleteResult) ProtoMessage()    {}
func (*ManageCompleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d447f7a82efed94, []int{14}
}

func (m *ManageCompleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManageCompleteResult.Unmarshal(m, b)
}
func (m *ManageCompleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManageCompleteResult.Marshal(b, m, deterministic)
}
func (m *ManageCompleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManageCompleteResult.Merge(m, src)
}
func (m *ManageCompleteResult) XXX_Size() int {
	return xxx_messageInfo_ManageCompleteResult.Size(m)
}
func (m *ManageCompleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ManageCompleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_ManageCompleteResult proto.InternalMessageInfo

func (m *ManageCompleteResult) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ManageCompleteResult) GetTimeAt() string {
	if m != nil {
		return m.TimeAt
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.OrderState", OrderState_name, OrderState_value)
	proto.RegisterType((*RestaurantStatistics)(nil), "pb.RestaurantStatistics")
	proto.RegisterType((*RestaurantName)(nil), "pb.RestaurantName")
	proto.RegisterType((*Restaurant)(nil), "pb.Restaurant")
	proto.RegisterType((*Menus)(nil), "pb.Menus")
	proto.RegisterType((*Menu)(nil), "pb.Menu")
	proto.RegisterType((*Categories)(nil), "pb.Categories")
	proto.RegisterType((*Auth)(nil), "pb.Auth")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*File)(nil), "pb.File")
	proto.RegisterType((*OrderField)(nil), "pb.OrderField")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*OrderResult)(nil), "pb.OrderResult")
	proto.RegisterType((*ManageCookingResult)(nil), "pb.ManageCookingResult")
	proto.RegisterType((*ManageDelieverResult)(nil), "pb.ManageDelieverResult")
	proto.RegisterType((*ManageCompleteResult)(nil), "pb.ManageCompleteResult")
}

func init() {
	proto.RegisterFile("Server.proto", fileDescriptor_7d447f7a82efed94)
}

var fileDescriptor_7d447f7a82efed94 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6a, 0xdb, 0x4a,
	0x14, 0x3d, 0xb2, 0x24, 0xdb, 0xda, 0x0e, 0xc6, 0xcc, 0x09, 0xe7, 0xe8, 0x9c, 0xde, 0x8c, 0x28,
	0xc5, 0x94, 0x92, 0x87, 0x14, 0xfa, 0x5c, 0x91, 0x10, 0x48, 0x21, 0xbd, 0x4c, 0x68, 0xdf, 0x4a,
	0x18, 0x4b, 0x1b, 0x7b, 0xa8, 0xa5, 0x11, 0xa3, 0x91, 0x4d, 0x7f, 0xa2, 0xff, 0xd0, 0x7f, 0xe9,
	0x87, 0x95, 0xb9, 0x48, 0x72, 0xdb, 0x34, 0x4d, 0xfb, 0x36, 0xfb, 0xb6, 0xf6, 0x5e, 0x6b, 0xf6,
	0x48, 0x70, 0x70, 0x89, 0x72, 0x8b, 0xf2, 0xa8, 0x92, 0x42, 0x09, 0x32, 0xa8, 0x96, 0xc9, 0x73,
	0x38, 0xa4, 0x58, 0x2b, 0xd6, 0x48, 0x56, 0xaa, 0x4b, 0xc5, 0x14, 0xaf, 0x15, 0xcf, 0x6a, 0xb2,
	0x80, 0xb0, 0x64, 0x05, 0xd6, 0xb1, 0x37, 0xf7, 0x17, 0x93, 0x63, 0x72, 0x54, 0x2d, 0x8f, 0xfa,
	0xc4, 0x97, 0xac, 0x40, 0x6a, 0x13, 0x92, 0x67, 0x30, 0xfd, 0x36, 0x40, 0x08, 0x04, 0x15, 0x53,
	0x6b, 0x53, 0x1a, 0x51, 0x73, 0x26, 0x33, 0xf0, 0x1b, 0xb9, 0x89, 0x07, 0x73, 0x6f, 0x11, 0x51,
	0x7d, 0x4c, 0x3e, 0x79, 0x00, 0x7d, 0x21, 0x79, 0x04, 0x81, 0xc6, 0x8b, 0xbd, 0xb9, 0xf7, 0x93,
	0x7e, 0x26, 0x4e, 0xe6, 0x30, 0xc9, 0xb1, 0xce, 0x24, 0xaf, 0x14, 0x17, 0xa5, 0x03, 0xdc, 0x77,
	0xb5, 0xad, 0xfc, 0xae, 0x15, 0x39, 0x84, 0x90, 0x17, 0x6c, 0x85, 0x71, 0x60, 0x7c, 0xd6, 0x20,
	0x53, 0x18, 0xf0, 0x3c, 0x0e, 0xe7, 0xde, 0xc2, 0xa7, 0x03, 0x9e, 0x27, 0x6f, 0x20, 0xbc, 0xc0,
	0xb2, 0x31, 0xdc, 0xc5, 0xae, 0x44, 0x79, 0xc3, 0x2c, 0x36, 0x81, 0xdc, 0x87, 0xb0, 0xd0, 0x25,
	0xf1, 0xc0, 0xa8, 0x34, 0xd6, 0x99, 0x1a, 0x83, 0x5a, 0x77, 0xf2, 0xc5, 0x83, 0x40, 0xdb, 0xae,
	0x97, 0xd7, 0xf6, 0xd2, 0x12, 0x19, 0xb6, 0x76, 0xfc, 0x6b, 0x99, 0xf9, 0x3f, 0x32, 0x9b, 0xc3,
	0xa4, 0x92, 0x3c, 0xc3, 0xb4, 0x10, 0x4d, 0xa9, 0x0c, 0x1b, 0x9f, 0xee, 0xbb, 0xc8, 0xff, 0x30,
	0xce, 0x1a, 0x29, 0xb1, 0xcc, 0x3e, 0x1a, 0x66, 0x11, 0xed, 0xec, 0x5e, 0x85, 0xd1, 0xbe, 0x0a,
	0x09, 0x8c, 0x84, 0x41, 0xaf, 0x63, 0xf8, 0x8e, 0x44, 0x1b, 0x48, 0xde, 0x01, 0x9c, 0x30, 0x85,
	0x2b, 0x21, 0x39, 0xfe, 0x9e, 0x3c, 0x90, 0x75, 0x75, 0x46, 0xa3, 0x88, 0xee, 0x79, 0x92, 0x14,
	0x82, 0xb4, 0x51, 0x6b, 0x3d, 0x75, 0x25, 0xc5, 0x96, 0xe7, 0x0e, 0x34, 0xa2, 0x9d, 0x4d, 0xee,
	0x01, 0xd4, 0x28, 0xb7, 0x3c, 0xc3, 0x2b, 0x9e, 0x3b, 0xbd, 0x22, 0xe7, 0x39, 0xcf, 0x13, 0x0a,
	0xc1, 0xdb, 0x1a, 0x25, 0xb9, 0x0b, 0x01, 0x6b, 0xcc, 0xce, 0x79, 0x2d, 0x07, 0x0d, 0x4d, 0x8d,
	0xf7, 0x5a, 0xb9, 0x63, 0x18, 0xb1, 0x3c, 0x97, 0x58, 0xd7, 0x4e, 0xea, 0xd6, 0x4c, 0xde, 0x43,
	0x70, 0xc6, 0x37, 0xd8, 0x55, 0x79, 0x7b, 0x55, 0x04, 0x82, 0x82, 0xf7, 0x48, 0xfa, 0x4c, 0xee,
	0x40, 0x94, 0x49, 0x64, 0x0a, 0xaf, 0x98, 0x72, 0x58, 0x63, 0xeb, 0x48, 0x55, 0xbb, 0x8d, 0x41,
	0xbf, 0xf8, 0x4b, 0x80, 0x57, 0x32, 0x47, 0x79, 0xc6, 0x71, 0x63, 0x36, 0x41, 0xef, 0x8a, 0xdb,
	0x0d, 0x73, 0xd6, 0x37, 0x95, 0x99, 0x1b, 0xd6, 0x5d, 0x42, 0x6a, 0x8d, 0x6e, 0x1c, 0x7f, 0x6f,
	0x9c, 0x7f, 0x60, 0xc8, 0xf6, 0x97, 0xc1, 0x59, 0xc9, 0xe7, 0x01, 0x84, 0xa6, 0x89, 0xbe, 0xad,
	0xdd, 0x1a, 0xe5, 0x4d, 0x0f, 0xcb, 0x26, 0xe8, 0x49, 0x77, 0x6b, 0xd1, 0x3e, 0xd1, 0xdd, 0x5a,
	0x90, 0xff, 0x60, 0x2c, 0x34, 0x48, 0xcf, 0x6b, 0x64, 0xec, 0x54, 0xe9, 0x6b, 0xc9, 0x84, 0xf8,
	0xc0, 0xcb, 0x95, 0x0e, 0x5a, 0x76, 0x91, 0xf3, 0xd8, 0x70, 0x8e, 0x1b, 0xbe, 0xb5, 0xb5, 0x76,
	0x13, 0x23, 0xe7, 0x49, 0x15, 0x79, 0x00, 0x93, 0x4c, 0x14, 0xd5, 0x06, 0xad, 0x66, 0x43, 0x13,
	0x87, 0xd6, 0x95, 0x2a, 0xf2, 0x10, 0xc2, 0x5a, 0x31, 0x65, 0x77, 0x75, 0x7a, 0x3c, 0xd5, 0x53,
	0x1b, 0x3e, 0xfa, 0x13, 0x85, 0xd4, 0x06, 0xdd, 0xab, 0x1a, 0x77, 0xaf, 0xea, 0x09, 0x44, 0x55,
	0x23, 0xb3, 0x35, 0xab, 0x31, 0x77, 0xdb, 0xdc, 0x57, 0x1a, 0xb9, 0x69, 0x9f, 0x90, 0x2c, 0x60,
	0x62, 0x02, 0x14, 0xeb, 0x66, 0xa3, 0x7a, 0xb2, 0xdd, 0x43, 0xb5, 0x64, 0xcf, 0xf3, 0xe4, 0x1c,
	0xfe, 0xbe, 0x60, 0x25, 0x5b, 0xe1, 0x89, 0x25, 0xf8, 0xcb, 0x0a, 0xf2, 0x2f, 0x8c, 0x14, 0x2f,
	0x0c, 0x39, 0xab, 0xe7, 0x50, 0x9b, 0xa9, 0x4a, 0x5e, 0xc0, 0xa1, 0x85, 0x3a, 0xc5, 0x0d, 0xc7,
	0xed, 0x2d, 0xba, 0xdf, 0x02, 0xeb, 0xc4, 0x09, 0xf7, 0xe7, 0x58, 0x8f, 0x53, 0xb7, 0x94, 0x46,
	0x5f, 0x32, 0x81, 0xd1, 0x6b, 0x2c, 0x73, 0x5e, 0xae, 0x66, 0x7f, 0x69, 0xc3, 0xf1, 0x9e, 0x79,
	0xda, 0x38, 0xb5, 0xd7, 0x38, 0x1b, 0x90, 0x03, 0x18, 0xb7, 0xad, 0x67, 0xfe, 0x72, 0x68, 0xfe,
	0x2a, 0x4f, 0xbf, 0x06, 0x00, 0x00, 0xff, 0xff, 0x47, 0x87, 0xf7, 0x14, 0x65, 0x06, 0x00, 0x00,
}
