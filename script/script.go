package script

import "unicode"

// Sogdian: 42 codepoints
var Sogdian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10F30, Hi: 10F59, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nyiakeng_Puachue_Hmong: 71 codepoints
var Nyiakeng_Puachue_Hmong = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 1E100, Hi: 1E12C, Stride: 1 },
    unicode.Range32{Lo: 1E130, Hi: 1E13D, Stride: 1 },
    unicode.Range32{Lo: 1E140, Hi: 1E149, Stride: 1 },
    unicode.Range32{Lo: 1E14E, Hi: 1E14F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Uyghur: 26 codepoints
var Old_Uyghur = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10F70, Hi: 10F89, Stride: 1 },
  },
  LatinOffset: 0,
}

// Syriac: 88 codepoints
var Syriac = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0700, Hi: 070D, Stride: 1 },
    unicode.Range16{Lo: 070F, Hi: 074A, Stride: 1 },
    unicode.Range16{Lo: 074D, Hi: 074F, Stride: 1 },
    unicode.Range16{Lo: 0860, Hi: 086A, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tibetan: 207 codepoints
var Tibetan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0F00, Hi: 0F47, Stride: 1 },
    unicode.Range16{Lo: 0F49, Hi: 0F6C, Stride: 1 },
    unicode.Range16{Lo: 0F71, Hi: 0F97, Stride: 1 },
    unicode.Range16{Lo: 0F99, Hi: 0FBC, Stride: 1 },
    unicode.Range16{Lo: 0FBE, Hi: 0FCC, Stride: 1 },
    unicode.Range16{Lo: 0FCE, Hi: 0FD4, Stride: 1 },
    unicode.Range16{Lo: 0FD9, Hi: 0FDA, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Meetei_Mayek: 79 codepoints
var Meetei_Mayek = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: AAE0, Hi: AAF6, Stride: 1 },
    unicode.Range16{Lo: ABC0, Hi: ABED, Stride: 1 },
    unicode.Range16{Lo: ABF0, Hi: ABF9, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Osmanya: 40 codepoints
var Osmanya = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10480, Hi: 1049D, Stride: 1 },
    unicode.Range32{Lo: 104A0, Hi: 104A9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lepcha: 74 codepoints
var Lepcha = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 1C00, Hi: 1C37, Stride: 1 },
    unicode.Range16{Lo: 1C3B, Hi: 1C49, Stride: 1 },
    unicode.Range16{Lo: 1C4D, Hi: 1C4F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Ahom: 65 codepoints
var Ahom = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 11700, Hi: 1171A, Stride: 1 },
    unicode.Range32{Lo: 1171D, Hi: 1172B, Stride: 1 },
    unicode.Range32{Lo: 11730, Hi: 11746, Stride: 1 },
  },
  LatinOffset: 0,
}

// Newa: 97 codepoints
var Newa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11400, Hi: 1145B, Stride: 1 },
    unicode.Range32{Lo: 1145D, Hi: 11461, Stride: 1 },
  },
  LatinOffset: 0,
}

// Medefaidrin: 91 codepoints
var Medefaidrin = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 16E40, Hi: 16E9A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Common: 8301 codepoints
var Common = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 82 */
    unicode.Range16{Lo: 0000, Hi: 0040, Stride: 1 },
    unicode.Range16{Lo: 005B, Hi: 0060, Stride: 1 },
    unicode.Range16{Lo: 007B, Hi: 00A9, Stride: 1 },
    unicode.Range16{Lo: 00AB, Hi: 00B9, Stride: 1 },
    unicode.Range16{Lo: 00BB, Hi: 00BF, Stride: 1 },
    unicode.Range16{Lo: 00D7, Hi: 00F7, Stride: 32 },
    unicode.Range16{Lo: 02B9, Hi: 02DF, Stride: 1 },
    unicode.Range16{Lo: 02E5, Hi: 02E9, Stride: 1 },
    unicode.Range16{Lo: 02EC, Hi: 02FF, Stride: 1 },
    unicode.Range16{Lo: 0374, Hi: 037E, Stride: 10 },
    unicode.Range16{Lo: 0385, Hi: 0387, Stride: 2 },
    unicode.Range16{Lo: 0605, Hi: 060C, Stride: 7 },
    unicode.Range16{Lo: 061B, Hi: 061F, Stride: 4 },
    unicode.Range16{Lo: 0640, Hi: 06DD, Stride: 157 },
    unicode.Range16{Lo: 08E2, Hi: 0964, Stride: 130 },
    unicode.Range16{Lo: 0965, Hi: 0E3F, Stride: 1242 },
    unicode.Range16{Lo: 0FD5, Hi: 0FD8, Stride: 1 },
    unicode.Range16{Lo: 10FB, Hi: 16EB, Stride: 1520 },
    unicode.Range16{Lo: 16EC, Hi: 16ED, Stride: 1 },
    unicode.Range16{Lo: 1735, Hi: 1736, Stride: 1 },
    unicode.Range16{Lo: 1802, Hi: 1803, Stride: 1 },
    unicode.Range16{Lo: 1805, Hi: 1CD3, Stride: 1230 },
    unicode.Range16{Lo: 1CE1, Hi: 1CE9, Stride: 8 },
    unicode.Range16{Lo: 1CEA, Hi: 1CEC, Stride: 1 },
    unicode.Range16{Lo: 1CEE, Hi: 1CF3, Stride: 1 },
    unicode.Range16{Lo: 1CF5, Hi: 1CF7, Stride: 1 },
    unicode.Range16{Lo: 1CFA, Hi: 2000, Stride: 774 },
    unicode.Range16{Lo: 2001, Hi: 200B, Stride: 1 },
    unicode.Range16{Lo: 200E, Hi: 2064, Stride: 1 },
    unicode.Range16{Lo: 2066, Hi: 2070, Stride: 1 },
    unicode.Range16{Lo: 2074, Hi: 207E, Stride: 1 },
    unicode.Range16{Lo: 2080, Hi: 208E, Stride: 1 },
    unicode.Range16{Lo: 20A0, Hi: 20C0, Stride: 1 },
    unicode.Range16{Lo: 2100, Hi: 2125, Stride: 1 },
    unicode.Range16{Lo: 2127, Hi: 2129, Stride: 1 },
    unicode.Range16{Lo: 212C, Hi: 2131, Stride: 1 },
    unicode.Range16{Lo: 2133, Hi: 214D, Stride: 1 },
    unicode.Range16{Lo: 214F, Hi: 215F, Stride: 1 },
    unicode.Range16{Lo: 2189, Hi: 218B, Stride: 1 },
    unicode.Range16{Lo: 2190, Hi: 2426, Stride: 1 },
    unicode.Range16{Lo: 2440, Hi: 244A, Stride: 1 },
    unicode.Range16{Lo: 2460, Hi: 27FF, Stride: 1 },
    unicode.Range16{Lo: 2900, Hi: 2B73, Stride: 1 },
    unicode.Range16{Lo: 2B76, Hi: 2B95, Stride: 1 },
    unicode.Range16{Lo: 2B97, Hi: 2BFF, Stride: 1 },
    unicode.Range16{Lo: 2E00, Hi: 2E5D, Stride: 1 },
    unicode.Range16{Lo: 2FF0, Hi: 2FFB, Stride: 1 },
    unicode.Range16{Lo: 3000, Hi: 3004, Stride: 1 },
    unicode.Range16{Lo: 3006, Hi: 3008, Stride: 2 },
    unicode.Range16{Lo: 3009, Hi: 3020, Stride: 1 },
    unicode.Range16{Lo: 3030, Hi: 3037, Stride: 1 },
    unicode.Range16{Lo: 303C, Hi: 303F, Stride: 1 },
    unicode.Range16{Lo: 309B, Hi: 309C, Stride: 1 },
    unicode.Range16{Lo: 30A0, Hi: 30FB, Stride: 91 },
    unicode.Range16{Lo: 30FC, Hi: 3190, Stride: 148 },
    unicode.Range16{Lo: 3191, Hi: 319F, Stride: 1 },
    unicode.Range16{Lo: 31C0, Hi: 31E3, Stride: 1 },
    unicode.Range16{Lo: 3220, Hi: 325F, Stride: 1 },
    unicode.Range16{Lo: 327F, Hi: 32CF, Stride: 1 },
    unicode.Range16{Lo: 32FF, Hi: 3358, Stride: 89 },
    unicode.Range16{Lo: 3359, Hi: 33FF, Stride: 1 },
    unicode.Range16{Lo: 4DC0, Hi: 4DFF, Stride: 1 },
    unicode.Range16{Lo: A700, Hi: A721, Stride: 1 },
    unicode.Range16{Lo: A788, Hi: A78A, Stride: 1 },
    unicode.Range16{Lo: A830, Hi: A839, Stride: 1 },
    unicode.Range16{Lo: A92E, Hi: A9CF, Stride: 161 },
    unicode.Range16{Lo: AB5B, Hi: AB6A, Stride: 15 },
    unicode.Range16{Lo: AB6B, Hi: FD3E, Stride: 20947 },
    unicode.Range16{Lo: FD3F, Hi: FE10, Stride: 209 },
    unicode.Range16{Lo: FE11, Hi: FE19, Stride: 1 },
    unicode.Range16{Lo: FE30, Hi: FE52, Stride: 1 },
    unicode.Range16{Lo: FE54, Hi: FE66, Stride: 1 },
    unicode.Range16{Lo: FE68, Hi: FE6B, Stride: 1 },
    unicode.Range16{Lo: FEFF, Hi: FF01, Stride: 2 },
    unicode.Range16{Lo: FF02, Hi: FF20, Stride: 1 },
    unicode.Range16{Lo: FF3B, Hi: FF40, Stride: 1 },
    unicode.Range16{Lo: FF5B, Hi: FF65, Stride: 1 },
    unicode.Range16{Lo: FF70, Hi: FF9E, Stride: 46 },
    unicode.Range16{Lo: FF9F, Hi: FFE0, Stride: 65 },
    unicode.Range16{Lo: FFE1, Hi: FFE6, Stride: 1 },
    unicode.Range16{Lo: FFE8, Hi: FFEE, Stride: 1 },
    unicode.Range16{Lo: FFF9, Hi: FFFD, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 82 */
    unicode.Range32{Lo: 10100, Hi: 10102, Stride: 1 },
    unicode.Range32{Lo: 10107, Hi: 10133, Stride: 1 },
    unicode.Range32{Lo: 10137, Hi: 1013F, Stride: 1 },
    unicode.Range32{Lo: 10190, Hi: 1019C, Stride: 1 },
    unicode.Range32{Lo: 101D0, Hi: 101FC, Stride: 1 },
    unicode.Range32{Lo: 102E1, Hi: 102FB, Stride: 1 },
    unicode.Range32{Lo: 1BCA0, Hi: 1BCA3, Stride: 1 },
    unicode.Range32{Lo: 1CF50, Hi: 1CFC3, Stride: 1 },
    unicode.Range32{Lo: 1D000, Hi: 1D0F5, Stride: 1 },
    unicode.Range32{Lo: 1D100, Hi: 1D126, Stride: 1 },
    unicode.Range32{Lo: 1D129, Hi: 1D166, Stride: 1 },
    unicode.Range32{Lo: 1D16A, Hi: 1D17A, Stride: 1 },
    unicode.Range32{Lo: 1D183, Hi: 1D184, Stride: 1 },
    unicode.Range32{Lo: 1D18C, Hi: 1D1A9, Stride: 1 },
    unicode.Range32{Lo: 1D1AE, Hi: 1D1EA, Stride: 1 },
    unicode.Range32{Lo: 1D2C0, Hi: 1D2D3, Stride: 1 },
    unicode.Range32{Lo: 1D2E0, Hi: 1D2F3, Stride: 1 },
    unicode.Range32{Lo: 1D300, Hi: 1D356, Stride: 1 },
    unicode.Range32{Lo: 1D360, Hi: 1D378, Stride: 1 },
    unicode.Range32{Lo: 1D400, Hi: 1D454, Stride: 1 },
    unicode.Range32{Lo: 1D456, Hi: 1D49C, Stride: 1 },
    unicode.Range32{Lo: 1D49E, Hi: 1D49F, Stride: 1 },
    unicode.Range32{Lo: 1D4A2, Hi: 1D4A5, Stride: 3 },
    unicode.Range32{Lo: 1D4A6, Hi: 1D4A9, Stride: 3 },
    unicode.Range32{Lo: 1D4AA, Hi: 1D4AC, Stride: 1 },
    unicode.Range32{Lo: 1D4AE, Hi: 1D4B9, Stride: 1 },
    unicode.Range32{Lo: 1D4BB, Hi: 1D4BD, Stride: 2 },
    unicode.Range32{Lo: 1D4BE, Hi: 1D4C3, Stride: 1 },
    unicode.Range32{Lo: 1D4C5, Hi: 1D505, Stride: 1 },
    unicode.Range32{Lo: 1D507, Hi: 1D50A, Stride: 1 },
    unicode.Range32{Lo: 1D50D, Hi: 1D514, Stride: 1 },
    unicode.Range32{Lo: 1D516, Hi: 1D51C, Stride: 1 },
    unicode.Range32{Lo: 1D51E, Hi: 1D539, Stride: 1 },
    unicode.Range32{Lo: 1D53B, Hi: 1D53E, Stride: 1 },
    unicode.Range32{Lo: 1D540, Hi: 1D544, Stride: 1 },
    unicode.Range32{Lo: 1D546, Hi: 1D54A, Stride: 4 },
    unicode.Range32{Lo: 1D54B, Hi: 1D550, Stride: 1 },
    unicode.Range32{Lo: 1D552, Hi: 1D6A5, Stride: 1 },
    unicode.Range32{Lo: 1D6A8, Hi: 1D7CB, Stride: 1 },
    unicode.Range32{Lo: 1D7CE, Hi: 1D7FF, Stride: 1 },
    unicode.Range32{Lo: 1EC71, Hi: 1ECB4, Stride: 1 },
    unicode.Range32{Lo: 1ED01, Hi: 1ED3D, Stride: 1 },
    unicode.Range32{Lo: 1F000, Hi: 1F02B, Stride: 1 },
    unicode.Range32{Lo: 1F030, Hi: 1F093, Stride: 1 },
    unicode.Range32{Lo: 1F0A0, Hi: 1F0AE, Stride: 1 },
    unicode.Range32{Lo: 1F0B1, Hi: 1F0BF, Stride: 1 },
    unicode.Range32{Lo: 1F0C1, Hi: 1F0CF, Stride: 1 },
    unicode.Range32{Lo: 1F0D1, Hi: 1F0F5, Stride: 1 },
    unicode.Range32{Lo: 1F100, Hi: 1F1AD, Stride: 1 },
    unicode.Range32{Lo: 1F1E6, Hi: 1F1FF, Stride: 1 },
    unicode.Range32{Lo: 1F201, Hi: 1F202, Stride: 1 },
    unicode.Range32{Lo: 1F210, Hi: 1F23B, Stride: 1 },
    unicode.Range32{Lo: 1F240, Hi: 1F248, Stride: 1 },
    unicode.Range32{Lo: 1F250, Hi: 1F251, Stride: 1 },
    unicode.Range32{Lo: 1F260, Hi: 1F265, Stride: 1 },
    unicode.Range32{Lo: 1F300, Hi: 1F6D7, Stride: 1 },
    unicode.Range32{Lo: 1F6DC, Hi: 1F6EC, Stride: 1 },
    unicode.Range32{Lo: 1F6F0, Hi: 1F6FC, Stride: 1 },
    unicode.Range32{Lo: 1F700, Hi: 1F776, Stride: 1 },
    unicode.Range32{Lo: 1F77B, Hi: 1F7D9, Stride: 1 },
    unicode.Range32{Lo: 1F7E0, Hi: 1F7EB, Stride: 1 },
    unicode.Range32{Lo: 1F7F0, Hi: 1F800, Stride: 16 },
    unicode.Range32{Lo: 1F801, Hi: 1F80B, Stride: 1 },
    unicode.Range32{Lo: 1F810, Hi: 1F847, Stride: 1 },
    unicode.Range32{Lo: 1F850, Hi: 1F859, Stride: 1 },
    unicode.Range32{Lo: 1F860, Hi: 1F887, Stride: 1 },
    unicode.Range32{Lo: 1F890, Hi: 1F8AD, Stride: 1 },
    unicode.Range32{Lo: 1F8B0, Hi: 1F8B1, Stride: 1 },
    unicode.Range32{Lo: 1F900, Hi: 1FA53, Stride: 1 },
    unicode.Range32{Lo: 1FA60, Hi: 1FA6D, Stride: 1 },
    unicode.Range32{Lo: 1FA70, Hi: 1FA7C, Stride: 1 },
    unicode.Range32{Lo: 1FA80, Hi: 1FA88, Stride: 1 },
    unicode.Range32{Lo: 1FA90, Hi: 1FABD, Stride: 1 },
    unicode.Range32{Lo: 1FABF, Hi: 1FAC5, Stride: 1 },
    unicode.Range32{Lo: 1FACE, Hi: 1FADB, Stride: 1 },
    unicode.Range32{Lo: 1FAE0, Hi: 1FAE8, Stride: 1 },
    unicode.Range32{Lo: 1FAF0, Hi: 1FAF8, Stride: 1 },
    unicode.Range32{Lo: 1FB00, Hi: 1FB92, Stride: 1 },
    unicode.Range32{Lo: 1FB94, Hi: 1FBCA, Stride: 1 },
    unicode.Range32{Lo: 1FBF0, Hi: 1FBF9, Stride: 1 },
    unicode.Range32{Lo: E0001, Hi: E0020, Stride: 31 },
    unicode.Range32{Lo: E0021, Hi: E007F, Stride: 1 },
  },
  LatinOffset: 6,
}

