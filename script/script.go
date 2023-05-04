package script

import "unicode"

// Adlam: 88 codepoints
var Adlam = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x01E900, Hi: 0x01E94B, Stride: 1 },
    unicode.Range32{Lo: 0x01E950, Hi: 0x01E959, Stride: 1 },
    unicode.Range32{Lo: 0x01E95E, Hi: 0x01E95F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Ahom: 65 codepoints
var Ahom = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x011700, Hi: 0x01171A, Stride: 1 },
    unicode.Range32{Lo: 0x01171D, Hi: 0x01172B, Stride: 1 },
    unicode.Range32{Lo: 0x011730, Hi: 0x011746, Stride: 1 },
  },
  LatinOffset: 0,
}

// Anatolian_Hieroglyphs: 583 codepoints
var Anatolian_Hieroglyphs = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x014400, Hi: 0x014646, Stride: 1 },
  },
  LatinOffset: 0,
}

// Arabic: 1368 codepoints
var Arabic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 22 */
    unicode.Range16{Lo: 0x0600, Hi: 0x0604, Stride: 1 },
    unicode.Range16{Lo: 0x0606, Hi: 0x060B, Stride: 1 },
    unicode.Range16{Lo: 0x060D, Hi: 0x061A, Stride: 1 },
    unicode.Range16{Lo: 0x061C, Hi: 0x061E, Stride: 1 },
    unicode.Range16{Lo: 0x0620, Hi: 0x063F, Stride: 1 },
    unicode.Range16{Lo: 0x0641, Hi: 0x064A, Stride: 1 },
    unicode.Range16{Lo: 0x0656, Hi: 0x066F, Stride: 1 },
    unicode.Range16{Lo: 0x0671, Hi: 0x06DC, Stride: 1 },
    unicode.Range16{Lo: 0x06DE, Hi: 0x06FF, Stride: 1 },
    unicode.Range16{Lo: 0x0750, Hi: 0x077F, Stride: 1 },
    unicode.Range16{Lo: 0x0870, Hi: 0x088E, Stride: 1 },
    unicode.Range16{Lo: 0x0890, Hi: 0x0891, Stride: 1 },
    unicode.Range16{Lo: 0x0898, Hi: 0x08E1, Stride: 1 },
    unicode.Range16{Lo: 0x08E3, Hi: 0x08FF, Stride: 1 },
    unicode.Range16{Lo: 0xFB50, Hi: 0xFBC2, Stride: 1 },
    unicode.Range16{Lo: 0xFBD3, Hi: 0xFD3D, Stride: 1 },
    unicode.Range16{Lo: 0xFD40, Hi: 0xFD8F, Stride: 1 },
    unicode.Range16{Lo: 0xFD92, Hi: 0xFDC7, Stride: 1 },
    unicode.Range16{Lo: 0xFDCF, Hi: 0xFDF0, Stride: 33 },
    unicode.Range16{Lo: 0xFDF1, Hi: 0xFDFF, Stride: 1 },
    unicode.Range16{Lo: 0xFE70, Hi: 0xFE74, Stride: 1 },
    unicode.Range16{Lo: 0xFE76, Hi: 0xFEFC, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 27 */
    unicode.Range32{Lo: 0x010E60, Hi: 0x010E7E, Stride: 1 },
    unicode.Range32{Lo: 0x010EFD, Hi: 0x010EFF, Stride: 1 },
    unicode.Range32{Lo: 0x01EE00, Hi: 0x01EE03, Stride: 1 },
    unicode.Range32{Lo: 0x01EE05, Hi: 0x01EE1F, Stride: 1 },
    unicode.Range32{Lo: 0x01EE21, Hi: 0x01EE22, Stride: 1 },
    unicode.Range32{Lo: 0x01EE24, Hi: 0x01EE27, Stride: 3 },
    unicode.Range32{Lo: 0x01EE29, Hi: 0x01EE32, Stride: 1 },
    unicode.Range32{Lo: 0x01EE34, Hi: 0x01EE37, Stride: 1 },
    unicode.Range32{Lo: 0x01EE39, Hi: 0x01EE3B, Stride: 2 },
    unicode.Range32{Lo: 0x01EE42, Hi: 0x01EE47, Stride: 5 },
    unicode.Range32{Lo: 0x01EE49, Hi: 0x01EE4D, Stride: 2 },
    unicode.Range32{Lo: 0x01EE4E, Hi: 0x01EE4F, Stride: 1 },
    unicode.Range32{Lo: 0x01EE51, Hi: 0x01EE52, Stride: 1 },
    unicode.Range32{Lo: 0x01EE54, Hi: 0x01EE57, Stride: 3 },
    unicode.Range32{Lo: 0x01EE59, Hi: 0x01EE61, Stride: 2 },
    unicode.Range32{Lo: 0x01EE62, Hi: 0x01EE64, Stride: 2 },
    unicode.Range32{Lo: 0x01EE67, Hi: 0x01EE6A, Stride: 1 },
    unicode.Range32{Lo: 0x01EE6C, Hi: 0x01EE72, Stride: 1 },
    unicode.Range32{Lo: 0x01EE74, Hi: 0x01EE77, Stride: 1 },
    unicode.Range32{Lo: 0x01EE79, Hi: 0x01EE7C, Stride: 1 },
    unicode.Range32{Lo: 0x01EE7E, Hi: 0x01EE80, Stride: 2 },
    unicode.Range32{Lo: 0x01EE81, Hi: 0x01EE89, Stride: 1 },
    unicode.Range32{Lo: 0x01EE8B, Hi: 0x01EE9B, Stride: 1 },
    unicode.Range32{Lo: 0x01EEA1, Hi: 0x01EEA3, Stride: 1 },
    unicode.Range32{Lo: 0x01EEA5, Hi: 0x01EEA9, Stride: 1 },
    unicode.Range32{Lo: 0x01EEAB, Hi: 0x01EEBB, Stride: 1 },
    unicode.Range32{Lo: 0x01EEF0, Hi: 0x01EEF1, Stride: 1 },
  },
  LatinOffset: 0,
}

// Armenian: 96 codepoints
var Armenian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0x0531, Hi: 0x0556, Stride: 1 },
    unicode.Range16{Lo: 0x0559, Hi: 0x058A, Stride: 1 },
    unicode.Range16{Lo: 0x058D, Hi: 0x058F, Stride: 1 },
    unicode.Range16{Lo: 0xFB13, Hi: 0xFB17, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Avestan: 61 codepoints
var Avestan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010B00, Hi: 0x010B35, Stride: 1 },
    unicode.Range32{Lo: 0x010B39, Hi: 0x010B3F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Balinese: 124 codepoints
var Balinese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1B00, Hi: 0x1B4C, Stride: 1 },
    unicode.Range16{Lo: 0x1B50, Hi: 0x1B7E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Bamum: 657 codepoints
var Bamum = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0xA6A0, Hi: 0xA6F7, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x016800, Hi: 0x016A38, Stride: 1 },
  },
  LatinOffset: 0,
}

// Bassa_Vah: 36 codepoints
var Bassa_Vah = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x016AD0, Hi: 0x016AED, Stride: 1 },
    unicode.Range32{Lo: 0x016AF0, Hi: 0x016AF5, Stride: 1 },
  },
  LatinOffset: 0,
}

// Batak: 56 codepoints
var Batak = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1BC0, Hi: 0x1BF3, Stride: 1 },
    unicode.Range16{Lo: 0x1BFC, Hi: 0x1BFF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Bengali: 96 codepoints
var Bengali = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0x0980, Hi: 0x0983, Stride: 1 },
    unicode.Range16{Lo: 0x0985, Hi: 0x098C, Stride: 1 },
    unicode.Range16{Lo: 0x098F, Hi: 0x0990, Stride: 1 },
    unicode.Range16{Lo: 0x0993, Hi: 0x09A8, Stride: 1 },
    unicode.Range16{Lo: 0x09AA, Hi: 0x09B0, Stride: 1 },
    unicode.Range16{Lo: 0x09B2, Hi: 0x09B6, Stride: 4 },
    unicode.Range16{Lo: 0x09B7, Hi: 0x09B9, Stride: 1 },
    unicode.Range16{Lo: 0x09BC, Hi: 0x09C4, Stride: 1 },
    unicode.Range16{Lo: 0x09C7, Hi: 0x09C8, Stride: 1 },
    unicode.Range16{Lo: 0x09CB, Hi: 0x09CE, Stride: 1 },
    unicode.Range16{Lo: 0x09D7, Hi: 0x09DC, Stride: 5 },
    unicode.Range16{Lo: 0x09DD, Hi: 0x09DF, Stride: 2 },
    unicode.Range16{Lo: 0x09E0, Hi: 0x09E3, Stride: 1 },
    unicode.Range16{Lo: 0x09E6, Hi: 0x09FE, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Bhaiksuki: 97 codepoints
var Bhaiksuki = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 0x011C00, Hi: 0x011C08, Stride: 1 },
    unicode.Range32{Lo: 0x011C0A, Hi: 0x011C36, Stride: 1 },
    unicode.Range32{Lo: 0x011C38, Hi: 0x011C45, Stride: 1 },
    unicode.Range32{Lo: 0x011C50, Hi: 0x011C6C, Stride: 1 },
  },
  LatinOffset: 0,
}

// Bopomofo: 77 codepoints
var Bopomofo = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x02EA, Hi: 0x02EB, Stride: 1 },
    unicode.Range16{Lo: 0x3105, Hi: 0x312F, Stride: 1 },
    unicode.Range16{Lo: 0x31A0, Hi: 0x31BF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Brahmi: 115 codepoints
var Brahmi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x011000, Hi: 0x01104D, Stride: 1 },
    unicode.Range32{Lo: 0x011052, Hi: 0x011075, Stride: 1 },
    unicode.Range32{Lo: 0x01107F, Hi: 0x01107F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Braille: 256 codepoints
var Braille = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x2800, Hi: 0x28FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Buginese: 30 codepoints
var Buginese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1A00, Hi: 0x1A1B, Stride: 1 },
    unicode.Range16{Lo: 0x1A1E, Hi: 0x1A1F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Buhid: 20 codepoints
var Buhid = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x1740, Hi: 0x1753, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Canadian_Aboriginal: 726 codepoints
var Canadian_Aboriginal = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1400, Hi: 0x167F, Stride: 1 },
    unicode.Range16{Lo: 0x18B0, Hi: 0x18F5, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011AB0, Hi: 0x011ABF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Carian: 49 codepoints
var Carian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x0102A0, Hi: 0x0102D0, Stride: 1 },
  },
  LatinOffset: 0,
}

// Caucasian_Albanian: 53 codepoints
var Caucasian_Albanian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010530, Hi: 0x010563, Stride: 1 },
    unicode.Range32{Lo: 0x01056F, Hi: 0x01056F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Chakma: 71 codepoints
var Chakma = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011100, Hi: 0x011134, Stride: 1 },
    unicode.Range32{Lo: 0x011136, Hi: 0x011147, Stride: 1 },
  },
  LatinOffset: 0,
}

// Cham: 83 codepoints
var Cham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0xAA00, Hi: 0xAA36, Stride: 1 },
    unicode.Range16{Lo: 0xAA40, Hi: 0xAA4D, Stride: 1 },
    unicode.Range16{Lo: 0xAA50, Hi: 0xAA59, Stride: 1 },
    unicode.Range16{Lo: 0xAA5C, Hi: 0xAA5F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Cherokee: 172 codepoints
var Cherokee = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x13A0, Hi: 0x13F5, Stride: 1 },
    unicode.Range16{Lo: 0x13F8, Hi: 0x13FD, Stride: 1 },
    unicode.Range16{Lo: 0xAB70, Hi: 0xABBF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Chorasmian: 28 codepoints
var Chorasmian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010FB0, Hi: 0x010FCB, Stride: 1 },
  },
  LatinOffset: 0,
}

