package files

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type FileService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewFileService(requester core.Requester) (r *FileService) {
	r = &FileService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedFileService struct {
	Files *FileService
}

func (r *PreloadedFileService) Init(service *FileService) {
	r.Files = service
}

func NewPreloadedFileService(service *FileService) (r *PreloadedFileService) {
	r = &PreloadedFileService{}
	r.Init(service)
	return
}

//
type File struct {
	// The time the File was created.
	CreatedAt *string `json:"created_at"`
	// The File's identifier.
	Id *string `json:"id"`
	// What the File will be used for. We may add additional possible values for this
	// enum over time; your application should be able to handle such additions
	// gracefully.
	Purpose *FilePurpose `json:"purpose"`
	// A description of the File.
	Description *string `json:"description"`
	// Whether the File was generated by Increase or by you and sent to Increase.
	Direction *FileDirection `json:"direction"`
	// The filename that was provided upon upload or generated by Increase.
	Filename *string `json:"filename"`
	// A URL from where the File can be downloaded at this point in time. The location
	// of this URL may change over time.
	DownloadURL *string `json:"download_url"`
	// A constant representing the object's type. For this resource it will always be
	// `file`.
	Type *FileType `json:"type"`
}

// The time the File was created.
func (r *File) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The File's identifier.
func (r *File) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// What the File will be used for. We may add additional possible values for this
// enum over time; your application should be able to handle such additions
// gracefully.
func (r *File) GetPurpose() (Purpose FilePurpose) {
	if r != nil && r.Purpose != nil {
		Purpose = *r.Purpose
	}
	return
}

// A description of the File.
func (r *File) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// Whether the File was generated by Increase or by you and sent to Increase.
func (r *File) GetDirection() (Direction FileDirection) {
	if r != nil && r.Direction != nil {
		Direction = *r.Direction
	}
	return
}

// The filename that was provided upon upload or generated by Increase.
func (r *File) GetFilename() (Filename string) {
	if r != nil && r.Filename != nil {
		Filename = *r.Filename
	}
	return
}

// A URL from where the File can be downloaded at this point in time. The location
// of this URL may change over time.
func (r *File) GetDownloadURL() (DownloadURL string) {
	if r != nil && r.DownloadURL != nil {
		DownloadURL = *r.DownloadURL
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `file`.
func (r *File) GetType() (Type FileType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type FilePurpose string

const (
	FilePurposeCheckImageFront            FilePurpose = "check_image_front"
	FilePurposeCheckImageBack             FilePurpose = "check_image_back"
	FilePurposeForm_1099Int               FilePurpose = "form_1099_int"
	FilePurposeFormSs_4                   FilePurpose = "form_ss_4"
	FilePurposeIdentityDocument           FilePurpose = "identity_document"
	FilePurposeIncreaseStatement          FilePurpose = "increase_statement"
	FilePurposeOther                      FilePurpose = "other"
	FilePurposeTrustFormationDocument     FilePurpose = "trust_formation_document"
	FilePurposeDigitalWalletArtwork       FilePurpose = "digital_wallet_artwork"
	FilePurposeDigitalWalletAppIcon       FilePurpose = "digital_wallet_app_icon"
	FilePurposeEntitySupplementalDocument FilePurpose = "entity_supplemental_document"
)

type FileDirection string

const (
	FileDirectionToIncrease   FileDirection = "to_increase"
	FileDirectionFromIncrease FileDirection = "from_increase"
)

type FileType string

const (
	FileTypeFile FileType = "file"
)

type CreateAFileParameters struct {
	// The file contents. This should follow the specifications of
	// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
	// transfers for the multipart/form-data protocol.
	File *string `json:"file"`
	// The description you choose to give the File.
	Description *string `json:"description,omitempty"`
	// What the File will be used for in Increase's systems.
	Purpose *CreateAFileParametersPurpose `json:"purpose"`
}

// The file contents. This should follow the specifications of
// [RFC 7578](https://datatracker.ietf.org/doc/html/rfc7578) which defines file
// transfers for the multipart/form-data protocol.
func (r *CreateAFileParameters) GetFile() (File string) {
	if r != nil && r.File != nil {
		File = *r.File
	}
	return
}

// The description you choose to give the File.
func (r *CreateAFileParameters) GetDescription() (Description string) {
	if r != nil && r.Description != nil {
		Description = *r.Description
	}
	return
}

// What the File will be used for in Increase's systems.
func (r *CreateAFileParameters) GetPurpose() (Purpose CreateAFileParametersPurpose) {
	if r != nil && r.Purpose != nil {
		Purpose = *r.Purpose
	}
	return
}

type CreateAFileParametersPurpose string

const (
	CreateAFileParametersPurposeCheckImageFront            CreateAFileParametersPurpose = "check_image_front"
	CreateAFileParametersPurposeCheckImageBack             CreateAFileParametersPurpose = "check_image_back"
	CreateAFileParametersPurposeFormSs_4                   CreateAFileParametersPurpose = "form_ss_4"
	CreateAFileParametersPurposeIdentityDocument           CreateAFileParametersPurpose = "identity_document"
	CreateAFileParametersPurposeOther                      CreateAFileParametersPurpose = "other"
	CreateAFileParametersPurposeTrustFormationDocument     CreateAFileParametersPurpose = "trust_formation_document"
	CreateAFileParametersPurposeDigitalWalletArtwork       CreateAFileParametersPurpose = "digital_wallet_artwork"
	CreateAFileParametersPurposeDigitalWalletAppIcon       CreateAFileParametersPurpose = "digital_wallet_app_icon"
	CreateAFileParametersPurposeEntitySupplementalDocument CreateAFileParametersPurpose = "entity_supplemental_document"
)

type ListFilesQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     *int                     `query:"limit"`
	CreatedAt *ListFilesQueryCreatedAt `query:"created_at"`
	Purpose   *ListFilesQueryPurpose   `query:"purpose"`
}

// Return the page of entries after this one.
func (r *ListFilesQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListFilesQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *ListFilesQuery) GetCreatedAt() (CreatedAt ListFilesQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *ListFilesQuery) GetPurpose() (Purpose ListFilesQueryPurpose) {
	if r != nil && r.Purpose != nil {
		Purpose = *r.Purpose
	}
	return
}

type ListFilesQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListFilesQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListFilesQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListFilesQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListFilesQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

type ListFilesQueryPurpose struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In *[]ListFilesQueryPurposeIn `json:"in,omitempty"`
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *ListFilesQueryPurpose) GetIn() (In []ListFilesQueryPurposeIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

type ListFilesQueryPurposeIn string

const (
	ListFilesQueryPurposeInCheckImageFront            ListFilesQueryPurposeIn = "check_image_front"
	ListFilesQueryPurposeInCheckImageBack             ListFilesQueryPurposeIn = "check_image_back"
	ListFilesQueryPurposeInForm_1099Int               ListFilesQueryPurposeIn = "form_1099_int"
	ListFilesQueryPurposeInFormSs_4                   ListFilesQueryPurposeIn = "form_ss_4"
	ListFilesQueryPurposeInIdentityDocument           ListFilesQueryPurposeIn = "identity_document"
	ListFilesQueryPurposeInIncreaseStatement          ListFilesQueryPurposeIn = "increase_statement"
	ListFilesQueryPurposeInOther                      ListFilesQueryPurposeIn = "other"
	ListFilesQueryPurposeInTrustFormationDocument     ListFilesQueryPurposeIn = "trust_formation_document"
	ListFilesQueryPurposeInDigitalWalletArtwork       ListFilesQueryPurposeIn = "digital_wallet_artwork"
	ListFilesQueryPurposeInDigitalWalletAppIcon       ListFilesQueryPurposeIn = "digital_wallet_app_icon"
	ListFilesQueryPurposeInEntitySupplementalDocument ListFilesQueryPurposeIn = "entity_supplemental_document"
)

//
type FileList struct {
	// The contents of the list.
	Data *[]File `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *FileList) GetData() (Data []File) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *FileList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *FileService) Create(ctx context.Context, body *CreateAFileParameters, opts ...*core.RequestOpts) (res *File, err error) {
	err = r.post(
		ctx,
		"/files",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

// To upload a file to Increase, you'll need to send a request of Content-Type
// `multipart/form-data`. The request should contain the file you would like to
// upload, as well as the parameters for creating a file.
func (r *PreloadedFileService) Create(ctx context.Context, body *CreateAFileParameters, opts ...*core.RequestOpts) (res *File, err error) {
	err = r.Files.post(
		ctx,
		"/files",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *FileService) Retrieve(ctx context.Context, file_id string, opts ...*core.RequestOpts) (res *File, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/files/%s", file_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedFileService) Retrieve(ctx context.Context, file_id string, opts ...*core.RequestOpts) (res *File, err error) {
	err = r.Files.get(
		ctx,
		fmt.Sprintf("/files/%s", file_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type FilesPage struct {
	*pagination.Page[File]
}

func (r *FilesPage) File() *File {
	return r.Current()
}

func (r *FilesPage) GetNextPage() (*FilesPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &FilesPage{page}, nil
	}
}

func (r *FileService) List(ctx context.Context, query *ListFilesQuery, opts ...*core.RequestOpts) (res *FilesPage, err error) {
	page := &FilesPage{
		Page: &pagination.Page[File]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/files",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedFileService) List(ctx context.Context, query *ListFilesQuery, opts ...*core.RequestOpts) (res *FilesPage, err error) {
	page := &FilesPage{
		Page: &pagination.Page[File]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/files",
			},
			Requester: r.Files.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
