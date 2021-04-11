// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.6
// source: movie.proto

package movie0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PerPage     uint64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	CurrentPage uint64 `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	LastPage    uint64 `protobuf:"varint,4,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetPerPage() uint64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *Pagination) GetCurrentPage() uint64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetLastPage() uint64 {
	if x != nil {
		return x.LastPage
	}
	return 0
}

type Writer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Writer) Reset() {
	*x = Writer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Writer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Writer) ProtoMessage() {}

func (x *Writer) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Writer.ProtoReflect.Descriptor instead.
func (*Writer) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *Writer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Writer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Value  float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Value  uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{3}
}

func (x *Currency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Currency) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   uint64 `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbId string `protobuf:"bytes,3,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{4}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() uint64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Movie) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *Movie) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Movie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type MovieDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       uint64                 `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
	Rated      string                 `protobuf:"bytes,3,opt,name=rated,proto3" json:"rated,omitempty"`
	Released   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=released,proto3" json:"released,omitempty"`
	Runtime    *durationpb.Duration   `protobuf:"bytes,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genres     []string               `protobuf:"bytes,6,rep,name=genres,proto3" json:"genres,omitempty"`
	Director   string                 `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`
	Writers    []*Writer              `protobuf:"bytes,8,rep,name=writers,proto3" json:"writers,omitempty"`
	Actors     []string               `protobuf:"bytes,9,rep,name=actors,proto3" json:"actors,omitempty"`
	Plot       string                 `protobuf:"bytes,10,opt,name=plot,proto3" json:"plot,omitempty"`
	Languages  []string               `protobuf:"bytes,11,rep,name=languages,proto3" json:"languages,omitempty"`
	Countries  []string               `protobuf:"bytes,12,rep,name=countries,proto3" json:"countries,omitempty"`
	Awards     string                 `protobuf:"bytes,13,opt,name=awards,proto3" json:"awards,omitempty"`
	Poster     string                 `protobuf:"bytes,14,opt,name=poster,proto3" json:"poster,omitempty"`
	Ratings    []*Rating              `protobuf:"bytes,15,rep,name=ratings,proto3" json:"ratings,omitempty"`
	Metascore  float32                `protobuf:"fixed32,16,opt,name=metascore,proto3" json:"metascore,omitempty"`
	ImdbRating float32                `protobuf:"fixed32,17,opt,name=imdb_rating,json=imdbRating,proto3" json:"imdb_rating,omitempty"`
	ImdbVotes  float32                `protobuf:"fixed32,18,opt,name=imdb_votes,json=imdbVotes,proto3" json:"imdb_votes,omitempty"`
	ImdbId     string                 `protobuf:"bytes,19,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Type       string                 `protobuf:"bytes,20,opt,name=type,proto3" json:"type,omitempty"`
	Dvd        *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=dvd,proto3" json:"dvd,omitempty"`
	BoxOffice  *Currency              `protobuf:"bytes,22,opt,name=box_office,json=boxOffice,proto3" json:"box_office,omitempty"`
	Production []string               `protobuf:"bytes,23,rep,name=production,proto3" json:"production,omitempty"`
	Website    string                 `protobuf:"bytes,24,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *MovieDetail) Reset() {
	*x = MovieDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetail) ProtoMessage() {}

func (x *MovieDetail) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetail.ProtoReflect.Descriptor instead.
func (*MovieDetail) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{5}
}

func (x *MovieDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetail) GetYear() uint64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *MovieDetail) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *MovieDetail) GetReleased() *timestamppb.Timestamp {
	if x != nil {
		return x.Released
	}
	return nil
}

func (x *MovieDetail) GetRuntime() *durationpb.Duration {
	if x != nil {
		return x.Runtime
	}
	return nil
}

func (x *MovieDetail) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *MovieDetail) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieDetail) GetWriters() []*Writer {
	if x != nil {
		return x.Writers
	}
	return nil
}

func (x *MovieDetail) GetActors() []string {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *MovieDetail) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *MovieDetail) GetLanguages() []string {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *MovieDetail) GetCountries() []string {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *MovieDetail) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *MovieDetail) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *MovieDetail) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *MovieDetail) GetMetascore() float32 {
	if x != nil {
		return x.Metascore
	}
	return 0
}

func (x *MovieDetail) GetImdbRating() float32 {
	if x != nil {
		return x.ImdbRating
	}
	return 0
}

func (x *MovieDetail) GetImdbVotes() float32 {
	if x != nil {
		return x.ImdbVotes
	}
	return 0
}

func (x *MovieDetail) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *MovieDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetail) GetDvd() *timestamppb.Timestamp {
	if x != nil {
		return x.Dvd
	}
	return nil
}

func (x *MovieDetail) GetBoxOffice() *Currency {
	if x != nil {
		return x.BoxOffice
	}
	return nil
}

func (x *MovieDetail) GetProduction() []string {
	if x != nil {
		return x.Production
	}
	return nil
}

func (x *MovieDetail) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

type MoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Page   uint64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *MoviesRequest) Reset() {
	*x = MoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviesRequest) ProtoMessage() {}

func (x *MoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviesRequest.ProtoReflect.Descriptor instead.
func (*MoviesRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{6}
}

func (x *MoviesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *MoviesRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type MoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response   bool        `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Message    string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data       []*Movie    `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *Pagination `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *MoviesResponse) Reset() {
	*x = MoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviesResponse) ProtoMessage() {}