// Common: 8301 codepoints
var Common = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 82 */
    unicode.Range16{Lo: 0x0000, Hi: 0x0040, Stride: 1 },
    unicode.Range16{Lo: 0x005B, Hi: 0x0060, Stride: 1 },
    unicode.Range16{Lo: 0x007B, Hi: 0x00A9, Stride: 1 },
    unicode.Range16{Lo: 0x00AB, Hi: 0x00B9, Stride: 1 },
    unicode.Range16{Lo: 0x00BB, Hi: 0x00BF, Stride: 1 },
    unicode.Range16{Lo: 0x00D7, Hi: 0x00F7, Stride: 32 },
    unicode.Range16{Lo: 0x02B9, Hi: 0x02DF, Stride: 1 },
    unicode.Range16{Lo: 0x02E5, Hi: 0x02E9, Stride: 1 },
    unicode.Range16{Lo: 0x02EC, Hi: 0x02FF, Stride: 1 },
    unicode.Range16{Lo: 0x0374, Hi: 0x037E, Stride: 10 },
    unicode.Range16{Lo: 0x0385, Hi: 0x0387, Stride: 2 },
    unicode.Range16{Lo: 0x0605, Hi: 0x060C, Stride: 7 },
    unicode.Range16{Lo: 0x061B, Hi: 0x061F, Stride: 4 },
    unicode.Range16{Lo: 0x0640, Hi: 0x06DD, Stride: 157 },
    unicode.Range16{Lo: 0x08E2, Hi: 0x0964, Stride: 130 },
    unicode.Range16{Lo: 0x0965, Hi: 0x0E3F, Stride: 1242 },
    unicode.Range16{Lo: 0x0FD5, Hi: 0x0FD8, Stride: 1 },
    unicode.Range16{Lo: 0x10FB, Hi: 0x16EB, Stride: 1520 },
    unicode.Range16{Lo: 0x16EC, Hi: 0x16ED, Stride: 1 },
    unicode.Range16{Lo: 0x1735, Hi: 0x1736, Stride: 1 },
    unicode.Range16{Lo: 0x1802, Hi: 0x1803, Stride: 1 },
    unicode.Range16{Lo: 0x1805, Hi: 0x1CD3, Stride: 1230 },
    unicode.Range16{Lo: 0x1CE1, Hi: 0x1CE9, Stride: 8 },
    unicode.Range16{Lo: 0x1CEA, Hi: 0x1CEC, Stride: 1 },
    unicode.Range16{Lo: 0x1CEE, Hi: 0x1CF3, Stride: 1 },
    unicode.Range16{Lo: 0x1CF5, Hi: 0x1CF7, Stride: 1 },
    unicode.Range16{Lo: 0x1CFA, Hi: 0x2000, Stride: 774 },
    unicode.Range16{Lo: 0x2001, Hi: 0x200B, Stride: 1 },
    unicode.Range16{Lo: 0x200E, Hi: 0x2064, Stride: 1 },
    unicode.Range16{Lo: 0x2066, Hi: 0x2070, Stride: 1 },
    unicode.Range16{Lo: 0x2074, Hi: 0x207E, Stride: 1 },
    unicode.Range16{Lo: 0x2080, Hi: 0x208E, Stride: 1 },
    unicode.Range16{Lo: 0x20A0, Hi: 0x20C0, Stride: 1 },
    unicode.Range16{Lo: 0x2100, Hi: 0x2125, Stride: 1 },
    unicode.Range16{Lo: 0x2127, Hi: 0x2129, Stride: 1 },
    unicode.Range16{Lo: 0x212C, Hi: 0x2131, Stride: 1 },
    unicode.Range16{Lo: 0x2133, Hi: 0x214D, Stride: 1 },
    unicode.Range16{Lo: 0x214F, Hi: 0x215F, Stride: 1 },
    unicode.Range16{Lo: 0x2189, Hi: 0x218B, Stride: 1 },
    unicode.Range16{Lo: 0x2190, Hi: 0x2426, Stride: 1 },
    unicode.Range16{Lo: 0x2440, Hi: 0x244A, Stride: 1 },
    unicode.Range16{Lo: 0x2460, Hi: 0x27FF, Stride: 1 },
    unicode.Range16{Lo: 0x2900, Hi: 0x2B73, Stride: 1 },
    unicode.Range16{Lo: 0x2B76, Hi: 0x2B95, Stride: 1 },
    unicode.Range16{Lo: 0x2B97, Hi: 0x2BFF, Stride: 1 },
    unicode.Range16{Lo: 0x2E00, Hi: 0x2E5D, Stride: 1 },
    unicode.Range16{Lo: 0x2FF0, Hi: 0x2FFB, Stride: 1 },
    unicode.Range16{Lo: 0x3000, Hi: 0x3004, Stride: 1 },
    unicode.Range16{Lo: 0x3006, Hi: 0x3008, Stride: 2 },
    unicode.Range16{Lo: 0x3009, Hi: 0x3020, Stride: 1 },
    unicode.Range16{Lo: 0x3030, Hi: 0x3037, Stride: 1 },
    unicode.Range16{Lo: 0x303C, Hi: 0x303F, Stride: 1 },
    unicode.Range16{Lo: 0x309B, Hi: 0x309C, Stride: 1 },
    unicode.Range16{Lo: 0x30A0, Hi: 0x30FB, Stride: 91 },
    unicode.Range16{Lo: 0x30FC, Hi: 0x3190, Stride: 148 },
    unicode.Range16{Lo: 0x3191, Hi: 0x319F, Stride: 1 },
    unicode.Range16{Lo: 0x31C0, Hi: 0x31E3, Stride: 1 },
    unicode.Range16{Lo: 0x3220, Hi: 0x325F, Stride: 1 },
    unicode.Range16{Lo: 0x327F, Hi: 0x32CF, Stride: 1 },
    unicode.Range16{Lo: 0x32FF, Hi: 0x3358, Stride: 89 },
    unicode.Range16{Lo: 0x3359, Hi: 0x33FF, Stride: 1 },
    unicode.Range16{Lo: 0x4DC0, Hi: 0x4DFF, Stride: 1 },
    unicode.Range16{Lo: 0xA700, Hi: 0xA721, Stride: 1 },
    unicode.Range16{Lo: 0xA788, Hi: 0xA78A, Stride: 1 },
    unicode.Range16{Lo: 0xA830, Hi: 0xA839, Stride: 1 },
    unicode.Range16{Lo: 0xA92E, Hi: 0xA9CF, Stride: 161 },
    unicode.Range16{Lo: 0xAB5B, Hi: 0xAB6A, Stride: 15 },
    unicode.Range16{Lo: 0xAB6B, Hi: 0xFD3E, Stride: 20947 },
    unicode.Range16{Lo: 0xFD3F, Hi: 0xFE10, Stride: 209 },
    unicode.Range16{Lo: 0xFE11, Hi: 0xFE19, Stride: 1 },
    unicode.Range16{Lo: 0xFE30, Hi: 0xFE52, Stride: 1 },
    unicode.Range16{Lo: 0xFE54, Hi: 0xFE66, Stride: 1 },
    unicode.Range16{Lo: 0xFE68, Hi: 0xFE6B, Stride: 1 },
    unicode.Range16{Lo: 0xFEFF, Hi: 0xFF01, Stride: 2 },
    unicode.Range16{Lo: 0xFF02, Hi: 0xFF20, Stride: 1 },
    unicode.Range16{Lo: 0xFF3B, Hi: 0xFF40, Stride: 1 },
    unicode.Range16{Lo: 0xFF5B, Hi: 0xFF65, Stride: 1 },
    unicode.Range16{Lo: 0xFF70, Hi: 0xFF9E, Stride: 46 },
    unicode.Range16{Lo: 0xFF9F, Hi: 0xFFE0, Stride: 65 },
    unicode.Range16{Lo: 0xFFE1, Hi: 0xFFE6, Stride: 1 },
    unicode.Range16{Lo: 0xFFE8, Hi: 0xFFEE, Stride: 1 },
    unicode.Range16{Lo: 0xFFF9, Hi: 0xFFFD, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 82 */
    unicode.Range32{Lo: 0x010100, Hi: 0x010102, Stride: 1 },
    unicode.Range32{Lo: 0x010107, Hi: 0x010133, Stride: 1 },
    unicode.Range32{Lo: 0x010137, Hi: 0x01013F, Stride: 1 },
    unicode.Range32{Lo: 0x010190, Hi: 0x01019C, Stride: 1 },
    unicode.Range32{Lo: 0x0101D0, Hi: 0x0101FC, Stride: 1 },
    unicode.Range32{Lo: 0x0102E1, Hi: 0x0102FB, Stride: 1 },
    unicode.Range32{Lo: 0x01BCA0, Hi: 0x01BCA3, Stride: 1 },
    unicode.Range32{Lo: 0x01CF50, Hi: 0x01CFC3, Stride: 1 },
    unicode.Range32{Lo: 0x01D000, Hi: 0x01D0F5, Stride: 1 },
    unicode.Range32{Lo: 0x01D100, Hi: 0x01D126, Stride: 1 },
    unicode.Range32{Lo: 0x01D129, Hi: 0x01D166, Stride: 1 },
    unicode.Range32{Lo: 0x01D16A, Hi: 0x01D17A, Stride: 1 },
    unicode.Range32{Lo: 0x01D183, Hi: 0x01D184, Stride: 1 },
    unicode.Range32{Lo: 0x01D18C, Hi: 0x01D1A9, Stride: 1 },
    unicode.Range32{Lo: 0x01D1AE, Hi: 0x01D1EA, Stride: 1 },
    unicode.Range32{Lo: 0x01D2C0, Hi: 0x01D2D3, Stride: 1 },
    unicode.Range32{Lo: 0x01D2E0, Hi: 0x01D2F3, Stride: 1 },
    unicode.Range32{Lo: 0x01D300, Hi: 0x01D356, Stride: 1 },
    unicode.Range32{Lo: 0x01D360, Hi: 0x01D378, Stride: 1 },
    unicode.Range32{Lo: 0x01D400, Hi: 0x01D454, Stride: 1 },
    unicode.Range32{Lo: 0x01D456, Hi: 0x01D49C, Stride: 1 },
    unicode.Range32{Lo: 0x01D49E, Hi: 0x01D49F, Stride: 1 },
    unicode.Range32{Lo: 0x01D4A2, Hi: 0x01D4A5, Stride: 3 },
    unicode.Range32{Lo: 0x01D4A6, Hi: 0x01D4A9, Stride: 3 },
    unicode.Range32{Lo: 0x01D4AA, Hi: 0x01D4AC, Stride: 1 },
    unicode.Range32{Lo: 0x01D4AE, Hi: 0x01D4B9, Stride: 1 },
    unicode.Range32{Lo: 0x01D4BB, Hi: 0x01D4BD, Stride: 2 },
    unicode.Range32{Lo: 0x01D4BE, Hi: 0x01D4C3, Stride: 1 },
    unicode.Range32{Lo: 0x01D4C5, Hi: 0x01D505, Stride: 1 },
    unicode.Range32{Lo: 0x01D507, Hi: 0x01D50A, Stride: 1 },
    unicode.Range32{Lo: 0x01D50D, Hi: 0x01D514, Stride: 1 },
    unicode.Range32{Lo: 0x01D516, Hi: 0x01D51C, Stride: 1 },
    unicode.Range32{Lo: 0x01D51E, Hi: 0x01D539, Stride: 1 },
    unicode.Range32{Lo: 0x01D53B, Hi: 0x01D53E, Stride: 1 },
    unicode.Range32{Lo: 0x01D540, Hi: 0x01D544, Stride: 1 },
    unicode.Range32{Lo: 0x01D546, Hi: 0x01D54A, Stride: 4 },
    unicode.Range32{Lo: 0x01D54B, Hi: 0x01D550, Stride: 1 },
    unicode.Range32{Lo: 0x01D552, Hi: 0x01D6A5, Stride: 1 },
    unicode.Range32{Lo: 0x01D6A8, Hi: 0x01D7CB, Stride: 1 },
    unicode.Range32{Lo: 0x01D7CE, Hi: 0x01D7FF, Stride: 1 },
    unicode.Range32{Lo: 0x01EC71, Hi: 0x01ECB4, Stride: 1 },
    unicode.Range32{Lo: 0x01ED01, Hi: 0x01ED3D, Stride: 1 },
    unicode.Range32{Lo: 0x01F000, Hi: 0x01F02B, Stride: 1 },
    unicode.Range32{Lo: 0x01F030, Hi: 0x01F093, Stride: 1 },
    unicode.Range32{Lo: 0x01F0A0, Hi: 0x01F0AE, Stride: 1 },
    unicode.Range32{Lo: 0x01F0B1, Hi: 0x01F0BF, Stride: 1 },
    unicode.Range32{Lo: 0x01F0C1, Hi: 0x01F0CF, Stride: 1 },
    unicode.Range32{Lo: 0x01F0D1, Hi: 0x01F0F5, Stride: 1 },
    unicode.Range32{Lo: 0x01F100, Hi: 0x01F1AD, Stride: 1 },
    unicode.Range32{Lo: 0x01F1E6, Hi: 0x01F1FF, Stride: 1 },
    unicode.Range32{Lo: 0x01F201, Hi: 0x01F202, Stride: 1 },
    unicode.Range32{Lo: 0x01F210, Hi: 0x01F23B, Stride: 1 },
    unicode.Range32{Lo: 0x01F240, Hi: 0x01F248, Stride: 1 },
    unicode.Range32{Lo: 0x01F250, Hi: 0x01F251, Stride: 1 },
    unicode.Range32{Lo: 0x01F260, Hi: 0x01F265, Stride: 1 },
    unicode.Range32{Lo: 0x01F300, Hi: 0x01F6D7, Stride: 1 },
    unicode.Range32{Lo: 0x01F6DC, Hi: 0x01F6EC, Stride: 1 },
    unicode.Range32{Lo: 0x01F6F0, Hi: 0x01F6FC, Stride: 1 },
    unicode.Range32{Lo: 0x01F700, Hi: 0x01F776, Stride: 1 },
    unicode.Range32{Lo: 0x01F77B, Hi: 0x01F7D9, Stride: 1 },
    unicode.Range32{Lo: 0x01F7E0, Hi: 0x01F7EB, Stride: 1 },
    unicode.Range32{Lo: 0x01F7F0, Hi: 0x01F800, Stride: 16 },
    unicode.Range32{Lo: 0x01F801, Hi: 0x01F80B, Stride: 1 },
    unicode.Range32{Lo: 0x01F810, Hi: 0x01F847, Stride: 1 },
    unicode.Range32{Lo: 0x01F850, Hi: 0x01F859, Stride: 1 },
    unicode.Range32{Lo: 0x01F860, Hi: 0x01F887, Stride: 1 },
    unicode.Range32{Lo: 0x01F890, Hi: 0x01F8AD, Stride: 1 },
    unicode.Range32{Lo: 0x01F8B0, Hi: 0x01F8B1, Stride: 1 },
    unicode.Range32{Lo: 0x01F900, Hi: 0x01FA53, Stride: 1 },
    unicode.Range32{Lo: 0x01FA60, Hi: 0x01FA6D, Stride: 1 },
    unicode.Range32{Lo: 0x01FA70, Hi: 0x01FA7C, Stride: 1 },
    unicode.Range32{Lo: 0x01FA80, Hi: 0x01FA88, Stride: 1 },
    unicode.Range32{Lo: 0x01FA90, Hi: 0x01FABD, Stride: 1 },
    unicode.Range32{Lo: 0x01FABF, Hi: 0x01FAC5, Stride: 1 },
    unicode.Range32{Lo: 0x01FACE, Hi: 0x01FADB, Stride: 1 },
    unicode.Range32{Lo: 0x01FAE0, Hi: 0x01FAE8, Stride: 1 },
    unicode.Range32{Lo: 0x01FAF0, Hi: 0x01FAF8, Stride: 1 },
    unicode.Range32{Lo: 0x01FB00, Hi: 0x01FB92, Stride: 1 },
    unicode.Range32{Lo: 0x01FB94, Hi: 0x01FBCA, Stride: 1 },
    unicode.Range32{Lo: 0x01FBF0, Hi: 0x01FBF9, Stride: 1 },
    unicode.Range32{Lo: 0x0E0001, Hi: 0x0E0020, Stride: 31 },
    unicode.Range32{Lo: 0x0E0021, Hi: 0x0E007F, Stride: 1 },
  },
  LatinOffset: 6,
}