// Thaana: 50 codepoints
var Thaana = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 0780, Hi: 07B1, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Kannada: 91 codepoints
var Kannada = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 13 */
    unicode.Range16{Lo: 0C80, Hi: 0C8C, Stride: 1 },
    unicode.Range16{Lo: 0C8E, Hi: 0C90, Stride: 1 },
    unicode.Range16{Lo: 0C92, Hi: 0CA8, Stride: 1 },
    unicode.Range16{Lo: 0CAA, Hi: 0CB3, Stride: 1 },
    unicode.Range16{Lo: 0CB5, Hi: 0CB9, Stride: 1 },
    unicode.Range16{Lo: 0CBC, Hi: 0CC4, Stride: 1 },
    unicode.Range16{Lo: 0CC6, Hi: 0CC8, Stride: 1 },
    unicode.Range16{Lo: 0CCA, Hi: 0CCD, Stride: 1 },
    unicode.Range16{Lo: 0CD5, Hi: 0CD6, Stride: 1 },
    unicode.Range16{Lo: 0CDD, Hi: 0CDE, Stride: 1 },
    unicode.Range16{Lo: 0CE0, Hi: 0CE3, Stride: 1 },
    unicode.Range16{Lo: 0CE6, Hi: 0CEF, Stride: 1 },
    unicode.Range16{Lo: 0CF1, Hi: 0CF3, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Phags_Pa: 56 codepoints
var Phags_Pa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: A840, Hi: A877, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Multani: 38 codepoints
var Multani = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 11280, Hi: 11286, Stride: 1 },
    unicode.Range32{Lo: 11288, Hi: 1128A, Stride: 2 },
    unicode.Range32{Lo: 1128B, Hi: 1128D, Stride: 1 },
    unicode.Range32{Lo: 1128F, Hi: 1129D, Stride: 1 },
    unicode.Range32{Lo: 1129F, Hi: 112A9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Cypriot: 55 codepoints
var Cypriot = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 10800, Hi: 10805, Stride: 1 },
    unicode.Range32{Lo: 10808, Hi: 1080A, Stride: 2 },
    unicode.Range32{Lo: 1080B, Hi: 10835, Stride: 1 },
    unicode.Range32{Lo: 10837, Hi: 10838, Stride: 1 },
    unicode.Range32{Lo: 1083C, Hi: 1083F, Stride: 3 },
  },
  LatinOffset: 0,
}

// New_Tai_Lue: 83 codepoints
var New_Tai_Lue = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 1980, Hi: 19AB, Stride: 1 },
    unicode.Range16{Lo: 19B0, Hi: 19C9, Stride: 1 },
    unicode.Range16{Lo: 19D0, Hi: 19DA, Stride: 1 },
    unicode.Range16{Lo: 19DE, Hi: 19DF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tifinagh: 59 codepoints
var Tifinagh = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 2D30, Hi: 2D67, Stride: 1 },
    unicode.Range16{Lo: 2D6F, Hi: 2D70, Stride: 1 },
    unicode.Range16{Lo: 2D7F, Hi: 2D7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Inscriptional_Parthian: 30 codepoints
var Inscriptional_Parthian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10B40, Hi: 10B55, Stride: 1 },
    unicode.Range32{Lo: 10B58, Hi: 10B5F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Psalter_Pahlavi: 29 codepoints
var Psalter_Pahlavi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 10B80, Hi: 10B91, Stride: 1 },
    unicode.Range32{Lo: 10B99, Hi: 10B9C, Stride: 1 },
    unicode.Range32{Lo: 10BA9, Hi: 10BAF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hiragana: 381 codepoints
var Hiragana = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 3041, Hi: 3096, Stride: 1 },
    unicode.Range16{Lo: 309D, Hi: 309F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 1B001, Hi: 1B11F, Stride: 1 },
    unicode.Range32{Lo: 1B132, Hi: 1B150, Stride: 30 },
    unicode.Range32{Lo: 1B151, Hi: 1B152, Stride: 1 },
    unicode.Range32{Lo: 1F200, Hi: 1F200, Stride: 1 },
  },
  LatinOffset: 0,
}

