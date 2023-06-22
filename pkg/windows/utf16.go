package windows

import "golang.org/x/sys/windows"

func (s *Utf16) String() string {
	return windows.UTF16PtrToString(s.PwStr)
}
