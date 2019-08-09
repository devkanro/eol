package support

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

var UTF8 Encoding = &internalEncoding{unicode.UTF8}

var UTF16LE Encoding = &internalEncoding{unicode.UTF16(unicode.LittleEndian, unicode.UseBOM)}

var UTF16BE Encoding = &internalEncoding{unicode.UTF16(unicode.BigEndian, unicode.UseBOM)}

var UTF16 = UTF16LE

var GBK Encoding = &internalEncoding{simplifiedchinese.GBK}

var GB18030 Encoding = &internalEncoding{simplifiedchinese.GB18030}

var Big5 Encoding = &internalEncoding{traditionalchinese.Big5}

var ShiftJIS Encoding = &internalEncoding{japanese.ShiftJIS}

var EUCKR Encoding = &internalEncoding{korean.EUCKR}