// Coptic: 137 codepoints
var Coptic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x03E2, Hi: 0x03EF, Stride: 1 },
    unicode.Range16{Lo: 0x2C80, Hi: 0x2CF3, Stride: 1 },
    unicode.Range16{Lo: 0x2CF9, Hi: 0x2CFF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Cuneiform: 1234 codepoints
var Cuneiform = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 0x012000, Hi: 0x012399, Stride: 1 },
    unicode.Range32{Lo: 0x012400, Hi: 0x01246E, Stride: 1 },
    unicode.Range32{Lo: 0x012470, Hi: 0x012474, Stride: 1 },
    unicode.Range32{Lo: 0x012480, Hi: 0x012543, Stride: 1 },
  },
  LatinOffset: 0,
}

// Cypriot: 55 codepoints
var Cypriot = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 0x010800, Hi: 0x010805, Stride: 1 },
    unicode.Range32{Lo: 0x010808, Hi: 0x01080A, Stride: 2 },
    unicode.Range32{Lo: 0x01080B, Hi: 0x010835, Stride: 1 },
    unicode.Range32{Lo: 0x010837, Hi: 0x010838, Stride: 1 },
    unicode.Range32{Lo: 0x01083C, Hi: 0x01083F, Stride: 3 },
  },
  LatinOffset: 0,
}

// Cypro_Minoan: 99 codepoints
var Cypro_Minoan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x012F90, Hi: 0x012FF2, Stride: 1 },
  },
  LatinOffset: 0,
}

// Cyrillic: 506 codepoints
var Cyrillic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0x0400, Hi: 0x0484, Stride: 1 },
    unicode.Range16{Lo: 0x0487, Hi: 0x052F, Stride: 1 },
    unicode.Range16{Lo: 0x1C80, Hi: 0x1C88, Stride: 1 },
    unicode.Range16{Lo: 0x1D2B, Hi: 0x1D78, Stride: 77 },
    unicode.Range16{Lo: 0x2DE0, Hi: 0x2DFF, Stride: 1 },
    unicode.Range16{Lo: 0xA640, Hi: 0xA69F, Stride: 1 },
    unicode.Range16{Lo: 0xFE2E, Hi: 0xFE2F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x01E030, Hi: 0x01E06D, Stride: 1 },
    unicode.Range32{Lo: 0x01E08F, Hi: 0x01E08F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Deseret: 80 codepoints
var Deseret = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010400, Hi: 0x01044F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Devanagari: 164 codepoints
var Devanagari = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0x0900, Hi: 0x0950, Stride: 1 },
    unicode.Range16{Lo: 0x0955, Hi: 0x0963, Stride: 1 },
    unicode.Range16{Lo: 0x0966, Hi: 0x097F, Stride: 1 },
    unicode.Range16{Lo: 0xA8E0, Hi: 0xA8FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011B00, Hi: 0x011B09, Stride: 1 },
  },
  LatinOffset: 0,
}

// Dives_Akuru: 72 codepoints
var Dives_Akuru = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 8 */
    unicode.Range32{Lo: 0x011900, Hi: 0x011906, Stride: 1 },
    unicode.Range32{Lo: 0x011909, Hi: 0x01190C, Stride: 3 },
    unicode.Range32{Lo: 0x01190D, Hi: 0x011913, Stride: 1 },
    unicode.Range32{Lo: 0x011915, Hi: 0x011916, Stride: 1 },
    unicode.Range32{Lo: 0x011918, Hi: 0x011935, Stride: 1 },
    unicode.Range32{Lo: 0x011937, Hi: 0x011938, Stride: 1 },
    unicode.Range32{Lo: 0x01193B, Hi: 0x011946, Stride: 1 },
    unicode.Range32{Lo: 0x011950, Hi: 0x011959, Stride: 1 },
  },
  LatinOffset: 0,
}

// Dogra: 60 codepoints
var Dogra = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011800, Hi: 0x01183B, Stride: 1 },
  },
  LatinOffset: 0,
}

// Duployan: 143 codepoints
var Duployan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 0x01BC00, Hi: 0x01BC6A, Stride: 1 },
    unicode.Range32{Lo: 0x01BC70, Hi: 0x01BC7C, Stride: 1 },
    unicode.Range32{Lo: 0x01BC80, Hi: 0x01BC88, Stride: 1 },
    unicode.Range32{Lo: 0x01BC90, Hi: 0x01BC99, Stride: 1 },
    unicode.Range32{Lo: 0x01BC9C, Hi: 0x01BC9F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Egyptian_Hieroglyphs: 1110 codepoints
var Egyptian_Hieroglyphs = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x013000, Hi: 0x013455, Stride: 1 },
  },
  LatinOffset: 0,
}

// Elbasan: 40 codepoints
var Elbasan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010500, Hi: 0x010527, Stride: 1 },
  },
  LatinOffset: 0,
}

// Elymaic: 23 codepoints
var Elymaic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010FE0, Hi: 0x010FF6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Ethiopic: 523 codepoints
var Ethiopic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 32 */
    unicode.Range16{Lo: 0x1200, Hi: 0x1248, Stride: 1 },
    unicode.Range16{Lo: 0x124A, Hi: 0x124D, Stride: 1 },
    unicode.Range16{Lo: 0x1250, Hi: 0x1256, Stride: 1 },
    unicode.Range16{Lo: 0x1258, Hi: 0x125A, Stride: 2 },
    unicode.Range16{Lo: 0x125B, Hi: 0x125D, Stride: 1 },
    unicode.Range16{Lo: 0x1260, Hi: 0x1288, Stride: 1 },
    unicode.Range16{Lo: 0x128A, Hi: 0x128D, Stride: 1 },
    unicode.Range16{Lo: 0x1290, Hi: 0x12B0, Stride: 1 },
    unicode.Range16{Lo: 0x12B2, Hi: 0x12B5, Stride: 1 },
    unicode.Range16{Lo: 0x12B8, Hi: 0x12BE, Stride: 1 },
    unicode.Range16{Lo: 0x12C0, Hi: 0x12C2, Stride: 2 },
    unicode.Range16{Lo: 0x12C3, Hi: 0x12C5, Stride: 1 },
    unicode.Range16{Lo: 0x12C8, Hi: 0x12D6, Stride: 1 },
    unicode.Range16{Lo: 0x12D8, Hi: 0x1310, Stride: 1 },
    unicode.Range16{Lo: 0x1312, Hi: 0x1315, Stride: 1 },
    unicode.Range16{Lo: 0x1318, Hi: 0x135A, Stride: 1 },
    unicode.Range16{Lo: 0x135D, Hi: 0x137C, Stride: 1 },
    unicode.Range16{Lo: 0x1380, Hi: 0x1399, Stride: 1 },
    unicode.Range16{Lo: 0x2D80, Hi: 0x2D96, Stride: 1 },
    unicode.Range16{Lo: 0x2DA0, Hi: 0x2DA6, Stride: 1 },
    unicode.Range16{Lo: 0x2DA8, Hi: 0x2DAE, Stride: 1 },
    unicode.Range16{Lo: 0x2DB0, Hi: 0x2DB6, Stride: 1 },
    unicode.Range16{Lo: 0x2DB8, Hi: 0x2DBE, Stride: 1 },
    unicode.Range16{Lo: 0x2DC0, Hi: 0x2DC6, Stride: 1 },
    unicode.Range16{Lo: 0x2DC8, Hi: 0x2DCE, Stride: 1 },
    unicode.Range16{Lo: 0x2DD0, Hi: 0x2DD6, Stride: 1 },
    unicode.Range16{Lo: 0x2DD8, Hi: 0x2DDE, Stride: 1 },
    unicode.Range16{Lo: 0xAB01, Hi: 0xAB06, Stride: 1 },
    unicode.Range16{Lo: 0xAB09, Hi: 0xAB0E, Stride: 1 },
    unicode.Range16{Lo: 0xAB11, Hi: 0xAB16, Stride: 1 },
    unicode.Range16{Lo: 0xAB20, Hi: 0xAB26, Stride: 1 },
    unicode.Range16{Lo: 0xAB28, Hi: 0xAB2E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 0x01E7E0, Hi: 0x01E7E6, Stride: 1 },
    unicode.Range32{Lo: 0x01E7E8, Hi: 0x01E7EB, Stride: 1 },
    unicode.Range32{Lo: 0x01E7ED, Hi: 0x01E7EE, Stride: 1 },
    unicode.Range32{Lo: 0x01E7F0, Hi: 0x01E7FE, Stride: 1 },
  },
  LatinOffset: 0,
}

