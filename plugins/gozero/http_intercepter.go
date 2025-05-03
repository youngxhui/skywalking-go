package gozero

type HttpInterceptor struct {}

	func (h *HttpInterceptor)Name() string {
return "go-zero"
	}
	func (h *HttpInterceptor)BasePackage() string {
return "github.com/zeromicro/go-zero"
	}
	func (h *HttpInterceptor)VersionChecker(version string) bool {
return strings.HasPrefix(version, "v1.")
	}
	func (h *HttpInterceptor)Points() []*Point {

	}
	func (h *HttpInterceptor) FS() *embed.FS {
return &fs
	}