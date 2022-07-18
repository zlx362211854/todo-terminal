#!/usr/bin/env node
var child_process = require("child_process");
const os = require('os');
var path = require('path')
var args = process.argv.slice(2);

var argsWithQuotes = args.map(function (a) {
    if (/\ /.test(a)) {
        return `"${a}"`
    }
    return a
})

var argStr = ' ' + argsWithQuotes.join(' ')
try {
    if (os.type() == 'Windows_NT') {
        //windows
        child_process.execSync(path.resolve(__dirname + '/main.exe') + argStr, { stdio: 'inherit' })
    } else {
        child_process.execSync(path.resolve(__dirname + '/main') + argStr, { stdio: 'inherit' })
    }
} catch (err) {
    console.error("An error occurred. Please try again")
}