// Georgian: 173 codepoints
var Georgian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 8 */
    unicode.Range16{Lo: 0x10A0, Hi: 0x10C5, Stride: 1 },
    unicode.Range16{Lo: 0x10C7, Hi: 0x10CD, Stride: 6 },
    unicode.Range16{Lo: 0x10D0, Hi: 0x10FA, Stride: 1 },
    unicode.Range16{Lo: 0x10FC, Hi: 0x10FF, Stride: 1 },
    unicode.Range16{Lo: 0x1C90, Hi: 0x1CBA, Stride: 1 },
    unicode.Range16{Lo: 0x1CBD, Hi: 0x1CBF, Stride: 1 },
    unicode.Range16{Lo: 0x2D00, Hi: 0x2D25, Stride: 1 },
    unicode.Range16{Lo: 0x2D27, Hi: 0x2D2D, Stride: 6 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Glagolitic: 134 codepoints
var Glagolitic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x2C00, Hi: 0x2C5F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 0x01E000, Hi: 0x01E006, Stride: 1 },
    unicode.Range32{Lo: 0x01E008, Hi: 0x01E018, Stride: 1 },
    unicode.Range32{Lo: 0x01E01B, Hi: 0x01E021, Stride: 1 },
    unicode.Range32{Lo: 0x01E023, Hi: 0x01E024, Stride: 1 },
    unicode.Range32{Lo: 0x01E026, Hi: 0x01E02A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gothic: 27 codepoints
var Gothic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010330, Hi: 0x01034A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Grantha: 85 codepoints
var Grantha = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 14 */
    unicode.Range32{Lo: 0x011300, Hi: 0x011303, Stride: 1 },
    unicode.Range32{Lo: 0x011305, Hi: 0x01130C, Stride: 1 },
    unicode.Range32{Lo: 0x01130F, Hi: 0x011310, Stride: 1 },
    unicode.Range32{Lo: 0x011313, Hi: 0x011328, Stride: 1 },
    unicode.Range32{Lo: 0x01132A, Hi: 0x011330, Stride: 1 },
    unicode.Range32{Lo: 0x011332, Hi: 0x011333, Stride: 1 },
    unicode.Range32{Lo: 0x011335, Hi: 0x011339, Stride: 1 },
    unicode.Range32{Lo: 0x01133C, Hi: 0x011344, Stride: 1 },
    unicode.Range32{Lo: 0x011347, Hi: 0x011348, Stride: 1 },
    unicode.Range32{Lo: 0x01134B, Hi: 0x01134D, Stride: 1 },
    unicode.Range32{Lo: 0x011350, Hi: 0x011357, Stride: 7 },
    unicode.Range32{Lo: 0x01135D, Hi: 0x011363, Stride: 1 },
    unicode.Range32{Lo: 0x011366, Hi: 0x01136C, Stride: 1 },
    unicode.Range32{Lo: 0x011370, Hi: 0x011374, Stride: 1 },
  },
  LatinOffset: 0,
}