func (x *MoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviesResponse.ProtoReflect.Descriptor instead.
func (*MoviesResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{7}
}

func (x *MoviesResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

func (x *MoviesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MoviesResponse) GetData() []*Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MoviesResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type MovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MovieRequest) Reset() {
	*x = MovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRequest) ProtoMessage() {}

func (x *MovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRequest.ProtoReflect.Descriptor instead.
func (*MovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{8}
}

func (x *MovieRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response   bool         `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
	Message    string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data       *MovieDetail `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Pagination *Pagination  `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *MovieResponse) Reset() {
	*x = MovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResponse) ProtoMessage() {}

func (x *MovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResponse.ProtoReflect.Descriptor instead.
func (*MovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{9}
}

func (x *MovieResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

func (x *MovieResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MovieResponse) GetData() *MovieDetail {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MovieResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x30, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x76, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d,
	0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22,
	0xfe, 0x05, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x36, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x28, 0x0a, 0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x52, 0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d,
	0x64, 0x62, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x6d, 0x64, 0x62, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d,
	0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x64, 0x76, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x03, 0x64, 0x76, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x62, 0x6f, 0x78, 0x5f, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x30, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x62, 0x6f, 0x78,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x22, 0x3b, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x9d, 0x01,
	0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x0a,
	0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa2, 0x01,
	0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x87, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x6a, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x30, 0x32, 0x5f, 0x6f, 0x6d, 0x64, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_movie_proto_goTypes = []interface{}{
	(*Pagination)(nil),            // 0: movie0.Pagination
	(*Writer)(nil),                // 1: movie0.Writer
	(*Rating)(nil),                // 2: movie0.Rating
	(*Currency)(nil),              // 3: movie0.Currency
	(*Movie)(nil),                 // 4: movie0.Movie
	(*MovieDetail)(nil),           // 5: movie0.MovieDetail
	(*MoviesRequest)(nil),         // 6: movie0.MoviesRequest
	(*MoviesResponse)(nil),        // 7: movie0.MoviesResponse
	(*MovieRequest)(nil),          // 8: movie0.MovieRequest
	(*MovieResponse)(nil),         // 9: movie0.MovieResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 11: google.protobuf.Duration
}
var file_movie_proto_depIdxs = []int32{
	10, // 0: movie0.MovieDetail.released:type_name -> google.protobuf.Timestamp
	11, // 1: movie0.MovieDetail.runtime:type_name -> google.protobuf.Duration
	1,  // 2: movie0.MovieDetail.writers:type_name -> movie0.Writer
	2,  // 3: movie0.MovieDetail.ratings:type_name -> movie0.Rating
	10, // 4: movie0.MovieDetail.dvd:type_name -> google.protobuf.Timestamp
	3,  // 5: movie0.MovieDetail.box_office:type_name -> movie0.Currency
	4,  // 6: movie0.MoviesResponse.data:type_name -> movie0.Movie
	0,  // 7: movie0.MoviesResponse.pagination:type_name -> movie0.Pagination
	5,  // 8: movie0.MovieResponse.data:type_name -> movie0.MovieDetail
	0,  // 9: movie0.MovieResponse.pagination:type_name -> movie0.Pagination
	6,  // 10: movie0.MovieService.GetMovies:input_type -> movie0.MoviesRequest
	8,  // 11: movie0.MovieService.GetMovie:input_type -> movie0.MovieRequest
	7,  // 12: movie0.MovieService.GetMovies:output_type -> movie0.MoviesResponse
	9,  // 13: movie0.MovieService.GetMovie:output_type -> movie0.MovieResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Writer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoviesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoviesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}