// Shavian: 48 codepoints
var Shavian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10450, Hi: 1047F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Vai: 300 codepoints
var Vai = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: A500, Hi: A62B, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Linear_B: 211 codepoints
var Linear_B = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 7 */
    unicode.Range32{Lo: 10000, Hi: 1000B, Stride: 1 },
    unicode.Range32{Lo: 1000D, Hi: 10026, Stride: 1 },
    unicode.Range32{Lo: 10028, Hi: 1003A, Stride: 1 },
    unicode.Range32{Lo: 1003C, Hi: 1003D, Stride: 1 },
    unicode.Range32{Lo: 1003F, Hi: 1004D, Stride: 1 },
    unicode.Range32{Lo: 10050, Hi: 1005D, Stride: 1 },
    unicode.Range32{Lo: 10080, Hi: 100FA, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tai_Viet: 72 codepoints
var Tai_Viet = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: AA80, Hi: AAC2, Stride: 1 },
    unicode.Range16{Lo: AADB, Hi: AADF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Cham: 83 codepoints
var Cham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: AA00, Hi: AA36, Stride: 1 },
    unicode.Range16{Lo: AA40, Hi: AA4D, Stride: 1 },
    unicode.Range16{Lo: AA50, Hi: AA59, Stride: 1 },
    unicode.Range16{Lo: AA5C, Hi: AA5F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Bamum: 657 codepoints
var Bamum = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: A6A0, Hi: A6F7, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 16800, Hi: 16A38, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Sogdian: 40 codepoints
var Old_Sogdian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10F00, Hi: 10F27, Stride: 1 },
  },
  LatinOffset: 0,
}

// Bengali: 96 codepoints
var Bengali = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0980, Hi: 0983, Stride: 1 },
    unicode.Range16{Lo: 0985, Hi: 098C, Stride: 1 },
    unicode.Range16{Lo: 098F, Hi: 0990, Stride: 1 },
    unicode.Range16{Lo: 0993, Hi: 09A8, Stride: 1 },
    unicode.Range16{Lo: 09AA, Hi: 09B0, Stride: 1 },
    unicode.Range16{Lo: 09B2, Hi: 09B6, Stride: 4 },
    unicode.Range16{Lo: 09B7, Hi: 09B9, Stride: 1 },
    unicode.Range16{Lo: 09BC, Hi: 09C4, Stride: 1 },
    unicode.Range16{Lo: 09C7, Hi: 09C8, Stride: 1 },
    unicode.Range16{Lo: 09CB, Hi: 09CE, Stride: 1 },
    unicode.Range16{Lo: 09D7, Hi: 09DC, Stride: 5 },
    unicode.Range16{Lo: 09DD, Hi: 09DF, Stride: 2 },
    unicode.Range16{Lo: 09E0, Hi: 09E3, Stride: 1 },
    unicode.Range16{Lo: 09E6, Hi: 09FE, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Ethiopic: 523 codepoints
var Ethiopic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 32 */
    unicode.Range16{Lo: 1200, Hi: 1248, Stride: 1 },
    unicode.Range16{Lo: 124A, Hi: 124D, Stride: 1 },
    unicode.Range16{Lo: 1250, Hi: 1256, Stride: 1 },
    unicode.Range16{Lo: 1258, Hi: 125A, Stride: 2 },
    unicode.Range16{Lo: 125B, Hi: 125D, Stride: 1 },
    unicode.Range16{Lo: 1260, Hi: 1288, Stride: 1 },
    unicode.Range16{Lo: 128A, Hi: 128D, Stride: 1 },
    unicode.Range16{Lo: 1290, Hi: 12B0, Stride: 1 },
    unicode.Range16{Lo: 12B2, Hi: 12B5, Stride: 1 },
    unicode.Range16{Lo: 12B8, Hi: 12BE, Stride: 1 },
    unicode.Range16{Lo: 12C0, Hi: 12C2, Stride: 2 },
    unicode.Range16{Lo: 12C3, Hi: 12C5, Stride: 1 },
    unicode.Range16{Lo: 12C8, Hi: 12D6, Stride: 1 },
    unicode.Range16{Lo: 12D8, Hi: 1310, Stride: 1 },
    unicode.Range16{Lo: 1312, Hi: 1315, Stride: 1 },
    unicode.Range16{Lo: 1318, Hi: 135A, Stride: 1 },
    unicode.Range16{Lo: 135D, Hi: 137C, Stride: 1 },
    unicode.Range16{Lo: 1380, Hi: 1399, Stride: 1 },
    unicode.Range16{Lo: 2D80, Hi: 2D96, Stride: 1 },
    unicode.Range16{Lo: 2DA0, Hi: 2DA6, Stride: 1 },
    unicode.Range16{Lo: 2DA8, Hi: 2DAE, Stride: 1 },
    unicode.Range16{Lo: 2DB0, Hi: 2DB6, Stride: 1 },
    unicode.Range16{Lo: 2DB8, Hi: 2DBE, Stride: 1 },
    unicode.Range16{Lo: 2DC0, Hi: 2DC6, Stride: 1 },
    unicode.Range16{Lo: 2DC8, Hi: 2DCE, Stride: 1 },
    unicode.Range16{Lo: 2DD0, Hi: 2DD6, Stride: 1 },
    unicode.Range16{Lo: 2DD8, Hi: 2DDE, Stride: 1 },
    unicode.Range16{Lo: AB01, Hi: AB06, Stride: 1 },
    unicode.Range16{Lo: AB09, Hi: AB0E, Stride: 1 },
    unicode.Range16{Lo: AB11, Hi: AB16, Stride: 1 },
    unicode.Range16{Lo: AB20, Hi: AB26, Stride: 1 },
    unicode.Range16{Lo: AB28, Hi: AB2E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 1E7E0, Hi: 1E7E6, Stride: 1 },
    unicode.Range32{Lo: 1E7E8, Hi: 1E7EB, Stride: 1 },
    unicode.Range32{Lo: 1E7ED, Hi: 1E7EE, Stride: 1 },
    unicode.Range32{Lo: 1E7F0, Hi: 1E7FE, Stride: 1 },
  },
  LatinOffset: 0,
}

// Phoenician: 29 codepoints
var Phoenician = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10900, Hi: 1091B, Stride: 1 },
    unicode.Range32{Lo: 1091F, Hi: 1091F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Mandaic: 29 codepoints
var Mandaic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0840, Hi: 085B, Stride: 1 },
    unicode.Range16{Lo: 085E, Hi: 085E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Miao: 149 codepoints
var Miao = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 16F00, Hi: 16F4A, Stride: 1 },
    unicode.Range32{Lo: 16F4F, Hi: 16F87, Stride: 1 },
    unicode.Range32{Lo: 16F8F, Hi: 16F9F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Mahajani: 39 codepoints
var Mahajani = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11150, Hi: 11176, Stride: 1 },
  },
  LatinOffset: 0,
}

// Khudawadi: 69 codepoints
var Khudawadi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 112B0, Hi: 112EA, Stride: 1 },
    unicode.Range32{Lo: 112F0, Hi: 112F9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hatran: 26 codepoints
var Hatran = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 108E0, Hi: 108F2, Stride: 1 },
    unicode.Range32{Lo: 108F4, Hi: 108F5, Stride: 1 },
    unicode.Range32{Lo: 108FB, Hi: 108FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Mongolian: 168 codepoints
var Mongolian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 5 */
    unicode.Range16{Lo: 1800, Hi: 1801, Stride: 1 },
    unicode.Range16{Lo: 1804, Hi: 1806, Stride: 2 },
    unicode.Range16{Lo: 1807, Hi: 1819, Stride: 1 },
    unicode.Range16{Lo: 1820, Hi: 1878, Stride: 1 },
    unicode.Range16{Lo: 1880, Hi: 18AA, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11660, Hi: 1166C, Stride: 1 },
  },
  LatinOffset: 0,
}

// Yi: 1220 codepoints
var Yi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: A000, Hi: A48C, Stride: 1 },
    unicode.Range16{Lo: A490, Hi: A4C6, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Old_Italic: 39 codepoints
var Old_Italic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10300, Hi: 10323, Stride: 1 },
    unicode.Range32{Lo: 1032D, Hi: 1032F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hebrew: 134 codepoints
var Hebrew = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 9 */
    unicode.Range16{Lo: 0591, Hi: 05C7, Stride: 1 },
    unicode.Range16{Lo: 05D0, Hi: 05EA, Stride: 1 },
    unicode.Range16{Lo: 05EF, Hi: 05F4, Stride: 1 },
    unicode.Range16{Lo: FB1D, Hi: FB36, Stride: 1 },
    unicode.Range16{Lo: FB38, Hi: FB3C, Stride: 1 },
    unicode.Range16{Lo: FB3E, Hi: FB40, Stride: 2 },
    unicode.Range16{Lo: FB41, Hi: FB43, Stride: 2 },
    unicode.Range16{Lo: FB44, Hi: FB46, Stride: 2 },
    unicode.Range16{Lo: FB47, Hi: FB4F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Ugaritic: 31 codepoints
var Ugaritic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10380, Hi: 1039D, Stride: 1 },
    unicode.Range32{Lo: 1039F, Hi: 1039F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Samaritan: 61 codepoints
var Samaritan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0800, Hi: 082D, Stride: 1 },
    unicode.Range16{Lo: 0830, Hi: 083E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Old_South_Arabian: 32 codepoints
var Old_South_Arabian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10A60, Hi: 10A7F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Linear_A: 341 codepoints
var Linear_A = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 10600, Hi: 10736, Stride: 1 },
    unicode.Range32{Lo: 10740, Hi: 10755, Stride: 1 },
    unicode.Range32{Lo: 10760, Hi: 10767, Stride: 1 },
  },
  LatinOffset: 0,
}