// Greek: 518 codepoints
var Greek = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 29 */
    unicode.Range16{Lo: 0x0370, Hi: 0x0373, Stride: 1 },
    unicode.Range16{Lo: 0x0375, Hi: 0x0377, Stride: 1 },
    unicode.Range16{Lo: 0x037A, Hi: 0x037D, Stride: 1 },
    unicode.Range16{Lo: 0x037F, Hi: 0x0384, Stride: 5 },
    unicode.Range16{Lo: 0x0386, Hi: 0x0388, Stride: 2 },
    unicode.Range16{Lo: 0x0389, Hi: 0x038A, Stride: 1 },
    unicode.Range16{Lo: 0x038C, Hi: 0x038E, Stride: 2 },
    unicode.Range16{Lo: 0x038F, Hi: 0x03A1, Stride: 1 },
    unicode.Range16{Lo: 0x03A3, Hi: 0x03E1, Stride: 1 },
    unicode.Range16{Lo: 0x03F0, Hi: 0x03FF, Stride: 1 },
    unicode.Range16{Lo: 0x1D26, Hi: 0x1D2A, Stride: 1 },
    unicode.Range16{Lo: 0x1D5D, Hi: 0x1D61, Stride: 1 },
    unicode.Range16{Lo: 0x1D66, Hi: 0x1D6A, Stride: 1 },
    unicode.Range16{Lo: 0x1DBF, Hi: 0x1F00, Stride: 321 },
    unicode.Range16{Lo: 0x1F01, Hi: 0x1F15, Stride: 1 },
    unicode.Range16{Lo: 0x1F18, Hi: 0x1F1D, Stride: 1 },
    unicode.Range16{Lo: 0x1F20, Hi: 0x1F45, Stride: 1 },
    unicode.Range16{Lo: 0x1F48, Hi: 0x1F4D, Stride: 1 },
    unicode.Range16{Lo: 0x1F50, Hi: 0x1F57, Stride: 1 },
    unicode.Range16{Lo: 0x1F59, Hi: 0x1F5F, Stride: 2 },
    unicode.Range16{Lo: 0x1F60, Hi: 0x1F7D, Stride: 1 },
    unicode.Range16{Lo: 0x1F80, Hi: 0x1FB4, Stride: 1 },
    unicode.Range16{Lo: 0x1FB6, Hi: 0x1FC4, Stride: 1 },
    unicode.Range16{Lo: 0x1FC6, Hi: 0x1FD3, Stride: 1 },
    unicode.Range16{Lo: 0x1FD6, Hi: 0x1FDB, Stride: 1 },
    unicode.Range16{Lo: 0x1FDD, Hi: 0x1FEF, Stride: 1 },
    unicode.Range16{Lo: 0x1FF2, Hi: 0x1FF4, Stride: 1 },
    unicode.Range16{Lo: 0x1FF6, Hi: 0x1FFE, Stride: 1 },
    unicode.Range16{Lo: 0x2126, Hi: 0xAB65, Stride: 35391 },
  },
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x010140, Hi: 0x01018E, Stride: 1 },
    unicode.Range32{Lo: 0x0101A0, Hi: 0x01D200, Stride: 53344 },
    unicode.Range32{Lo: 0x01D201, Hi: 0x01D245, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gujarati: 91 codepoints
var Gujarati = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0x0A81, Hi: 0x0A83, Stride: 1 },
    unicode.Range16{Lo: 0x0A85, Hi: 0x0A8D, Stride: 1 },
    unicode.Range16{Lo: 0x0A8F, Hi: 0x0A91, Stride: 1 },
    unicode.Range16{Lo: 0x0A93, Hi: 0x0AA8, Stride: 1 },
    unicode.Range16{Lo: 0x0AAA, Hi: 0x0AB0, Stride: 1 },
    unicode.Range16{Lo: 0x0AB2, Hi: 0x0AB3, Stride: 1 },
    unicode.Range16{Lo: 0x0AB5, Hi: 0x0AB9, Stride: 1 },
    unicode.Range16{Lo: 0x0ABC, Hi: 0x0AC5, Stride: 1 },
    unicode.Range16{Lo: 0x0AC7, Hi: 0x0AC9, Stride: 1 },
    unicode.Range16{Lo: 0x0ACB, Hi: 0x0ACD, Stride: 1 },
    unicode.Range16{Lo: 0x0AD0, Hi: 0x0AE0, Stride: 16 },
    unicode.Range16{Lo: 0x0AE1, Hi: 0x0AE3, Stride: 1 },
    unicode.Range16{Lo: 0x0AE6, Hi: 0x0AF1, Stride: 1 },
    unicode.Range16{Lo: 0x0AF9, Hi: 0x0AFF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Gunjala_Gondi: 63 codepoints
var Gunjala_Gondi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 6 */
    unicode.Range32{Lo: 0x011D60, Hi: 0x011D65, Stride: 1 },
    unicode.Range32{Lo: 0x011D67, Hi: 0x011D68, Stride: 1 },
    unicode.Range32{Lo: 0x011D6A, Hi: 0x011D8E, Stride: 1 },
    unicode.Range32{Lo: 0x011D90, Hi: 0x011D91, Stride: 1 },
    unicode.Range32{Lo: 0x011D93, Hi: 0x011D98, Stride: 1 },
    unicode.Range32{Lo: 0x011DA0, Hi: 0x011DA9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gurmukhi: 80 codepoints
var Gurmukhi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 16 */
    unicode.Range16{Lo: 0x0A01, Hi: 0x0A03, Stride: 1 },
    unicode.Range16{Lo: 0x0A05, Hi: 0x0A0A, Stride: 1 },
    unicode.Range16{Lo: 0x0A0F, Hi: 0x0A10, Stride: 1 },
    unicode.Range16{Lo: 0x0A13, Hi: 0x0A28, Stride: 1 },
    unicode.Range16{Lo: 0x0A2A, Hi: 0x0A30, Stride: 1 },
    unicode.Range16{Lo: 0x0A32, Hi: 0x0A33, Stride: 1 },
    unicode.Range16{Lo: 0x0A35, Hi: 0x0A36, Stride: 1 },
    unicode.Range16{Lo: 0x0A38, Hi: 0x0A39, Stride: 1 },
    unicode.Range16{Lo: 0x0A3C, Hi: 0x0A3E, Stride: 2 },
    unicode.Range16{Lo: 0x0A3F, Hi: 0x0A42, Stride: 1 },
    unicode.Range16{Lo: 0x0A47, Hi: 0x0A48, Stride: 1 },
    unicode.Range16{Lo: 0x0A4B, Hi: 0x0A4D, Stride: 1 },
    unicode.Range16{Lo: 0x0A51, Hi: 0x0A59, Stride: 8 },
    unicode.Range16{Lo: 0x0A5A, Hi: 0x0A5C, Stride: 1 },
    unicode.Range16{Lo: 0x0A5E, Hi: 0x0A66, Stride: 8 },
    unicode.Range16{Lo: 0x0A67, Hi: 0x0A76, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Han: 98408 codepoints
var Han = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 10 */
    unicode.Range16{Lo: 0x2E80, Hi: 0x2E99, Stride: 1 },
    unicode.Range16{Lo: 0x2E9B, Hi: 0x2EF3, Stride: 1 },
    unicode.Range16{Lo: 0x2F00, Hi: 0x2FD5, Stride: 1 },
    unicode.Range16{Lo: 0x3005, Hi: 0x3007, Stride: 2 },
    unicode.Range16{Lo: 0x3021, Hi: 0x3029, Stride: 1 },
    unicode.Range16{Lo: 0x3038, Hi: 0x303B, Stride: 1 },
    unicode.Range16{Lo: 0x3400, Hi: 0x4DBF, Stride: 1 },
    unicode.Range16{Lo: 0x4E00, Hi: 0x9FFF, Stride: 1 },
    unicode.Range16{Lo: 0xF900, Hi: 0xFA6D, Stride: 1 },
    unicode.Range16{Lo: 0xFA70, Hi: 0xFAD9, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 10 */
    unicode.Range32{Lo: 0x016FE2, Hi: 0x016FE3, Stride: 1 },
    unicode.Range32{Lo: 0x016FF0, Hi: 0x016FF1, Stride: 1 },
    unicode.Range32{Lo: 0x020000, Hi: 0x02A6DF, Stride: 1 },
    unicode.Range32{Lo: 0x02A700, Hi: 0x02B739, Stride: 1 },
    unicode.Range32{Lo: 0x02B740, Hi: 0x02B81D, Stride: 1 },
    unicode.Range32{Lo: 0x02B820, Hi: 0x02CEA1, Stride: 1 },
    unicode.Range32{Lo: 0x02CEB0, Hi: 0x02EBE0, Stride: 1 },
    unicode.Range32{Lo: 0x02F800, Hi: 0x02FA1D, Stride: 1 },
    unicode.Range32{Lo: 0x030000, Hi: 0x03134A, Stride: 1 },
    unicode.Range32{Lo: 0x031350, Hi: 0x0323AF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hangul: 11739 codepoints
var Hangul = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0x1100, Hi: 0x11FF, Stride: 1 },
    unicode.Range16{Lo: 0x302E, Hi: 0x302F, Stride: 1 },
    unicode.Range16{Lo: 0x3131, Hi: 0x318E, Stride: 1 },
    unicode.Range16{Lo: 0x3200, Hi: 0x321E, Stride: 1 },
    unicode.Range16{Lo: 0x3260, Hi: 0x327E, Stride: 1 },
    unicode.Range16{Lo: 0xA960, Hi: 0xA97C, Stride: 1 },
    unicode.Range16{Lo: 0xAC00, Hi: 0xD7A3, Stride: 1 },
    unicode.Range16{Lo: 0xD7B0, Hi: 0xD7C6, Stride: 1 },
    unicode.Range16{Lo: 0xD7CB, Hi: 0xD7FB, Stride: 1 },
    unicode.Range16{Lo: 0xFFA0, Hi: 0xFFBE, Stride: 1 },
    unicode.Range16{Lo: 0xFFC2, Hi: 0xFFC7, Stride: 1 },
    unicode.Range16{Lo: 0xFFCA, Hi: 0xFFCF, Stride: 1 },
    unicode.Range16{Lo: 0xFFD2, Hi: 0xFFD7, Stride: 1 },
    unicode.Range16{Lo: 0xFFDA, Hi: 0xFFDC, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Hanifi_Rohingya: 50 codepoints
var Hanifi_Rohingya = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010D00, Hi: 0x010D27, Stride: 1 },
    unicode.Range32{Lo: 0x010D30, Hi: 0x010D39, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hanunoo: 21 codepoints
var Hanunoo = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x1720, Hi: 0x1734, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Hatran: 26 codepoints
var Hatran = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x0108E0, Hi: 0x0108F2, Stride: 1 },
    unicode.Range32{Lo: 0x0108F4, Hi: 0x0108F5, Stride: 1 },
    unicode.Range32{Lo: 0x0108FB, Hi: 0x0108FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hebrew: 134 codepoints
var Hebrew = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 9 */
    unicode.Range16{Lo: 0x0591, Hi: 0x05C7, Stride: 1 },
    unicode.Range16{Lo: 0x05D0, Hi: 0x05EA, Stride: 1 },
    unicode.Range16{Lo: 0x05EF, Hi: 0x05F4, Stride: 1 },
    unicode.Range16{Lo: 0xFB1D, Hi: 0xFB36, Stride: 1 },
    unicode.Range16{Lo: 0xFB38, Hi: 0xFB3C, Stride: 1 },
    unicode.Range16{Lo: 0xFB3E, Hi: 0xFB40, Stride: 2 },
    unicode.Range16{Lo: 0xFB41, Hi: 0xFB43, Stride: 2 },
    unicode.Range16{Lo: 0xFB44, Hi: 0xFB46, Stride: 2 },
    unicode.Range16{Lo: 0xFB47, Hi: 0xFB4F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Hiragana: 381 codepoints
var Hiragana = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x3041, Hi: 0x3096, Stride: 1 },
    unicode.Range16{Lo: 0x309D, Hi: 0x309F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 0x01B001, Hi: 0x01B11F, Stride: 1 },
    unicode.Range32{Lo: 0x01B132, Hi: 0x01B150, Stride: 30 },
    unicode.Range32{Lo: 0x01B151, Hi: 0x01B152, Stride: 1 },
    unicode.Range32{Lo: 0x01F200, Hi: 0x01F200, Stride: 1 },
  },
  LatinOffset: 0,
}

// Imperial_Aramaic: 31 codepoints
var Imperial_Aramaic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010840, Hi: 0x010855, Stride: 1 },
    unicode.Range32{Lo: 0x010857, Hi: 0x01085F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Inherited: 657 codepoints
var Inherited = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 18 */
    unicode.Range16{Lo: 0x0300, Hi: 0x036F, Stride: 1 },
    unicode.Range16{Lo: 0x0485, Hi: 0x0486, Stride: 1 },
    unicode.Range16{Lo: 0x064B, Hi: 0x0655, Stride: 1 },
    unicode.Range16{Lo: 0x0670, Hi: 0x0951, Stride: 737 },
    unicode.Range16{Lo: 0x0952, Hi: 0x0954, Stride: 1 },
    unicode.Range16{Lo: 0x1AB0, Hi: 0x1ACE, Stride: 1 },
    unicode.Range16{Lo: 0x1CD0, Hi: 0x1CD2, Stride: 1 },
    unicode.Range16{Lo: 0x1CD4, Hi: 0x1CE0, Stride: 1 },
    unicode.Range16{Lo: 0x1CE2, Hi: 0x1CE8, Stride: 1 },
    unicode.Range16{Lo: 0x1CED, Hi: 0x1CF4, Stride: 7 },
    unicode.Range16{Lo: 0x1CF8, Hi: 0x1CF9, Stride: 1 },
    unicode.Range16{Lo: 0x1DC0, Hi: 0x1DFF, Stride: 1 },
    unicode.Range16{Lo: 0x200C, Hi: 0x200D, Stride: 1 },
    unicode.Range16{Lo: 0x20D0, Hi: 0x20F0, Stride: 1 },
    unicode.Range16{Lo: 0x302A, Hi: 0x302D, Stride: 1 },
    unicode.Range16{Lo: 0x3099, Hi: 0x309A, Stride: 1 },
    unicode.Range16{Lo: 0xFE00, Hi: 0xFE0F, Stride: 1 },
    unicode.Range16{Lo: 0xFE20, Hi: 0xFE2D, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 9 */
    unicode.Range32{Lo: 0x0101FD, Hi: 0x0102E0, Stride: 227 },
    unicode.Range32{Lo: 0x01133B, Hi: 0x01CF00, Stride: 48069 },
    unicode.Range32{Lo: 0x01CF01, Hi: 0x01CF2D, Stride: 1 },
    unicode.Range32{Lo: 0x01CF30, Hi: 0x01CF46, Stride: 1 },
    unicode.Range32{Lo: 0x01D167, Hi: 0x01D169, Stride: 1 },
    unicode.Range32{Lo: 0x01D17B, Hi: 0x01D182, Stride: 1 },
    unicode.Range32{Lo: 0x01D185, Hi: 0x01D18B, Stride: 1 },
    unicode.Range32{Lo: 0x01D1AA, Hi: 0x01D1AD, Stride: 1 },
    unicode.Range32{Lo: 0x0E0100, Hi: 0x0E01EF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Inscriptional_Pahlavi: 27 codepoints
var Inscriptional_Pahlavi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010B60, Hi: 0x010B72, Stride: 1 },
    unicode.Range32{Lo: 0x010B78, Hi: 0x010B7F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Inscriptional_Parthian: 30 codepoints
var Inscriptional_Parthian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010B40, Hi: 0x010B55, Stride: 1 },
    unicode.Range32{Lo: 0x010B58, Hi: 0x010B5F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Javanese: 90 codepoints
var Javanese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0xA980, Hi: 0xA9CD, Stride: 1 },
    unicode.Range16{Lo: 0xA9D0, Hi: 0xA9D9, Stride: 1 },
    unicode.Range16{Lo: 0xA9DE, Hi: 0xA9DF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Kaithi: 68 codepoints
var Kaithi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011080, Hi: 0x0110C2, Stride: 1 },
    unicode.Range32{Lo: 0x0110CD, Hi: 0x0110CD, Stride: 1 },
  },
  LatinOffset: 0,
}

// Kannada: 91 codepoints
var Kannada = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 13 */
    unicode.Range16{Lo: 0x0C80, Hi: 0x0C8C, Stride: 1 },
    unicode.Range16{Lo: 0x0C8E, Hi: 0x0C90, Stride: 1 },
    unicode.Range16{Lo: 0x0C92, Hi: 0x0CA8, Stride: 1 },
    unicode.Range16{Lo: 0x0CAA, Hi: 0x0CB3, Stride: 1 },
    unicode.Range16{Lo: 0x0CB5, Hi: 0x0CB9, Stride: 1 },
    unicode.Range16{Lo: 0x0CBC, Hi: 0x0CC4, Stride: 1 },
    unicode.Range16{Lo: 0x0CC6, Hi: 0x0CC8, Stride: 1 },
    unicode.Range16{Lo: 0x0CCA, Hi: 0x0CCD, Stride: 1 },
    unicode.Range16{Lo: 0x0CD5, Hi: 0x0CD6, Stride: 1 },
    unicode.Range16{Lo: 0x0CDD, Hi: 0x0CDE, Stride: 1 },
    unicode.Range16{Lo: 0x0CE0, Hi: 0x0CE3, Stride: 1 },
    unicode.Range16{Lo: 0x0CE6, Hi: 0x0CEF, Stride: 1 },
    unicode.Range16{Lo: 0x0CF1, Hi: 0x0CF3, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Katakana: 321 codepoints
var Katakana = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0x30A1, Hi: 0x30FA, Stride: 1 },
    unicode.Range16{Lo: 0x30FD, Hi: 0x30FF, Stride: 1 },
    unicode.Range16{Lo: 0x31F0, Hi: 0x31FF, Stride: 1 },
    unicode.Range16{Lo: 0x32D0, Hi: 0x32FE, Stride: 1 },
    unicode.Range16{Lo: 0x3300, Hi: 0x3357, Stride: 1 },
    unicode.Range16{Lo: 0xFF66, Hi: 0xFF6F, Stride: 1 },
    unicode.Range16{Lo: 0xFF71, Hi: 0xFF9D, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 7 */
    unicode.Range32{Lo: 0x01AFF0, Hi: 0x01AFF3, Stride: 1 },
    unicode.Range32{Lo: 0x01AFF5, Hi: 0x01AFFB, Stride: 1 },
    unicode.Range32{Lo: 0x01AFFD, Hi: 0x01AFFE, Stride: 1 },
    unicode.Range32{Lo: 0x01B000, Hi: 0x01B120, Stride: 288 },
    unicode.Range32{Lo: 0x01B121, Hi: 0x01B122, Stride: 1 },
    unicode.Range32{Lo: 0x01B155, Hi: 0x01B164, Stride: 15 },
    unicode.Range32{Lo: 0x01B165, Hi: 0x01B167, Stride: 1 },
  },
  LatinOffset: 0,
}

// Kawi: 86 codepoints
var Kawi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x011F00, Hi: 0x011F10, Stride: 1 },
    unicode.Range32{Lo: 0x011F12, Hi: 0x011F3A, Stride: 1 },
    unicode.Range32{Lo: 0x011F3E, Hi: 0x011F59, Stride: 1 },
  },
  LatinOffset: 0,
}

