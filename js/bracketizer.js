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
            // this will also catch the case of a closeing bracket preceding all opening brackets
            if (closeBrackets > openBrackets) {
                return false
            }
        }
        
        // if, after checking the entire string, opening brackets == close brackets - return true, or clean
        // else we know that not all brackets are matched in the string - return false, or dirty
        return openBrackets == closeBrackets
    }
}