// Latin: 1481 codepoints
var Latin = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 32 */
    unicode.Range16{Lo: 0041, Hi: 005A, Stride: 1 },
    unicode.Range16{Lo: 0061, Hi: 007A, Stride: 1 },
    unicode.Range16{Lo: 00AA, Hi: 00BA, Stride: 16 },
    unicode.Range16{Lo: 00C0, Hi: 00D6, Stride: 1 },
    unicode.Range16{Lo: 00D8, Hi: 00F6, Stride: 1 },
    unicode.Range16{Lo: 00F8, Hi: 00FF, Stride: 1 },
    unicode.Range16{Lo: 0100, Hi: 02B8, Stride: 1 },
    unicode.Range16{Lo: 02E0, Hi: 02E4, Stride: 1 },
    unicode.Range16{Lo: 1D00, Hi: 1D25, Stride: 1 },
    unicode.Range16{Lo: 1D2C, Hi: 1D5C, Stride: 1 },
    unicode.Range16{Lo: 1D62, Hi: 1D65, Stride: 1 },
    unicode.Range16{Lo: 1D6B, Hi: 1D77, Stride: 1 },
    unicode.Range16{Lo: 1D79, Hi: 1DBE, Stride: 1 },
    unicode.Range16{Lo: 1E00, Hi: 1EFF, Stride: 1 },
    unicode.Range16{Lo: 2071, Hi: 207F, Stride: 14 },
    unicode.Range16{Lo: 2090, Hi: 209C, Stride: 1 },
    unicode.Range16{Lo: 212A, Hi: 212B, Stride: 1 },
    unicode.Range16{Lo: 2132, Hi: 214E, Stride: 28 },
    unicode.Range16{Lo: 2160, Hi: 2188, Stride: 1 },
    unicode.Range16{Lo: 2C60, Hi: 2C7F, Stride: 1 },
    unicode.Range16{Lo: A722, Hi: A787, Stride: 1 },
    unicode.Range16{Lo: A78B, Hi: A7CA, Stride: 1 },
    unicode.Range16{Lo: A7D0, Hi: A7D1, Stride: 1 },
    unicode.Range16{Lo: A7D3, Hi: A7D5, Stride: 2 },
    unicode.Range16{Lo: A7D6, Hi: A7D9, Stride: 1 },
    unicode.Range16{Lo: A7F2, Hi: A7FF, Stride: 1 },
    unicode.Range16{Lo: AB30, Hi: AB5A, Stride: 1 },
    unicode.Range16{Lo: AB5C, Hi: AB64, Stride: 1 },
    unicode.Range16{Lo: AB66, Hi: AB69, Stride: 1 },
    unicode.Range16{Lo: FB00, Hi: FB06, Stride: 1 },
    unicode.Range16{Lo: FF21, Hi: FF3A, Stride: 1 },
    unicode.Range16{Lo: FF41, Hi: FF5A, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 10780, Hi: 10785, Stride: 1 },
    unicode.Range32{Lo: 10787, Hi: 107B0, Stride: 1 },
    unicode.Range32{Lo: 107B2, Hi: 107BA, Stride: 1 },
    unicode.Range32{Lo: 1DF00, Hi: 1DF1E, Stride: 1 },
    unicode.Range32{Lo: 1DF25, Hi: 1DF2A, Stride: 1 },
  },
  LatinOffset: 6,
}

// Cyrillic: 506 codepoints
var Cyrillic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0400, Hi: 0484, Stride: 1 },
    unicode.Range16{Lo: 0487, Hi: 052F, Stride: 1 },
    unicode.Range16{Lo: 1C80, Hi: 1C88, Stride: 1 },
    unicode.Range16{Lo: 1D2B, Hi: 1D78, Stride: 77 },
    unicode.Range16{Lo: 2DE0, Hi: 2DFF, Stride: 1 },
    unicode.Range16{Lo: A640, Hi: A69F, Stride: 1 },
    unicode.Range16{Lo: FE2E, Hi: FE2F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 1E030, Hi: 1E06D, Stride: 1 },
    unicode.Range32{Lo: 1E08F, Hi: 1E08F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Armenian: 96 codepoints
var Armenian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0531, Hi: 0556, Stride: 1 },
    unicode.Range16{Lo: 0559, Hi: 058A, Stride: 1 },
    unicode.Range16{Lo: 058D, Hi: 058F, Stride: 1 },
    unicode.Range16{Lo: FB13, Hi: FB17, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Palmyrene: 32 codepoints
var Palmyrene = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10860, Hi: 1087F, Stride: 1 },
  },
  LatinOffset: 0,
}

// SignWriting: 672 codepoints
var SignWriting = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 1D800, Hi: 1DA8B, Stride: 1 },
    unicode.Range32{Lo: 1DA9B, Hi: 1DA9F, Stride: 1 },
    unicode.Range32{Lo: 1DAA1, Hi: 1DAAF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Marchen: 68 codepoints
var Marchen = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 11C70, Hi: 11C8F, Stride: 1 },
    unicode.Range32{Lo: 11C92, Hi: 11CA7, Stride: 1 },
    unicode.Range32{Lo: 11CA9, Hi: 11CB6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Inscriptional_Pahlavi: 27 codepoints
var Inscriptional_Pahlavi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10B60, Hi: 10B72, Stride: 1 },
    unicode.Range32{Lo: 10B78, Hi: 10B7F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Bhaiksuki: 97 codepoints
var Bhaiksuki = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 11C00, Hi: 11C08, Stride: 1 },
    unicode.Range32{Lo: 11C0A, Hi: 11C36, Stride: 1 },
    unicode.Range32{Lo: 11C38, Hi: 11C45, Stride: 1 },
    unicode.Range32{Lo: 11C50, Hi: 11C6C, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lao: 83 codepoints
var Lao = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 11 */
    unicode.Range16{Lo: 0E81, Hi: 0E82, Stride: 1 },
    unicode.Range16{Lo: 0E84, Hi: 0E86, Stride: 2 },
    unicode.Range16{Lo: 0E87, Hi: 0E8A, Stride: 1 },
    unicode.Range16{Lo: 0E8C, Hi: 0EA3, Stride: 1 },
    unicode.Range16{Lo: 0EA5, Hi: 0EA7, Stride: 2 },
    unicode.Range16{Lo: 0EA8, Hi: 0EBD, Stride: 1 },
    unicode.Range16{Lo: 0EC0, Hi: 0EC4, Stride: 1 },
    unicode.Range16{Lo: 0EC6, Hi: 0EC8, Stride: 2 },
    unicode.Range16{Lo: 0EC9, Hi: 0ECE, Stride: 1 },
    unicode.Range16{Lo: 0ED0, Hi: 0ED9, Stride: 1 },
    unicode.Range16{Lo: 0EDC, Hi: 0EDF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Myanmar: 223 codepoints
var Myanmar = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 1000, Hi: 109F, Stride: 1 },
    unicode.Range16{Lo: A9E0, Hi: A9FE, Stride: 1 },
    unicode.Range16{Lo: AA60, Hi: AA7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Syloti_Nagri: 45 codepoints
var Syloti_Nagri = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: A800, Hi: A82C, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Hanunoo: 21 codepoints
var Hanunoo = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 1720, Hi: 1734, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Buginese: 30 codepoints
var Buginese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1A00, Hi: 1A1B, Stride: 1 },
    unicode.Range16{Lo: 1A1E, Hi: 1A1F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Chakma: 71 codepoints
var Chakma = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11100, Hi: 11134, Stride: 1 },
    unicode.Range32{Lo: 11136, Hi: 11147, Stride: 1 },
  },
  LatinOffset: 0,
}

// Modi: 79 codepoints
var Modi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11600, Hi: 11644, Stride: 1 },
    unicode.Range32{Lo: 11650, Hi: 11659, Stride: 1 },
  },
  LatinOffset: 0,
}

// Siddham: 92 codepoints
var Siddham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11580, Hi: 115B5, Stride: 1 },
    unicode.Range32{Lo: 115B8, Hi: 115DD, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tangsa: 89 codepoints
var Tangsa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 16A70, Hi: 16ABE, Stride: 1 },
    unicode.Range32{Lo: 16AC0, Hi: 16AC9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Rejang: 37 codepoints
var Rejang = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: A930, Hi: A953, Stride: 1 },
    unicode.Range16{Lo: A95F, Hi: A95F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Carian: 49 codepoints
var Carian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 102A0, Hi: 102D0, Stride: 1 },
  },
  LatinOffset: 0,
}