// Kayah_Li: 47 codepoints
var Kayah_Li = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0xA900, Hi: 0xA92D, Stride: 1 },
    unicode.Range16{Lo: 0xA92F, Hi: 0xA92F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Kharoshthi: 68 codepoints
var Kharoshthi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 8 */
    unicode.Range32{Lo: 0x010A00, Hi: 0x010A03, Stride: 1 },
    unicode.Range32{Lo: 0x010A05, Hi: 0x010A06, Stride: 1 },
    unicode.Range32{Lo: 0x010A0C, Hi: 0x010A13, Stride: 1 },
    unicode.Range32{Lo: 0x010A15, Hi: 0x010A17, Stride: 1 },
    unicode.Range32{Lo: 0x010A19, Hi: 0x010A35, Stride: 1 },
    unicode.Range32{Lo: 0x010A38, Hi: 0x010A3A, Stride: 1 },
    unicode.Range32{Lo: 0x010A3F, Hi: 0x010A48, Stride: 1 },
    unicode.Range32{Lo: 0x010A50, Hi: 0x010A58, Stride: 1 },
  },
  LatinOffset: 0,
}

// Khitan_Small_Script: 471 codepoints
var Khitan_Small_Script = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x016FE4, Hi: 0x018B00, Stride: 6940 },
    unicode.Range32{Lo: 0x018B01, Hi: 0x018CD5, Stride: 1 },
  },
  LatinOffset: 0,
}

// Khmer: 146 codepoints
var Khmer = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0x1780, Hi: 0x17DD, Stride: 1 },
    unicode.Range16{Lo: 0x17E0, Hi: 0x17E9, Stride: 1 },
    unicode.Range16{Lo: 0x17F0, Hi: 0x17F9, Stride: 1 },
    unicode.Range16{Lo: 0x19E0, Hi: 0x19FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Khojki: 65 codepoints
var Khojki = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011200, Hi: 0x011211, Stride: 1 },
    unicode.Range32{Lo: 0x011213, Hi: 0x011241, Stride: 1 },
  },
  LatinOffset: 0,
}

// Khudawadi: 69 codepoints
var Khudawadi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x0112B0, Hi: 0x0112EA, Stride: 1 },
    unicode.Range32{Lo: 0x0112F0, Hi: 0x0112F9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lao: 83 codepoints
var Lao = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 11 */
    unicode.Range16{Lo: 0x0E81, Hi: 0x0E82, Stride: 1 },
    unicode.Range16{Lo: 0x0E84, Hi: 0x0E86, Stride: 2 },
    unicode.Range16{Lo: 0x0E87, Hi: 0x0E8A, Stride: 1 },
    unicode.Range16{Lo: 0x0E8C, Hi: 0x0EA3, Stride: 1 },
    unicode.Range16{Lo: 0x0EA5, Hi: 0x0EA7, Stride: 2 },
    unicode.Range16{Lo: 0x0EA8, Hi: 0x0EBD, Stride: 1 },
    unicode.Range16{Lo: 0x0EC0, Hi: 0x0EC4, Stride: 1 },
    unicode.Range16{Lo: 0x0EC6, Hi: 0x0EC8, Stride: 2 },
    unicode.Range16{Lo: 0x0EC9, Hi: 0x0ECE, Stride: 1 },
    unicode.Range16{Lo: 0x0ED0, Hi: 0x0ED9, Stride: 1 },
    unicode.Range16{Lo: 0x0EDC, Hi: 0x0EDF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Latin: 1481 codepoints
var Latin = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 32 */
    unicode.Range16{Lo: 0x0041, Hi: 0x005A, Stride: 1 },
    unicode.Range16{Lo: 0x0061, Hi: 0x007A, Stride: 1 },
    unicode.Range16{Lo: 0x00AA, Hi: 0x00BA, Stride: 16 },
    unicode.Range16{Lo: 0x00C0, Hi: 0x00D6, Stride: 1 },
    unicode.Range16{Lo: 0x00D8, Hi: 0x00F6, Stride: 1 },
    unicode.Range16{Lo: 0x00F8, Hi: 0x00FF, Stride: 1 },
    unicode.Range16{Lo: 0x0100, Hi: 0x02B8, Stride: 1 },
    unicode.Range16{Lo: 0x02E0, Hi: 0x02E4, Stride: 1 },
    unicode.Range16{Lo: 0x1D00, Hi: 0x1D25, Stride: 1 },
    unicode.Range16{Lo: 0x1D2C, Hi: 0x1D5C, Stride: 1 },
    unicode.Range16{Lo: 0x1D62, Hi: 0x1D65, Stride: 1 },
    unicode.Range16{Lo: 0x1D6B, Hi: 0x1D77, Stride: 1 },
    unicode.Range16{Lo: 0x1D79, Hi: 0x1DBE, Stride: 1 },
    unicode.Range16{Lo: 0x1E00, Hi: 0x1EFF, Stride: 1 },
    unicode.Range16{Lo: 0x2071, Hi: 0x207F, Stride: 14 },
    unicode.Range16{Lo: 0x2090, Hi: 0x209C, Stride: 1 },
    unicode.Range16{Lo: 0x212A, Hi: 0x212B, Stride: 1 },
    unicode.Range16{Lo: 0x2132, Hi: 0x214E, Stride: 28 },
    unicode.Range16{Lo: 0x2160, Hi: 0x2188, Stride: 1 },
    unicode.Range16{Lo: 0x2C60, Hi: 0x2C7F, Stride: 1 },
    unicode.Range16{Lo: 0xA722, Hi: 0xA787, Stride: 1 },
    unicode.Range16{Lo: 0xA78B, Hi: 0xA7CA, Stride: 1 },
    unicode.Range16{Lo: 0xA7D0, Hi: 0xA7D1, Stride: 1 },
    unicode.Range16{Lo: 0xA7D3, Hi: 0xA7D5, Stride: 2 },
    unicode.Range16{Lo: 0xA7D6, Hi: 0xA7D9, Stride: 1 },
    unicode.Range16{Lo: 0xA7F2, Hi: 0xA7FF, Stride: 1 },
    unicode.Range16{Lo: 0xAB30, Hi: 0xAB5A, Stride: 1 },
    unicode.Range16{Lo: 0xAB5C, Hi: 0xAB64, Stride: 1 },
    unicode.Range16{Lo: 0xAB66, Hi: 0xAB69, Stride: 1 },
    unicode.Range16{Lo: 0xFB00, Hi: 0xFB06, Stride: 1 },
    unicode.Range16{Lo: 0xFF21, Hi: 0xFF3A, Stride: 1 },
    unicode.Range16{Lo: 0xFF41, Hi: 0xFF5A, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 0x010780, Hi: 0x010785, Stride: 1 },
    unicode.Range32{Lo: 0x010787, Hi: 0x0107B0, Stride: 1 },
    unicode.Range32{Lo: 0x0107B2, Hi: 0x0107BA, Stride: 1 },
    unicode.Range32{Lo: 0x01DF00, Hi: 0x01DF1E, Stride: 1 },
    unicode.Range32{Lo: 0x01DF25, Hi: 0x01DF2A, Stride: 1 },
  },
  LatinOffset: 6,
}

// Lepcha: 74 codepoints
var Lepcha = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x1C00, Hi: 0x1C37, Stride: 1 },
    unicode.Range16{Lo: 0x1C3B, Hi: 0x1C49, Stride: 1 },
    unicode.Range16{Lo: 0x1C4D, Hi: 0x1C4F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Limbu: 68 codepoints
var Limbu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 5 */
    unicode.Range16{Lo: 0x1900, Hi: 0x191E, Stride: 1 },
    unicode.Range16{Lo: 0x1920, Hi: 0x192B, Stride: 1 },
    unicode.Range16{Lo: 0x1930, Hi: 0x193B, Stride: 1 },
    unicode.Range16{Lo: 0x1940, Hi: 0x1944, Stride: 4 },
    unicode.Range16{Lo: 0x1945, Hi: 0x194F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Linear_A: 341 codepoints
var Linear_A = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x010600, Hi: 0x010736, Stride: 1 },
    unicode.Range32{Lo: 0x010740, Hi: 0x010755, Stride: 1 },
    unicode.Range32{Lo: 0x010760, Hi: 0x010767, Stride: 1 },
  },
  LatinOffset: 0,
}

// Linear_B: 211 codepoints
var Linear_B = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 7 */
    unicode.Range32{Lo: 0x010000, Hi: 0x01000B, Stride: 1 },
    unicode.Range32{Lo: 0x01000D, Hi: 0x010026, Stride: 1 },
    unicode.Range32{Lo: 0x010028, Hi: 0x01003A, Stride: 1 },
    unicode.Range32{Lo: 0x01003C, Hi: 0x01003D, Stride: 1 },
    unicode.Range32{Lo: 0x01003F, Hi: 0x01004D, Stride: 1 },
    unicode.Range32{Lo: 0x010050, Hi: 0x01005D, Stride: 1 },
    unicode.Range32{Lo: 0x010080, Hi: 0x0100FA, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lisu: 49 codepoints
var Lisu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0xA4D0, Hi: 0xA4FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011FB0, Hi: 0x011FB0, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lycian: 29 codepoints
var Lycian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010280, Hi: 0x01029C, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lydian: 27 codepoints
var Lydian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010920, Hi: 0x010939, Stride: 1 },
    unicode.Range32{Lo: 0x01093F, Hi: 0x01093F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Mahajani: 39 codepoints
var Mahajani = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011150, Hi: 0x011176, Stride: 1 },
  },
  LatinOffset: 0,
}

// Makasar: 25 codepoints
var Makasar = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011EE0, Hi: 0x011EF8, Stride: 1 },
  },
  LatinOffset: 0,
}

// Malayalam: 118 codepoints
var Malayalam = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0x0D00, Hi: 0x0D0C, Stride: 1 },
    unicode.Range16{Lo: 0x0D0E, Hi: 0x0D10, Stride: 1 },
    unicode.Range16{Lo: 0x0D12, Hi: 0x0D44, Stride: 1 },
    unicode.Range16{Lo: 0x0D46, Hi: 0x0D48, Stride: 1 },
    unicode.Range16{Lo: 0x0D4A, Hi: 0x0D4F, Stride: 1 },
    unicode.Range16{Lo: 0x0D54, Hi: 0x0D63, Stride: 1 },
    unicode.Range16{Lo: 0x0D66, Hi: 0x0D7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Mandaic: 29 codepoints
var Mandaic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x0840, Hi: 0x085B, Stride: 1 },
    unicode.Range16{Lo: 0x085E, Hi: 0x085E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Manichaean: 51 codepoints
var Manichaean = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010AC0, Hi: 0x010AE6, Stride: 1 },
    unicode.Range32{Lo: 0x010AEB, Hi: 0x010AF6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Marchen: 68 codepoints
var Marchen = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x011C70, Hi: 0x011C8F, Stride: 1 },
    unicode.Range32{Lo: 0x011C92, Hi: 0x011CA7, Stride: 1 },
    unicode.Range32{Lo: 0x011CA9, Hi: 0x011CB6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Masaram_Gondi: 75 codepoints
var Masaram_Gondi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 7 */
    unicode.Range32{Lo: 0x011D00, Hi: 0x011D06, Stride: 1 },
    unicode.Range32{Lo: 0x011D08, Hi: 0x011D09, Stride: 1 },
    unicode.Range32{Lo: 0x011D0B, Hi: 0x011D36, Stride: 1 },
    unicode.Range32{Lo: 0x011D3A, Hi: 0x011D3C, Stride: 2 },
    unicode.Range32{Lo: 0x011D3D, Hi: 0x011D3F, Stride: 2 },
    unicode.Range32{Lo: 0x011D40, Hi: 0x011D47, Stride: 1 },
    unicode.Range32{Lo: 0x011D50, Hi: 0x011D59, Stride: 1 },
  },
  LatinOffset: 0,
}

