module.exports = {
    // sweepForBrackets - evaluates a string for appropriately paired brackets, if not paired == false, if paired == true
    sweepForBrackets: (doWeBracket) => {

        let openBrackets = 0
        let closeBrackets = 0

        for (let s of doWeBracket) {
            if (s == "{") {
                openBrackets++
            } else if (s == "}") {
                closeBrackets++
            }
            // if there are ever more closing brackets than opening we know the string is dirty
            if (closeBrackets > openBrackets) {
                return false
            }
        }
        return openBrackets == closeBrackets
    }
}