// Brahmi: 115 codepoints
var Brahmi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 11000, Hi: 1104D, Stride: 1 },
    unicode.Range32{Lo: 11052, Hi: 11075, Stride: 1 },
    unicode.Range32{Lo: 1107F, Hi: 1107F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Meroitic_Hieroglyphs: 32 codepoints
var Meroitic_Hieroglyphs = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10980, Hi: 1099F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Osage: 72 codepoints
var Osage = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 104B0, Hi: 104D3, Stride: 1 },
    unicode.Range32{Lo: 104D8, Hi: 104FB, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Permic: 43 codepoints
var Old_Permic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10350, Hi: 1037A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Kawi: 86 codepoints
var Kawi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 11F00, Hi: 11F10, Stride: 1 },
    unicode.Range32{Lo: 11F12, Hi: 11F3A, Stride: 1 },
    unicode.Range32{Lo: 11F3E, Hi: 11F59, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nag_Mundari: 42 codepoints
var Nag_Mundari = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 1E4D0, Hi: 1E4F9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Thai: 86 codepoints
var Thai = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 0E01, Hi: 0E3A, Stride: 1 },
    unicode.Range16{Lo: 0E40, Hi: 0E5B, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Imperial_Aramaic: 31 codepoints
var Imperial_Aramaic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10840, Hi: 10855, Stride: 1 },
    unicode.Range32{Lo: 10857, Hi: 1085F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_North_Arabian: 32 codepoints
var Old_North_Arabian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10A80, Hi: 10A9F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Inherited: 657 codepoints
var Inherited = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 18 */
    unicode.Range16{Lo: 0300, Hi: 036F, Stride: 1 },
    unicode.Range16{Lo: 0485, Hi: 0486, Stride: 1 },
    unicode.Range16{Lo: 064B, Hi: 0655, Stride: 1 },
    unicode.Range16{Lo: 0670, Hi: 0951, Stride: 737 },
    unicode.Range16{Lo: 0952, Hi: 0954, Stride: 1 },
    unicode.Range16{Lo: 1AB0, Hi: 1ACE, Stride: 1 },
    unicode.Range16{Lo: 1CD0, Hi: 1CD2, Stride: 1 },
    unicode.Range16{Lo: 1CD4, Hi: 1CE0, Stride: 1 },
    unicode.Range16{Lo: 1CE2, Hi: 1CE8, Stride: 1 },
    unicode.Range16{Lo: 1CED, Hi: 1CF4, Stride: 7 },
    unicode.Range16{Lo: 1CF8, Hi: 1CF9, Stride: 1 },
    unicode.Range16{Lo: 1DC0, Hi: 1DFF, Stride: 1 },
    unicode.Range16{Lo: 200C, Hi: 200D, Stride: 1 },
    unicode.Range16{Lo: 20D0, Hi: 20F0, Stride: 1 },
    unicode.Range16{Lo: 302A, Hi: 302D, Stride: 1 },
    unicode.Range16{Lo: 3099, Hi: 309A, Stride: 1 },
    unicode.Range16{Lo: FE00, Hi: FE0F, Stride: 1 },
    unicode.Range16{Lo: FE20, Hi: FE2D, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 9 */
    unicode.Range32{Lo: 101FD, Hi: 102E0, Stride: 227 },
    unicode.Range32{Lo: 1133B, Hi: 1CF00, Stride: 48069 },
    unicode.Range32{Lo: 1CF01, Hi: 1CF2D, Stride: 1 },
    unicode.Range32{Lo: 1CF30, Hi: 1CF46, Stride: 1 },
    unicode.Range32{Lo: 1D167, Hi: 1D169, Stride: 1 },
    unicode.Range32{Lo: 1D17B, Hi: 1D182, Stride: 1 },
    unicode.Range32{Lo: 1D185, Hi: 1D18B, Stride: 1 },
    unicode.Range32{Lo: 1D1AA, Hi: 1D1AD, Stride: 1 },
    unicode.Range32{Lo: E0100, Hi: E01EF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Duployan: 143 codepoints
var Duployan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 1BC00, Hi: 1BC6A, Stride: 1 },
    unicode.Range32{Lo: 1BC70, Hi: 1BC7C, Stride: 1 },
    unicode.Range32{Lo: 1BC80, Hi: 1BC88, Stride: 1 },
    unicode.Range32{Lo: 1BC90, Hi: 1BC99, Stride: 1 },
    unicode.Range32{Lo: 1BC9C, Hi: 1BC9F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Elymaic: 23 codepoints
var Elymaic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10FE0, Hi: 10FF6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tamil: 123 codepoints
var Tamil = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 15 */
    unicode.Range16{Lo: 0B82, Hi: 0B83, Stride: 1 },
    unicode.Range16{Lo: 0B85, Hi: 0B8A, Stride: 1 },
    unicode.Range16{Lo: 0B8E, Hi: 0B90, Stride: 1 },
    unicode.Range16{Lo: 0B92, Hi: 0B95, Stride: 1 },
    unicode.Range16{Lo: 0B99, Hi: 0B9A, Stride: 1 },
    unicode.Range16{Lo: 0B9C, Hi: 0B9E, Stride: 2 },
    unicode.Range16{Lo: 0B9F, Hi: 0BA3, Stride: 4 },
    unicode.Range16{Lo: 0BA4, Hi: 0BA8, Stride: 4 },
    unicode.Range16{Lo: 0BA9, Hi: 0BAA, Stride: 1 },
    unicode.Range16{Lo: 0BAE, Hi: 0BB9, Stride: 1 },
    unicode.Range16{Lo: 0BBE, Hi: 0BC2, Stride: 1 },
    unicode.Range16{Lo: 0BC6, Hi: 0BC8, Stride: 1 },
    unicode.Range16{Lo: 0BCA, Hi: 0BCD, Stride: 1 },
    unicode.Range16{Lo: 0BD0, Hi: 0BD7, Stride: 7 },
    unicode.Range16{Lo: 0BE6, Hi: 0BFA, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11FC0, Hi: 11FF1, Stride: 1 },
    unicode.Range32{Lo: 11FFF, Hi: 11FFF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Ogham: 29 codepoints
var Ogham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 1680, Hi: 169C, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Deseret: 80 codepoints
var Deseret = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10400, Hi: 1044F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Saurashtra: 82 codepoints
var Saurashtra = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: A880, Hi: A8C5, Stride: 1 },
    unicode.Range16{Lo: A8CE, Hi: A8D9, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Lisu: 49 codepoints
var Lisu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: A4D0, Hi: A4FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11FB0, Hi: 11FB0, Stride: 1 },
  },
  LatinOffset: 0,
}

// Yezidi: 47 codepoints
var Yezidi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 10E80, Hi: 10EA9, Stride: 1 },
    unicode.Range32{Lo: 10EAB, Hi: 10EAD, Stride: 1 },
    unicode.Range32{Lo: 10EB0, Hi: 10EB1, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tagalog: 23 codepoints
var Tagalog = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1700, Hi: 1715, Stride: 1 },
    unicode.Range16{Lo: 171F, Hi: 171F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Buhid: 20 codepoints
var Buhid = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 1740, Hi: 1753, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Old_Persian: 50 codepoints
var Old_Persian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 103A0, Hi: 103C3, Stride: 1 },
    unicode.Range32{Lo: 103C8, Hi: 103D5, Stride: 1 },
  },
  LatinOffset: 0,
}

// Warang_Citi: 84 codepoints
var Warang_Citi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 118A0, Hi: 118F2, Stride: 1 },
    unicode.Range32{Lo: 118FF, Hi: 118FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Cherokee: 172 codepoints
var Cherokee = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 13A0, Hi: 13F5, Stride: 1 },
    unicode.Range16{Lo: 13F8, Hi: 13FD, Stride: 1 },
    unicode.Range16{Lo: AB70, Hi: ABBF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Meroitic_Cursive: 90 codepoints
var Meroitic_Cursive = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 109A0, Hi: 109B7, Stride: 1 },
    unicode.Range32{Lo: 109BC, Hi: 109CF, Stride: 1 },
    unicode.Range32{Lo: 109D2, Hi: 109FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nabataean: 40 codepoints
var Nabataean = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10880, Hi: 1089E, Stride: 1 },
    unicode.Range32{Lo: 108A7, Hi: 108AF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Batak: 56 codepoints
var Batak = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1BC0, Hi: 1BF3, Stride: 1 },
    unicode.Range16{Lo: 1BFC, Hi: 1BFF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Manichaean: 51 codepoints
var Manichaean = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10AC0, Hi: 10AE6, Stride: 1 },
    unicode.Range32{Lo: 10AEB, Hi: 10AF6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sinhala: 111 codepoints
var Sinhala = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 12 */
    unicode.Range16{Lo: 0D81, Hi: 0D83, Stride: 1 },
    unicode.Range16{Lo: 0D85, Hi: 0D96, Stride: 1 },
    unicode.Range16{Lo: 0D9A, Hi: 0DB1, Stride: 1 },
    unicode.Range16{Lo: 0DB3, Hi: 0DBB, Stride: 1 },
    unicode.Range16{Lo: 0DBD, Hi: 0DC0, Stride: 3 },
    unicode.Range16{Lo: 0DC1, Hi: 0DC6, Stride: 1 },
    unicode.Range16{Lo: 0DCA, Hi: 0DCF, Stride: 5 },
    unicode.Range16{Lo: 0DD0, Hi: 0DD4, Stride: 1 },
    unicode.Range16{Lo: 0DD6, Hi: 0DD8, Stride: 2 },
    unicode.Range16{Lo: 0DD9, Hi: 0DDF, Stride: 1 },
    unicode.Range16{Lo: 0DE6, Hi: 0DEF, Stride: 1 },
    unicode.Range16{Lo: 0DF2, Hi: 0DF4, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 111E1, Hi: 111F4, Stride: 1 },
  },
  LatinOffset: 0,
}

// Canadian_Aboriginal: 726 codepoints
var Canadian_Aboriginal = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1400, Hi: 167F, Stride: 1 },
    unicode.Range16{Lo: 18B0, Hi: 18F5, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11AB0, Hi: 11ABF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Kaithi: 68 codepoints
var Kaithi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11080, Hi: 110C2, Stride: 1 },
    unicode.Range32{Lo: 110CD, Hi: 110CD, Stride: 1 },
  },
  LatinOffset: 0,
}

// Elbasan: 40 codepoints
var Elbasan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10500, Hi: 10527, Stride: 1 },
  },
  LatinOffset: 0,
}

// Dogra: 60 codepoints
var Dogra = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11800, Hi: 1183B, Stride: 1 },
  },
  LatinOffset: 0,
}

// Arabic: 1368 codepoints
var Arabic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 22 */
    unicode.Range16{Lo: 0600, Hi: 0604, Stride: 1 },
    unicode.Range16{Lo: 0606, Hi: 060B, Stride: 1 },
    unicode.Range16{Lo: 060D, Hi: 061A, Stride: 1 },
    unicode.Range16{Lo: 061C, Hi: 061E, Stride: 1 },
    unicode.Range16{Lo: 0620, Hi: 063F, Stride: 1 },
    unicode.Range16{Lo: 0641, Hi: 064A, Stride: 1 },
    unicode.Range16{Lo: 0656, Hi: 066F, Stride: 1 },
    unicode.Range16{Lo: 0671, Hi: 06DC, Stride: 1 },
    unicode.Range16{Lo: 06DE, Hi: 06FF, Stride: 1 },
    unicode.Range16{Lo: 0750, Hi: 077F, Stride: 1 },
    unicode.Range16{Lo: 0870, Hi: 088E, Stride: 1 },
    unicode.Range16{Lo: 0890, Hi: 0891, Stride: 1 },
    unicode.Range16{Lo: 0898, Hi: 08E1, Stride: 1 },
    unicode.Range16{Lo: 08E3, Hi: 08FF, Stride: 1 },
    unicode.Range16{Lo: FB50, Hi: FBC2, Stride: 1 },
    unicode.Range16{Lo: FBD3, Hi: FD3D, Stride: 1 },
    unicode.Range16{Lo: FD40, Hi: FD8F, Stride: 1 },
    unicode.Range16{Lo: FD92, Hi: FDC7, Stride: 1 },
    unicode.Range16{Lo: FDCF, Hi: FDF0, Stride: 33 },
    unicode.Range16{Lo: FDF1, Hi: FDFF, Stride: 1 },
    unicode.Range16{Lo: FE70, Hi: FE74, Stride: 1 },
    unicode.Range16{Lo: FE76, Hi: FEFC, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 27 */
    unicode.Range32{Lo: 10E60, Hi: 10E7E, Stride: 1 },
    unicode.Range32{Lo: 10EFD, Hi: 10EFF, Stride: 1 },
    unicode.Range32{Lo: 1EE00, Hi: 1EE03, Stride: 1 },
    unicode.Range32{Lo: 1EE05, Hi: 1EE1F, Stride: 1 },
    unicode.Range32{Lo: 1EE21, Hi: 1EE22, Stride: 1 },
    unicode.Range32{Lo: 1EE24, Hi: 1EE27, Stride: 3 },
    unicode.Range32{Lo: 1EE29, Hi: 1EE32, Stride: 1 },
    unicode.Range32{Lo: 1EE34, Hi: 1EE37, Stride: 1 },
    unicode.Range32{Lo: 1EE39, Hi: 1EE3B, Stride: 2 },
    unicode.Range32{Lo: 1EE42, Hi: 1EE47, Stride: 5 },
    unicode.Range32{Lo: 1EE49, Hi: 1EE4D, Stride: 2 },
    unicode.Range32{Lo: 1EE4E, Hi: 1EE4F, Stride: 1 },
    unicode.Range32{Lo: 1EE51, Hi: 1EE52, Stride: 1 },
    unicode.Range32{Lo: 1EE54, Hi: 1EE57, Stride: 3 },
    unicode.Range32{Lo: 1EE59, Hi: 1EE61, Stride: 2 },
    unicode.Range32{Lo: 1EE62, Hi: 1EE64, Stride: 2 },
    unicode.Range32{Lo: 1EE67, Hi: 1EE6A, Stride: 1 },
    unicode.Range32{Lo: 1EE6C, Hi: 1EE72, Stride: 1 },
    unicode.Range32{Lo: 1EE74, Hi: 1EE77, Stride: 1 },
    unicode.Range32{Lo: 1EE79, Hi: 1EE7C, Stride: 1 },
    unicode.Range32{Lo: 1EE7E, Hi: 1EE80, Stride: 2 },
    unicode.Range32{Lo: 1EE81, Hi: 1EE89, Stride: 1 },
    unicode.Range32{Lo: 1EE8B, Hi: 1EE9B, Stride: 1 },
    unicode.Range32{Lo: 1EEA1, Hi: 1EEA3, Stride: 1 },
    unicode.Range32{Lo: 1EEA5, Hi: 1EEA9, Stride: 1 },
    unicode.Range32{Lo: 1EEAB, Hi: 1EEBB, Stride: 1 },
    unicode.Range32{Lo: 1EEF0, Hi: 1EEF1, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sharada: 96 codepoints
var Sharada = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11180, Hi: 111DF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Caucasian_Albanian: 53 codepoints
var Caucasian_Albanian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10530, Hi: 10563, Stride: 1 },
    unicode.Range32{Lo: 1056F, Hi: 1056F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Bassa_Vah: 36 codepoints
var Bassa_Vah = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 16AD0, Hi: 16AED, Stride: 1 },
    unicode.Range32{Lo: 16AF0, Hi: 16AF5, Stride: 1 },
  },
  LatinOffset: 0,
}