// Medefaidrin: 91 codepoints
var Medefaidrin = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x016E40, Hi: 0x016E9A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Meetei_Mayek: 79 codepoints
var Meetei_Mayek = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0xAAE0, Hi: 0xAAF6, Stride: 1 },
    unicode.Range16{Lo: 0xABC0, Hi: 0xABED, Stride: 1 },
    unicode.Range16{Lo: 0xABF0, Hi: 0xABF9, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Mende_Kikakui: 213 codepoints
var Mende_Kikakui = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x01E800, Hi: 0x01E8C4, Stride: 1 },
    unicode.Range32{Lo: 0x01E8C7, Hi: 0x01E8D6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Meroitic_Cursive: 90 codepoints
var Meroitic_Cursive = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x0109A0, Hi: 0x0109B7, Stride: 1 },
    unicode.Range32{Lo: 0x0109BC, Hi: 0x0109CF, Stride: 1 },
    unicode.Range32{Lo: 0x0109D2, Hi: 0x0109FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Meroitic_Hieroglyphs: 32 codepoints
var Meroitic_Hieroglyphs = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010980, Hi: 0x01099F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Miao: 149 codepoints
var Miao = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x016F00, Hi: 0x016F4A, Stride: 1 },
    unicode.Range32{Lo: 0x016F4F, Hi: 0x016F87, Stride: 1 },
    unicode.Range32{Lo: 0x016F8F, Hi: 0x016F9F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Modi: 79 codepoints
var Modi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011600, Hi: 0x011644, Stride: 1 },
    unicode.Range32{Lo: 0x011650, Hi: 0x011659, Stride: 1 },
  },
  LatinOffset: 0,
}

// Mongolian: 168 codepoints
var Mongolian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 5 */
    unicode.Range16{Lo: 0x1800, Hi: 0x1801, Stride: 1 },
    unicode.Range16{Lo: 0x1804, Hi: 0x1806, Stride: 2 },
    unicode.Range16{Lo: 0x1807, Hi: 0x1819, Stride: 1 },
    unicode.Range16{Lo: 0x1820, Hi: 0x1878, Stride: 1 },
    unicode.Range16{Lo: 0x1880, Hi: 0x18AA, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011660, Hi: 0x01166C, Stride: 1 },
  },
  LatinOffset: 0,
}

// Mro: 43 codepoints
var Mro = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x016A40, Hi: 0x016A5E, Stride: 1 },
    unicode.Range32{Lo: 0x016A60, Hi: 0x016A69, Stride: 1 },
    unicode.Range32{Lo: 0x016A6E, Hi: 0x016A6F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Multani: 38 codepoints
var Multani = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 0x011280, Hi: 0x011286, Stride: 1 },
    unicode.Range32{Lo: 0x011288, Hi: 0x01128A, Stride: 2 },
    unicode.Range32{Lo: 0x01128B, Hi: 0x01128D, Stride: 1 },
    unicode.Range32{Lo: 0x01128F, Hi: 0x01129D, Stride: 1 },
    unicode.Range32{Lo: 0x01129F, Hi: 0x0112A9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Myanmar: 223 codepoints
var Myanmar = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x1000, Hi: 0x109F, Stride: 1 },
    unicode.Range16{Lo: 0xA9E0, Hi: 0xA9FE, Stride: 1 },
    unicode.Range16{Lo: 0xAA60, Hi: 0xAA7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Nabataean: 40 codepoints
var Nabataean = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010880, Hi: 0x01089E, Stride: 1 },
    unicode.Range32{Lo: 0x0108A7, Hi: 0x0108AF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nag_Mundari: 42 codepoints
var Nag_Mundari = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x01E4D0, Hi: 0x01E4F9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nandinagari: 65 codepoints
var Nandinagari = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x0119A0, Hi: 0x0119A7, Stride: 1 },
    unicode.Range32{Lo: 0x0119AA, Hi: 0x0119D7, Stride: 1 },
    unicode.Range32{Lo: 0x0119DA, Hi: 0x0119E4, Stride: 1 },
  },
  LatinOffset: 0,
}

// New_Tai_Lue: 83 codepoints
var New_Tai_Lue = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0x1980, Hi: 0x19AB, Stride: 1 },
    unicode.Range16{Lo: 0x19B0, Hi: 0x19C9, Stride: 1 },
    unicode.Range16{Lo: 0x19D0, Hi: 0x19DA, Stride: 1 },
    unicode.Range16{Lo: 0x19DE, Hi: 0x19DF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Newa: 97 codepoints
var Newa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011400, Hi: 0x01145B, Stride: 1 },
    unicode.Range32{Lo: 0x01145D, Hi: 0x011461, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nko: 62 codepoints
var Nko = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x07C0, Hi: 0x07FA, Stride: 1 },
    unicode.Range16{Lo: 0x07FD, Hi: 0x07FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Nushu: 397 codepoints
var Nushu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x016FE1, Hi: 0x01B170, Stride: 16783 },
    unicode.Range32{Lo: 0x01B171, Hi: 0x01B2FB, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nyiakeng_Puachue_Hmong: 71 codepoints
var Nyiakeng_Puachue_Hmong = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 0x01E100, Hi: 0x01E12C, Stride: 1 },
    unicode.Range32{Lo: 0x01E130, Hi: 0x01E13D, Stride: 1 },
    unicode.Range32{Lo: 0x01E140, Hi: 0x01E149, Stride: 1 },
    unicode.Range32{Lo: 0x01E14E, Hi: 0x01E14F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Ogham: 29 codepoints
var Ogham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x1680, Hi: 0x169C, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Ol_Chiki: 48 codepoints
var Ol_Chiki = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x1C50, Hi: 0x1C7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Old_Hungarian: 108 codepoints
var Old_Hungarian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x010C80, Hi: 0x010CB2, Stride: 1 },
    unicode.Range32{Lo: 0x010CC0, Hi: 0x010CF2, Stride: 1 },
    unicode.Range32{Lo: 0x010CFA, Hi: 0x010CFF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Italic: 39 codepoints
var Old_Italic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010300, Hi: 0x010323, Stride: 1 },
    unicode.Range32{Lo: 0x01032D, Hi: 0x01032F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_North_Arabian: 32 codepoints
var Old_North_Arabian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010A80, Hi: 0x010A9F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Permic: 43 codepoints
var Old_Permic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010350, Hi: 0x01037A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Persian: 50 codepoints
var Old_Persian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x0103A0, Hi: 0x0103C3, Stride: 1 },
    unicode.Range32{Lo: 0x0103C8, Hi: 0x0103D5, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Sogdian: 40 codepoints
var Old_Sogdian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010F00, Hi: 0x010F27, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_South_Arabian: 32 codepoints
var Old_South_Arabian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010A60, Hi: 0x010A7F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Turkic: 73 codepoints
var Old_Turkic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010C00, Hi: 0x010C48, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Uyghur: 26 codepoints
var Old_Uyghur = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010F70, Hi: 0x010F89, Stride: 1 },
  },
  LatinOffset: 0,
}

// Oriya: 91 codepoints
var Oriya = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0x0B01, Hi: 0x0B03, Stride: 1 },
    unicode.Range16{Lo: 0x0B05, Hi: 0x0B0C, Stride: 1 },
    unicode.Range16{Lo: 0x0B0F, Hi: 0x0B10, Stride: 1 },
    unicode.Range16{Lo: 0x0B13, Hi: 0x0B28, Stride: 1 },
    unicode.Range16{Lo: 0x0B2A, Hi: 0x0B30, Stride: 1 },
    unicode.Range16{Lo: 0x0B32, Hi: 0x0B33, Stride: 1 },
    unicode.Range16{Lo: 0x0B35, Hi: 0x0B39, Stride: 1 },
    unicode.Range16{Lo: 0x0B3C, Hi: 0x0B44, Stride: 1 },
    unicode.Range16{Lo: 0x0B47, Hi: 0x0B48, Stride: 1 },
    unicode.Range16{Lo: 0x0B4B, Hi: 0x0B4D, Stride: 1 },
    unicode.Range16{Lo: 0x0B55, Hi: 0x0B57, Stride: 1 },
    unicode.Range16{Lo: 0x0B5C, Hi: 0x0B5D, Stride: 1 },
    unicode.Range16{Lo: 0x0B5F, Hi: 0x0B63, Stride: 1 },
    unicode.Range16{Lo: 0x0B66, Hi: 0x0B77, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Osage: 72 codepoints
var Osage = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x0104B0, Hi: 0x0104D3, Stride: 1 },
    unicode.Range32{Lo: 0x0104D8, Hi: 0x0104FB, Stride: 1 },
  },
  LatinOffset: 0,
}

// Osmanya: 40 codepoints
var Osmanya = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010480, Hi: 0x01049D, Stride: 1 },
    unicode.Range32{Lo: 0x0104A0, Hi: 0x0104A9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Pahawh_Hmong: 127 codepoints
var Pahawh_Hmong = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 0x016B00, Hi: 0x016B45, Stride: 1 },
    unicode.Range32{Lo: 0x016B50, Hi: 0x016B59, Stride: 1 },
    unicode.Range32{Lo: 0x016B5B, Hi: 0x016B61, Stride: 1 },
    unicode.Range32{Lo: 0x016B63, Hi: 0x016B77, Stride: 1 },
    unicode.Range32{Lo: 0x016B7D, Hi: 0x016B8F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Palmyrene: 32 codepoints
var Palmyrene = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010860, Hi: 0x01087F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Pau_Cin_Hau: 57 codepoints
var Pau_Cin_Hau = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011AC0, Hi: 0x011AF8, Stride: 1 },
  },
  LatinOffset: 0,
}

// Phags_Pa: 56 codepoints
var Phags_Pa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0xA840, Hi: 0xA877, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Phoenician: 29 codepoints
var Phoenician = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010900, Hi: 0x01091B, Stride: 1 },
    unicode.Range32{Lo: 0x01091F, Hi: 0x01091F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Psalter_Pahlavi: 29 codepoints
var Psalter_Pahlavi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x010B80, Hi: 0x010B91, Stride: 1 },
    unicode.Range32{Lo: 0x010B99, Hi: 0x010B9C, Stride: 1 },
    unicode.Range32{Lo: 0x010BA9, Hi: 0x010BAF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Rejang: 37 codepoints
var Rejang = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0xA930, Hi: 0xA953, Stride: 1 },
    unicode.Range16{Lo: 0xA95F, Hi: 0xA95F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Runic: 86 codepoints
var Runic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x16A0, Hi: 0x16EA, Stride: 1 },
    unicode.Range16{Lo: 0x16EE, Hi: 0x16F8, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Samaritan: 61 codepoints
var Samaritan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x0800, Hi: 0x082D, Stride: 1 },
    unicode.Range16{Lo: 0x0830, Hi: 0x083E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Saurashtra: 82 codepoints
var Saurashtra = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0xA880, Hi: 0xA8C5, Stride: 1 },
    unicode.Range16{Lo: 0xA8CE, Hi: 0xA8D9, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Sharada: 96 codepoints
var Sharada = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011180, Hi: 0x0111DF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Shavian: 48 codepoints
var Shavian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010450, Hi: 0x01047F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Siddham: 92 codepoints
var Siddham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011580, Hi: 0x0115B5, Stride: 1 },
    unicode.Range32{Lo: 0x0115B8, Hi: 0x0115DD, Stride: 1 },
  },
  LatinOffset: 0,
}

