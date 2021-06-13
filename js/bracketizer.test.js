const bracketizer = require("./bracketizer");

const bracketizeUs = [
    "{}",  // 0
    "}{",  // 1
    "{{}", // 2
    "",    // 3
    "     ",    // 4
    "F{{DF}SDF{SDFSD}SDF{SDF}SDFSD{FSD}FSD{FS{DFSD}SD{FSDF}SDF}SD}FSD{FSDF}",        // 5
    "F{{DF}SDF{SDFSD}SDF{SDF}SD{F{S{D{FSD}F{SD{FS{DFSD}SD{FSDF}SDF}SD}FSD{FS}D}F}}", // 6
    "  F{{ DF}SDF{SD FSD} SDF{SDF}SDFSD{    FSD}FSD{FS{DFSD} SD{FSDF}SD F}SD}FSD{FSD F} ",        // 7
    "  F{{DF}SDF{SDFSD}SDF{SDF}S   D{F{S{D{FSD} F{SD{FS{DFSD}SD{FSDF }SDF}SD}FSD{ FS}D}F}}  ", // 8
]

// 0
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[0])).toBeTruthy()
})

// 1
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[1])).toBeFalsy()
})

// 2
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[2])).toBeFalsy()
})

// 3
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[3])).toBeTruthy()
})

// 4
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[4])).toBeTruthy()
})

// 5
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[5])).toBeTruthy()
})

// 6
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[6])).toBeFalsy()
})

// 7
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[7])).toBeTruthy()
})

// 8
test('checks if string is clean or dirty', () => {
    expect(bracketizer.sweepForBrackets(bracketizeUs[8])).toBeFalsy()
})