// Makasar: 25 codepoints
var Makasar = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11EE0, Hi: 11EF8, Stride: 1 },
  },
  LatinOffset: 0,
}

// Cypro_Minoan: 99 codepoints
var Cypro_Minoan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 12F90, Hi: 12FF2, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gothic: 27 codepoints
var Gothic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10330, Hi: 1034A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Lycian: 29 codepoints
var Lycian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10280, Hi: 1029C, Stride: 1 },
  },
  LatinOffset: 0,
}

// Egyptian_Hieroglyphs: 1110 codepoints
var Egyptian_Hieroglyphs = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 13000, Hi: 13455, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tangut: 6914 codepoints
var Tangut = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 16FE0, Hi: 17000, Stride: 32 },
    unicode.Range32{Lo: 17001, Hi: 187F7, Stride: 1 },
    unicode.Range32{Lo: 18800, Hi: 18AFF, Stride: 1 },
    unicode.Range32{Lo: 18D00, Hi: 18D08, Stride: 1 },
  },
  LatinOffset: 0,
}

// Dives_Akuru: 72 codepoints
var Dives_Akuru = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 8 */
    unicode.Range32{Lo: 11900, Hi: 11906, Stride: 1 },
    unicode.Range32{Lo: 11909, Hi: 1190C, Stride: 3 },
    unicode.Range32{Lo: 1190D, Hi: 11913, Stride: 1 },
    unicode.Range32{Lo: 11915, Hi: 11916, Stride: 1 },
    unicode.Range32{Lo: 11918, Hi: 11935, Stride: 1 },
    unicode.Range32{Lo: 11937, Hi: 11938, Stride: 1 },
    unicode.Range32{Lo: 1193B, Hi: 11946, Stride: 1 },
    unicode.Range32{Lo: 11950, Hi: 11959, Stride: 1 },
  },
  LatinOffset: 0,
}

// Khitan_Small_Script: 471 codepoints
var Khitan_Small_Script = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 16FE4, Hi: 18B00, Stride: 6940 },
    unicode.Range32{Lo: 18B01, Hi: 18CD5, Stride: 1 },
  },
  LatinOffset: 0,
}

// Glagolitic: 134 codepoints
var Glagolitic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 2C00, Hi: 2C5F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 1E000, Hi: 1E006, Stride: 1 },
    unicode.Range32{Lo: 1E008, Hi: 1E018, Stride: 1 },
    unicode.Range32{Lo: 1E01B, Hi: 1E021, Stride: 1 },
    unicode.Range32{Lo: 1E023, Hi: 1E024, Stride: 1 },
    unicode.Range32{Lo: 1E026, Hi: 1E02A, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tai_Tham: 127 codepoints
var Tai_Tham = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 5 */
    unicode.Range16{Lo: 1A20, Hi: 1A5E, Stride: 1 },
    unicode.Range16{Lo: 1A60, Hi: 1A7C, Stride: 1 },
    unicode.Range16{Lo: 1A7F, Hi: 1A89, Stride: 1 },
    unicode.Range16{Lo: 1A90, Hi: 1A99, Stride: 1 },
    unicode.Range16{Lo: 1AA0, Hi: 1AAD, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tirhuta: 82 codepoints
var Tirhuta = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11480, Hi: 114C7, Stride: 1 },
    unicode.Range32{Lo: 114D0, Hi: 114D9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Takri: 68 codepoints
var Takri = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11680, Hi: 116B9, Stride: 1 },
    unicode.Range32{Lo: 116C0, Hi: 116C9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Hanifi_Rohingya: 50 codepoints
var Hanifi_Rohingya = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10D00, Hi: 10D27, Stride: 1 },
    unicode.Range32{Lo: 10D30, Hi: 10D39, Stride: 1 },
  },
  LatinOffset: 0,
}

// Chorasmian: 28 codepoints
var Chorasmian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10FB0, Hi: 10FCB, Stride: 1 },
  },
  LatinOffset: 0,
}

// Devanagari: 164 codepoints
var Devanagari = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 0900, Hi: 0950, Stride: 1 },
    unicode.Range16{Lo: 0955, Hi: 0963, Stride: 1 },
    unicode.Range16{Lo: 0966, Hi: 097F, Stride: 1 },
    unicode.Range16{Lo: A8E0, Hi: A8FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11B00, Hi: 11B09, Stride: 1 },
  },
  LatinOffset: 0,
}

// Georgian: 173 codepoints
var Georgian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 8 */
    unicode.Range16{Lo: 10A0, Hi: 10C5, Stride: 1 },
    unicode.Range16{Lo: 10C7, Hi: 10CD, Stride: 6 },
    unicode.Range16{Lo: 10D0, Hi: 10FA, Stride: 1 },
    unicode.Range16{Lo: 10FC, Hi: 10FF, Stride: 1 },
    unicode.Range16{Lo: 1C90, Hi: 1CBA, Stride: 1 },
    unicode.Range16{Lo: 1CBD, Hi: 1CBF, Stride: 1 },
    unicode.Range16{Lo: 2D00, Hi: 2D25, Stride: 1 },
    unicode.Range16{Lo: 2D27, Hi: 2D2D, Stride: 6 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Hangul: 11739 codepoints
var Hangul = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 1100, Hi: 11FF, Stride: 1 },
    unicode.Range16{Lo: 302E, Hi: 302F, Stride: 1 },
    unicode.Range16{Lo: 3131, Hi: 318E, Stride: 1 },
    unicode.Range16{Lo: 3200, Hi: 321E, Stride: 1 },
    unicode.Range16{Lo: 3260, Hi: 327E, Stride: 1 },
    unicode.Range16{Lo: A960, Hi: A97C, Stride: 1 },
    unicode.Range16{Lo: AC00, Hi: D7A3, Stride: 1 },
    unicode.Range16{Lo: D7B0, Hi: D7C6, Stride: 1 },
    unicode.Range16{Lo: D7CB, Hi: D7FB, Stride: 1 },
    unicode.Range16{Lo: FFA0, Hi: FFBE, Stride: 1 },
    unicode.Range16{Lo: FFC2, Hi: FFC7, Stride: 1 },
    unicode.Range16{Lo: FFCA, Hi: FFCF, Stride: 1 },
    unicode.Range16{Lo: FFD2, Hi: FFD7, Stride: 1 },
    unicode.Range16{Lo: FFDA, Hi: FFDC, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Mende_Kikakui: 213 codepoints
var Mende_Kikakui = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 1E800, Hi: 1E8C4, Stride: 1 },
    unicode.Range32{Lo: 1E8C7, Hi: 1E8D6, Stride: 1 },
  },
  LatinOffset: 0,
}

// Adlam: 88 codepoints
var Adlam = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 1E900, Hi: 1E94B, Stride: 1 },
    unicode.Range32{Lo: 1E950, Hi: 1E959, Stride: 1 },
    unicode.Range32{Lo: 1E95E, Hi: 1E95F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Limbu: 68 codepoints
var Limbu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 5 */
    unicode.Range16{Lo: 1900, Hi: 191E, Stride: 1 },
    unicode.Range16{Lo: 1920, Hi: 192B, Stride: 1 },
    unicode.Range16{Lo: 1930, Hi: 193B, Stride: 1 },
    unicode.Range16{Lo: 1940, Hi: 1944, Stride: 4 },
    unicode.Range16{Lo: 1945, Hi: 194F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Balinese: 124 codepoints
var Balinese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1B00, Hi: 1B4C, Stride: 1 },
    unicode.Range16{Lo: 1B50, Hi: 1B7E, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Lydian: 27 codepoints
var Lydian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10920, Hi: 10939, Stride: 1 },
    unicode.Range32{Lo: 1093F, Hi: 1093F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Avestan: 61 codepoints
var Avestan = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 10B00, Hi: 10B35, Stride: 1 },
    unicode.Range32{Lo: 10B39, Hi: 10B3F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Old_Hungarian: 108 codepoints
var Old_Hungarian = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 10C80, Hi: 10CB2, Stride: 1 },
    unicode.Range32{Lo: 10CC0, Hi: 10CF2, Stride: 1 },
    unicode.Range32{Lo: 10CFA, Hi: 10CFF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Soyombo: 83 codepoints
var Soyombo = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11A50, Hi: 11AA2, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gunjala_Gondi: 63 codepoints
var Gunjala_Gondi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 6 */
    unicode.Range32{Lo: 11D60, Hi: 11D65, Stride: 1 },
    unicode.Range32{Lo: 11D67, Hi: 11D68, Stride: 1 },
    unicode.Range32{Lo: 11D6A, Hi: 11D8E, Stride: 1 },
    unicode.Range32{Lo: 11D90, Hi: 11D91, Stride: 1 },
    unicode.Range32{Lo: 11D93, Hi: 11D98, Stride: 1 },
    unicode.Range32{Lo: 11DA0, Hi: 11DA9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Tagbanwa: 18 codepoints
var Tagbanwa = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 1760, Hi: 176C, Stride: 1 },
    unicode.Range16{Lo: 176E, Hi: 1770, Stride: 1 },
    unicode.Range16{Lo: 1772, Hi: 1773, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Kharoshthi: 68 codepoints
var Kharoshthi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 8 */
    unicode.Range32{Lo: 10A00, Hi: 10A03, Stride: 1 },
    unicode.Range32{Lo: 10A05, Hi: 10A06, Stride: 1 },
    unicode.Range32{Lo: 10A0C, Hi: 10A13, Stride: 1 },
    unicode.Range32{Lo: 10A15, Hi: 10A17, Stride: 1 },
    unicode.Range32{Lo: 10A19, Hi: 10A35, Stride: 1 },
    unicode.Range32{Lo: 10A38, Hi: 10A3A, Stride: 1 },
    unicode.Range32{Lo: 10A3F, Hi: 10A48, Stride: 1 },
    unicode.Range32{Lo: 10A50, Hi: 10A58, Stride: 1 },
  },
  LatinOffset: 0,
}

// Sundanese: 72 codepoints
var Sundanese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1B80, Hi: 1BBF, Stride: 1 },
    unicode.Range16{Lo: 1CC0, Hi: 1CC7, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Kayah_Li: 47 codepoints
var Kayah_Li = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: A900, Hi: A92D, Stride: 1 },
    unicode.Range16{Lo: A92F, Hi: A92F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Old_Turkic: 73 codepoints
var Old_Turkic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 10C00, Hi: 10C48, Stride: 1 },
  },
  LatinOffset: 0,
}