// SignWriting: 672 codepoints
var SignWriting = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x01D800, Hi: 0x01DA8B, Stride: 1 },
    unicode.Range32{Lo: 0x01DA9B, Hi: 0x01DA9F, Stride: 1 },
    unicode.Range32{Lo: 0x01DAA1, Hi: 0x01DAAF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sinhala: 111 codepoints
var Sinhala = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 12 */
    unicode.Range16{Lo: 0x0D81, Hi: 0x0D83, Stride: 1 },
    unicode.Range16{Lo: 0x0D85, Hi: 0x0D96, Stride: 1 },
    unicode.Range16{Lo: 0x0D9A, Hi: 0x0DB1, Stride: 1 },
    unicode.Range16{Lo: 0x0DB3, Hi: 0x0DBB, Stride: 1 },
    unicode.Range16{Lo: 0x0DBD, Hi: 0x0DC0, Stride: 3 },
    unicode.Range16{Lo: 0x0DC1, Hi: 0x0DC6, Stride: 1 },
    unicode.Range16{Lo: 0x0DCA, Hi: 0x0DCF, Stride: 5 },
    unicode.Range16{Lo: 0x0DD0, Hi: 0x0DD4, Stride: 1 },
    unicode.Range16{Lo: 0x0DD6, Hi: 0x0DD8, Stride: 2 },
    unicode.Range16{Lo: 0x0DD9, Hi: 0x0DDF, Stride: 1 },
    unicode.Range16{Lo: 0x0DE6, Hi: 0x0DEF, Stride: 1 },
    unicode.Range16{Lo: 0x0DF2, Hi: 0x0DF4, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x0111E1, Hi: 0x0111F4, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sogdian: 42 codepoints
var Sogdian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x010F30, Hi: 0x010F59, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sora_Sompeng: 35 codepoints
var Sora_Sompeng = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x0110D0, Hi: 0x0110E8, Stride: 1 },
    unicode.Range32{Lo: 0x0110F0, Hi: 0x0110F9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Soyombo: 83 codepoints
var Soyombo = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011A50, Hi: 0x011AA2, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sundanese: 72 codepoints
var Sundanese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1B80, Hi: 0x1BBF, Stride: 1 },
    unicode.Range16{Lo: 0x1CC0, Hi: 0x1CC7, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Syloti_Nagri: 45 codepoints
var Syloti_Nagri = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0xA800, Hi: 0xA82C, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Syriac: 88 codepoints
var Syriac = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0x0700, Hi: 0x070D, Stride: 1 },
    unicode.Range16{Lo: 0x070F, Hi: 0x074A, Stride: 1 },
    unicode.Range16{Lo: 0x074D, Hi: 0x074F, Stride: 1 },
    unicode.Range16{Lo: 0x0860, Hi: 0x086A, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tagalog: 23 codepoints
var Tagalog = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1700, Hi: 0x1715, Stride: 1 },
    unicode.Range16{Lo: 0x171F, Hi: 0x171F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tagbanwa: 18 codepoints
var Tagbanwa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x1760, Hi: 0x176C, Stride: 1 },
    unicode.Range16{Lo: 0x176E, Hi: 0x1770, Stride: 1 },
    unicode.Range16{Lo: 0x1772, Hi: 0x1773, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tai_Le: 35 codepoints
var Tai_Le = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x1950, Hi: 0x196D, Stride: 1 },
    unicode.Range16{Lo: 0x1970, Hi: 0x1974, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tai_Tham: 127 codepoints
var Tai_Tham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 5 */
    unicode.Range16{Lo: 0x1A20, Hi: 0x1A5E, Stride: 1 },
    unicode.Range16{Lo: 0x1A60, Hi: 0x1A7C, Stride: 1 },
    unicode.Range16{Lo: 0x1A7F, Hi: 0x1A89, Stride: 1 },
    unicode.Range16{Lo: 0x1A90, Hi: 0x1A99, Stride: 1 },
    unicode.Range16{Lo: 0x1AA0, Hi: 0x1AAD, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tai_Viet: 72 codepoints
var Tai_Viet = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0xAA80, Hi: 0xAAC2, Stride: 1 },
    unicode.Range16{Lo: 0xAADB, Hi: 0xAADF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Takri: 68 codepoints
var Takri = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011680, Hi: 0x0116B9, Stride: 1 },
    unicode.Range32{Lo: 0x0116C0, Hi: 0x0116C9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tamil: 123 codepoints
var Tamil = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 15 */
    unicode.Range16{Lo: 0x0B82, Hi: 0x0B83, Stride: 1 },
    unicode.Range16{Lo: 0x0B85, Hi: 0x0B8A, Stride: 1 },
    unicode.Range16{Lo: 0x0B8E, Hi: 0x0B90, Stride: 1 },
    unicode.Range16{Lo: 0x0B92, Hi: 0x0B95, Stride: 1 },
    unicode.Range16{Lo: 0x0B99, Hi: 0x0B9A, Stride: 1 },
    unicode.Range16{Lo: 0x0B9C, Hi: 0x0B9E, Stride: 2 },
    unicode.Range16{Lo: 0x0B9F, Hi: 0x0BA3, Stride: 4 },
    unicode.Range16{Lo: 0x0BA4, Hi: 0x0BA8, Stride: 4 },
    unicode.Range16{Lo: 0x0BA9, Hi: 0x0BAA, Stride: 1 },
    unicode.Range16{Lo: 0x0BAE, Hi: 0x0BB9, Stride: 1 },
    unicode.Range16{Lo: 0x0BBE, Hi: 0x0BC2, Stride: 1 },
    unicode.Range16{Lo: 0x0BC6, Hi: 0x0BC8, Stride: 1 },
    unicode.Range16{Lo: 0x0BCA, Hi: 0x0BCD, Stride: 1 },
    unicode.Range16{Lo: 0x0BD0, Hi: 0x0BD7, Stride: 7 },
    unicode.Range16{Lo: 0x0BE6, Hi: 0x0BFA, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011FC0, Hi: 0x011FF1, Stride: 1 },
    unicode.Range32{Lo: 0x011FFF, Hi: 0x011FFF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tangsa: 89 codepoints
var Tangsa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x016A70, Hi: 0x016ABE, Stride: 1 },
    unicode.Range32{Lo: 0x016AC0, Hi: 0x016AC9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tangut: 6914 codepoints
var Tangut = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 0x016FE0, Hi: 0x017000, Stride: 32 },
    unicode.Range32{Lo: 0x017001, Hi: 0x0187F7, Stride: 1 },
    unicode.Range32{Lo: 0x018800, Hi: 0x018AFF, Stride: 1 },
    unicode.Range32{Lo: 0x018D00, Hi: 0x018D08, Stride: 1 },
  },
  LatinOffset: 0,
}

// Telugu: 100 codepoints
var Telugu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 13 */
    unicode.Range16{Lo: 0x0C00, Hi: 0x0C0C, Stride: 1 },
    unicode.Range16{Lo: 0x0C0E, Hi: 0x0C10, Stride: 1 },
    unicode.Range16{Lo: 0x0C12, Hi: 0x0C28, Stride: 1 },
    unicode.Range16{Lo: 0x0C2A, Hi: 0x0C39, Stride: 1 },
    unicode.Range16{Lo: 0x0C3C, Hi: 0x0C44, Stride: 1 },
    unicode.Range16{Lo: 0x0C46, Hi: 0x0C48, Stride: 1 },
    unicode.Range16{Lo: 0x0C4A, Hi: 0x0C4D, Stride: 1 },
    unicode.Range16{Lo: 0x0C55, Hi: 0x0C56, Stride: 1 },
    unicode.Range16{Lo: 0x0C58, Hi: 0x0C5A, Stride: 1 },
    unicode.Range16{Lo: 0x0C5D, Hi: 0x0C60, Stride: 3 },
    unicode.Range16{Lo: 0x0C61, Hi: 0x0C63, Stride: 1 },
    unicode.Range16{Lo: 0x0C66, Hi: 0x0C6F, Stride: 1 },
    unicode.Range16{Lo: 0x0C77, Hi: 0x0C7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Thaana: 50 codepoints
var Thaana = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0x0780, Hi: 0x07B1, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Thai: 86 codepoints
var Thai = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0x0E01, Hi: 0x0E3A, Stride: 1 },
    unicode.Range16{Lo: 0x0E40, Hi: 0x0E5B, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tibetan: 207 codepoints
var Tibetan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0x0F00, Hi: 0x0F47, Stride: 1 },
    unicode.Range16{Lo: 0x0F49, Hi: 0x0F6C, Stride: 1 },
    unicode.Range16{Lo: 0x0F71, Hi: 0x0F97, Stride: 1 },
    unicode.Range16{Lo: 0x0F99, Hi: 0x0FBC, Stride: 1 },
    unicode.Range16{Lo: 0x0FBE, Hi: 0x0FCC, Stride: 1 },
    unicode.Range16{Lo: 0x0FCE, Hi: 0x0FD4, Stride: 1 },
    unicode.Range16{Lo: 0x0FD9, Hi: 0x0FDA, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tifinagh: 59 codepoints
var Tifinagh = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 0x2D30, Hi: 0x2D67, Stride: 1 },
    unicode.Range16{Lo: 0x2D6F, Hi: 0x2D70, Stride: 1 },
    unicode.Range16{Lo: 0x2D7F, Hi: 0x2D7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tirhuta: 82 codepoints
var Tirhuta = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x011480, Hi: 0x0114C7, Stride: 1 },
    unicode.Range32{Lo: 0x0114D0, Hi: 0x0114D9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Toto: 31 codepoints
var Toto = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x01E290, Hi: 0x01E2AE, Stride: 1 },
  },
  LatinOffset: 0,
}

// Ugaritic: 31 codepoints
var Ugaritic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x010380, Hi: 0x01039D, Stride: 1 },
    unicode.Range32{Lo: 0x01039F, Hi: 0x01039F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Vai: 300 codepoints
var Vai = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0xA500, Hi: 0xA62B, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Vithkuqi: 70 codepoints
var Vithkuqi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 8 */
    unicode.Range32{Lo: 0x010570, Hi: 0x01057A, Stride: 1 },
    unicode.Range32{Lo: 0x01057C, Hi: 0x01058A, Stride: 1 },
    unicode.Range32{Lo: 0x01058C, Hi: 0x010592, Stride: 1 },
    unicode.Range32{Lo: 0x010594, Hi: 0x010595, Stride: 1 },
    unicode.Range32{Lo: 0x010597, Hi: 0x0105A1, Stride: 1 },
    unicode.Range32{Lo: 0x0105A3, Hi: 0x0105B1, Stride: 1 },
    unicode.Range32{Lo: 0x0105B3, Hi: 0x0105B9, Stride: 1 },
    unicode.Range32{Lo: 0x0105BB, Hi: 0x0105BC, Stride: 1 },
  },
  LatinOffset: 0,
}

// Wancho: 59 codepoints
var Wancho = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x01E2C0, Hi: 0x01E2F9, Stride: 1 },
    unicode.Range32{Lo: 0x01E2FF, Hi: 0x01E2FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Warang_Citi: 84 codepoints
var Warang_Citi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 0x0118A0, Hi: 0x0118F2, Stride: 1 },
    unicode.Range32{Lo: 0x0118FF, Hi: 0x0118FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Yezidi: 47 codepoints
var Yezidi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 0x010E80, Hi: 0x010EA9, Stride: 1 },
    unicode.Range32{Lo: 0x010EAB, Hi: 0x010EAD, Stride: 1 },
    unicode.Range32{Lo: 0x010EB0, Hi: 0x010EB1, Stride: 1 },
  },
  LatinOffset: 0,
}

// Yi: 1220 codepoints
var Yi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0xA000, Hi: 0xA48C, Stride: 1 },
    unicode.Range16{Lo: 0xA490, Hi: 0xA4C6, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Zanabazar_Square: 72 codepoints
var Zanabazar_Square = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 0x011A00, Hi: 0x011A47, Stride: 1 },
  },
  LatinOffset: 0,
}
