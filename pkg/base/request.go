package base

import "net/http"

// Request represents a download request with all necessary metadata
type Request struct {
	// URL is the target download URL
	URL string `json:"url"`
	// Extra contains protocol-specific extra information
	Extra interface{} `json:"extra,omitempty"`
	// Labels are user-defined key-value pairs for categorizing requests
	Labels map[string]string `json:"labels,omitempty"`
	// ConnectTimeout is the timeout in seconds for establishing a connection
	ConnectTimeout int `json:"connectTimeout,omitempty"`
}

// Resource represents a downloadable resource resolved from a Request
type Resource struct {
	// Name is the suggested filename for the downloaded resource
	Name string `json:"name"`
	// Size is the total size of the resource in bytes, 0 if unknown
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests (resumable downloads)
	Range bool `json:"range"`
	// Files contains the list of files within this resource (e.g., for torrents)
	Files []*FileInfo `json:"files,omitempty"`
	// Hash is an optional content hash for integrity verification
	Hash string `json:"hash,omitempty"`
}

// FileInfo describes a single file within a multi-file resource
type FileInfo struct {
	// Name is the filename
	Name string `json:"name"`
	// Path is the relative directory path within the resource
	Path string `json:"path"`
	// Size is the file size in bytes
	Size int64 `json:"size"`
	// Selected indicates whether this file should be downloaded
	Selected bool `json:"selected"`
}

// Options holds configuration options for a download task
type Options struct {
	// Name overrides the default filename for the download
	Name string `json:"name,omitempty"`
	// Path is the local directory where the file will be saved
	Path string `json:"path,omitempty"`
	// SelectFiles specifies which files to download in a multi-file resource
	SelectFiles []int `json:"selectFiles,omitempty"`
	// Extra contains protocol-specific download options
	Extra interface{} `json:"extra,omitempty"`
	// Connections is the number of concurrent connections to use
	Connections int `json:"connections,omitempty"`
}

// Fetcher defines the interface that all download protocol implementations must satisfy
type Fetcher interface {
	// Resolve inspects the request and returns metadata about the downloadable resource
	Resolve(req *Request) (*Resource, error)
	// Create initializes the fetcher with the given resource and options
	Create(res *Resource, opts *Options) error
	// Start begins the download process
	Start() error
	// Pause suspends the ongoing download
	Pause() error
	// Continue resumes a paused download
	Continue() error
	// Close releases all resources held by the fetcher
	Close() error
	// Progress returns the current download progress
	Progress() Progress
}

// Progress tracks the download progress of a task
type Progress struct {
	// Used is the number of bytes downloaded so far
	Used int64 `json:"used"`
	// Speed is the current download speed in bytes per second
	Speed int64 `json:"speed"`
	// Downloaded is the total bytes downloaded in the current session
	Downloaded int64 `json:"downloaded"`
}

// HttpRequestConfig holds HTTP-specific configuration for a download
type HttpRequestConfig struct {
	// Header contains additional HTTP headers to send with the request
	Header http.Header `json:"header,omitempty"`
	// UserAgent overrides the default User-Agent header
	UserAgent string `json:"userAgent,omitempty"`
}