// Khojki: 65 codepoints
var Khojki = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 11200, Hi: 11211, Stride: 1 },
    unicode.Range32{Lo: 11213, Hi: 11241, Stride: 1 },
  },
  LatinOffset: 0,
}

// Wancho: 59 codepoints
var Wancho = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 1E2C0, Hi: 1E2F9, Stride: 1 },
    unicode.Range32{Lo: 1E2FF, Hi: 1E2FF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gujarati: 91 codepoints
var Gujarati = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0A81, Hi: 0A83, Stride: 1 },
    unicode.Range16{Lo: 0A85, Hi: 0A8D, Stride: 1 },
    unicode.Range16{Lo: 0A8F, Hi: 0A91, Stride: 1 },
    unicode.Range16{Lo: 0A93, Hi: 0AA8, Stride: 1 },
    unicode.Range16{Lo: 0AAA, Hi: 0AB0, Stride: 1 },
    unicode.Range16{Lo: 0AB2, Hi: 0AB3, Stride: 1 },
    unicode.Range16{Lo: 0AB5, Hi: 0AB9, Stride: 1 },
    unicode.Range16{Lo: 0ABC, Hi: 0AC5, Stride: 1 },
    unicode.Range16{Lo: 0AC7, Hi: 0AC9, Stride: 1 },
    unicode.Range16{Lo: 0ACB, Hi: 0ACD, Stride: 1 },
    unicode.Range16{Lo: 0AD0, Hi: 0AE0, Stride: 16 },
    unicode.Range16{Lo: 0AE1, Hi: 0AE3, Stride: 1 },
    unicode.Range16{Lo: 0AE6, Hi: 0AF1, Stride: 1 },
    unicode.Range16{Lo: 0AF9, Hi: 0AFF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Oriya: 91 codepoints
var Oriya = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 14 */
    unicode.Range16{Lo: 0B01, Hi: 0B03, Stride: 1 },
    unicode.Range16{Lo: 0B05, Hi: 0B0C, Stride: 1 },
    unicode.Range16{Lo: 0B0F, Hi: 0B10, Stride: 1 },
    unicode.Range16{Lo: 0B13, Hi: 0B28, Stride: 1 },
    unicode.Range16{Lo: 0B2A, Hi: 0B30, Stride: 1 },
    unicode.Range16{Lo: 0B32, Hi: 0B33, Stride: 1 },
    unicode.Range16{Lo: 0B35, Hi: 0B39, Stride: 1 },
    unicode.Range16{Lo: 0B3C, Hi: 0B44, Stride: 1 },
    unicode.Range16{Lo: 0B47, Hi: 0B48, Stride: 1 },
    unicode.Range16{Lo: 0B4B, Hi: 0B4D, Stride: 1 },
    unicode.Range16{Lo: 0B55, Hi: 0B57, Stride: 1 },
    unicode.Range16{Lo: 0B5C, Hi: 0B5D, Stride: 1 },
    unicode.Range16{Lo: 0B5F, Hi: 0B63, Stride: 1 },
    unicode.Range16{Lo: 0B66, Hi: 0B77, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Coptic: 137 codepoints
var Coptic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 03E2, Hi: 03EF, Stride: 1 },
    unicode.Range16{Lo: 2C80, Hi: 2CF3, Stride: 1 },
    unicode.Range16{Lo: 2CF9, Hi: 2CFF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Masaram_Gondi: 75 codepoints
var Masaram_Gondi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 7 */
    unicode.Range32{Lo: 11D00, Hi: 11D06, Stride: 1 },
    unicode.Range32{Lo: 11D08, Hi: 11D09, Stride: 1 },
    unicode.Range32{Lo: 11D0B, Hi: 11D36, Stride: 1 },
    unicode.Range32{Lo: 11D3A, Hi: 11D3C, Stride: 2 },
    unicode.Range32{Lo: 11D3D, Hi: 11D3F, Stride: 2 },
    unicode.Range32{Lo: 11D40, Hi: 11D47, Stride: 1 },
    unicode.Range32{Lo: 11D50, Hi: 11D59, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nandinagari: 65 codepoints
var Nandinagari = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 119A0, Hi: 119A7, Stride: 1 },
    unicode.Range32{Lo: 119AA, Hi: 119D7, Stride: 1 },
    unicode.Range32{Lo: 119DA, Hi: 119E4, Stride: 1 },
  },
  LatinOffset: 0,
}

// Bopomofo: 77 codepoints
var Bopomofo = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: 02EA, Hi: 02EB, Stride: 1 },
    unicode.Range16{Lo: 3105, Hi: 312F, Stride: 1 },
    unicode.Range16{Lo: 31A0, Hi: 31BF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Tai_Le: 35 codepoints
var Tai_Le = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 1950, Hi: 196D, Stride: 1 },
    unicode.Range16{Lo: 1970, Hi: 1974, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Braille: 256 codepoints
var Braille = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 2800, Hi: 28FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Telugu: 100 codepoints
var Telugu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 13 */
    unicode.Range16{Lo: 0C00, Hi: 0C0C, Stride: 1 },
    unicode.Range16{Lo: 0C0E, Hi: 0C10, Stride: 1 },
    unicode.Range16{Lo: 0C12, Hi: 0C28, Stride: 1 },
    unicode.Range16{Lo: 0C2A, Hi: 0C39, Stride: 1 },
    unicode.Range16{Lo: 0C3C, Hi: 0C44, Stride: 1 },
    unicode.Range16{Lo: 0C46, Hi: 0C48, Stride: 1 },
    unicode.Range16{Lo: 0C4A, Hi: 0C4D, Stride: 1 },
    unicode.Range16{Lo: 0C55, Hi: 0C56, Stride: 1 },
    unicode.Range16{Lo: 0C58, Hi: 0C5A, Stride: 1 },
    unicode.Range16{Lo: 0C5D, Hi: 0C60, Stride: 3 },
    unicode.Range16{Lo: 0C61, Hi: 0C63, Stride: 1 },
    unicode.Range16{Lo: 0C66, Hi: 0C6F, Stride: 1 },
    unicode.Range16{Lo: 0C77, Hi: 0C7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Sora_Sompeng: 35 codepoints
var Sora_Sompeng = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 110D0, Hi: 110E8, Stride: 1 },
    unicode.Range32{Lo: 110F0, Hi: 110F9, Stride: 1 },
  },
  LatinOffset: 0,
}

// Grantha: 85 codepoints
var Grantha = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 14 */
    unicode.Range32{Lo: 11300, Hi: 11303, Stride: 1 },
    unicode.Range32{Lo: 11305, Hi: 1130C, Stride: 1 },
    unicode.Range32{Lo: 1130F, Hi: 11310, Stride: 1 },
    unicode.Range32{Lo: 11313, Hi: 11328, Stride: 1 },
    unicode.Range32{Lo: 1132A, Hi: 11330, Stride: 1 },
    unicode.Range32{Lo: 11332, Hi: 11333, Stride: 1 },
    unicode.Range32{Lo: 11335, Hi: 11339, Stride: 1 },
    unicode.Range32{Lo: 1133C, Hi: 11344, Stride: 1 },
    unicode.Range32{Lo: 11347, Hi: 11348, Stride: 1 },
    unicode.Range32{Lo: 1134B, Hi: 1134D, Stride: 1 },
    unicode.Range32{Lo: 11350, Hi: 11357, Stride: 7 },
    unicode.Range32{Lo: 1135D, Hi: 11363, Stride: 1 },
    unicode.Range32{Lo: 11366, Hi: 1136C, Stride: 1 },
    unicode.Range32{Lo: 11370, Hi: 11374, Stride: 1 },
  },
  LatinOffset: 0,
}

// Anatolian_Hieroglyphs: 583 codepoints
var Anatolian_Hieroglyphs = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 14400, Hi: 14646, Stride: 1 },
  },
  LatinOffset: 0,
}

// Nushu: 397 codepoints
var Nushu = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 2 */
    unicode.Range32{Lo: 16FE1, Hi: 1B170, Stride: 16783 },
    unicode.Range32{Lo: 1B171, Hi: 1B2FB, Stride: 1 },
  },
  LatinOffset: 0,
}

// Gurmukhi: 80 codepoints
var Gurmukhi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 16 */
    unicode.Range16{Lo: 0A01, Hi: 0A03, Stride: 1 },
    unicode.Range16{Lo: 0A05, Hi: 0A0A, Stride: 1 },
    unicode.Range16{Lo: 0A0F, Hi: 0A10, Stride: 1 },
    unicode.Range16{Lo: 0A13, Hi: 0A28, Stride: 1 },
    unicode.Range16{Lo: 0A2A, Hi: 0A30, Stride: 1 },
    unicode.Range16{Lo: 0A32, Hi: 0A33, Stride: 1 },
    unicode.Range16{Lo: 0A35, Hi: 0A36, Stride: 1 },
    unicode.Range16{Lo: 0A38, Hi: 0A39, Stride: 1 },
    unicode.Range16{Lo: 0A3C, Hi: 0A3E, Stride: 2 },
    unicode.Range16{Lo: 0A3F, Hi: 0A42, Stride: 1 },
    unicode.Range16{Lo: 0A47, Hi: 0A48, Stride: 1 },
    unicode.Range16{Lo: 0A4B, Hi: 0A4D, Stride: 1 },
    unicode.Range16{Lo: 0A51, Hi: 0A59, Stride: 8 },
    unicode.Range16{Lo: 0A5A, Hi: 0A5C, Stride: 1 },
    unicode.Range16{Lo: 0A5E, Hi: 0A66, Stride: 8 },
    unicode.Range16{Lo: 0A67, Hi: 0A76, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Cuneiform: 1234 codepoints
var Cuneiform = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 4 */
    unicode.Range32{Lo: 12000, Hi: 12399, Stride: 1 },
    unicode.Range32{Lo: 12400, Hi: 1246E, Stride: 1 },
    unicode.Range32{Lo: 12470, Hi: 12474, Stride: 1 },
    unicode.Range32{Lo: 12480, Hi: 12543, Stride: 1 },
  },
  LatinOffset: 0,
}

// Ol_Chiki: 48 codepoints
var Ol_Chiki = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 1 */
    unicode.Range16{Lo: 1C50, Hi: 1C7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Runic: 86 codepoints
var Runic = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 16A0, Hi: 16EA, Stride: 1 },
    unicode.Range16{Lo: 16EE, Hi: 16F8, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Khmer: 146 codepoints
var Khmer = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 4 */
    unicode.Range16{Lo: 1780, Hi: 17DD, Stride: 1 },
    unicode.Range16{Lo: 17E0, Hi: 17E9, Stride: 1 },
    unicode.Range16{Lo: 17F0, Hi: 17F9, Stride: 1 },
    unicode.Range16{Lo: 19E0, Hi: 19FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Katakana: 321 codepoints
var Katakana = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 30A1, Hi: 30FA, Stride: 1 },
    unicode.Range16{Lo: 30FD, Hi: 30FF, Stride: 1 },
    unicode.Range16{Lo: 31F0, Hi: 31FF, Stride: 1 },
    unicode.Range16{Lo: 32D0, Hi: 32FE, Stride: 1 },
    unicode.Range16{Lo: 3300, Hi: 3357, Stride: 1 },
    unicode.Range16{Lo: FF66, Hi: FF6F, Stride: 1 },
    unicode.Range16{Lo: FF71, Hi: FF9D, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 7 */
    unicode.Range32{Lo: 1AFF0, Hi: 1AFF3, Stride: 1 },
    unicode.Range32{Lo: 1AFF5, Hi: 1AFFB, Stride: 1 },
    unicode.Range32{Lo: 1AFFD, Hi: 1AFFE, Stride: 1 },
    unicode.Range32{Lo: 1B000, Hi: 1B120, Stride: 288 },
    unicode.Range32{Lo: 1B121, Hi: 1B122, Stride: 1 },
    unicode.Range32{Lo: 1B155, Hi: 1B164, Stride: 15 },
    unicode.Range32{Lo: 1B165, Hi: 1B167, Stride: 1 },
  },
  LatinOffset: 0,
}

// Pahawh_Hmong: 127 codepoints
var Pahawh_Hmong = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 5 */
    unicode.Range32{Lo: 16B00, Hi: 16B45, Stride: 1 },
    unicode.Range32{Lo: 16B50, Hi: 16B59, Stride: 1 },
    unicode.Range32{Lo: 16B5B, Hi: 16B61, Stride: 1 },
    unicode.Range32{Lo: 16B63, Hi: 16B77, Stride: 1 },
    unicode.Range32{Lo: 16B7D, Hi: 16B8F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Zanabazar_Square: 72 codepoints
var Zanabazar_Square = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11A00, Hi: 11A47, Stride: 1 },
  },
  LatinOffset: 0,
}

// Pau_Cin_Hau: 57 codepoints
var Pau_Cin_Hau = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 11AC0, Hi: 11AF8, Stride: 1 },
  },
  LatinOffset: 0,
}

// Toto: 31 codepoints
var Toto = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 1 */
    unicode.Range32{Lo: 1E290, Hi: 1E2AE, Stride: 1 },
  },
  LatinOffset: 0,
}

// Vithkuqi: 70 codepoints
var Vithkuqi = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 8 */
    unicode.Range32{Lo: 10570, Hi: 1057A, Stride: 1 },
    unicode.Range32{Lo: 1057C, Hi: 1058A, Stride: 1 },
    unicode.Range32{Lo: 1058C, Hi: 10592, Stride: 1 },
    unicode.Range32{Lo: 10594, Hi: 10595, Stride: 1 },
    unicode.Range32{Lo: 10597, Hi: 105A1, Stride: 1 },
    unicode.Range32{Lo: 105A3, Hi: 105B1, Stride: 1 },
    unicode.Range32{Lo: 105B3, Hi: 105B9, Stride: 1 },
    unicode.Range32{Lo: 105BB, Hi: 105BC, Stride: 1 },
  },
  LatinOffset: 0,
}

// Malayalam: 118 codepoints
var Malayalam = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 7 */
    unicode.Range16{Lo: 0D00, Hi: 0D0C, Stride: 1 },
    unicode.Range16{Lo: 0D0E, Hi: 0D10, Stride: 1 },
    unicode.Range16{Lo: 0D12, Hi: 0D44, Stride: 1 },
    unicode.Range16{Lo: 0D46, Hi: 0D48, Stride: 1 },
    unicode.Range16{Lo: 0D4A, Hi: 0D4F, Stride: 1 },
    unicode.Range16{Lo: 0D54, Hi: 0D63, Stride: 1 },
    unicode.Range16{Lo: 0D66, Hi: 0D7F, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Nko: 62 codepoints
var Nko = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 2 */
    unicode.Range16{Lo: 07C0, Hi: 07FA, Stride: 1 },
    unicode.Range16{Lo: 07FD, Hi: 07FF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}

// Mro: 43 codepoints
var Mro = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 0 */},
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 16A40, Hi: 16A5E, Stride: 1 },
    unicode.Range32{Lo: 16A60, Hi: 16A69, Stride: 1 },
    unicode.Range32{Lo: 16A6E, Hi: 16A6F, Stride: 1 },
  },
  LatinOffset: 0,
}

// Greek: 518 codepoints
var Greek = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 29 */
    unicode.Range16{Lo: 0370, Hi: 0373, Stride: 1 },
    unicode.Range16{Lo: 0375, Hi: 0377, Stride: 1 },
    unicode.Range16{Lo: 037A, Hi: 037D, Stride: 1 },
    unicode.Range16{Lo: 037F, Hi: 0384, Stride: 5 },
    unicode.Range16{Lo: 0386, Hi: 0388, Stride: 2 },
    unicode.Range16{Lo: 0389, Hi: 038A, Stride: 1 },
    unicode.Range16{Lo: 038C, Hi: 038E, Stride: 2 },
    unicode.Range16{Lo: 038F, Hi: 03A1, Stride: 1 },
    unicode.Range16{Lo: 03A3, Hi: 03E1, Stride: 1 },
    unicode.Range16{Lo: 03F0, Hi: 03FF, Stride: 1 },
    unicode.Range16{Lo: 1D26, Hi: 1D2A, Stride: 1 },
    unicode.Range16{Lo: 1D5D, Hi: 1D61, Stride: 1 },
    unicode.Range16{Lo: 1D66, Hi: 1D6A, Stride: 1 },
    unicode.Range16{Lo: 1DBF, Hi: 1F00, Stride: 321 },
    unicode.Range16{Lo: 1F01, Hi: 1F15, Stride: 1 },
    unicode.Range16{Lo: 1F18, Hi: 1F1D, Stride: 1 },
    unicode.Range16{Lo: 1F20, Hi: 1F45, Stride: 1 },
    unicode.Range16{Lo: 1F48, Hi: 1F4D, Stride: 1 },
    unicode.Range16{Lo: 1F50, Hi: 1F57, Stride: 1 },
    unicode.Range16{Lo: 1F59, Hi: 1F5F, Stride: 2 },
    unicode.Range16{Lo: 1F60, Hi: 1F7D, Stride: 1 },
    unicode.Range16{Lo: 1F80, Hi: 1FB4, Stride: 1 },
    unicode.Range16{Lo: 1FB6, Hi: 1FC4, Stride: 1 },
    unicode.Range16{Lo: 1FC6, Hi: 1FD3, Stride: 1 },
    unicode.Range16{Lo: 1FD6, Hi: 1FDB, Stride: 1 },
    unicode.Range16{Lo: 1FDD, Hi: 1FEF, Stride: 1 },
    unicode.Range16{Lo: 1FF2, Hi: 1FF4, Stride: 1 },
    unicode.Range16{Lo: 1FF6, Hi: 1FFE, Stride: 1 },
    unicode.Range16{Lo: 2126, Hi: AB65, Stride: 35391 },
  },
  R32: []unicode.Range32{ /* 3 */
    unicode.Range32{Lo: 10140, Hi: 1018E, Stride: 1 },
    unicode.Range32{Lo: 101A0, Hi: 1D200, Stride: 53344 },
    unicode.Range32{Lo: 1D201, Hi: 1D245, Stride: 1 },
  },
  LatinOffset: 0,
}

// Han: 98408 codepoints
var Han = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 10 */
    unicode.Range16{Lo: 2E80, Hi: 2E99, Stride: 1 },
    unicode.Range16{Lo: 2E9B, Hi: 2EF3, Stride: 1 },
    unicode.Range16{Lo: 2F00, Hi: 2FD5, Stride: 1 },
    unicode.Range16{Lo: 3005, Hi: 3007, Stride: 2 },
    unicode.Range16{Lo: 3021, Hi: 3029, Stride: 1 },
    unicode.Range16{Lo: 3038, Hi: 303B, Stride: 1 },
    unicode.Range16{Lo: 3400, Hi: 4DBF, Stride: 1 },
    unicode.Range16{Lo: 4E00, Hi: 9FFF, Stride: 1 },
    unicode.Range16{Lo: F900, Hi: FA6D, Stride: 1 },
    unicode.Range16{Lo: FA70, Hi: FAD9, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 10 */
    unicode.Range32{Lo: 16FE2, Hi: 16FE3, Stride: 1 },
    unicode.Range32{Lo: 16FF0, Hi: 16FF1, Stride: 1 },
    unicode.Range32{Lo: 20000, Hi: 2A6DF, Stride: 1 },
    unicode.Range32{Lo: 2A700, Hi: 2B739, Stride: 1 },
    unicode.Range32{Lo: 2B740, Hi: 2B81D, Stride: 1 },
    unicode.Range32{Lo: 2B820, Hi: 2CEA1, Stride: 1 },
    unicode.Range32{Lo: 2CEB0, Hi: 2EBE0, Stride: 1 },
    unicode.Range32{Lo: 2F800, Hi: 2FA1D, Stride: 1 },
    unicode.Range32{Lo: 30000, Hi: 3134A, Stride: 1 },
    unicode.Range32{Lo: 31350, Hi: 323AF, Stride: 1 },
  },
  LatinOffset: 0,
}

// Javanese: 90 codepoints
var Javanese = &unicode.RangeTable{
  R16: []unicode.Range16{ /* 3 */
    unicode.Range16{Lo: A980, Hi: A9CD, Stride: 1 },
    unicode.Range16{Lo: A9D0, Hi: A9D9, Stride: 1 },
    unicode.Range16{Lo: A9DE, Hi: A9DF, Stride: 1 },
  },
  R32: []unicode.Range32{ /* 0 */},
  LatinOffset: 0,
}
